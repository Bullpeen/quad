// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: quadlek/plugins/comics/v1/comics.proto

package v1

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

type Templates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls []string `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
}

func (x *Templates) Reset() {
	*x = Templates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quadlek_plugins_comics_v1_comics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Templates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Templates) ProtoMessage() {}

func (x *Templates) ProtoReflect() protoreflect.Message {
	mi := &file_quadlek_plugins_comics_v1_comics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Templates.ProtoReflect.Descriptor instead.
func (*Templates) Descriptor() ([]byte, []int) {
	return file_quadlek_plugins_comics_v1_comics_proto_rawDescGZIP(), []int{0}
}

func (x *Templates) GetUrls() []string {
	if x != nil {
		return x.Urls
	}
	return nil
}

var File_quadlek_plugins_comics_v1_comics_proto protoreflect.FileDescriptor

var file_quadlek_plugins_comics_v1_comics_proto_rawDesc = []byte{
	0x0a, 0x26, 0x71, 0x75, 0x61, 0x64, 0x6c, 0x65, 0x6b, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x69,
	0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x71, 0x75, 0x61, 0x64, 0x6c, 0x65,
	0x6b, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x69, 0x63, 0x73,
	0x2e, 0x76, 0x31, 0x22, 0x1f, 0x0a, 0x09, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x72, 0x6c, 0x73, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6a, 0x69, 0x72, 0x77, 0x69, 0x6e, 0x2f, 0x71, 0x75, 0x61, 0x64, 0x6c, 0x65,
	0x6b, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_quadlek_plugins_comics_v1_comics_proto_rawDescOnce sync.Once
	file_quadlek_plugins_comics_v1_comics_proto_rawDescData = file_quadlek_plugins_comics_v1_comics_proto_rawDesc
)

func file_quadlek_plugins_comics_v1_comics_proto_rawDescGZIP() []byte {
	file_quadlek_plugins_comics_v1_comics_proto_rawDescOnce.Do(func() {
		file_quadlek_plugins_comics_v1_comics_proto_rawDescData = protoimpl.X.CompressGZIP(file_quadlek_plugins_comics_v1_comics_proto_rawDescData)
	})
	return file_quadlek_plugins_comics_v1_comics_proto_rawDescData
}

var file_quadlek_plugins_comics_v1_comics_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_quadlek_plugins_comics_v1_comics_proto_goTypes = []interface{}{
	(*Templates)(nil), // 0: quadlek.plugins.comics.v1.Templates
}
var file_quadlek_plugins_comics_v1_comics_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_quadlek_plugins_comics_v1_comics_proto_init() }
func file_quadlek_plugins_comics_v1_comics_proto_init() {
	if File_quadlek_plugins_comics_v1_comics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_quadlek_plugins_comics_v1_comics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Templates); i {
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
			RawDescriptor: file_quadlek_plugins_comics_v1_comics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_quadlek_plugins_comics_v1_comics_proto_goTypes,
		DependencyIndexes: file_quadlek_plugins_comics_v1_comics_proto_depIdxs,
		MessageInfos:      file_quadlek_plugins_comics_v1_comics_proto_msgTypes,
	}.Build()
	File_quadlek_plugins_comics_v1_comics_proto = out.File
	file_quadlek_plugins_comics_v1_comics_proto_rawDesc = nil
	file_quadlek_plugins_comics_v1_comics_proto_goTypes = nil
	file_quadlek_plugins_comics_v1_comics_proto_depIdxs = nil
}
