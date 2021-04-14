package chorus

import (
	"context"
	"encoding/json"

	pb "github.com/third-light/go-grpc-client/protobuf"
	"google.golang.org/grpc"
)

// InternalApiClient provides access to the Chorus internal API via
// gRPC. This is provided without support, and a stern warning that
// the internal API is subject to change.
type InternalApiClient struct {
	pb pb.InternalApiClient
}

func NewInternalApiClient(cc *grpc.ClientConn) *InternalApiClient {
	return &InternalApiClient{
		pb: pb.NewInternalApiClient(cc),
	}
}

func (c *InternalApiClient) Do(method string, payload interface{}) (*pb.InternalApiResponse, error) {
	return c.DoWithContext(context.Background(), method, payload)
}

func (c *InternalApiClient) DoWithContext(ctx context.Context, method string, payload interface{}) (*pb.InternalApiResponse, error) {
	var s string
	switch payload.(type) {
	case string:
		s = payload.(string)
	case []byte:
		s = string(payload.([]byte))
	default:
		b, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		s = string(b)
	}
	return c.pb.Do(ctx, &pb.InternalApiRequest{Method: method, Payload: s})
}
