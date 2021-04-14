// This is a generated file, please do not edit by hand
package chorus

import (
	"context"

	pb "github.com/third-light/go-grpc-client/protobuf"
	"google.golang.org/grpc"
)

// SharedLinkDetails details about a shared link.
type SharedLinkDetails struct {
	// CreatedDate the date/time that this object was created
	CreatedDate DateTime
	// Description a description for this Shared Link, which will get passed on to the link.
	Description string
	// FolderId ID of the special collection that backs this shared link.
	FolderId string
	// Id unique ID of this object.
	Id string
	// ModifiedDate the date/time that this object was last changed
	ModifiedDate DateTime
	// Name a convenient name for this Shared Link.
	Name string
	// Urls for the shared link
	Urls []string
}

// SharedLinksRequests
type SharedLinksRequests struct {
}

// SharedLinksRequestsGet
type SharedLinksRequestsGet struct {
	// SharedLinkId the id, or array thereof, of the Shared Links of which to get the details.
	SharedLinkId string
}

// SharedLinksRequestsGetMulti
type SharedLinksRequestsGetMulti struct {
	// SharedLinkIds the id, or array thereof, of the Shared Links of which to get the details.
	SharedLinkIds []string
}

// SharedLinksResponses
type SharedLinksResponses struct {
}

// SharedLinksResponsesGetMulti
type SharedLinksResponsesGetMulti struct {
	Response []SharedLinkDetails
}

// SharedLinksClient is the SharedLinks client
type SharedLinksClient struct {
	pb pb.SharedLinksClient
}

// NewSharedLinksClient returns a new SharedLinksClient created from supplied gRPC ClientConn
func NewSharedLinksClient(cc *grpc.ClientConn) *SharedLinksClient {
	return &SharedLinksClient{
		pb: pb.NewSharedLinksClient(cc),
	}
}

// Get the details for the requested shared links.
func (c *SharedLinksClient) Get(in *SharedLinksRequestsGet) (*SharedLinkDetails, error) {
	return c.GetWithContext(context.Background(), in)
}

// GetWithContext the details for the requested shared links.
func (c *SharedLinksClient) GetWithContext(ctx context.Context, in *SharedLinksRequestsGet) (*SharedLinkDetails, error) {
	out, err := c.pb.Get(ctx, convertFromSharedLinksRequestsGetPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToSharedLinkDetailsPtr(out), nil
}

// GetMulti get the details for the requested shared links.
func (c *SharedLinksClient) GetMulti(in *SharedLinksRequestsGetMulti) (*SharedLinksResponsesGetMulti, error) {
	return c.GetMultiWithContext(context.Background(), in)
}

// GetMultiWithContext get the details for the requested shared links.
func (c *SharedLinksClient) GetMultiWithContext(ctx context.Context, in *SharedLinksRequestsGetMulti) (*SharedLinksResponsesGetMulti, error) {
	out, err := c.pb.GetMulti(ctx, convertFromSharedLinksRequestsGetMultiPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToSharedLinksResponsesGetMultiPtr(out), nil
}
