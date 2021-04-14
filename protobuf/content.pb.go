// This is a protobuf specifications for the Content module of the Third Light gRPC API for Chorus.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: content.proto

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

// Representation of a type of content in Chorus.
type ContentItemType int32

const (
	// This is a file
	ContentItemType_FILE ContentItemType = 0
	// This is a link to a file.
	ContentItemType_FILE_LINK ContentItemType = 1
	// This is a folder.
	ContentItemType_FOLDER ContentItemType = 2
	// This is a link to a folder.
	ContentItemType_FOLDER_LINK ContentItemType = 3
	// This is a file derivative.
	ContentItemType_DERIVATIVE ContentItemType = 4
)

// Enum value maps for ContentItemType.
var (
	ContentItemType_name = map[int32]string{
		0: "FILE",
		1: "FILE_LINK",
		2: "FOLDER",
		3: "FOLDER_LINK",
		4: "DERIVATIVE",
	}
	ContentItemType_value = map[string]int32{
		"FILE":        0,
		"FILE_LINK":   1,
		"FOLDER":      2,
		"FOLDER_LINK": 3,
		"DERIVATIVE":  4,
	}
)

func (x ContentItemType) Enum() *ContentItemType {
	p := new(ContentItemType)
	*p = x
	return p
}

func (x ContentItemType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContentItemType) Descriptor() protoreflect.EnumDescriptor {
	return file_content_proto_enumTypes[0].Descriptor()
}

func (ContentItemType) Type() protoreflect.EnumType {
	return &file_content_proto_enumTypes[0]
}

func (x ContentItemType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContentItemType.Descriptor instead.
func (ContentItemType) EnumDescriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{0}
}

// One of multiple content specific details of an item in Chorus.
type ContentItemDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Details:
	//	*ContentItemDetails_Derivative
	//	*ContentItemDetails_File
	//	*ContentItemDetails_FileLink
	//	*ContentItemDetails_Folder
	//	*ContentItemDetails_FolderLink
	Details isContentItemDetails_Details `protobuf_oneof:"details"`
}

func (x *ContentItemDetails) Reset() {
	*x = ContentItemDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentItemDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentItemDetails) ProtoMessage() {}

func (x *ContentItemDetails) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentItemDetails.ProtoReflect.Descriptor instead.
func (*ContentItemDetails) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{0}
}

func (m *ContentItemDetails) GetDetails() isContentItemDetails_Details {
	if m != nil {
		return m.Details
	}
	return nil
}

func (x *ContentItemDetails) GetDerivative() *FileDetails {
	if x, ok := x.GetDetails().(*ContentItemDetails_Derivative); ok {
		return x.Derivative
	}
	return nil
}

func (x *ContentItemDetails) GetFile() *FileDetails {
	if x, ok := x.GetDetails().(*ContentItemDetails_File); ok {
		return x.File
	}
	return nil
}

func (x *ContentItemDetails) GetFileLink() *FileLinkDetails {
	if x, ok := x.GetDetails().(*ContentItemDetails_FileLink); ok {
		return x.FileLink
	}
	return nil
}

func (x *ContentItemDetails) GetFolder() *FolderDetails {
	if x, ok := x.GetDetails().(*ContentItemDetails_Folder); ok {
		return x.Folder
	}
	return nil
}

func (x *ContentItemDetails) GetFolderLink() *FolderLinkDetails {
	if x, ok := x.GetDetails().(*ContentItemDetails_FolderLink); ok {
		return x.FolderLink
	}
	return nil
}

type isContentItemDetails_Details interface {
	isContentItemDetails_Details()
}

type ContentItemDetails_Derivative struct {
	Derivative *FileDetails `protobuf:"bytes,5,opt,name=derivative,proto3,oneof"`
}

type ContentItemDetails_File struct {
	File *FileDetails `protobuf:"bytes,1,opt,name=file,proto3,oneof"`
}

type ContentItemDetails_FileLink struct {
	FileLink *FileLinkDetails `protobuf:"bytes,3,opt,name=fileLink,proto3,oneof"`
}

type ContentItemDetails_Folder struct {
	Folder *FolderDetails `protobuf:"bytes,2,opt,name=folder,proto3,oneof"`
}

type ContentItemDetails_FolderLink struct {
	FolderLink *FolderLinkDetails `protobuf:"bytes,4,opt,name=folderLink,proto3,oneof"`
}

func (*ContentItemDetails_Derivative) isContentItemDetails_Details() {}

func (*ContentItemDetails_File) isContentItemDetails_Details() {}

func (*ContentItemDetails_FileLink) isContentItemDetails_Details() {}

func (*ContentItemDetails_Folder) isContentItemDetails_Details() {}

func (*ContentItemDetails_FolderLink) isContentItemDetails_Details() {}

// Representation of a type of an item of content in Chorus.
type ContentItemWithDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Provides type content specific details for the item in question
	Details *ContentItemDetails `protobuf:"bytes,4,opt,name=details,proto3" json:"details,omitempty"`
	// Unique ID of this object.
	Id   string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type ContentItemType `protobuf:"varint,3,opt,name=type,proto3,enum=chorus.ContentItemType" json:"type,omitempty"`
}

