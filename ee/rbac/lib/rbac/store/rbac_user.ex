defmodule Rbac.Store.RbacUser do
  require Logger

  import Ecto.Query

  alias Rbac.Repo.{RbacUser, Subject}
  alias Ecto.Multi

  @spec fetch_by_oidc_id(String.t()) :: {:ok, RbacUser.t()} | {:error, :not_found}
  def fetch_by_oidc_id(oidc_user_id) do
    RbacUser
    |> join(:inner, [u], o in assoc(u, :oidc_users))
    |> join(:inner, [u, o], s in assoc(u, :subject))
    |> where([u, o, _], o.oidc_user_id == ^oidc_user_id)
    |> select([u, o, s], %{u | name: s.name})
    |> Rbac.Repo.one()
    |> case do
      nil -> {:error, :not_found}
      user -> {:ok, user}
    end
  end

  @spec fetch_by_email(String.t()) :: {:ok, RbacUser.t()} | {:error, :not_found}
  def fetch_by_email(email) do
    RbacUser
    |> where([u], u.email == ^email)
    |> join(:inner, [u], s in assoc(u, :subject))
    |> select([u, s], %{u | name: s.name})
    |> Rbac.Repo.one()
    |> case do
      nil -> {:error, :not_found}
      user -> {:ok, user}
    end
  end

  @spec fetch(Ecto.UUID.t()) :: RbacUser.t() | nil
  def fetch(user_id) do
    RbacUser
    |> where([user], user.id == ^user_id)
    |> join(:inner, [u], s in assoc(u, :subject))
    |> select([u, s], %{u | name: s.name})
    |> Rbac.Repo.one()
  end

  @spec fetch_users_without_oidc_connection(page :: non_neg_integer(), limit :: non_neg_integer()) ::
          {page :: non_neg_integer, [RbacUser.t()]}
  def fetch_users_without_oidc_connection(page \\ 1, per_page \\ 1000) do
    users =
      RbacUser
      |> join(:left, [u], o in assoc(u, :oidc_users))
      |> join(:inner, [u, _], s in assoc(u, :subject))
      |> where([_, o, _], is_nil(o.oidc_user_id))
      |> select([u, _, s], %{u | name: s.name})
      |> order_by([u], desc: u.inserted_at, desc: u.id)
      |> offset(^per_page * (^page - 1))
      |> limit(^per_page)
      |> Rbac.Repo.all()

    {page, users}
  end

  @spec create(Ecto.UUID.t(), String.t(), String.t()) :: :ok | :error
  def create(user_id, email, name) do
    subject_changeset = Subject.changeset(%Subject{}, %{id: user_id, name: name, type: "user"})
    user_changeset = RbacUser.changeset(%RbacUser{}, %{id: user_id, email: email})

    Multi.new()
    |> Multi.insert(:subject, subject_changeset)
    |> Multi.insert(:user, user_changeset)
    |> handle_transaction(user_id, "create_rbac_user")
  end

  @spec update(Ecto.UUID.t(), map()) :: :ok | :error
  def update(user_id, params \\ %{}) do
    subject = Subject |> where([s], s.id == ^user_id) |> Rbac.Repo.one!()
    rbac_user = RbacUser |> where([u], u.id == ^user_id) |> Rbac.Repo.one!()

    Multi.new()
    |> Multi.update(:subject, Subject.changeset(subject, params))
    |> Multi.update(:rbac_user, RbacUser.changeset(rbac_user, params))
    |> handle_transaction(user_id, "update_rbac_user")
  end

  @spec delete(Ecto.UUID.t()) :: :ok | :error
  def delete(user_id) do
    user_group_memberships =
      Rbac.Repo.UserGroupBinding |> where([binding], binding.user_id == ^user_id)

    Multi.new()
    |> Multi.run(:retract_al_roles, fn _repo, _changes ->
      {:ok, rbi} = Rbac.RoleBindingIdentification.new(user_id: user_id)
      Rbac.RoleManagement.retract_roles(rbi)
    end)
    |> Multi.delete_all(:delete_user_group_membershinps, user_group_memberships)
    |> Multi.delete_all(
      :delete_subject,
      Rbac.Repo.Subject |> where([subject], subject.id == ^user_id)
    )
    |> Multi.delete_all(
      :delete_user,
      Rbac.Repo.RbacUser |> where([user], user.id == ^user_id)
    )
    |> handle_transaction(user_id, "delete_rbac_user")
  end

  # Helper functions
  defp handle_transaction(transaction, user_id, action_name) do
    case Rbac.Repo.transaction(transaction, timeout: 60_000) do
      {:ok, _} ->
        Watchman.increment("#{action_name}.success")
        Logger.info(fn -> "Action #{action_name} successful for user #{user_id}" end)
        :ok

      error_msg ->
        Watchman.increment("#{action_name}.failure")

        Logger.error(fn ->
          "Action #{action_name} failed for user #{user_id}. Error #{inspect(error_msg)}"
        end)

        :error
    end
  end
end
