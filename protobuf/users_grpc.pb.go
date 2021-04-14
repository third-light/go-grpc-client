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

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersClient interface {
	// Create a new site user
	CreateSiteUser(ctx context.Context, in *UsersRequests_CreateSiteUser, opts ...grpc.CallOption) (*UserDetails, error)
	// Create a new spaces user
	CreateSpacesUser(ctx context.Context, in *UsersRequests_CreateSpacesUser, opts ...grpc.CallOption) (*UserDetails, error)
	// Returns details of the current user.
	// Note: this will return an error if the current user is not a Normal User.
	Current(ctx context.Context, in *UsersRequests_Current, opts ...grpc.CallOption) (*UserDetails, error)
	// Delete the specified user
	Delete(ctx context.Context, in *UsersRequests_Delete, opts ...grpc.CallOption) (*UsersResponses_Delete, error)
	// Get details about users.
	Get(ctx context.Context, in *UsersRequests_Get, opts ...grpc.CallOption) (*UserDetails, error)
	// Get details about users.
	GetMulti(ctx context.Context, in *UsersRequests_GetMulti, opts ...grpc.CallOption) (*UsersResponses_GetMulti, error)
	// Retrieve a list of all user IDs in the system
	GetSiteUsers(ctx context.Context, in *UsersRequests_GetSiteUsers, opts ...grpc.CallOption) (*UsersResponses_GetSiteUsers, error)
	// Retrieve a list of all user IDs in the provided space
	GetSpacesUsers(ctx context.Context, in *UsersRequests_GetSpacesUsers, opts ...grpc.CallOption) (*UsersResponses_GetSpacesUsers, error)
	// Update the details of an API key
	Set(ctx context.Context, in *UsersRequests_Set, opts ...grpc.CallOption) (*UserDetails, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) CreateSiteUser(ctx context.Context, in *UsersRequests_CreateSiteUser, opts ...grpc.CallOption) (*UserDetails, error) {
	out := new(UserDetails)
	err := c.cc.Invoke(ctx, "/chorus.Users/CreateSiteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) CreateSpacesUser(ctx context.Context, in *UsersRequests_CreateSpacesUser, opts ...grpc.CallOption) (*UserDetails, error) {
	out := new(UserDetails)
	err := c.cc.Invoke(ctx, "/chorus.Users/CreateSpacesUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Current(ctx context.Context, in *UsersRequests_Current, opts ...grpc.CallOption) (*UserDetails, error) {
	out := new(UserDetails)
	err := c.cc.Invoke(ctx, "/chorus.Users/Current", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Delete(ctx context.Context, in *UsersRequests_Delete, opts ...grpc.CallOption) (*UsersResponses_Delete, error) {
	out := new(UsersResponses_Delete)
	err := c.cc.Invoke(ctx, "/chorus.Users/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Get(ctx context.Context, in *UsersRequests_Get, opts ...grpc.CallOption) (*UserDetails, error) {
	out := new(UserDetails)
	err := c.cc.Invoke(ctx, "/chorus.Users/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetMulti(ctx context.Context, in *UsersRequests_GetMulti, opts ...grpc.CallOption) (*UsersResponses_GetMulti, error) {
	out := new(UsersResponses_GetMulti)
	err := c.cc.Invoke(ctx, "/chorus.Users/GetMulti", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetSiteUsers(ctx context.Context, in *UsersRequests_GetSiteUsers, opts ...grpc.CallOption) (*UsersResponses_GetSiteUsers, error) {
	out := new(UsersResponses_GetSiteUsers)
	err := c.cc.Invoke(ctx, "/chorus.Users/GetSiteUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) GetSpacesUsers(ctx context.Context, in *UsersRequests_GetSpacesUsers, opts ...grpc.CallOption) (*UsersResponses_GetSpacesUsers, error) {
	out := new(UsersResponses_GetSpacesUsers)
	err := c.cc.Invoke(ctx, "/chorus.Users/GetSpacesUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) Set(ctx context.Context, in *UsersRequests_Set, opts ...grpc.CallOption) (*UserDetails, error) {
	out := new(UserDetails)
	err := c.cc.Invoke(ctx, "/chorus.Users/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
