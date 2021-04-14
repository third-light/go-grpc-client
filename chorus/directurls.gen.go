// This is a generated file, please do not edit by hand
package chorus

import (
	"github.com/third-light/go-grpc-client/chorus/directurlsettingsfileformat"
	"github.com/third-light/go-grpc-client/chorus/directurlsettingsfitmode"
)

// DirectUrlSettings allows specification of the file returned via a Direct URL.
type DirectUrlSettings struct {
	// Blur sigma value for Gaussian blur
	Blur *int64
	// Crop details of a crop to apply
	Crop *DirectUrlSettingsCropSettings
	// Download if true, this changes the content-type to application/x-force-download and forces browsers to trigger a download.
	Download *bool
	// Dpi for output image
	Dpi *int64
	// Filename specify a filename to append to the URL (defaults to original file name)
	Filename *string
	// Fit the desired method of resolving aspect ratio differences between the requested input and output
	Fit *directurlsettingsfitmode.DirectUrlSettingsFitMode
	// Format code for requested file format
	Format *directurlsettingsfileformat.DirectUrlSettingsFileFormat
	// Height of output file in pixels
	Height *int64
	// Page the page number to download, should be zero-indexed (e.g. the first page is '0'). Except when downloading the original, this will default to '0'
	Page *int64
	// Quality a quality parameter for lossy formats like JPEG - accepted range 60-99, default is 85. Smaller parameters yield smaller filesize but lower fidelity
	Quality *int64
	// Rotate rotation angle in degrees
	Rotate *int64
	// Width of output file in pixels
	Width *int64
}

// DirectUrlSettingsCropSettings the settings to use for cropping a file.
type DirectUrlSettingsCropSettings struct {
	// Height height, in pixels, of the area to crop.
	Height int64
	// Width width, in pixels, of the area to crop.
	Width int64
	// X pixel on the `x` axis in which to start the crop.
	X int64
	// Y pixel on the `y` axis in which to start the crop.
	Y int64
}
