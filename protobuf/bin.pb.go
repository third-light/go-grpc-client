// This is a protobuf specifications for the Bin module of the Third Light gRPC API for Chorus.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: bin.proto

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

type BinRequests struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BinRequests) Reset() {
	*x = BinRequests{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinRequests) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinRequests) ProtoMessage() {}

func (x *BinRequests) ProtoReflect() protoreflect.Message {
	mi := &file_bin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinRequests.ProtoReflect.Descriptor instead.
func (*BinRequests) Descriptor() ([]byte, []int) {
	return file_bin_proto_rawDescGZIP(), []int{0}
}

type BinResponses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BinResponses) Reset() {
	*x = BinResponses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinResponses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinResponses) ProtoMessage() {}

func (x *BinResponses) ProtoReflect() protoreflect.Message {
	mi := &file_bin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinResponses.ProtoReflect.Descriptor instead.
func (*BinResponses) Descriptor() ([]byte, []int) {
	return file_bin_proto_rawDescGZIP(), []int{1}
}

type BinRequests_Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BinRequests_Empty) Reset() {
	*x = BinRequests_Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinRequests_Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinRequests_Empty) ProtoMessage() {}

func (x *BinRequests_Empty) ProtoReflect() protoreflect.Message {
	mi := &file_bin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinRequests_Empty.ProtoReflect.Descriptor instead.
func (*BinRequests_Empty) Descriptor() ([]byte, []int) {
	return file_bin_proto_rawDescGZIP(), []int{0, 0}
}

type BinRequests_Get struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BinRequests_Get) Reset() {
	*x = BinRequests_Get{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinRequests_Get) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinRequests_Get) ProtoMessage() {}

func (x *BinRequests_Get) ProtoReflect() protoreflect.Message {
	mi := &file_bin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinRequests_Get.ProtoReflect.Descriptor instead.
func (*BinRequests_Get) Descriptor() ([]byte, []int) {
	return file_bin_proto_rawDescGZIP(), []int{0, 1}
}

type BinResponses_Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BinResponses_Empty) Reset() {
	*x = BinResponses_Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinResponses_Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinResponses_Empty) ProtoMessage() {}

func (x *BinResponses_Empty) ProtoReflect() protoreflect.Message {
	mi := &file_bin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinResponses_Empty.ProtoReflect.Descriptor instead.
func (*BinResponses_Empty) Descriptor() ([]byte, []int) {
	return file_bin_proto_rawDescGZIP(), []int{1, 0}
}

var File_bin_proto protoreflect.FileDescriptor

var file_bin_proto_rawDesc = []byte{
	0x0a, 0x09, 0x62, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x68, 0x6f,
	0x72, 0x75, 0x73, 0x1a, 0x0d, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x0b, 0x42, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x1a, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x05, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x22, 0x17, 0x0a, 0x0c, 0x42, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x73, 0x1a, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x80, 0x01, 0x0a, 0x03, 0x42,
	0x69, 0x6e, 0x12, 0x40, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x19, 0x2e, 0x63, 0x68,
	0x6f, 0x72, 0x75, 0x73, 0x2e, 0x42, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x75, 0x73, 0x2e,
	0x42, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x63, 0x68,
	0x6f, 0x72, 0x75, 0x73, 0x2e, 0x42, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x75, 0x73, 0x2e, 0x46, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_bin_proto_rawDescOnce sync.Once
	file_bin_proto_rawDescData = file_bin_proto_rawDesc
)

func file_bin_proto_rawDescGZIP() []byte {
	file_bin_proto_rawDescOnce.Do(func() {
		file_bin_proto_rawDescData = protoimpl.X.CompressGZIP(file_bin_proto_rawDescData)
	})
	return file_bin_proto_rawDescData
}

var file_bin_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_bin_proto_goTypes = []interface{}{
	(*BinRequests)(nil),        // 0: chorus.BinRequests
	(*BinResponses)(nil),       // 1: chorus.BinResponses
	(*BinRequests_Empty)(nil),  // 2: chorus.BinRequests.Empty
	(*BinRequests_Get)(nil),    // 3: chorus.BinRequests.Get
	(*BinResponses_Empty)(nil), // 4: chorus.BinResponses.Empty
	(*FolderDetails)(nil),      // 5: chorus.FolderDetails
}
var file_bin_proto_depIdxs = []int32{
	2, // 0: chorus.Bin.Empty:input_type -> chorus.BinRequests.Empty
	3, // 1: chorus.Bin.Get:input_type -> chorus.BinRequests.Get
	4, // 2: chorus.Bin.Empty:output_type -> chorus.BinResponses.Empty
	5, // 3: chorus.Bin.Get:output_type -> chorus.FolderDetails
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bin_proto_init() }
func file_bin_proto_init() {
	if File_bin_proto != nil {
		return
	}
	file_folders_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinRequests); i {
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
		file_bin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinResponses); i {
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
		file_bin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinRequests_Empty); i {
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
		file_bin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinRequests_Get); i {
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
		file_bin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinResponses_Empty); i {
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
			RawDescriptor: file_bin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bin_proto_goTypes,
		DependencyIndexes: file_bin_proto_depIdxs,
		MessageInfos:      file_bin_proto_msgTypes,
	}.Build()
	File_bin_proto = out.File
	file_bin_proto_rawDesc = nil
	file_bin_proto_goTypes = nil
	file_bin_proto_depIdxs = nil
}