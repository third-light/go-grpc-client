// This is a protobuf specifications for the Spaces module of the Third Light gRPC API for Chorus.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: spaces.proto

package protobuf

import (
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// Details about a Space in Chorus.
type SpaceDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// a URL to the current avatar
	AvatarUrl string `protobuf:"bytes,51,opt,name=avatarUrl,proto3" json:"avatarUrl,omitempty"`
	// The date/time that this object was created
	CreatedDate *DateTime `protobuf:"bytes,102,opt,name=createdDate,proto3" json:"createdDate,omitempty"`
	// Description of this Space.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Unique ID of this object.
	Id string `protobuf:"bytes,101,opt,name=id,proto3" json:"id,omitempty"`
	// The date/time that this object was last changed
	ModifiedDate *DateTime `protobuf:"bytes,103,opt,name=modifiedDate,proto3" json:"modifiedDate,omitempty"`
	// Name of this Space.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ID of the parent Space, or null if this is a top-level Space.
	ParentId *wrappers.StringValue `protobuf:"bytes,3,opt,name=parentId,proto3" json:"parentId,omitempty"`
}

func (x *SpaceDetails) Reset() {
	*x = SpaceDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaces_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpaceDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpaceDetails) ProtoMessage() {}

func (x *SpaceDetails) ProtoReflect() protoreflect.Message {
	mi := &file_spaces_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpaceDetails.ProtoReflect.Descriptor instead.
func (*SpaceDetails) Descriptor() ([]byte, []int) {
	return file_spaces_proto_rawDescGZIP(), []int{0}
}

func (x *SpaceDetails) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *SpaceDetails) GetCreatedDate() *DateTime {
	if x != nil {
		return x.CreatedDate
	}
	return nil
}

func (x *SpaceDetails) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SpaceDetails) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SpaceDetails) GetModifiedDate() *DateTime {
	if x != nil {
		return x.ModifiedDate
	}
	return nil
}

func (x *SpaceDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SpaceDetails) GetParentId() *wrappers.StringValue {
	if x != nil {
		return x.ParentId
	}
	return nil
}

type SpacesRequests struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SpacesRequests) Reset() {
	*x = SpacesRequests{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaces_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpacesRequests) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpacesRequests) ProtoMessage() {}

func (x *SpacesRequests) ProtoReflect() protoreflect.Message {
	mi := &file_spaces_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpacesRequests.ProtoReflect.Descriptor instead.
func (*SpacesRequests) Descriptor() ([]byte, []int) {
	return file_spaces_proto_rawDescGZIP(), []int{1}
}

type SpacesResponses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SpacesResponses) Reset() {
	*x = SpacesResponses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaces_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpacesResponses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpacesResponses) ProtoMessage() {}

func (x *SpacesResponses) ProtoReflect() protoreflect.Message {
	mi := &file_spaces_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpacesResponses.ProtoReflect.Descriptor instead.
func (*SpacesResponses) Descriptor() ([]byte, []int) {
	return file_spaces_proto_rawDescGZIP(), []int{2}
}

type SpacesRequests_Get struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the requested Space
	SpaceId string `protobuf:"bytes,1,opt,name=spaceId,proto3" json:"spaceId,omitempty"`
}

func (x *SpacesRequests_Get) Reset() {
	*x = SpacesRequests_Get{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaces_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpacesRequests_Get) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpacesRequests_Get) ProtoMessage() {}

func (x *SpacesRequests_Get) ProtoReflect() protoreflect.Message {
	mi := &file_spaces_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpacesRequests_Get.ProtoReflect.Descriptor instead.
func (*SpacesRequests_Get) Descriptor() ([]byte, []int) {
	return file_spaces_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SpacesRequests_Get) GetSpaceId() string {
	if x != nil {
		return x.SpaceId
	}
	return ""
}

type SpacesRequests_GetMulti struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the requested Space
	SpaceIds []string `protobuf:"bytes,1,rep,name=spaceIds,proto3" json:"spaceIds,omitempty"`
}

func (x *SpacesRequests_GetMulti) Reset() {
	*x = SpacesRequests_GetMulti{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaces_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpacesRequests_GetMulti) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpacesRequests_GetMulti) ProtoMessage() {}

func (x *SpacesRequests_GetMulti) ProtoReflect() protoreflect.Message {
	mi := &file_spaces_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpacesRequests_GetMulti.ProtoReflect.Descriptor instead.
func (*SpacesRequests_GetMulti) Descriptor() ([]byte, []int) {
	return file_spaces_proto_rawDescGZIP(), []int{1, 1}
}

func (x *SpacesRequests_GetMulti) GetSpaceIds() []string {
	if x != nil {
		return x.SpaceIds
	}
	return nil
}

type SpacesResponses_GetMulti struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response []*SpaceDetails `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty"`
}

