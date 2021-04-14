// This is a generated file, please do not edit by hand
package chorus

import (
	"context"

	"github.com/third-light/go-grpc-client/chorus/folderdetailstype"
	pb "github.com/third-light/go-grpc-client/protobuf"
	"google.golang.org/grpc"
)

// BinRequests
type BinRequests struct {
}

// BinRequestsEmpty
type BinRequestsEmpty struct {
}

// BinRequestsGet
type BinRequestsGet struct {
}

// BinResponses
type BinResponses struct {
}

// BinResponsesEmpty
type BinResponsesEmpty struct {
}

// FolderDetails details about a Chorus folder.
type FolderDetails struct {
	// Context the identifier for the containing context
	Context string
	// CreatedDate the date/time that this object was created
	CreatedDate DateTime
	// Description the description of the folder.
	Description string
	// Id unique ID of this object.
	Id string
	// ModifiedDate the date/time that this object was last changed
	ModifiedDate DateTime
	Name         string
	// ParentId the ID of the parent of the folder, or null if there is no parent
	ParentId *string
	// Type the type of the folder.
	Type folderdetailstype.FolderDetailsType
}

// BinClient is the Bin client
type BinClient struct {
	pb pb.BinClient
}

// NewBinClient returns a new BinClient created from supplied gRPC ClientConn
func NewBinClient(cc *grpc.ClientConn) *BinClient {
	return &BinClient{
		pb: pb.NewBinClient(cc),
	}
}

// Empty get details about the user's bin
func (c *BinClient) Empty(in *BinRequestsEmpty) (*BinResponsesEmpty, error) {
	return c.EmptyWithContext(context.Background(), in)
}

// EmptyWithContext get details about the user's bin
func (c *BinClient) EmptyWithContext(ctx context.Context, in *BinRequestsEmpty) (*BinResponsesEmpty, error) {
	out, err := c.pb.Empty(ctx, convertFromBinRequestsEmptyPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToBinResponsesEmptyPtr(out), nil
}

// Get details about the user's bin
func (c *BinClient) Get(in *BinRequestsGet) (*FolderDetails, error) {
	return c.GetWithContext(context.Background(), in)
}

// GetWithContext details about the user's bin
func (c *BinClient) GetWithContext(ctx context.Context, in *BinRequestsGet) (*FolderDetails, error) {
	out, err := c.pb.Get(ctx, convertFromBinRequestsGetPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToFolderDetailsPtr(out), nil
}
