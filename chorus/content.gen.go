// This is a generated file, please do not edit by hand
package chorus

import (
	"context"

	"github.com/third-light/go-grpc-client/chorus/contentitemtype"
	"github.com/third-light/go-grpc-client/chorus/filetype"
	pb "github.com/third-light/go-grpc-client/protobuf"
	"google.golang.org/grpc"
)

// FileDetails details about a file in Chorus.
type FileDetails struct {
	// CanViewOriginal if the file is a derivative, whether the current user can see the original file
	CanViewOriginal bool
	// CreatedDate the date/time that this object was created
	CreatedDate DateTime
	// Filename the name of the file
	Filename      string
	FileSizeBytes int64
	// FileType the type of the file
	FileType filetype.FileType
	// Id unique ID of this object.
	Id string
	// IsDerivative whether the file is a derivative or not.
	IsDerivative bool
	Media        MediaDetails
	// ModifiedDate the date/time that this object was last changed
	ModifiedDate DateTime
	// OriginalAssetId if the file is a derivative, the guid for the original file
	OriginalAssetId *string
	// OwnerId unique ID of the object's owning user or space
	OwnerId *string
	// ParentId the GUID of the folder containing this file
	ParentId string
	// RevisionNumber the revision number of the file
	RevisionNumber int64
	// Thumbnails populated keys depend on the get options (e.g. withExtraThumbs)
	Thumbnails Thumbnails
}

// FileLinkDetails details about a file link in Chorus.
type FileLinkDetails struct {
	// CreatedDate the date/time that this object was created
	CreatedDate DateTime
	// Id unique ID of this object.
	Id string
	// LinkName the name of the linked file.
	LinkName string
	// ModifiedDate the date/time that this object was last changed
	ModifiedDate DateTime
	// OriginalFileId the ID of the linked file.
	OriginalFileId string
	// OriginalFileType the type of the linked file.
	OriginalFileType filetype.FileType
}

// FolderLinkDetails details about a folder link in Chorus.
type FolderLinkDetails struct {
	// CreatedDate the date/time that this object was created
	CreatedDate DateTime
	// Id unique ID of this object.
	Id string
	// LinkName the name of the linked folder.
	LinkName string
	// ModifiedDate the date/time that this object was last changed
	ModifiedDate DateTime
	// OriginalFolderId the ID of the linked folder.
	OriginalFolderId string
}

// ContentItemDetails one of multiple content specific details of an item in Chorus.
type ContentItemDetails interface {
	getProto() *pb.ContentItemDetails
}

func (item FileDetails) getProto() *pb.ContentItemDetails {
	return &pb.ContentItemDetails{
		Details: &pb.ContentItemDetails_Derivative{
			Derivative: convertFromFileDetails(item),
		},
	}
}

func (item FileLinkDetails) getProto() *pb.ContentItemDetails {
	return &pb.ContentItemDetails{
		Details: &pb.ContentItemDetails_FileLink{
			FileLink: convertFromFileLinkDetails(item),
		},
	}
}

func (item FolderDetails) getProto() *pb.ContentItemDetails {
	return &pb.ContentItemDetails{
		Details: &pb.ContentItemDetails_Folder{
			Folder: convertFromFolderDetails(item),
		},
	}
}

func (item FolderLinkDetails) getProto() *pb.ContentItemDetails {
	return &pb.ContentItemDetails{
		Details: &pb.ContentItemDetails_FolderLink{
			FolderLink: convertFromFolderLinkDetails(item),
		},
	}
}

// ContentItemWithDetails representation of a type of an item of content in Chorus.
type ContentItemWithDetails struct {
	// Details provides type content specific details for the item in question
	Details ContentItemDetails
	// Id unique ID of this object.
	Id   string
	Name string
	Type contentitemtype.ContentItemType
}

// ContentRequests
type ContentRequests struct {
}

// ContentRequestsDelete
type ContentRequestsDelete struct {
	// ItemId ID of the item to delete
	ItemId string
}

// ContentRequestsGet
type ContentRequestsGet struct {
	// ItemId the ID of an item
	ItemId string
}

// ContentRequestsGetMulti
type ContentRequestsGetMulti struct {
	// ItemIds the IDs of the items
	ItemIds []string
}

// ContentResponses
type ContentResponses struct {
}

// ContentResponsesDelete
type ContentResponsesDelete struct {
}

// ContentResponsesGetMulti
type ContentResponsesGetMulti struct {
	Response []ContentItemWithDetails
}

// ContentClient is the Content client
type ContentClient struct {
	pb pb.ContentClient
}

// NewContentClient returns a new ContentClient created from supplied gRPC ClientConn
func NewContentClient(cc *grpc.ClientConn) *ContentClient {
	return &ContentClient{
		pb: pb.NewContentClient(cc),
	}
}

// Delete an item
func (c *ContentClient) Delete(in *ContentRequestsDelete) (*ContentResponsesDelete, error) {
	return c.DeleteWithContext(context.Background(), in)
}

// DeleteWithContext an item
func (c *ContentClient) DeleteWithContext(ctx context.Context, in *ContentRequestsDelete) (*ContentResponsesDelete, error) {
	out, err := c.pb.Delete(ctx, convertFromContentRequestsDeletePtr(in))
	if err != nil {
		return nil, err
	}
	return convertToContentResponsesDeletePtr(out), nil
}

// Get details about an item.
func (c *ContentClient) Get(in *ContentRequestsGet) (*ContentItemWithDetails, error) {
	return c.GetWithContext(context.Background(), in)
}

// GetWithContext details about an item.
func (c *ContentClient) GetWithContext(ctx context.Context, in *ContentRequestsGet) (*ContentItemWithDetails, error) {
	out, err := c.pb.Get(ctx, convertFromContentRequestsGetPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToContentItemWithDetailsPtr(out), nil
}

// GetMulti get details about an item.
func (c *ContentClient) GetMulti(in *ContentRequestsGetMulti) (*ContentResponsesGetMulti, error) {
	return c.GetMultiWithContext(context.Background(), in)
}

// GetMultiWithContext get details about an item.
func (c *ContentClient) GetMultiWithContext(ctx context.Context, in *ContentRequestsGetMulti) (*ContentResponsesGetMulti, error) {
	out, err := c.pb.GetMulti(ctx, convertFromContentRequestsGetMultiPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToContentResponsesGetMultiPtr(out), nil
}