func (x *ContentItemWithDetails) Reset() {
	*x = ContentItemWithDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentItemWithDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentItemWithDetails) ProtoMessage() {}

func (x *ContentItemWithDetails) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentItemWithDetails.ProtoReflect.Descriptor instead.
func (*ContentItemWithDetails) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{1}
}

func (x *ContentItemWithDetails) GetDetails() *ContentItemDetails {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *ContentItemWithDetails) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ContentItemWithDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ContentItemWithDetails) GetType() ContentItemType {
	if x != nil {
		return x.Type
	}
	return ContentItemType_FILE
}

type ContentRequests struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ContentRequests) Reset() {
	*x = ContentRequests{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentRequests) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentRequests) ProtoMessage() {}

func (x *ContentRequests) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentRequests.ProtoReflect.Descriptor instead.
func (*ContentRequests) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{2}
}

type ContentResponses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ContentResponses) Reset() {
	*x = ContentResponses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentResponses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentResponses) ProtoMessage() {}

func (x *ContentResponses) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentResponses.ProtoReflect.Descriptor instead.
func (*ContentResponses) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{3}
}

type ContentRequests_Delete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the item to delete
	ItemId string `protobuf:"bytes,1,opt,name=itemId,proto3" json:"itemId,omitempty"`
}

func (x *ContentRequests_Delete) Reset() {
	*x = ContentRequests_Delete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentRequests_Delete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentRequests_Delete) ProtoMessage() {}

func (x *ContentRequests_Delete) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentRequests_Delete.ProtoReflect.Descriptor instead.
func (*ContentRequests_Delete) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ContentRequests_Delete) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

type ContentRequests_Get struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of an item
	ItemId string `protobuf:"bytes,1,opt,name=itemId,proto3" json:"itemId,omitempty"`
}

func (x *ContentRequests_Get) Reset() {
	*x = ContentRequests_Get{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentRequests_Get) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentRequests_Get) ProtoMessage() {}

func (x *ContentRequests_Get) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentRequests_Get.ProtoReflect.Descriptor instead.
func (*ContentRequests_Get) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{2, 1}
}

func (x *ContentRequests_Get) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

type ContentRequests_GetMulti struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The IDs of the items
	ItemIds []string `protobuf:"bytes,1,rep,name=itemIds,proto3" json:"itemIds,omitempty"`
}

func (x *ContentRequests_GetMulti) Reset() {
	*x = ContentRequests_GetMulti{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentRequests_GetMulti) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentRequests_GetMulti) ProtoMessage() {}

func (x *ContentRequests_GetMulti) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentRequests_GetMulti.ProtoReflect.Descriptor instead.
func (*ContentRequests_GetMulti) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{2, 2}
}

func (x *ContentRequests_GetMulti) GetItemIds() []string {
	if x != nil {
		return x.ItemIds
	}
	return nil
}

type ContentResponses_Delete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ContentResponses_Delete) Reset() {
	*x = ContentResponses_Delete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentResponses_Delete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentResponses_Delete) ProtoMessage() {}

func (x *ContentResponses_Delete) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentResponses_Delete.ProtoReflect.Descriptor instead.
func (*ContentResponses_Delete) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{3, 0}
}

type ContentResponses_GetMulti struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response []*ContentItemWithDetails `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty"`
}

func (x *ContentResponses_GetMulti) Reset() {
	*x = ContentResponses_GetMulti{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentResponses_GetMulti) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentResponses_GetMulti) ProtoMessage() {}

func (x *ContentResponses_GetMulti) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentResponses_GetMulti.ProtoReflect.Descriptor instead.
func (*ContentResponses_GetMulti) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{3, 1}
}

func (x *ContentResponses_GetMulti) GetResponse() []*ContentItemWithDetails {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_content_proto protoreflect.FileDescriptor

var file_content_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x63, 0x68, 0x6f, 0x72, 0x75, 0x73, 0x1a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x02, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x35, 0x0a, 0x0a, 0x64, 0x65,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x63, 0x68, 0x6f, 0x72, 0x75, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x64, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x12, 0x29, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x75, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x48, 0x00, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x35, 0x0a, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x63, 0x68, 0x6f, 0x72, 0x75, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x6b,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4c,
	0x69, 0x6e, 0x6b, 0x12, 0x2f, 0x0a, 0x06, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x75, 0x73, 0x2e, 0x46, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x48, 0x00, 0x52, 0x06, 0x66, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0a, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x4c, 0x69,
	0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x75,
	0x73, 0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x6e,
	0x6b, 0x42, 0x09, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x9f, 0x01, 0x0a,
	0x16, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x57, 0x69, 0x74, 0x68,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x75,
	0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x17, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x75, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x78,
	0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x1a, 0x20, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65,
	0x6d, 0x49, 0x64, 0x1a, 0x1d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x74,
	0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d,
	0x49, 0x64, 0x1a, 0x24, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x12, 0x18,
	0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x73, 0x22, 0x64, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x1a, 0x08, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x1a, 0x46, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x12, 0x3a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x75, 0x73, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x57,
	0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x46,
	0x49, 0x4c, 0x45, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x4f,
	0x4c, 0x44, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x4f, 0x4c, 0x44, 0x45, 0x52,
	0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x45, 0x52, 0x49, 0x56,
	0x41, 0x54, 0x49, 0x56, 0x45, 0x10, 0x04, 0x32, 0xef, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x4b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1e, 0x2e,
	0x63, 0x68, 0x6f, 0x72, 0x75, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x1a, 0x1f, 0x2e,
	0x63, 0x68, 0x6f, 0x72, 0x75, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x00,
	0x12, 0x44, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x75, 0x73,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x75, 0x73, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x12, 0x20, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x75, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x1a, 0x21, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x75, 0x73, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_content_proto_rawDescOnce sync.Once
	file_content_proto_rawDescData = file_content_proto_rawDesc
)

