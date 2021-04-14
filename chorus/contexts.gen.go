// This is a generated file, please do not edit by hand
package chorus

import (
	"context"

	"github.com/third-light/go-grpc-client/chorus/contexttype"
	pb "github.com/third-light/go-grpc-client/protobuf"
	"google.golang.org/grpc"
)

// ContextDetails details about a "context". A context is the root of a section of the site.
type ContextDetails struct {
	// AvatarUrl the URL for the avatar of the context.
	AvatarUrl string
	// BackingFolderId ID of the folder that is the home of this context. This will be null if the current user cannot view the folders.
	BackingFolderId *string
	// Id the ID of the context.
	Id string
	// Name the name of the context.
	Name string
	// OwnerId ID of the Space or User that this is the context of. This will be null either if you cannot see the Space or User, or this is
	// a context for the site.
	OwnerId *string
	// ParentId the ID of the parent context.
	ParentId string
	// Type the type of the context.
	Type contexttype.ContextType
}

// ContextsRequests
type ContextsRequests struct {
}

// ContextsRequestsGet
type ContextsRequestsGet struct {
	// ContextId an array of contexts for which details should be returned.
	ContextId string
}

// ContextsRequestsGetChildren
type ContextsRequestsGetChildren struct {
	// ContextId the parent context for which to get children
	ContextId string
}

// ContextsRequestsGetMulti
type ContextsRequestsGetMulti struct {
	// ContextIds an array of contexts for which details should be returned.
	ContextIds []string
}

// ContextsRequestsGetSharedLinks
type ContextsRequestsGetSharedLinks struct {
	// ContextId ID of the context for which to get shared links
	ContextId string
}

// ContextsResponses
type ContextsResponses struct {
}

// ContextsResponsesGetChildren
type ContextsResponsesGetChildren struct {
	Response []string
}

// ContextsResponsesGetMulti
type ContextsResponsesGetMulti struct {
	Response []ContextDetails
}

// ContextsResponsesGetSharedLinks
type ContextsResponsesGetSharedLinks struct {
	Response []SharedLinkDetails
}

// ContextsClient is the Contexts client
type ContextsClient struct {
	pb pb.ContextsClient
}

// NewContextsClient returns a new ContextsClient created from supplied gRPC ClientConn
func NewContextsClient(cc *grpc.ClientConn) *ContextsClient {
	return &ContextsClient{
		pb: pb.NewContextsClient(cc),
	}
}

// Get details about one or more contexts.
func (c *ContextsClient) Get(in *ContextsRequestsGet) (*ContextDetails, error) {
	return c.GetWithContext(context.Background(), in)
}

// GetWithContext details about one or more contexts.
func (c *ContextsClient) GetWithContext(ctx context.Context, in *ContextsRequestsGet) (*ContextDetails, error) {
	out, err := c.pb.Get(ctx, convertFromContextsRequestsGetPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToContextDetailsPtr(out), nil
}

// GetChildren get the logical areas visible to the current user
// This will return site-wide "domain" contexts, the user's own home directory ("me"),
// the home directories of any groups to which the user belongs, and the public folders
// of any other users or groups in the site.
// The context IDs returned are suitable for use with Folders.GetTopLevelFolders and suchlike
func (c *ContextsClient) GetChildren(in *ContextsRequestsGetChildren) (*ContextsResponsesGetChildren, error) {
	return c.GetChildrenWithContext(context.Background(), in)
}

// GetChildrenWithContext get the logical areas visible to the current user
// This will return site-wide "domain" contexts, the user's own home directory ("me"),
// the home directories of any groups to which the user belongs, and the public folders
// of any other users or groups in the site.
// The context IDs returned are suitable for use with Folders.GetTopLevelFolders and suchlike
func (c *ContextsClient) GetChildrenWithContext(ctx context.Context, in *ContextsRequestsGetChildren) (*ContextsResponsesGetChildren, error) {
	out, err := c.pb.GetChildren(ctx, convertFromContextsRequestsGetChildrenPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToContextsResponsesGetChildrenPtr(out), nil
}

// GetMulti get details about one or more contexts.
func (c *ContextsClient) GetMulti(in *ContextsRequestsGetMulti) (*ContextsResponsesGetMulti, error) {
	return c.GetMultiWithContext(context.Background(), in)
}

// GetMultiWithContext get details about one or more contexts.
func (c *ContextsClient) GetMultiWithContext(ctx context.Context, in *ContextsRequestsGetMulti) (*ContextsResponsesGetMulti, error) {
	out, err := c.pb.GetMulti(ctx, convertFromContextsRequestsGetMultiPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToContextsResponsesGetMultiPtr(out), nil
}

// GetSharedLinks get details about the shared links for a context. This is only available for Users and Spaces.
func (c *ContextsClient) GetSharedLinks(in *ContextsRequestsGetSharedLinks) (*ContextsResponsesGetSharedLinks, error) {
	return c.GetSharedLinksWithContext(context.Background(), in)
}

// GetSharedLinksWithContext get details about the shared links for a context. This is only available for Users and Spaces.
func (c *ContextsClient) GetSharedLinksWithContext(ctx context.Context, in *ContextsRequestsGetSharedLinks) (*ContextsResponsesGetSharedLinks, error) {
	out, err := c.pb.GetSharedLinks(ctx, convertFromContextsRequestsGetSharedLinksPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToContextsResponsesGetSharedLinksPtr(out), nil
}
