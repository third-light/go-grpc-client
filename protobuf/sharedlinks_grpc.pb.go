// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protobuf

import (
	context "context"

	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SharedLinksClient is the client API for SharedLinks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SharedLinksClient interface {
	// Get the details for the requested shared links.
	Get(ctx context.Context, in *SharedLinksRequests_Get, opts ...grpc.CallOption) (*SharedLinkDetails, error)
	// Get the details for the requested shared links.
	GetMulti(ctx context.Context, in *SharedLinksRequests_GetMulti, opts ...grpc.CallOption) (*SharedLinksResponses_GetMulti, error)
}

type sharedLinksClient struct {
	cc grpc.ClientConnInterface
}

func NewSharedLinksClient(cc grpc.ClientConnInterface) SharedLinksClient {
	return &sharedLinksClient{cc}
}

func (c *sharedLinksClient) Get(ctx context.Context, in *SharedLinksRequests_Get, opts ...grpc.CallOption) (*SharedLinkDetails, error) {
	out := new(SharedLinkDetails)
	err := c.cc.Invoke(ctx, "/chorus.SharedLinks/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharedLinksClient) GetMulti(ctx context.Context, in *SharedLinksRequests_GetMulti, opts ...grpc.CallOption) (*SharedLinksResponses_GetMulti, error) {
	out := new(SharedLinksResponses_GetMulti)
	err := c.cc.Invoke(ctx, "/chorus.SharedLinks/GetMulti", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
