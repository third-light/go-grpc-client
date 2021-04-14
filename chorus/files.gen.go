// This is a generated file, please do not edit by hand
package chorus

import (
	"context"

	pb "github.com/third-light/go-grpc-client/protobuf"
	"google.golang.org/grpc"
)

// AudioDetails audio specific details about a file in Chorus.
type AudioDetails struct {
	// DurationSecs audio duration in seconds
	DurationSecs float64
}

// DocumentDetails document specific details about a file in Chorus.
type DocumentDetails struct {
	// Height document height
	Height int64
	// PageCount document page count
	PageCount int64
	// Width document width
	Width int64
}

// FilesRequests
type FilesRequests struct {
}

// FilesRequestsAttach
type FilesRequestsAttach struct {
	// AttachedFileIds the IDs of the File(s) to attach
	AttachedFileIds []string
	// FileId the ID of the File to which the attachment should be added
	FileId string
}

// FilesRequestsDeleteDirectUrl
type FilesRequestsDeleteDirectUrl struct {
	// FileId the ID of the file for which to revoke the Direct URLs.
	FileId string
}

// FilesRequestsDetach
type FilesRequestsDetach struct {
	// AttachedFileId the ID of the File to remove as an attachment
	AttachedFileId string
	// FileId the ID of the File from which the attachment should be removed
	FileId string
}

// FilesRequestsDetachAll
type FilesRequestsDetachAll struct {
	// FileId the file to update
	FileId string
}

// FilesRequestsGet
type FilesRequestsGet struct {
	// FileId the ID of the file
	FileId string
}

// FilesRequestsGetAttachments
type FilesRequestsGetAttachments struct {
	// FileId the ID of the File for which to list attachments
	FileId string
}

// FilesRequestsGetDirectUrl
type FilesRequestsGetDirectUrl struct {
	// FileId ID of the file for which to get the Direct URL.
	FileId string
	// Settings to control how the returned file looks. Supplying no
	// settings will return a link to the original file.
	Settings DirectUrlSettings
}

// FilesRequestsGetLink
type FilesRequestsGetLink struct {
	// LinkId the ID of the file link
	LinkId string
}

// FilesRequestsGetLinkMulti
type FilesRequestsGetLinkMulti struct {
	// LinkIds the IDs of the file links
	LinkIds []string
}

// FilesRequestsGetMetadata
type FilesRequestsGetMetadata struct {
	// FileId the files to get metadata for
	FileId string
}

// FilesRequestsGetMulti
type FilesRequestsGetMulti struct {
	// FileIds the IDs of the files
	FileIds []string
}

// FilesRequestsGetTemporaryDirectUrl
type FilesRequestsGetTemporaryDirectUrl struct {
	// FileId ID of the file for which to get the Direct URL
	FileId string
	// Settings to control how the returned file looks. Supplying no
	// settings will return a link to the original file.
	Settings DirectUrlSettings
}

// FilesResponses
type FilesResponses struct {
}

// FilesResponsesAttach
type FilesResponsesAttach struct {
	Response []FileDetails
}

// FilesResponsesDeleteDirectUrl
type FilesResponsesDeleteDirectUrl struct {
}

// FilesResponsesDetach
type FilesResponsesDetach struct {
	Response []FileDetails
}

// FilesResponsesDetachAll
type FilesResponsesDetachAll struct {
}

// FilesResponsesGetAttachments
type FilesResponsesGetAttachments struct {
	Response []FileDetails
}

// FilesResponsesGetDirectUrl
type FilesResponsesGetDirectUrl struct {
	Response string
}

// FilesResponsesGetLinkMulti
type FilesResponsesGetLinkMulti struct {
	Response []FileLinkDetails
}

// FilesResponsesGetMulti
type FilesResponsesGetMulti struct {
	Response []FileDetails
}

// FilesResponsesGetTemporaryDirectUrl
type FilesResponsesGetTemporaryDirectUrl struct {
	Response string
}

// GenericDetails generic details about a file in Chorus.
type GenericDetails struct {
}

// ImageDetails image specific details about a file in Chorus.
type ImageDetails struct {
	// Height image height in pixels
	Height int64
	// Width image width in pixels
	Width int64
}

// VideoDetails video specific details about a file in Chorus.
type VideoDetails struct {
	// DurationSecs video duration in seconds
	DurationSecs float64
	// Framerate video framerate in frames per second
	Framerate float64
	// Height video height in pixels
	Height int64
	// Width video width in pixels
	Width int64
}

// MediaDetails one of multiple media specific details of a file in Chorus.
type MediaDetails interface {
	getProto() *pb.MediaDetails
}

func (item AudioDetails) getProto() *pb.MediaDetails {
	return &pb.MediaDetails{
		Details: &pb.MediaDetails_Audio{
			Audio: convertFromAudioDetails(item),
		},
	}
}

