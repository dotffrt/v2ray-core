// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: features/dns/config.proto

package dns

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

type QueryStrategy int32

const (
	QueryStrategy_USE_IP     QueryStrategy = 0
	QueryStrategy_USE_IP4    QueryStrategy = 1
	QueryStrategy_USE_IP6    QueryStrategy = 2
	QueryStrategy_PREFER_IP4 QueryStrategy = 3
	QueryStrategy_PREFER_IP6 QueryStrategy = 4
)

// Enum value maps for QueryStrategy.
var (
	QueryStrategy_name = map[int32]string{
		0: "USE_IP",
		1: "USE_IP4",
		2: "USE_IP6",
		3: "PREFER_IP4",
		4: "PREFER_IP6",
	}
	QueryStrategy_value = map[string]int32{
		"USE_IP":     0,
		"USE_IP4":    1,
		"USE_IP6":    2,
		"PREFER_IP4": 3,
		"PREFER_IP6": 4,
	}
)

func (x QueryStrategy) Enum() *QueryStrategy {
	p := new(QueryStrategy)
	*p = x
	return p
}

func (x QueryStrategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QueryStrategy) Descriptor() protoreflect.EnumDescriptor {
	return file_features_dns_config_proto_enumTypes[0].Descriptor()
}

func (QueryStrategy) Type() protoreflect.EnumType {
	return &file_features_dns_config_proto_enumTypes[0]
}

func (x QueryStrategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QueryStrategy.Descriptor instead.
func (QueryStrategy) EnumDescriptor() ([]byte, []int) {
	return file_features_dns_config_proto_rawDescGZIP(), []int{0}
}

var File_features_dns_config_proto protoreflect.FileDescriptor

var file_features_dns_config_proto_rawDesc = []byte{
	0x0a, 0x19, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2f, 0x64, 0x6e, 0x73, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x76, 0x32, 0x72,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x2e, 0x64, 0x6e, 0x73, 0x2a, 0x55, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x53, 0x45, 0x5f, 0x49, 0x50, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x53, 0x45, 0x5f, 0x49, 0x50, 0x34, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x53, 0x45, 0x5f, 0x49, 0x50, 0x36, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x50,
	0x52, 0x45, 0x46, 0x45, 0x52, 0x5f, 0x49, 0x50, 0x34, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x50,
	0x52, 0x45, 0x46, 0x45, 0x52, 0x5f, 0x49, 0x50, 0x36, 0x10, 0x04, 0x42, 0x66, 0x0a, 0x1b, 0x63,
	0x6f, 0x6d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x64, 0x6e, 0x73, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x32, 0x66, 0x6c, 0x79, 0x2f, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x35, 0x2f, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x2f, 0x64, 0x6e, 0x73, 0xaa, 0x02, 0x17, 0x56, 0x32, 0x52, 0x61,
	0x79, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e,
	0x44, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_features_dns_config_proto_rawDescOnce sync.Once
	file_features_dns_config_proto_rawDescData = file_features_dns_config_proto_rawDesc
)

func file_features_dns_config_proto_rawDescGZIP() []byte {
	file_features_dns_config_proto_rawDescOnce.Do(func() {
		file_features_dns_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_features_dns_config_proto_rawDescData)
	})
	return file_features_dns_config_proto_rawDescData
}

var file_features_dns_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_features_dns_config_proto_goTypes = []interface{}{
	(QueryStrategy)(0), // 0: v2ray.core.features.dns.QueryStrategy
}
var file_features_dns_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_features_dns_config_proto_init() }
func file_features_dns_config_proto_init() {
	if File_features_dns_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_features_dns_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_features_dns_config_proto_goTypes,
		DependencyIndexes: file_features_dns_config_proto_depIdxs,
		EnumInfos:         file_features_dns_config_proto_enumTypes,
	}.Build()
	File_features_dns_config_proto = out.File
	file_features_dns_config_proto_rawDesc = nil
	file_features_dns_config_proto_goTypes = nil
	file_features_dns_config_proto_depIdxs = nil
}
