// This is a generated file, please do not edit by hand
package directurlsettingsfitmode

import (
	pb "github.com/third-light/go-grpc-client/protobuf"
)

// DirectUrlSettingsFitMode representation of the 'fit mode' for a requested crop. This is the strategy to use when the requested crop's dimensions
// don't match the output dimensions or aspect ratio.

type DirectUrlSettingsFitMode int32

// DirectUrlSettingsFitModePtr provides a quick way to convert an enum constant to pointer
// for structs with optional members
func DirectUrlSettingsFitModePtr(in DirectUrlSettingsFitMode) *DirectUrlSettingsFitMode { return &in }

// DirectUrlSettingsFitModeFromPtr provides a way to convert back from an enum pointer to a
// guaranteed DirectUrlSettingsFitMode value of some kind (0 is default)
func DirectUrlSettingsFitModeFromPtr(in *DirectUrlSettingsFitMode) DirectUrlSettingsFitMode {
	if in == nil {
		return 0
	}
	return *in
}

// DirectUrlSettingsFitMode constant values
const (
	SCALE   = DirectUrlSettingsFitMode(pb.DirectUrlSettings_SCALE)
	STRETCH = DirectUrlSettingsFitMode(pb.DirectUrlSettings_STRETCH)
	CROP    = DirectUrlSettingsFitMode(pb.DirectUrlSettings_CROP)
	PAD     = DirectUrlSettingsFitMode(pb.DirectUrlSettings_PAD)
)
