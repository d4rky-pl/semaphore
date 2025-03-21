// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.3.0
// source: loghub2.proto

package loghub2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TokenType int32

const (
	TokenType_PULL TokenType = 0
	TokenType_PUSH TokenType = 1
)

// Enum value maps for TokenType.
var (
	TokenType_name = map[int32]string{
		0: "PULL",
		1: "PUSH",
	}
	TokenType_value = map[string]int32{
		"PULL": 0,
		"PUSH": 1,
	}
)

func (x TokenType) Enum() *TokenType {
	p := new(TokenType)
	*p = x
	return p
}

func (x TokenType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenType) Descriptor() protoreflect.EnumDescriptor {
	return file_loghub2_proto_enumTypes[0].Descriptor()
}

func (TokenType) Type() protoreflect.EnumType {
	return &file_loghub2_proto_enumTypes[0]
}

func (x TokenType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenType.Descriptor instead.
func (TokenType) EnumDescriptor() ([]byte, []int) {
	return file_loghub2_proto_rawDescGZIP(), []int{0}
}

// Request for GenerateToken
// - job_id = [required] UUID of the job.
// - type   = the type of token, either to PULL logs or to PUSH logs; default is PULL.
type GenerateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId    string    `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	Type     TokenType `protobuf:"varint,2,opt,name=type,proto3,enum=InternalApi.Loghub2.TokenType" json:"type,omitempty"`
	Duration uint32    `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *GenerateTokenRequest) Reset() {
	*x = GenerateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loghub2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenRequest) ProtoMessage() {}

func (x *GenerateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loghub2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenRequest.ProtoReflect.Descriptor instead.
func (*GenerateTokenRequest) Descriptor() ([]byte, []int) {
	return file_loghub2_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateTokenRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *GenerateTokenRequest) GetType() TokenType {
	if x != nil {
		return x.Type
	}
	return TokenType_PULL
}

func (x *GenerateTokenRequest) GetDuration() uint32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

// Response for GenerateToken
// - token  = [required] JWT token generated for job_id
// - type   = [required] type of token generated, either PULL or PUSH
type GenerateTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string    `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Type  TokenType `protobuf:"varint,2,opt,name=type,proto3,enum=InternalApi.Loghub2.TokenType" json:"type,omitempty"`
}

func (x *GenerateTokenResponse) Reset() {
	*x = GenerateTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loghub2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenResponse) ProtoMessage() {}

func (x *GenerateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loghub2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenResponse.ProtoReflect.Descriptor instead.
func (*GenerateTokenResponse) Descriptor() ([]byte, []int) {
	return file_loghub2_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GenerateTokenResponse) GetType() TokenType {
	if x != nil {
		return x.Type
	}
	return TokenType_PULL
}

var File_loghub2_proto protoreflect.FileDescriptor

var file_loghub2_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6c, 0x6f, 0x67, 0x68, 0x75, 0x62, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67,
	0x68, 0x75, 0x62, 0x32, 0x22, 0x7d, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06,
	0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f,
	0x62, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1e, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x70, 0x69, 0x2e,
	0x4c, 0x6f, 0x67, 0x68, 0x75, 0x62, 0x32, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x61, 0x0a, 0x15, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1e, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x70, 0x69, 0x2e, 0x4c,
	0x6f, 0x67, 0x68, 0x75, 0x62, 0x32, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x1f, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x55, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x50, 0x55, 0x53, 0x48, 0x10, 0x01, 0x32, 0x71, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x68, 0x75,
	0x62, 0x32, 0x12, 0x66, 0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x29, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x70,
	0x69, 0x2e, 0x4c, 0x6f, 0x67, 0x68, 0x75, 0x62, 0x32, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67,
	0x68, 0x75, 0x62, 0x32, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x65,
	0x64, 0x74, 0x65, 0x78, 0x74, 0x2f, 0x61, 0x6c, 0x6c, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x6c, 0x66,
	0x5f, 0x68, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x68, 0x75, 0x62, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x68, 0x75, 0x62, 0x32, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_loghub2_proto_rawDescOnce sync.Once
	file_loghub2_proto_rawDescData = file_loghub2_proto_rawDesc
)

func file_loghub2_proto_rawDescGZIP() []byte {
	file_loghub2_proto_rawDescOnce.Do(func() {
		file_loghub2_proto_rawDescData = protoimpl.X.CompressGZIP(file_loghub2_proto_rawDescData)
	})
	return file_loghub2_proto_rawDescData
}

var file_loghub2_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_loghub2_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_loghub2_proto_goTypes = []any{
	(TokenType)(0),                // 0: InternalApi.Loghub2.TokenType
	(*GenerateTokenRequest)(nil),  // 1: InternalApi.Loghub2.GenerateTokenRequest
	(*GenerateTokenResponse)(nil), // 2: InternalApi.Loghub2.GenerateTokenResponse
}
var file_loghub2_proto_depIdxs = []int32{
	0, // 0: InternalApi.Loghub2.GenerateTokenRequest.type:type_name -> InternalApi.Loghub2.TokenType
	0, // 1: InternalApi.Loghub2.GenerateTokenResponse.type:type_name -> InternalApi.Loghub2.TokenType
	1, // 2: InternalApi.Loghub2.Loghub2.GenerateToken:input_type -> InternalApi.Loghub2.GenerateTokenRequest
	2, // 3: InternalApi.Loghub2.Loghub2.GenerateToken:output_type -> InternalApi.Loghub2.GenerateTokenResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_loghub2_proto_init() }
func file_loghub2_proto_init() {
	if File_loghub2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loghub2_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GenerateTokenRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_loghub2_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GenerateTokenResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_loghub2_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_loghub2_proto_goTypes,
		DependencyIndexes: file_loghub2_proto_depIdxs,
		EnumInfos:         file_loghub2_proto_enumTypes,
		MessageInfos:      file_loghub2_proto_msgTypes,
	}.Build()
	File_loghub2_proto = out.File
	file_loghub2_proto_rawDesc = nil
	file_loghub2_proto_goTypes = nil
	file_loghub2_proto_depIdxs = nil
}
