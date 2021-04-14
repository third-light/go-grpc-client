// This is a generated file, please do not edit by hand
package authdetailstype

import (
	pb "github.com/third-light/go-grpc-client/protobuf"
)

// AuthDetailsType represents the type of a normal user.

type AuthDetailsType int32

// AuthDetailsTypePtr provides a quick way to convert an enum constant to pointer
// for structs with optional members
func AuthDetailsTypePtr(in AuthDetailsType) *AuthDetailsType { return &in }

// AuthDetailsTypeFromPtr provides a way to convert back from an enum pointer to a
// guaranteed AuthDetailsType value of some kind (0 is default)
func AuthDetailsTypeFromPtr(in *AuthDetailsType) AuthDetailsType {
	if in == nil {
		return 0
	}
	return *in
}

// AuthDetailsType constant values
const (
	NotLoggedIn  = AuthDetailsType(pb.AuthDetails_NotLoggedIn)
	User         = AuthDetailsType(pb.AuthDetails_User)
	ExternalUser = AuthDetailsType(pb.AuthDetails_ExternalUser)
)
