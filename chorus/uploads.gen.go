// This is a generated file, please do not edit by hand
package chorus

import (
	"context"

	"github.com/third-light/go-grpc-client/chorus/uploaddetailsstatus"
	pb "github.com/third-light/go-grpc-client/protobuf"
	"google.golang.org/grpc"
)

// UploadDetails details of a created upload session.
type UploadDetails struct {
	// CancelledCount number of files that have been cancelled before processing.
	CancelledCount int64
	// CreatedDate the date/time that this object was created
	CreatedDate DateTime
	// DestinationFolderId ID of the destination folder for the upload.
	DestinationFolderId string
	// FailedCount number of files that have been failed to process.
	FailedCount int64
	// Id unique ID of this object.
	Id string
	// IsAcceptingFiles indication of whether the upload is still accepting files. This will be `false` after 'finish' has been
	// called, or the upload has been cleaned up by the system.
	IsAcceptingFiles bool
	// ModifiedDate the date/time that this object was last changed
	ModifiedDate DateTime
	// PostUrl the URL to which files can be uploaded to using HTTP POST requests. The whole file should be posted raw.
	PostUrl string
	// Status the current status of this upload batch.
	Status uploaddetailsstatus.UploadDetailsStatus
	// SuccessCount number of files that have been successfully processed.
	SuccessCount int64
	// TotalCount total number of files in the upload batch.
	TotalCount int64
	// UploadedIds iDs of files that have been successfully uploaded.
	UploadedIds []string
}

// UploadsRequests
type UploadsRequests struct {
}

// UploadsRequestsCreate
type UploadsRequestsCreate struct {
	// DestinationFolderId the ID of the container in which to upload
	DestinationFolderId string
}

// UploadsRequestsFinish
type UploadsRequestsFinish struct {
	// UploadId ID of the Upload to finish.
	UploadId string
}

// UploadsRequestsGet
type UploadsRequestsGet struct {
	// UploadId ID of the Upload for which to get the details.
	UploadId string
}

// UploadsClient is the Uploads client
type UploadsClient struct {
	pb pb.UploadsClient
}

// NewUploadsClient returns a new UploadsClient created from supplied gRPC ClientConn
func NewUploadsClient(cc *grpc.ClientConn) *UploadsClient {
	return &UploadsClient{
		pb: pb.NewUploadsClient(cc),
	}
}

// Create an upload session to allow uploading of files into Chorus. Once all files have been transferred, then
// the upload should be finished using the `finish` method.
func (c *UploadsClient) Create(in *UploadsRequestsCreate) (*UploadDetails, error) {
	return c.CreateWithContext(context.Background(), in)
}

// CreateWithContext an upload session to allow uploading of files into Chorus. Once all files have been transferred, then
// the upload should be finished using the `finish` method.
func (c *UploadsClient) CreateWithContext(ctx context.Context, in *UploadsRequestsCreate) (*UploadDetails, error) {
	out, err := c.pb.Create(ctx, convertFromUploadsRequestsCreatePtr(in))
	if err != nil {
		return nil, err
	}
	return convertToUploadDetailsPtr(out), nil
}

// Finish mark an upload as finished.
func (c *UploadsClient) Finish(in *UploadsRequestsFinish) (*UploadDetails, error) {
	return c.FinishWithContext(context.Background(), in)
}

// FinishWithContext mark an upload as finished.
func (c *UploadsClient) FinishWithContext(ctx context.Context, in *UploadsRequestsFinish) (*UploadDetails, error) {
	out, err := c.pb.Finish(ctx, convertFromUploadsRequestsFinishPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToUploadDetailsPtr(out), nil
}

// Get details about an upload.
func (c *UploadsClient) Get(in *UploadsRequestsGet) (*UploadDetails, error) {
	return c.GetWithContext(context.Background(), in)
}

// GetWithContext details about an upload.
func (c *UploadsClient) GetWithContext(ctx context.Context, in *UploadsRequestsGet) (*UploadDetails, error) {
	out, err := c.pb.Get(ctx, convertFromUploadsRequestsGetPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToUploadDetailsPtr(out), nil
}
