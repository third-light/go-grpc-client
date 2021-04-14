// This is a generated file, please do not edit by hand
package vocabmode

import (
	pb "github.com/third-light/go-grpc-client/protobuf"
)

// VocabMode represents the vocab mode

type VocabMode int32

// VocabModePtr provides a quick way to convert an enum constant to pointer
// for structs with optional members
func VocabModePtr(in VocabMode) *VocabMode { return &in }

// VocabModeFromPtr provides a way to convert back from an enum pointer to a
// guaranteed VocabMode value of some kind (0 is default)
func VocabModeFromPtr(in *VocabMode) VocabMode {
	if in == nil {
		return 0
	}
	return *in
}

// VocabMode constant values
const (
	DISABLED = VocabMode(pb.Vocab_DISABLED)
	INHERIT  = VocabMode(pb.Vocab_INHERIT)
	REPLACE  = VocabMode(pb.Vocab_REPLACE)
)
