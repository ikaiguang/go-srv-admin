// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: api/admin/v1/errors/admin.error.v1.proto

package adminerrorv1

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

// ERROR admin error
type ERROR int32

const (
	// ADMIN_UNKNOWN 未知错误
	ERROR_ADMIN_UNKNOWN ERROR = 0
	// ADMIN_NOT_EXIST 用户不存在
	ERROR_ADMIN_NOT_EXIST ERROR = 10001
	// ADMIN_EXIST 用户已存在
	ERROR_ADMIN_EXIST ERROR = 10002
	// ADMIN_NAME_INVALID 用户名不合法
	ERROR_ADMIN_NAME_INVALID ERROR = 10003
	// ADMIN_NAME_EXIST 用户名已存在
	ERROR_ADMIN_NAME_EXIST ERROR = 10004
	// ADMIN_PASSWORD_INVALID 用户密码不合法
	ERROR_ADMIN_PASSWORD_INVALID ERROR = 10005
	// ADMIN_PASSWORD_INCORRECT 用户密码不正确
	ERROR_ADMIN_PASSWORD_INCORRECT ERROR = 10006
	// ADMIN_TOKEN_INVALID 令牌已失效
	ERROR_ADMIN_TOKEN_INVALID ERROR = 10007
	// ADMIN_ACCOUNT_EXPIRE 账户已过期
	ERROR_ADMIN_ACCOUNT_EXPIRE ERROR = 10008
)

// Enum value maps for ERROR.
var (
	ERROR_name = map[int32]string{
		0:     "ADMIN_UNKNOWN",
		10001: "ADMIN_NOT_EXIST",
		10002: "ADMIN_EXIST",
		10003: "ADMIN_NAME_INVALID",
		10004: "ADMIN_NAME_EXIST",
		10005: "ADMIN_PASSWORD_INVALID",
		10006: "ADMIN_PASSWORD_INCORRECT",
		10007: "ADMIN_TOKEN_INVALID",
		10008: "ADMIN_ACCOUNT_EXPIRE",
	}
	ERROR_value = map[string]int32{
		"ADMIN_UNKNOWN":            0,
		"ADMIN_NOT_EXIST":          10001,
		"ADMIN_EXIST":              10002,
		"ADMIN_NAME_INVALID":       10003,
		"ADMIN_NAME_EXIST":         10004,
		"ADMIN_PASSWORD_INVALID":   10005,
		"ADMIN_PASSWORD_INCORRECT": 10006,
		"ADMIN_TOKEN_INVALID":      10007,
		"ADMIN_ACCOUNT_EXPIRE":     10008,
	}
)

func (x ERROR) Enum() *ERROR {
	p := new(ERROR)
	*p = x
	return p
}

func (x ERROR) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ERROR) Descriptor() protoreflect.EnumDescriptor {
	return file_api_admin_v1_errors_admin_error_v1_proto_enumTypes[0].Descriptor()
}

func (ERROR) Type() protoreflect.EnumType {
	return &file_api_admin_v1_errors_admin_error_v1_proto_enumTypes[0]
}

func (x ERROR) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ERROR.Descriptor instead.
func (ERROR) EnumDescriptor() ([]byte, []int) {
	return file_api_admin_v1_errors_admin_error_v1_proto_rawDescGZIP(), []int{0}
}

var File_api_admin_v1_errors_admin_error_v1_proto protoreflect.FileDescriptor

var file_api_admin_v1_errors_admin_error_v1_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x76, 0x31, 0x2a, 0xe3, 0x01, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x12, 0x11, 0x0a, 0x0d,
	0x41, 0x44, 0x4d, 0x49, 0x4e, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x14, 0x0a, 0x0f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49,
	0x53, 0x54, 0x10, 0x91, 0x4e, 0x12, 0x10, 0x0a, 0x0b, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x5f, 0x45,
	0x58, 0x49, 0x53, 0x54, 0x10, 0x92, 0x4e, 0x12, 0x17, 0x0a, 0x12, 0x41, 0x44, 0x4d, 0x49, 0x4e,
	0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x93, 0x4e,
	0x12, 0x15, 0x0a, 0x10, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x45,
	0x58, 0x49, 0x53, 0x54, 0x10, 0x94, 0x4e, 0x12, 0x1b, 0x0a, 0x16, 0x41, 0x44, 0x4d, 0x49, 0x4e,
	0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x10, 0x95, 0x4e, 0x12, 0x1d, 0x0a, 0x18, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x5f, 0x50, 0x41,
	0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x49, 0x4e, 0x43, 0x4f, 0x52, 0x52, 0x45, 0x43, 0x54,
	0x10, 0x96, 0x4e, 0x12, 0x18, 0x0a, 0x13, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x5f, 0x54, 0x4f, 0x4b,
	0x45, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x97, 0x4e, 0x12, 0x19, 0x0a,
	0x14, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x45,
	0x58, 0x50, 0x49, 0x52, 0x45, 0x10, 0x98, 0x4e, 0x42, 0x74, 0x0a, 0x16, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x76, 0x31, 0x42, 0x14, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x70, 0x69, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6b, 0x61, 0x69, 0x67, 0x75, 0x61, 0x6e, 0x67,
	0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x72, 0x76, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_admin_v1_errors_admin_error_v1_proto_rawDescOnce sync.Once
	file_api_admin_v1_errors_admin_error_v1_proto_rawDescData = file_api_admin_v1_errors_admin_error_v1_proto_rawDesc
)

func file_api_admin_v1_errors_admin_error_v1_proto_rawDescGZIP() []byte {
	file_api_admin_v1_errors_admin_error_v1_proto_rawDescOnce.Do(func() {
		file_api_admin_v1_errors_admin_error_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_admin_v1_errors_admin_error_v1_proto_rawDescData)
	})
	return file_api_admin_v1_errors_admin_error_v1_proto_rawDescData
}

var file_api_admin_v1_errors_admin_error_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_admin_v1_errors_admin_error_v1_proto_goTypes = []interface{}{
	(ERROR)(0), // 0: admin.api.adminerrorv1.ERROR
}
var file_api_admin_v1_errors_admin_error_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_admin_v1_errors_admin_error_v1_proto_init() }
func file_api_admin_v1_errors_admin_error_v1_proto_init() {
	if File_api_admin_v1_errors_admin_error_v1_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_admin_v1_errors_admin_error_v1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_admin_v1_errors_admin_error_v1_proto_goTypes,
		DependencyIndexes: file_api_admin_v1_errors_admin_error_v1_proto_depIdxs,
		EnumInfos:         file_api_admin_v1_errors_admin_error_v1_proto_enumTypes,
	}.Build()
	File_api_admin_v1_errors_admin_error_v1_proto = out.File
	file_api_admin_v1_errors_admin_error_v1_proto_rawDesc = nil
	file_api_admin_v1_errors_admin_error_v1_proto_goTypes = nil
	file_api_admin_v1_errors_admin_error_v1_proto_depIdxs = nil
}
