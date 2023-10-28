// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: keep_live.proto

package cow

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

type GeoCoordinate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *GeoCoordinate) Reset() {
	*x = GeoCoordinate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keep_live_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoCoordinate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoCoordinate) ProtoMessage() {}

func (x *GeoCoordinate) ProtoReflect() protoreflect.Message {
	mi := &file_keep_live_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoCoordinate.ProtoReflect.Descriptor instead.
func (*GeoCoordinate) Descriptor() ([]byte, []int) {
	return file_keep_live_proto_rawDescGZIP(), []int{0}
}

func (x *GeoCoordinate) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *GeoCoordinate) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type KeepAliveMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp string         `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Token     string         `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Geo       *GeoCoordinate `protobuf:"bytes,3,opt,name=geo,proto3" json:"geo,omitempty"`
	Weight    float64        `protobuf:"fixed64,4,opt,name=weight,proto3" json:"weight,omitempty"`
	Health    float64        `protobuf:"fixed64,5,opt,name=health,proto3" json:"health,omitempty"`
}

func (x *KeepAliveMsg) Reset() {
	*x = KeepAliveMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keep_live_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeepAliveMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeepAliveMsg) ProtoMessage() {}

func (x *KeepAliveMsg) ProtoReflect() protoreflect.Message {
	mi := &file_keep_live_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeepAliveMsg.ProtoReflect.Descriptor instead.
func (*KeepAliveMsg) Descriptor() ([]byte, []int) {
	return file_keep_live_proto_rawDescGZIP(), []int{1}
}

func (x *KeepAliveMsg) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *KeepAliveMsg) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *KeepAliveMsg) GetGeo() *GeoCoordinate {
	if x != nil {
		return x.Geo
	}
	return nil
}

func (x *KeepAliveMsg) GetWeight() float64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *KeepAliveMsg) GetHealth() float64 {
	if x != nil {
		return x.Health
	}
	return 0
}

var File_keep_live_proto protoreflect.FileDescriptor

var file_keep_live_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x66, 0x61, 0x72, 0x6d, 0x2e, 0x63, 0x6f, 0x77, 0x22, 0x49, 0x0a, 0x0d, 0x47,
	0x65, 0x6f, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x9d, 0x01, 0x0a, 0x0c, 0x4b, 0x65, 0x65, 0x70, 0x41,
	0x6c, 0x69, 0x76, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x29, 0x0a, 0x03, 0x67,
	0x65, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x66, 0x61, 0x72, 0x6d, 0x2e,
	0x63, 0x6f, 0x77, 0x2e, 0x47, 0x65, 0x6f, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x52, 0x03, 0x67, 0x65, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x42, 0x0a, 0x5a, 0x08, 0x66, 0x61, 0x72, 0x6d, 0x2f, 0x63,
	0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_keep_live_proto_rawDescOnce sync.Once
	file_keep_live_proto_rawDescData = file_keep_live_proto_rawDesc
)

func file_keep_live_proto_rawDescGZIP() []byte {
	file_keep_live_proto_rawDescOnce.Do(func() {
		file_keep_live_proto_rawDescData = protoimpl.X.CompressGZIP(file_keep_live_proto_rawDescData)
	})
	return file_keep_live_proto_rawDescData
}

var file_keep_live_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_keep_live_proto_goTypes = []interface{}{
	(*GeoCoordinate)(nil), // 0: farm.cow.GeoCoordinate
	(*KeepAliveMsg)(nil),  // 1: farm.cow.KeepAliveMsg
}
var file_keep_live_proto_depIdxs = []int32{
	0, // 0: farm.cow.KeepAliveMsg.geo:type_name -> farm.cow.GeoCoordinate
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_keep_live_proto_init() }
func file_keep_live_proto_init() {
	if File_keep_live_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_keep_live_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoCoordinate); i {
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
		file_keep_live_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeepAliveMsg); i {
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
			RawDescriptor: file_keep_live_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_keep_live_proto_goTypes,
		DependencyIndexes: file_keep_live_proto_depIdxs,
		MessageInfos:      file_keep_live_proto_msgTypes,
	}.Build()
	File_keep_live_proto = out.File
	file_keep_live_proto_rawDesc = nil
	file_keep_live_proto_goTypes = nil
	file_keep_live_proto_depIdxs = nil
}
