// This is a generated file, please do not edit by hand
package chorus

import (
	"context"

	pb "github.com/third-light/go-grpc-client/protobuf"
	"google.golang.org/grpc"
)

// FoldersRequests
type FoldersRequests struct {
}

// FoldersRequestsGet
type FoldersRequestsGet struct {
	// FolderId the of the folder
	FolderId string
}

// FoldersRequestsGetChildFiles
type FoldersRequestsGetChildFiles struct {
	// FolderId the ID of the folder to list
	FolderId string
}

// FoldersRequestsGetChildFolders
type FoldersRequestsGetChildFolders struct {
	// FolderId the ID of the folder to list
	FolderId string
}

// FoldersRequestsGetLink
type FoldersRequestsGetLink struct {
	// LinkId the ID of the folder link
	LinkId string
}

// FoldersRequestsGetLinkMulti
type FoldersRequestsGetLinkMulti struct {
	// LinkIds the IDs of the folder links
	LinkIds []string
}

// FoldersRequestsGetMulti
type FoldersRequestsGetMulti struct {
	// FolderIds the IDs of the folder
	FolderIds []string
}

// FoldersRequestsSet
type FoldersRequestsSet struct {
	// Details the folder details to set.
	Details SettableFolderDetails `fieldMask:"details"`
	// FolderId a folder ID for the folder of interest.
	FolderId string
}

// FoldersResponses
type FoldersResponses struct {
}

// FoldersResponsesGetChildFiles
type FoldersResponsesGetChildFiles struct {
	Response []FileDetails
}

// FoldersResponsesGetChildFolders
type FoldersResponsesGetChildFolders struct {
	Response []FolderDetails
}

// FoldersResponsesGetLinkMulti
type FoldersResponsesGetLinkMulti struct {
	Response []FolderLinkDetails
}

// FoldersResponsesGetMulti
type FoldersResponsesGetMulti struct {
	Response []FolderDetails
}

// SettableFolderDetails settings to use for updating a folder's details.
type SettableFolderDetails struct {
	// Description the description of the folder.
	Description *string `fieldMask:"description"`
	// Name the name of the folder.
	Name *string `fieldMask:"name"`
}

// FoldersClient is the Folders client
type FoldersClient struct {
	pb pb.FoldersClient
}

// NewFoldersClient returns a new FoldersClient created from supplied gRPC ClientConn
func NewFoldersClient(cc *grpc.ClientConn) *FoldersClient {
	return &FoldersClient{
		pb: pb.NewFoldersClient(cc),
	}
}

// Get details about a given folder.
func (c *FoldersClient) Get(in *FoldersRequestsGet) (*FolderDetails, error) {
	return c.GetWithContext(context.Background(), in)
}

// GetWithContext details about a given folder.
func (c *FoldersClient) GetWithContext(ctx context.Context, in *FoldersRequestsGet) (*FolderDetails, error) {
	out, err := c.pb.Get(ctx, convertFromFoldersRequestsGetPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToFolderDetailsPtr(out), nil
}

// GetChildFiles get all the files for a given parent folder.
func (c *FoldersClient) GetChildFiles(in *FoldersRequestsGetChildFiles) (*FoldersResponsesGetChildFiles, error) {
	return c.GetChildFilesWithContext(context.Background(), in)
}

// GetChildFilesWithContext get all the files for a given parent folder.
func (c *FoldersClient) GetChildFilesWithContext(ctx context.Context, in *FoldersRequestsGetChildFiles) (*FoldersResponsesGetChildFiles, error) {
	out, err := c.pb.GetChildFiles(ctx, convertFromFoldersRequestsGetChildFilesPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToFoldersResponsesGetChildFilesPtr(out), nil
}

// GetChildFolders get all the sub folders for a given parent folder.
func (c *FoldersClient) GetChildFolders(in *FoldersRequestsGetChildFolders) (*FoldersResponsesGetChildFolders, error) {
	return c.GetChildFoldersWithContext(context.Background(), in)
}

// GetChildFoldersWithContext get all the sub folders for a given parent folder.
func (c *FoldersClient) GetChildFoldersWithContext(ctx context.Context, in *FoldersRequestsGetChildFolders) (*FoldersResponsesGetChildFolders, error) {
	out, err := c.pb.GetChildFolders(ctx, convertFromFoldersRequestsGetChildFoldersPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToFoldersResponsesGetChildFoldersPtr(out), nil
}

// GetLink get details of a link to a folder
// If `linkId` is specified as an array, the return hash will be nested in a hash keyed by the link ID.
func (c *FoldersClient) GetLink(in *FoldersRequestsGetLink) (*FolderLinkDetails, error) {
	return c.GetLinkWithContext(context.Background(), in)
}

// GetLinkWithContext get details of a link to a folder
// If `linkId` is specified as an array, the return hash will be nested in a hash keyed by the link ID.
func (c *FoldersClient) GetLinkWithContext(ctx context.Context, in *FoldersRequestsGetLink) (*FolderLinkDetails, error) {
	out, err := c.pb.GetLink(ctx, convertFromFoldersRequestsGetLinkPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToFolderLinkDetailsPtr(out), nil
}

// GetLinkMulti get details of a link to a folder
// If `linkId` is specified as an array, the return hash will be nested in a hash keyed by the link ID.
func (c *FoldersClient) GetLinkMulti(in *FoldersRequestsGetLinkMulti) (*FoldersResponsesGetLinkMulti, error) {
	return c.GetLinkMultiWithContext(context.Background(), in)
}

// GetLinkMultiWithContext get details of a link to a folder
// If `linkId` is specified as an array, the return hash will be nested in a hash keyed by the link ID.
func (c *FoldersClient) GetLinkMultiWithContext(ctx context.Context, in *FoldersRequestsGetLinkMulti) (*FoldersResponsesGetLinkMulti, error) {
	out, err := c.pb.GetLinkMulti(ctx, convertFromFoldersRequestsGetLinkMultiPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToFoldersResponsesGetLinkMultiPtr(out), nil
}

// GetMulti get details about a given folder.
func (c *FoldersClient) GetMulti(in *FoldersRequestsGetMulti) (*FoldersResponsesGetMulti, error) {
	return c.GetMultiWithContext(context.Background(), in)
}

// GetMultiWithContext get details about a given folder.
func (c *FoldersClient) GetMultiWithContext(ctx context.Context, in *FoldersRequestsGetMulti) (*FoldersResponsesGetMulti, error) {
	out, err := c.pb.GetMulti(ctx, convertFromFoldersRequestsGetMultiPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToFoldersResponsesGetMultiPtr(out), nil
}

// Set details on a folder. Allows setting the search parameters for a Smart Folder.
func (c *FoldersClient) Set(in *FoldersRequestsSet) (*FolderDetails, error) {
	return c.SetWithContext(context.Background(), in)
}

// SetWithContext details on a folder. Allows setting the search parameters for a Smart Folder.
func (c *FoldersClient) SetWithContext(ctx context.Context, in *FoldersRequestsSet) (*FolderDetails, error) {
	out, err := c.pb.Set(ctx, convertFromFoldersRequestsSetPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToFolderDetailsPtr(out), nil
}