func file_content_proto_rawDescGZIP() []byte {
	file_content_proto_rawDescOnce.Do(func() {
		file_content_proto_rawDescData = protoimpl.X.CompressGZIP(file_content_proto_rawDescData)
	})
	return file_content_proto_rawDescData
}

var file_content_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_content_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_content_proto_goTypes = []interface{}{
	(ContentItemType)(0),              // 0: chorus.ContentItemType
	(*ContentItemDetails)(nil),        // 1: chorus.ContentItemDetails
	(*ContentItemWithDetails)(nil),    // 2: chorus.ContentItemWithDetails
	(*ContentRequests)(nil),           // 3: chorus.ContentRequests
	(*ContentResponses)(nil),          // 4: chorus.ContentResponses
	(*ContentRequests_Delete)(nil),    // 5: chorus.ContentRequests.Delete
	(*ContentRequests_Get)(nil),       // 6: chorus.ContentRequests.Get
	(*ContentRequests_GetMulti)(nil),  // 7: chorus.ContentRequests.GetMulti
	(*ContentResponses_Delete)(nil),   // 8: chorus.ContentResponses.Delete
	(*ContentResponses_GetMulti)(nil), // 9: chorus.ContentResponses.GetMulti
	(*FileDetails)(nil),               // 10: chorus.FileDetails
	(*FileLinkDetails)(nil),           // 11: chorus.FileLinkDetails
	(*FolderDetails)(nil),             // 12: chorus.FolderDetails
	(*FolderLinkDetails)(nil),         // 13: chorus.FolderLinkDetails
}
var file_content_proto_depIdxs = []int32{
	10, // 0: chorus.ContentItemDetails.derivative:type_name -> chorus.FileDetails
	10, // 1: chorus.ContentItemDetails.file:type_name -> chorus.FileDetails
	11, // 2: chorus.ContentItemDetails.fileLink:type_name -> chorus.FileLinkDetails
	12, // 3: chorus.ContentItemDetails.folder:type_name -> chorus.FolderDetails
	13, // 4: chorus.ContentItemDetails.folderLink:type_name -> chorus.FolderLinkDetails
	1,  // 5: chorus.ContentItemWithDetails.details:type_name -> chorus.ContentItemDetails
	0,  // 6: chorus.ContentItemWithDetails.type:type_name -> chorus.ContentItemType
	2,  // 7: chorus.ContentResponses.GetMulti.response:type_name -> chorus.ContentItemWithDetails
	5,  // 8: chorus.Content.Delete:input_type -> chorus.ContentRequests.Delete
	6,  // 9: chorus.Content.Get:input_type -> chorus.ContentRequests.Get
	7,  // 10: chorus.Content.GetMulti:input_type -> chorus.ContentRequests.GetMulti
	8,  // 11: chorus.Content.Delete:output_type -> chorus.ContentResponses.Delete
	2,  // 12: chorus.Content.Get:output_type -> chorus.ContentItemWithDetails
	9,  // 13: chorus.Content.GetMulti:output_type -> chorus.ContentResponses.GetMulti
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_content_proto_init() }
func file_content_proto_init() {
	if File_content_proto != nil {
		return
	}
	file_files_proto_init()
	file_folders_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_content_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentItemDetails); i {
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
		file_content_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentItemWithDetails); i {
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
		file_content_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentRequests); i {
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
		file_content_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentResponses); i {
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
		file_content_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentRequests_Delete); i {
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
		file_content_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentRequests_Get); i {
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
		file_content_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentRequests_GetMulti); i {
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
		file_content_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentResponses_Delete); i {
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
		file_content_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentResponses_GetMulti); i {
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
	file_content_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ContentItemDetails_Derivative)(nil),
		(*ContentItemDetails_File)(nil),
		(*ContentItemDetails_FileLink)(nil),
		(*ContentItemDetails_Folder)(nil),
		(*ContentItemDetails_FolderLink)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_content_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_content_proto_goTypes,
		DependencyIndexes: file_content_proto_depIdxs,
		EnumInfos:         file_content_proto_enumTypes,
		MessageInfos:      file_content_proto_msgTypes,
	}.Build()
	File_content_proto = out.File
	file_content_proto_rawDesc = nil
	file_content_proto_goTypes = nil
	file_content_proto_depIdxs = nil
}
