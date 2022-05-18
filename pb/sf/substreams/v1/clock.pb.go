// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: sf/substreams/v1/clock.proto

package pbsubstreams

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Clock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Number    uint64                 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Clock) Reset() {
	*x = Clock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sf_substreams_v1_clock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Clock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Clock) ProtoMessage() {}

func (x *Clock) ProtoReflect() protoreflect.Message {
	mi := &file_sf_substreams_v1_clock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Clock.ProtoReflect.Descriptor instead.
func (*Clock) Descriptor() ([]byte, []int) {
	return file_sf_substreams_v1_clock_proto_rawDescGZIP(), []int{0}
}

func (x *Clock) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Clock) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Clock) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_sf_substreams_v1_clock_proto protoreflect.FileDescriptor

var file_sf_substreams_v1_clock_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x66, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x73, 0x66, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x69, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x46, 0x5a, 0x44,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x66, 0x61, 0x73, 0x74, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x66, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sf_substreams_v1_clock_proto_rawDescOnce sync.Once
	file_sf_substreams_v1_clock_proto_rawDescData = file_sf_substreams_v1_clock_proto_rawDesc
)

func file_sf_substreams_v1_clock_proto_rawDescGZIP() []byte {
	file_sf_substreams_v1_clock_proto_rawDescOnce.Do(func() {
		file_sf_substreams_v1_clock_proto_rawDescData = protoimpl.X.CompressGZIP(file_sf_substreams_v1_clock_proto_rawDescData)
	})
	return file_sf_substreams_v1_clock_proto_rawDescData
}

var file_sf_substreams_v1_clock_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sf_substreams_v1_clock_proto_goTypes = []interface{}{
	(*Clock)(nil),                 // 0: sf.substreams.v1.Clock
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_sf_substreams_v1_clock_proto_depIdxs = []int32{
	1, // 0: sf.substreams.v1.Clock.timestamp:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sf_substreams_v1_clock_proto_init() }
func file_sf_substreams_v1_clock_proto_init() {
	if File_sf_substreams_v1_clock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sf_substreams_v1_clock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Clock); i {
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
			RawDescriptor: file_sf_substreams_v1_clock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sf_substreams_v1_clock_proto_goTypes,
		DependencyIndexes: file_sf_substreams_v1_clock_proto_depIdxs,
		MessageInfos:      file_sf_substreams_v1_clock_proto_msgTypes,
	}.Build()
	File_sf_substreams_v1_clock_proto = out.File
	file_sf_substreams_v1_clock_proto_rawDesc = nil
	file_sf_substreams_v1_clock_proto_goTypes = nil
	file_sf_substreams_v1_clock_proto_depIdxs = nil
}