func (item DocumentDetails) getProto() *pb.MediaDetails {
	return &pb.MediaDetails{
		Details: &pb.MediaDetails_Document{
			Document: convertFromDocumentDetails(item),
		},
	}
}

func (item GenericDetails) getProto() *pb.MediaDetails {
	return &pb.MediaDetails{
		Details: &pb.MediaDetails_Generic{
			Generic: convertFromGenericDetails(item),
		},
	}
}

func (item ImageDetails) getProto() *pb.MediaDetails {
	return &pb.MediaDetails{
		Details: &pb.MediaDetails_Image{
			Image: convertFromImageDetails(item),
		},
	}
}

func (item VideoDetails) getProto() *pb.MediaDetails {
	return &pb.MediaDetails{
		Details: &pb.MediaDetails_Video{
			Video: convertFromVideoDetails(item),
		},
	}
}

// ThumbDetails details about a thumbnail for a file.
type ThumbDetails struct {
	// Height in pixels
	Height int64
	// Url the URL to access the thumbnail
	Url string
	// Width in pixels
	Width int64
}

// Thumbnails a collection of available thumbnail sizes/types for a file.
type Thumbnails struct {
	Large  *ThumbDetails
	Medium *ThumbDetails
	Small  *ThumbDetails
}

// MetadataValueMap a map matching a metadata value to its values.
type MetadataValueMap struct {
	Values map[string]MetadataValue
}

// FilesClient is the Files client
type FilesClient struct {
	pb pb.FilesClient
}

// NewFilesClient returns a new FilesClient created from supplied gRPC ClientConn
func NewFilesClient(cc *grpc.ClientConn) *FilesClient {
	return &FilesClient{
		pb: pb.NewFilesClient(cc),
	}
}

// Attach one or more files to another file.
func (c *FilesClient) Attach(in *FilesRequestsAttach) (*FilesResponsesAttach, error) {
	return c.AttachWithContext(context.Background(), in)
}

