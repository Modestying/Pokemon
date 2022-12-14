// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: feature.proto

package proto

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

type Sex int32

const (
	Sex_Male            Sex = 0 //boy
	Sex_Femal           Sex = 1 //girl
	Sex_Hermaphroditism Sex = 2 //boy and girl
)

// Enum value maps for Sex.
var (
	Sex_name = map[int32]string{
		0: "Male",
		1: "Femal",
		2: "Hermaphroditism",
	}
	Sex_value = map[string]int32{
		"Male":            0,
		"Femal":           1,
		"Hermaphroditism": 2,
	}
)

func (x Sex) Enum() *Sex {
	p := new(Sex)
	*p = x
	return p
}

func (x Sex) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Sex) Descriptor() protoreflect.EnumDescriptor {
	return file_feature_proto_enumTypes[0].Descriptor()
}

func (Sex) Type() protoreflect.EnumType {
	return &file_feature_proto_enumTypes[0]
}

func (x Sex) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sex.Descriptor instead.
func (Sex) EnumDescriptor() ([]byte, []int) {
	return file_feature_proto_rawDescGZIP(), []int{0}
}

var File_feature_proto protoreflect.FileDescriptor

var file_feature_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2a, 0x2f, 0x0a, 0x03, 0x53, 0x65, 0x78, 0x12,
	0x08, 0x0a, 0x04, 0x4d, 0x61, 0x6c, 0x65, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x65, 0x6d,
	0x61, 0x6c, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x48, 0x65, 0x72, 0x6d, 0x61, 0x70, 0x68, 0x72,
	0x6f, 0x64, 0x69, 0x74, 0x69, 0x73, 0x6d, 0x10, 0x02, 0x42, 0x10, 0x5a, 0x0e, 0x50, 0x6f, 0x6b,
	0x65, 0x6e, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_feature_proto_rawDescOnce sync.Once
	file_feature_proto_rawDescData = file_feature_proto_rawDesc
)

func file_feature_proto_rawDescGZIP() []byte {
	file_feature_proto_rawDescOnce.Do(func() {
		file_feature_proto_rawDescData = protoimpl.X.CompressGZIP(file_feature_proto_rawDescData)
	})
	return file_feature_proto_rawDescData
}

var file_feature_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_feature_proto_goTypes = []interface{}{
	(Sex)(0), // 0: feature.Sex
}
var file_feature_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_feature_proto_init() }
func file_feature_proto_init() {
	if File_feature_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_feature_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_feature_proto_goTypes,
		DependencyIndexes: file_feature_proto_depIdxs,
		EnumInfos:         file_feature_proto_enumTypes,
	}.Build()
	File_feature_proto = out.File
	file_feature_proto_rawDesc = nil
	file_feature_proto_goTypes = nil
	file_feature_proto_depIdxs = nil
}
