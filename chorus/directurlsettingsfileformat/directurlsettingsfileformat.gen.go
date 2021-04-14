// This is a generated file, please do not edit by hand
package directurlsettingsfileformat

import (
	pb "github.com/third-light/go-grpc-client/protobuf"
)

// DirectUrlSettingsFileFormat representation of the possible file formats that can be requested for a Direct URL.

type DirectUrlSettingsFileFormat int32

// DirectUrlSettingsFileFormatPtr provides a quick way to convert an enum constant to pointer
// for structs with optional members
func DirectUrlSettingsFileFormatPtr(in DirectUrlSettingsFileFormat) *DirectUrlSettingsFileFormat {
	return &in
}

// DirectUrlSettingsFileFormatFromPtr provides a way to convert back from an enum pointer to a
// guaranteed DirectUrlSettingsFileFormat value of some kind (0 is default)
func DirectUrlSettingsFileFormatFromPtr(in *DirectUrlSettingsFileFormat) DirectUrlSettingsFileFormat {
	if in == nil {
		return 0
	}
	return *in
}

// DirectUrlSettingsFileFormat constant values
const (
	ORIGINAL = DirectUrlSettingsFileFormat(pb.DirectUrlSettings_ORIGINAL)
	JPG      = DirectUrlSettingsFileFormat(pb.DirectUrlSettings_JPG)
	GIF      = DirectUrlSettingsFileFormat(pb.DirectUrlSettings_GIF)
	PNG      = DirectUrlSettingsFileFormat(pb.DirectUrlSettings_PNG)
	TIF      = DirectUrlSettingsFileFormat(pb.DirectUrlSettings_TIF)
	WEB      = DirectUrlSettingsFileFormat(pb.DirectUrlSettings_WEB)
)