// AttachWithContext one or more files to another file.
func (c *FilesClient) AttachWithContext(ctx context.Context, in *FilesRequestsAttach) (*FilesResponsesAttach, error) {
	out, err := c.pb.Attach(ctx, convertFromFilesRequestsAttachPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToFilesResponsesAttachPtr(out), nil
}

// DeleteDirectUrl delete any existing Direct URL (temporary or permanent) for the given file.
func (c *FilesClient) DeleteDirectUrl(in *FilesRequestsDeleteDirectUrl) (*FilesResponsesDeleteDirectUrl, error) {
	return c.DeleteDirectUrlWithContext(context.Background(), in)
}

// DeleteDirectUrlWithContext delete any existing Direct URL (temporary or permanent) for the given file.
func (c *FilesClient) DeleteDirectUrlWithContext(ctx context.Context, in *FilesRequestsDeleteDirectUrl) (*FilesResponsesDeleteDirectUrl, error) {
	out, err := c.pb.DeleteDirectUrl(ctx, convertFromFilesRequestsDeleteDirectUrlPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToFilesResponsesDeleteDirectUrlPtr(out), nil
}

// Detach break the attachment between two files.
func (c *FilesClient) Detach(in *FilesRequestsDetach) (*FilesResponsesDetach, error) {
	return c.DetachWithContext(context.Background(), in)
}

// DetachWithContext break the attachment between two files.
func (c *FilesClient) DetachWithContext(ctx context.Context, in *FilesRequestsDetach) (*FilesResponsesDetach, error) {
	out, err := c.pb.Detach(ctx, convertFromFilesRequestsDetachPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToFilesResponsesDetachPtr(out), nil
}

// DetachAll remove all attached files from an asset.
func (c *FilesClient) DetachAll(in *FilesRequestsDetachAll) (*FilesResponsesDetachAll, error) {
	return c.DetachAllWithContext(context.Background(), in)
}

// DetachAllWithContext remove all attached files from an asset.
func (c *FilesClient) DetachAllWithContext(ctx context.Context, in *FilesRequestsDetachAll) (*FilesResponsesDetachAll, error) {
	out, err := c.pb.DetachAll(ctx, convertFromFilesRequestsDetachAllPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToFilesResponsesDetachAllPtr(out), nil
}

// Get retrieve details for a particular file.
func (c *FilesClient) Get(in *FilesRequestsGet) (*FileDetails, error) {
	return c.GetWithContext(context.Background(), in)
}

// GetWithContext retrieve details for a particular file.
func (c *FilesClient) GetWithContext(ctx context.Context, in *FilesRequestsGet) (*FileDetails, error) {
	out, err := c.pb.Get(ctx, convertFromFilesRequestsGetPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToFileDetailsPtr(out), nil
}

// GetAttachments get the files attached to a given one
func (c *FilesClient) GetAttachments(in *FilesRequestsGetAttachments) (*FilesResponsesGetAttachments, error) {
	return c.GetAttachmentsWithContext(context.Background(), in)
}

// GetAttachmentsWithContext get the files attached to a given one
func (c *FilesClient) GetAttachmentsWithContext(ctx context.Context, in *FilesRequestsGetAttachments) (*FilesResponsesGetAttachments, error) {
	out, err := c.pb.GetAttachments(ctx, convertFromFilesRequestsGetAttachmentsPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToFilesResponsesGetAttachmentsPtr(out), nil
}

// GetDirectUrl get a permanent Direct URL for the given file.
func (c *FilesClient) GetDirectUrl(in *FilesRequestsGetDirectUrl) (*FilesResponsesGetDirectUrl, error) {
	return c.GetDirectUrlWithContext(context.Background(), in)
}

// GetDirectUrlWithContext get a permanent Direct URL for the given file.
func (c *FilesClient) GetDirectUrlWithContext(ctx context.Context, in *FilesRequestsGetDirectUrl) (*FilesResponsesGetDirectUrl, error) {
	out, err := c.pb.GetDirectUrl(ctx, convertFromFilesRequestsGetDirectUrlPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToFilesResponsesGetDirectUrlPtr(out), nil
}

// GetLink retrieve details for a particular file link.
func (c *FilesClient) GetLink(in *FilesRequestsGetLink) (*FileLinkDetails, error) {
	return c.GetLinkWithContext(context.Background(), in)
}

// GetLinkWithContext retrieve details for a particular file link.
func (c *FilesClient) GetLinkWithContext(ctx context.Context, in *FilesRequestsGetLink) (*FileLinkDetails, error) {
	out, err := c.pb.GetLink(ctx, convertFromFilesRequestsGetLinkPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToFileLinkDetailsPtr(out), nil
}

// GetLinkMulti retrieve details for a particular file link.
func (c *FilesClient) GetLinkMulti(in *FilesRequestsGetLinkMulti) (*FilesResponsesGetLinkMulti, error) {
	return c.GetLinkMultiWithContext(context.Background(), in)
}

// GetLinkMultiWithContext retrieve details for a particular file link.
func (c *FilesClient) GetLinkMultiWithContext(ctx context.Context, in *FilesRequestsGetLinkMulti) (*FilesResponsesGetLinkMulti, error) {
	out, err := c.pb.GetLinkMulti(ctx, convertFromFilesRequestsGetLinkMultiPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToFilesResponsesGetLinkMultiPtr(out), nil
}

// GetMetadata gets the metadata for a file.
func (c *FilesClient) GetMetadata(in *FilesRequestsGetMetadata) (*MetadataValueMap, error) {
	return c.GetMetadataWithContext(context.Background(), in)
}

// GetMetadataWithContext gets the metadata for a file.
func (c *FilesClient) GetMetadataWithContext(ctx context.Context, in *FilesRequestsGetMetadata) (*MetadataValueMap, error) {
	out, err := c.pb.GetMetadata(ctx, convertFromFilesRequestsGetMetadataPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToMetadataValueMapPtr(out), nil
}

// GetMulti retrieve details for a particular file.
func (c *FilesClient) GetMulti(in *FilesRequestsGetMulti) (*FilesResponsesGetMulti, error) {
	return c.GetMultiWithContext(context.Background(), in)
}

// GetMultiWithContext retrieve details for a particular file.
func (c *FilesClient) GetMultiWithContext(ctx context.Context, in *FilesRequestsGetMulti) (*FilesResponsesGetMulti, error) {
	out, err := c.pb.GetMulti(ctx, convertFromFilesRequestsGetMultiPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToFilesResponsesGetMultiPtr(out), nil
}

// GetTemporaryDirectUrl get a temporary Direct URL for the given file.
func (c *FilesClient) GetTemporaryDirectUrl(in *FilesRequestsGetTemporaryDirectUrl) (*FilesResponsesGetTemporaryDirectUrl, error) {
	return c.GetTemporaryDirectUrlWithContext(context.Background(), in)
}

// GetTemporaryDirectUrlWithContext get a temporary Direct URL for the given file.
func (c *FilesClient) GetTemporaryDirectUrlWithContext(ctx context.Context, in *FilesRequestsGetTemporaryDirectUrl) (*FilesResponsesGetTemporaryDirectUrl, error) {
	out, err := c.pb.GetTemporaryDirectUrl(ctx, convertFromFilesRequestsGetTemporaryDirectUrlPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToFilesResponsesGetTemporaryDirectUrlPtr(out), nil
}
