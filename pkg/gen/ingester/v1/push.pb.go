// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: ingester/v1/push.proto

package ingestv1

import (
	v1 "github.com/grafana/fire/pkg/gen/push/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_ingester_v1_push_proto protoreflect.FileDescriptor

var file_ingester_v1_push_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75,
	0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x1a, 0x12, 0x70, 0x75, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x73,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x41, 0x0a, 0x08, 0x49, 0x6e, 0x67, 0x65, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x14, 0x2e, 0x70, 0x75,
	0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x95, 0x01, 0x0a, 0x0d, 0x63,
	0x6f, 0x6d, 0x2e, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x50, 0x75,
	0x73, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x61, 0x6e, 0x61, 0x2f, 0x66, 0x69,
	0x72, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x69, 0x6e, 0x67, 0x65, 0x73,
	0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x49, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x09, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15,
	0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_ingester_v1_push_proto_goTypes = []interface{}{
	(*v1.PushRequest)(nil),  // 0: push.v1.PushRequest
	(*v1.PushResponse)(nil), // 1: push.v1.PushResponse
}
var file_ingester_v1_push_proto_depIdxs = []int32{
	0, // 0: ingest.v1.Ingester.Push:input_type -> push.v1.PushRequest
	1, // 1: ingest.v1.Ingester.Push:output_type -> push.v1.PushResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ingester_v1_push_proto_init() }
func file_ingester_v1_push_proto_init() {
	if File_ingester_v1_push_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ingester_v1_push_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ingester_v1_push_proto_goTypes,
		DependencyIndexes: file_ingester_v1_push_proto_depIdxs,
	}.Build()
	File_ingester_v1_push_proto = out.File
	file_ingester_v1_push_proto_rawDesc = nil
	file_ingester_v1_push_proto_goTypes = nil
	file_ingester_v1_push_proto_depIdxs = nil
}