func (x *SpacesResponses_GetMulti) Reset() {
	*x = SpacesResponses_GetMulti{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaces_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpacesResponses_GetMulti) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpacesResponses_GetMulti) ProtoMessage() {}

func (x *SpacesResponses_GetMulti) ProtoReflect() protoreflect.Message {
	mi := &file_spaces_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpacesResponses_GetMulti.ProtoReflect.Descriptor instead.
func (*SpacesResponses_GetMulti) Descriptor() ([]byte, []int) {
	return file_spaces_proto_rawDescGZIP(), []int{2, 0}
}

func (x *SpacesResponses_GetMulti) GetResponse() []*SpaceDetails {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_spaces_proto protoreflect.FileDescriptor

var file_spaces_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x68, 0x6f, 0x72, 0x75, 0x73, 0x1a, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x96, 0x02, 0x0a, 0x0c, 0x53, 0x70, 0x61, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c,
	0x18, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72,
	0x6c, 0x12, 0x32, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x75, 0x73, 0x2e,
	0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x65, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x0c, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x63, 0x68, 0x6f, 0x72, 0x75, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52,
	0x0c, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x38, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x59, 0x0a, 0x0e, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x1a, 0x1f, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x1a, 0x26,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x73, 0x22, 0x4f, 0x0a, 0x0f, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x1a, 0x3c, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x12, 0x30, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x75, 0x73,
	0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x94, 0x01, 0x0a, 0x06, 0x53, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x12, 0x39, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x63, 0x68, 0x6f, 0x72,
	0x75, 0x73, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x75, 0x73, 0x2e, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x12, 0x4f, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x12, 0x1f, 0x2e, 0x63, 0x68, 0x6f, 0x72,
	0x75, 0x73, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x1a, 0x20, 0x2e, 0x63, 0x68, 0x6f,
	0x72, 0x75, 0x73, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x22, 0x00, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaces_proto_rawDescOnce sync.Once
	file_spaces_proto_rawDescData = file_spaces_proto_rawDesc
)

func file_spaces_proto_rawDescGZIP() []byte {
	file_spaces_proto_rawDescOnce.Do(func() {
		file_spaces_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaces_proto_rawDescData)
	})
	return file_spaces_proto_rawDescData
}

var file_spaces_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_spaces_proto_goTypes = []interface{}{
	(*SpaceDetails)(nil),             // 0: chorus.SpaceDetails
	(*SpacesRequests)(nil),           // 1: chorus.SpacesRequests
	(*SpacesResponses)(nil),          // 2: chorus.SpacesResponses
	(*SpacesRequests_Get)(nil),       // 3: chorus.SpacesRequests.Get
	(*SpacesRequests_GetMulti)(nil),  // 4: chorus.SpacesRequests.GetMulti
	(*SpacesResponses_GetMulti)(nil), // 5: chorus.SpacesResponses.GetMulti
	(*DateTime)(nil),                 // 6: chorus.DateTime
	(*wrappers.StringValue)(nil),     // 7: google.protobuf.StringValue
}
var file_spaces_proto_depIdxs = []int32{
	6, // 0: chorus.SpaceDetails.createdDate:type_name -> chorus.DateTime
	6, // 1: chorus.SpaceDetails.modifiedDate:type_name -> chorus.DateTime
	7, // 2: chorus.SpaceDetails.parentId:type_name -> google.protobuf.StringValue
	0, // 3: chorus.SpacesResponses.GetMulti.response:type_name -> chorus.SpaceDetails
	3, // 4: chorus.Spaces.Get:input_type -> chorus.SpacesRequests.Get
	4, // 5: chorus.Spaces.GetMulti:input_type -> chorus.SpacesRequests.GetMulti
	0, // 6: chorus.Spaces.Get:output_type -> chorus.SpaceDetails
	5, // 7: chorus.Spaces.GetMulti:output_type -> chorus.SpacesResponses.GetMulti
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_spaces_proto_init() }
func file_spaces_proto_init() {
	if File_spaces_proto != nil {
		return
	}
	file_core_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_spaces_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpaceDetails); i {
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
		file_spaces_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpacesRequests); i {
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
		file_spaces_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpacesResponses); i {
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
		file_spaces_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpacesRequests_Get); i {
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
		file_spaces_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpacesRequests_GetMulti); i {
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
		file_spaces_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpacesResponses_GetMulti); i {
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
			RawDescriptor: file_spaces_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaces_proto_goTypes,
		DependencyIndexes: file_spaces_proto_depIdxs,
		MessageInfos:      file_spaces_proto_msgTypes,
	}.Build()
	File_spaces_proto = out.File
	file_spaces_proto_rawDesc = nil
	file_spaces_proto_goTypes = nil
	file_spaces_proto_depIdxs = nil
}
