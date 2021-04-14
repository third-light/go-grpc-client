// This is a generated file, please do not edit by hand
package contexttype

import (
	pb "github.com/third-light/go-grpc-client/protobuf"
)

// ContextType type of the context.

type ContextType int32

// ContextTypePtr provides a quick way to convert an enum constant to pointer
// for structs with optional members
func ContextTypePtr(in ContextType) *ContextType { return &in }

// ContextTypeFromPtr provides a way to convert back from an enum pointer to a
// guaranteed ContextType value of some kind (0 is default)
func ContextTypeFromPtr(in *ContextType) ContextType {
	if in == nil {
		return 0
	}
	return *in
}

// ContextType constant values
const (
	Unknown      = ContextType(pb.ContextType_Unknown)
	PrivateSpace = ContextType(pb.ContextType_PrivateSpace)
	Space        = ContextType(pb.ContextType_Space)
	OtherUser    = ContextType(pb.ContextType_OtherUser)
	Site         = ContextType(pb.ContextType_Site)
)
