// This is a generated file, please do not edit by hand
package uploaddetailsstatus

import (
	pb "github.com/third-light/go-grpc-client/protobuf"
)

// UploadDetailsStatus status of an upload.

type UploadDetailsStatus int32

// UploadDetailsStatusPtr provides a quick way to convert an enum constant to pointer
// for structs with optional members
func UploadDetailsStatusPtr(in UploadDetailsStatus) *UploadDetailsStatus { return &in }

// UploadDetailsStatusFromPtr provides a way to convert back from an enum pointer to a
// guaranteed UploadDetailsStatus value of some kind (0 is default)
func UploadDetailsStatusFromPtr(in *UploadDetailsStatus) UploadDetailsStatus {
	if in == nil {
		return 0
	}
	return *in
}

// UploadDetailsStatus constant values
const (
	Ready      = UploadDetailsStatus(pb.UploadDetails_Ready)
	Started    = UploadDetailsStatus(pb.UploadDetails_Started)
	Queued     = UploadDetailsStatus(pb.UploadDetails_Queued)
	Processing = UploadDetailsStatus(pb.UploadDetails_Processing)
	Finished   = UploadDetailsStatus(pb.UploadDetails_Finished)
)
