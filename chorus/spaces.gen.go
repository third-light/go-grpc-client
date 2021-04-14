// This is a generated file, please do not edit by hand
package chorus

import (
	"context"

	pb "github.com/third-light/go-grpc-client/protobuf"
	"google.golang.org/grpc"
)

// SpaceDetails details about a Space in Chorus.
type SpaceDetails struct {
	// AvatarUrl a URL to the current avatar
	AvatarUrl string
	// CreatedDate the date/time that this object was created
	CreatedDate DateTime
	// Description of this Space.
	Description string
	// Id unique ID of this object.
	Id string
	// ModifiedDate the date/time that this object was last changed
	ModifiedDate DateTime
	// Name of this Space.
	Name string
	// ParentId ID of the parent Space, or null if this is a top-level Space.
	ParentId *string
}

// SpacesRequests
type SpacesRequests struct {
}

// SpacesRequestsGet
type SpacesRequestsGet struct {
	// SpaceId ID of the requested Space
	SpaceId string
}

// SpacesRequestsGetMulti
type SpacesRequestsGetMulti struct {
	// SpaceIds ID of the requested Space
	SpaceIds []string
}

// SpacesResponses
type SpacesResponses struct {
}

// SpacesResponsesGetMulti
type SpacesResponsesGetMulti struct {
	Response []SpaceDetails
}

// SpacesClient is the Spaces client
type SpacesClient struct {
	pb pb.SpacesClient
}

// NewSpacesClient returns a new SpacesClient created from supplied gRPC ClientConn
func NewSpacesClient(cc *grpc.ClientConn) *SpacesClient {
	return &SpacesClient{
		pb: pb.NewSpacesClient(cc),
	}
}

// Get the details about a Space.
func (c *SpacesClient) Get(in *SpacesRequestsGet) (*SpaceDetails, error) {
	return c.GetWithContext(context.Background(), in)
}

// GetWithContext the details about a Space.
func (c *SpacesClient) GetWithContext(ctx context.Context, in *SpacesRequestsGet) (*SpaceDetails, error) {
	out, err := c.pb.Get(ctx, convertFromSpacesRequestsGetPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToSpaceDetailsPtr(out), nil
}

// GetMulti get the details about a Space.
func (c *SpacesClient) GetMulti(in *SpacesRequestsGetMulti) (*SpacesResponsesGetMulti, error) {
	return c.GetMultiWithContext(context.Background(), in)
}

// GetMultiWithContext get the details about a Space.
func (c *SpacesClient) GetMultiWithContext(ctx context.Context, in *SpacesRequestsGetMulti) (*SpacesResponsesGetMulti, error) {
	out, err := c.pb.GetMulti(ctx, convertFromSpacesRequestsGetMultiPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToSpacesResponsesGetMultiPtr(out), nil
}
