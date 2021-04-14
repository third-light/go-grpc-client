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

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	// Returns details about the currently active user.
	GetAuthDetails(ctx context.Context, in *AuthRequests_GetAuthDetails, opts ...grpc.CallOption) (*AuthDetails, error)
	// Login to Chorus using a supplied username and password.
	// This method should only be used if you're allowing users to login using their Chorus username and password.
	// **The usernames and passwords must not be persistently stored and should be requested for each login. Use API
	// Keys, and `LoginWithKey` if you require persistent user access.**
	Login(ctx context.Context, in *AuthRequests_Login, opts ...grpc.CallOption) (*LoginDetails, error)
	// Login to Chorus using an API Key.
	LoginWithKey(ctx context.Context, in *AuthRequests_LoginWithKey, opts ...grpc.CallOption) (*LoginDetails, error)
	// Ends the IMS user session.
	Logout(ctx context.Context, in *AuthRequests_Logout, opts ...grpc.CallOption) (*AuthResponses_Logout, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) GetAuthDetails(ctx context.Context, in *AuthRequests_GetAuthDetails, opts ...grpc.CallOption) (*AuthDetails, error) {
	out := new(AuthDetails)
	err := c.cc.Invoke(ctx, "/chorus.Auth/GetAuthDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Login(ctx context.Context, in *AuthRequests_Login, opts ...grpc.CallOption) (*LoginDetails, error) {
	out := new(LoginDetails)
	err := c.cc.Invoke(ctx, "/chorus.Auth/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) LoginWithKey(ctx context.Context, in *AuthRequests_LoginWithKey, opts ...grpc.CallOption) (*LoginDetails, error) {
	out := new(LoginDetails)
	err := c.cc.Invoke(ctx, "/chorus.Auth/LoginWithKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Logout(ctx context.Context, in *AuthRequests_Logout, opts ...grpc.CallOption) (*AuthResponses_Logout, error) {
	out := new(AuthResponses_Logout)
	err := c.cc.Invoke(ctx, "/chorus.Auth/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}