// This is a generated file, please do not edit by hand
package contentitemtype

import (
	pb "github.com/third-light/go-grpc-client/protobuf"
)

// ContentItemType representation of a type of content in Chorus.

type ContentItemType int32

// ContentItemTypePtr provides a quick way to convert an enum constant to pointer
// for structs with optional members
func ContentItemTypePtr(in ContentItemType) *ContentItemType { return &in }

// ContentItemTypeFromPtr provides a way to convert back from an enum pointer to a
// guaranteed ContentItemType value of some kind (0 is default)
func ContentItemTypeFromPtr(in *ContentItemType) ContentItemType {
	if in == nil {
		return 0
	}
	return *in
}

// ContentItemType constant values
const (
	FILE        = ContentItemType(pb.ContentItemType_FILE)
	FILE_LINK   = ContentItemType(pb.ContentItemType_FILE_LINK)
	FOLDER      = ContentItemType(pb.ContentItemType_FOLDER)
	FOLDER_LINK = ContentItemType(pb.ContentItemType_FOLDER_LINK)
	DERIVATIVE  = ContentItemType(pb.ContentItemType_DERIVATIVE)
)
