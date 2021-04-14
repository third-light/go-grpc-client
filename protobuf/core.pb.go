// This is a protobuf specifications for the Core module of the Third Light gRPC API for Chorus.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: core.proto

package protobuf

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

type CoreRequests struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CoreRequests) Reset() {
	*x = CoreRequests{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoreRequests) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoreRequests) ProtoMessage() {}

func (x *CoreRequests) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoreRequests.ProtoReflect.Descriptor instead.
func (*CoreRequests) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{0}
}

// This is a representation of a date time.
type DateTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The date encoded as per RFC-3339 (a.k.a ATOM)
	Rfc3339 string `protobuf:"bytes,2,opt,name=rfc3339,proto3" json:"rfc3339,omitempty"`
	// this is the date/time represented as the number of seconds since the Unix Epoch (January 1 1970, 00:00:00 UTC)
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *DateTime) Reset() {
	*x = DateTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DateTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateTime) ProtoMessage() {}

func (x *DateTime) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateTime.ProtoReflect.Descriptor instead.
func (*DateTime) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{1}
}

func (x *DateTime) GetRfc3339() string {
	if x != nil {
		return x.Rfc3339
	}
	return ""
}

func (x *DateTime) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// Details of the Chorus server.
type EnvironmentDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The version of the title to use as the browser title (this often is longer and more verbose). If no specific
	// browser title has been configured, then this will be the same as the 'title'.
	BrowserTitle string `protobuf:"bytes,2,opt,name=browserTitle,proto3" json:"browserTitle,omitempty"`
	// The site's title. This is used for the shortcuts and other places referring to the site.
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// The current version of Chorus.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *EnvironmentDetails) Reset() {
	*x = EnvironmentDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnvironmentDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnvironmentDetails) ProtoMessage() {}

func (x *EnvironmentDetails) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnvironmentDetails.ProtoReflect.Descriptor instead.
func (*EnvironmentDetails) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{2}
}

func (x *EnvironmentDetails) GetBrowserTitle() string {
	if x != nil {
		return x.BrowserTitle
	}
	return ""
}

func (x *EnvironmentDetails) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EnvironmentDetails) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type CoreRequests_GetEnvironment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CoreRequests_GetEnvironment) Reset() {
	*x = CoreRequests_GetEnvironment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoreRequests_GetEnvironment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoreRequests_GetEnvironment) ProtoMessage() {}

func (x *CoreRequests_GetEnvironment) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoreRequests_GetEnvironment.ProtoReflect.Descriptor instead.
func (*CoreRequests_GetEnvironment) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{0, 0}
}

var File_core_proto protoreflect.FileDescriptor

var file_core_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x68,
	0x6f, 0x72, 0x75, 0x73, 0x22, 0x20, 0x0a, 0x0c, 0x43, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x1a, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x42, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x66, 0x63, 0x33, 0x33, 0x33, 0x39, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x66, 0x63, 0x33, 0x33, 0x33, 0x39, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x68, 0x0a, 0x12, 0x45, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x22, 0x0a, 0x0c, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x32, 0x5b, 0x0a, 0x04, 0x43, 0x6f, 0x72, 0x65, 0x12, 0x53, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23,
	0x2e, 0x63, 0x68, 0x6f, 0x72, 0x75, 0x73, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x75, 0x73, 0x2e, 0x45, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22,
	0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_proto_rawDescOnce sync.Once
	file_core_proto_rawDescData = file_core_proto_rawDesc
)

func file_core_proto_rawDescGZIP() []byte {
	file_core_proto_rawDescOnce.Do(func() {
		file_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_proto_rawDescData)
	})
	return file_core_proto_rawDescData
}

var file_core_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_core_proto_goTypes = []interface{}{
	(*CoreRequests)(nil),                // 0: chorus.CoreRequests
	(*DateTime)(nil),                    // 1: chorus.DateTime
	(*EnvironmentDetails)(nil),          // 2: chorus.EnvironmentDetails
	(*CoreRequests_GetEnvironment)(nil), // 3: chorus.CoreRequests.GetEnvironment
}
var file_core_proto_depIdxs = []int32{
	3, // 0: chorus.Core.GetEnvironment:input_type -> chorus.CoreRequests.GetEnvironment
	2, // 1: chorus.Core.GetEnvironment:output_type -> chorus.EnvironmentDetails
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_core_proto_init() }
func file_core_proto_init() {
	if File_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoreRequests); i {
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
		file_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DateTime); i {
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
		file_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnvironmentDetails); i {
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
		file_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoreRequests_GetEnvironment); i {
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
			RawDescriptor: file_core_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_core_proto_goTypes,
		DependencyIndexes: file_core_proto_depIdxs,
		MessageInfos:      file_core_proto_msgTypes,
	}.Build()
	File_core_proto = out.File
	file_core_proto_rawDesc = nil
	file_core_proto_goTypes = nil
	file_core_proto_depIdxs = nil
}