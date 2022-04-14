// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: examplesrv/v1/events.proto

package examplesrvv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExampleEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Example string                 `protobuf:"bytes,1,opt,name=example,proto3" json:"example,omitempty"`
	EventAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=event_at,json=eventAt,proto3" json:"event_at,omitempty"`
}

func (x *ExampleEvent) Reset() {
	*x = ExampleEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_examplesrv_v1_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExampleEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExampleEvent) ProtoMessage() {}

func (x *ExampleEvent) ProtoReflect() protoreflect.Message {
	mi := &file_examplesrv_v1_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExampleEvent.ProtoReflect.Descriptor instead.
func (*ExampleEvent) Descriptor() ([]byte, []int) {
	return file_examplesrv_v1_events_proto_rawDescGZIP(), []int{0}
}

func (x *ExampleEvent) GetExample() string {
	if x != nil {
		return x.Example
	}
	return ""
}

func (x *ExampleEvent) GetEventAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EventAt
	}
	return nil
}

var File_examplesrv_v1_events_proto protoreflect.FileDescriptor

var file_examplesrv_v1_events_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x72, 0x76, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x72,
	0x61, 0x6e, 0x77, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x72, 0x76, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x0c, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x07, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12,
	0x35, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x72, 0x61, 0x6e, 0x77, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x72, 0x76, 0x2f, 0x76, 0x31, 0x3b, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x72, 0x76, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_examplesrv_v1_events_proto_rawDescOnce sync.Once
	file_examplesrv_v1_events_proto_rawDescData = file_examplesrv_v1_events_proto_rawDesc
)

func file_examplesrv_v1_events_proto_rawDescGZIP() []byte {
	file_examplesrv_v1_events_proto_rawDescOnce.Do(func() {
		file_examplesrv_v1_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_examplesrv_v1_events_proto_rawDescData)
	})
	return file_examplesrv_v1_events_proto_rawDescData
}

var file_examplesrv_v1_events_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_examplesrv_v1_events_proto_goTypes = []interface{}{
	(*ExampleEvent)(nil),          // 0: aranw.examplesrv.v1.ExampleEvent
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_examplesrv_v1_events_proto_depIdxs = []int32{
	1, // 0: aranw.examplesrv.v1.ExampleEvent.event_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_examplesrv_v1_events_proto_init() }
func file_examplesrv_v1_events_proto_init() {
	if File_examplesrv_v1_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_examplesrv_v1_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExampleEvent); i {
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
			RawDescriptor: file_examplesrv_v1_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_examplesrv_v1_events_proto_goTypes,
		DependencyIndexes: file_examplesrv_v1_events_proto_depIdxs,
		MessageInfos:      file_examplesrv_v1_events_proto_msgTypes,
	}.Build()
	File_examplesrv_v1_events_proto = out.File
	file_examplesrv_v1_events_proto_rawDesc = nil
	file_examplesrv_v1_events_proto_goTypes = nil
	file_examplesrv_v1_events_proto_depIdxs = nil
}
