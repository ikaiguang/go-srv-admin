// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/admin/v1/common/admin.common.v1.proto

package admincommonv1

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

// AdminCommon AdminCommon enum
type AdminCommonEnum_AdminCommon int32

const (
	// UNSPECIFIED 未指定
	AdminCommonEnum_UNSPECIFIED AdminCommonEnum_AdminCommon = 0
)

// Enum value maps for AdminCommonEnum_AdminCommon.
var (
	AdminCommonEnum_AdminCommon_name = map[int32]string{
		0: "UNSPECIFIED",
	}
	AdminCommonEnum_AdminCommon_value = map[string]int32{
		"UNSPECIFIED": 0,
	}
)

func (x AdminCommonEnum_AdminCommon) Enum() *AdminCommonEnum_AdminCommon {
	p := new(AdminCommonEnum_AdminCommon)
	*p = x
	return p
}

func (x AdminCommonEnum_AdminCommon) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdminCommonEnum_AdminCommon) Descriptor() protoreflect.EnumDescriptor {
	return file_api_admin_v1_common_admin_common_v1_proto_enumTypes[0].Descriptor()
}

func (AdminCommonEnum_AdminCommon) Type() protoreflect.EnumType {
	return &file_api_admin_v1_common_admin_common_v1_proto_enumTypes[0]
}

func (x AdminCommonEnum_AdminCommon) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdminCommonEnum_AdminCommon.Descriptor instead.
func (AdminCommonEnum_AdminCommon) EnumDescriptor() ([]byte, []int) {
	return file_api_admin_v1_common_admin_common_v1_proto_rawDescGZIP(), []int{0, 0}
}

// AdminCommonEnum AdminCommonEnum enum
type AdminCommonEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AdminCommonEnum) Reset() {
	*x = AdminCommonEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_admin_v1_common_admin_common_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminCommonEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminCommonEnum) ProtoMessage() {}

func (x *AdminCommonEnum) ProtoReflect() protoreflect.Message {
	mi := &file_api_admin_v1_common_admin_common_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminCommonEnum.ProtoReflect.Descriptor instead.
func (*AdminCommonEnum) Descriptor() ([]byte, []int) {
	return file_api_admin_v1_common_admin_common_v1_proto_rawDescGZIP(), []int{0}
}

var File_api_admin_v1_common_admin_common_v1_proto protoreflect.FileDescriptor

var file_api_admin_v1_common_admin_common_v1_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0x22, 0x31, 0x0a, 0x0f, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x22, 0x1e, 0x0a,
	0x0b, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x0f, 0x0a, 0x0b,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x42, 0x82, 0x01,
	0x0a, 0x1d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0x42,
	0x1a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x70, 0x69, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x56, 0x31, 0x50, 0x01, 0x5a, 0x43, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6b, 0x61, 0x69, 0x67, 0x75,
	0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x72, 0x76, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_admin_v1_common_admin_common_v1_proto_rawDescOnce sync.Once
	file_api_admin_v1_common_admin_common_v1_proto_rawDescData = file_api_admin_v1_common_admin_common_v1_proto_rawDesc
)

func file_api_admin_v1_common_admin_common_v1_proto_rawDescGZIP() []byte {
	file_api_admin_v1_common_admin_common_v1_proto_rawDescOnce.Do(func() {
		file_api_admin_v1_common_admin_common_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_admin_v1_common_admin_common_v1_proto_rawDescData)
	})
	return file_api_admin_v1_common_admin_common_v1_proto_rawDescData
}

var file_api_admin_v1_common_admin_common_v1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_admin_v1_common_admin_common_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_admin_v1_common_admin_common_v1_proto_goTypes = []interface{}{
	(AdminCommonEnum_AdminCommon)(0), // 0: admin.api.admin.admincommonv1.AdminCommonEnum.AdminCommon
	(*AdminCommonEnum)(nil),          // 1: admin.api.admin.admincommonv1.AdminCommonEnum
}
var file_api_admin_v1_common_admin_common_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_admin_v1_common_admin_common_v1_proto_init() }
func file_api_admin_v1_common_admin_common_v1_proto_init() {
	if File_api_admin_v1_common_admin_common_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_admin_v1_common_admin_common_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminCommonEnum); i {
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
			RawDescriptor: file_api_admin_v1_common_admin_common_v1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_admin_v1_common_admin_common_v1_proto_goTypes,
		DependencyIndexes: file_api_admin_v1_common_admin_common_v1_proto_depIdxs,
		EnumInfos:         file_api_admin_v1_common_admin_common_v1_proto_enumTypes,
		MessageInfos:      file_api_admin_v1_common_admin_common_v1_proto_msgTypes,
	}.Build()
	File_api_admin_v1_common_admin_common_v1_proto = out.File
	file_api_admin_v1_common_admin_common_v1_proto_rawDesc = nil
	file_api_admin_v1_common_admin_common_v1_proto_goTypes = nil
	file_api_admin_v1_common_admin_common_v1_proto_depIdxs = nil
}
