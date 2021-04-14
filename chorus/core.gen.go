// This is a generated file, please do not edit by hand
package chorus

import (
	"context"

	pb "github.com/third-light/go-grpc-client/protobuf"
	"google.golang.org/grpc"
)

// CoreRequests
type CoreRequests struct {
}

// CoreRequestsGetEnvironment
type CoreRequestsGetEnvironment struct {
}

// DateTime this is a representation of a date time.
type DateTime struct {
	// Rfc3339 the date encoded as per RFC-3339 (a.k.a ATOM)
	Rfc3339 string
	// Timestamp this is the date/time represented as the number of seconds since the Unix Epoch (January 1 1970, 00:00:00 UTC)
	Timestamp int64
}

// EnvironmentDetails details of the Chorus server.
type EnvironmentDetails struct {
	// BrowserTitle the version of the title to use as the browser title (this often is longer and more verbose). If no specific
	// browser title has been configured, then this will be the same as the 'title'.
	BrowserTitle string
	// Title the site's title. This is used for the shortcuts and other places referring to the site.
	Title string
	// Version the current version of Chorus.
	Version string
}

// CoreClient is the Core client
type CoreClient struct {
	pb pb.CoreClient
}

// NewCoreClient returns a new CoreClient created from supplied gRPC ClientConn
func NewCoreClient(cc *grpc.ClientConn) *CoreClient {
	return &CoreClient{
		pb: pb.NewCoreClient(cc),
	}
}

// GetEnvironment returns the environment settings for the current user; useful for initial configuration
// of a UI or API based application.
func (c *CoreClient) GetEnvironment(in *CoreRequestsGetEnvironment) (*EnvironmentDetails, error) {
	return c.GetEnvironmentWithContext(context.Background(), in)
}

// GetEnvironmentWithContext returns the environment settings for the current user; useful for initial configuration
// of a UI or API based application.
func (c *CoreClient) GetEnvironmentWithContext(ctx context.Context, in *CoreRequestsGetEnvironment) (*EnvironmentDetails, error) {
	out, err := c.pb.GetEnvironment(ctx, convertFromCoreRequestsGetEnvironmentPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToEnvironmentDetailsPtr(out), nil
}
