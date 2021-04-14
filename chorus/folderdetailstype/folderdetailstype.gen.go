// This is a generated file, please do not edit by hand
package folderdetailstype

import (
	pb "github.com/third-light/go-grpc-client/protobuf"
)

// FolderDetailsType representation of the type of a folder in Chorus.

type FolderDetailsType int32

// FolderDetailsTypePtr provides a quick way to convert an enum constant to pointer
// for structs with optional members
func FolderDetailsTypePtr(in FolderDetailsType) *FolderDetailsType { return &in }

// FolderDetailsTypeFromPtr provides a way to convert back from an enum pointer to a
// guaranteed FolderDetailsType value of some kind (0 is default)
func FolderDetailsTypeFromPtr(in *FolderDetailsType) FolderDetailsType {
	if in == nil {
		return 0
	}
	return *in
}

// FolderDetailsType constant values
const (
	Container               = FolderDetailsType(pb.FolderDetails_Container)
	Folder                  = FolderDetailsType(pb.FolderDetails_Folder)
	SmartCollection         = FolderDetailsType(pb.FolderDetails_SmartCollection)
	Collection              = FolderDetailsType(pb.FolderDetails_Collection)
	RecycleBin              = FolderDetailsType(pb.FolderDetails_RecycleBin)
	ContextFolder           = FolderDetailsType(pb.FolderDetails_ContextFolder)
	SearchResults           = FolderDetailsType(pb.FolderDetails_SearchResults)
	DeletedFolder           = FolderDetailsType(pb.FolderDetails_DeletedFolder)
	SearchRoot              = FolderDetailsType(pb.FolderDetails_SearchRoot)
	ShareCollection         = FolderDetailsType(pb.FolderDetails_ShareCollection)
	SharedLinkCollection    = FolderDetailsType(pb.FolderDetails_SharedLinkCollection)
	SharesRoot              = FolderDetailsType(pb.FolderDetails_SharesRoot)
	SharedLinksRoot         = FolderDetailsType(pb.FolderDetails_SharedLinksRoot)
	IncomingAttachments     = FolderDetailsType(pb.FolderDetails_IncomingAttachments)
	OutgoingAttachments     = FolderDetailsType(pb.FolderDetails_OutgoingAttachments)
	WorkflowFolder          = FolderDetailsType(pb.FolderDetails_WorkflowFolder)
	WorkflowCollection      = FolderDetailsType(pb.FolderDetails_WorkflowCollection)
	WorkflowPrivateRoot     = FolderDetailsType(pb.FolderDetails_WorkflowPrivateRoot)
	DerivativeCollection    = FolderDetailsType(pb.FolderDetails_DerivativeCollection)
	SyncFolder              = FolderDetailsType(pb.FolderDetails_SyncFolder)
	PrivateCollection       = FolderDetailsType(pb.FolderDetails_PrivateCollection)
	WorkflowBatchCollection = FolderDetailsType(pb.FolderDetails_WorkflowBatchCollection)
)
