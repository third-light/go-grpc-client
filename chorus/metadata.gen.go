// This is a generated file, please do not edit by hand
package chorus

import (
	"context"

	"github.com/third-light/go-grpc-client/chorus/vocabmode"
	pb "github.com/third-light/go-grpc-client/protobuf"
	"google.golang.org/grpc"
)

// Field represents the details of a metadata field.
type Field struct {
	// DataType the data type of the metadata field
	DataType string
	// Description a description of the field
	Description string
	// DisplayType the display type of the field
	DisplayType string
	// ReadOnly whether the field is read only
	ReadOnly bool
	// SearchSettings the search settings for the metadata field
	SearchSettings FieldSearchSettings
	// SupportedFileTypes types of file supported by this metadata field
	SupportedFileTypes []string
	// TagName the tag name
	TagName string
}

// FieldSearchSettings represents the search details of a metadata field.
type FieldSearchSettings struct {
	// AdvancedSearch whether the field is available in advanced search
	AdvancedSearch bool
	// GeneralSearch whether the field is available in general search
	GeneralSearch bool
}

// MetadataRequests
type MetadataRequests struct {
}

// MetadataRequestsDeleteMetadataValues
type MetadataRequestsDeleteMetadataValues struct {
	// Details the value details to delete
	Details MetadataValue
	// ItemId the file or folder to delete the metadata for
	ItemId string
	// TagName the name of the tag for which to delete the metadata
	TagName string
}

// MetadataRequestsGetField
type MetadataRequestsGetField struct {
	// TagName the id of the field of which to get the details
	TagName string
}

// MetadataRequestsGetFieldMulti
type MetadataRequestsGetFieldMulti struct {
	// TagNames the id of the field of which to get the details
	TagNames []string
}

// MetadataRequestsGetMetadataFields
type MetadataRequestsGetMetadataFields struct {
	// ItemId the item to get metadata fields for
	ItemId string
}

// MetadataRequestsGetMetadataPanels
type MetadataRequestsGetMetadataPanels struct {
	// ItemId the item to get metadata panel IDs for
	ItemId string
}

// MetadataRequestsGetMetadataValues
type MetadataRequestsGetMetadataValues struct {
	// ItemId the file or folder to get metadata for
	ItemId string
	// TagName the name of the tag for which to retrieve the values
	TagName string
}

// MetadataRequestsGetPanel
type MetadataRequestsGetPanel struct {
	// PanelId the id of the panel for which to get the details
	PanelId string
}

// MetadataRequestsGetPanelMulti
type MetadataRequestsGetPanelMulti struct {
	// PanelIds the id of the panel for which to get the details
	PanelIds []string
}

// MetadataRequestsGetProfileVocab
type MetadataRequestsGetProfileVocab struct {
	// TagName the name of the tag for which to retrieve the vocab
	TagName string
}

// MetadataRequestsGetSiteVocab
type MetadataRequestsGetSiteVocab struct {
	// TagName the name of the tag for which to retrieve the vocab
	TagName string
}

// MetadataRequestsGetSpaceVocab
type MetadataRequestsGetSpaceVocab struct {
	// SpaceId the space for which to get the tag vocab
	SpaceId string
	// TagName the name of the tag for which to retrieve the vocab
	TagName string
}

// MetadataRequestsReplaceMetadataValues
type MetadataRequestsReplaceMetadataValues struct {
	// Details the value details to set
	Details MetadataValue
	// ItemId the file or folder to replace the metadata for
	ItemId string
	// TagName the name of the tag for which to update the metadata
	TagName string
}

// MetadataRequestsReplaceProfileVocab
type MetadataRequestsReplaceProfileVocab struct {
	// Details the vocab details to set
	Details Vocab
	// TagName the name of the tag for which to replace the vocab
	TagName string
}

// MetadataRequestsReplaceSiteVocab
type MetadataRequestsReplaceSiteVocab struct {
	// Details the vocab details to set
	Details Vocab
	// TagName the name of the tag for which to replace the vocab
	TagName string
}

// MetadataRequestsReplaceSpaceVocab
type MetadataRequestsReplaceSpaceVocab struct {
	// Details the vocab details to set
	Details Vocab
	// SpaceId the space for which to replace the tag vocab
	SpaceId string
	// TagName the name of the tag for which to replace the vocab
	TagName string
}

// MetadataRequestsUpdateMetadataValues
type MetadataRequestsUpdateMetadataValues struct {
	// Details the value details to set
	Details MetadataValue
	// ItemId the file or folder to update the metadata for
	ItemId string
	// TagName the name of the tag for which to update the metadata
	TagName string
}

// MetadataRequestsUpdateProfileVocab
type MetadataRequestsUpdateProfileVocab struct {
	// Details the vocab details to set
	Details Vocab
	// TagName the name of the tag for which to replace the vocab
	TagName string
}

// MetadataRequestsUpdateSiteVocab
type MetadataRequestsUpdateSiteVocab struct {
	// Details the vocab details to set
	Details Vocab
	// TagName the name of the tag for which to replace the vocab
	TagName string
}

// MetadataRequestsUpdateSpaceVocab
type MetadataRequestsUpdateSpaceVocab struct {
	// Details the vocab details to set
	Details Vocab
	// SpaceId the space for which to update the tag vocab
	SpaceId string
	// TagName the name of the tag for which to replace the vocab
	TagName string
}

// MetadataResponses
type MetadataResponses struct {
}

// MetadataResponsesGetFieldMulti
type MetadataResponsesGetFieldMulti struct {
	Response []Field
}

// MetadataResponsesGetMetadataFields
type MetadataResponsesGetMetadataFields struct {
	Response []string
}

// MetadataResponsesGetMetadataPanels
type MetadataResponsesGetMetadataPanels struct {
	Response []string
}

// MetadataResponsesGetPanelMulti
type MetadataResponsesGetPanelMulti struct {
	Response []Panel
}

// MetadataValue represents the value of a metadata field.
type MetadataValue interface {
	getProto() *pb.MetadataValue
}
type MetadataValueNumericType []int64

func (item MetadataValueNumericType) getProto() *pb.MetadataValue {
	return &pb.MetadataValue{
		Values: &pb.MetadataValue_Numeric{
			Numeric: &pb.MetadataValue_NumericType{
				Values: []int64(item),
			},
		},
	}
}

type MetadataValueTextType []string

func (item MetadataValueTextType) getProto() *pb.MetadataValue {
	return &pb.MetadataValue{
		Values: &pb.MetadataValue_Text{
			Text: &pb.MetadataValue_TextType{
				Values: []string(item),
			},
		},
	}
}

// Panel represents the details of a metadata panel.
type Panel struct {
	// Description a description of the panel
	Description string
	// PanelId the id of the panel
	PanelId string
	// TagNames names of the tags in this metadata panel
	TagNames []string
}

// Vocab represents the details of a metadata panel.
type Vocab struct {
	// Mode the mode of the vocab
	Mode vocabmode.VocabMode
	// TagName the name of the tag
	TagName string
	// Values names of the values in the vocab list
	Values []string
}

// MetadataClient is the Metadata client
type MetadataClient struct {
	pb pb.MetadataClient
}

// NewMetadataClient returns a new MetadataClient created from supplied gRPC ClientConn
func NewMetadataClient(cc *grpc.ClientConn) *MetadataClient {
	return &MetadataClient{
		pb: pb.NewMetadataClient(cc),
	}
}

// DeleteMetadataValues delete the metadata for the tag for the item.
func (c *MetadataClient) DeleteMetadataValues(in *MetadataRequestsDeleteMetadataValues) (MetadataValue, error) {
	return c.DeleteMetadataValuesWithContext(context.Background(), in)
}

// DeleteMetadataValuesWithContext delete the metadata for the tag for the item.
func (c *MetadataClient) DeleteMetadataValuesWithContext(ctx context.Context, in *MetadataRequestsDeleteMetadataValues) (MetadataValue, error) {
	out, err := c.pb.DeleteMetadataValues(ctx, convertFromMetadataRequestsDeleteMetadataValuesPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToMetadataValue(out), nil
}

// GetField get the details of the metadata field with the given ID
func (c *MetadataClient) GetField(in *MetadataRequestsGetField) (*Field, error) {
	return c.GetFieldWithContext(context.Background(), in)
}

// GetFieldWithContext get the details of the metadata field with the given ID
func (c *MetadataClient) GetFieldWithContext(ctx context.Context, in *MetadataRequestsGetField) (*Field, error) {
	out, err := c.pb.GetField(ctx, convertFromMetadataRequestsGetFieldPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToFieldPtr(out), nil
}

// GetFieldMulti get the details of the metadata field with the given ID
func (c *MetadataClient) GetFieldMulti(in *MetadataRequestsGetFieldMulti) (*MetadataResponsesGetFieldMulti, error) {
	return c.GetFieldMultiWithContext(context.Background(), in)
}

// GetFieldMultiWithContext get the details of the metadata field with the given ID
func (c *MetadataClient) GetFieldMultiWithContext(ctx context.Context, in *MetadataRequestsGetFieldMulti) (*MetadataResponsesGetFieldMulti, error) {
	out, err := c.pb.GetFieldMulti(ctx, convertFromMetadataRequestsGetFieldMultiPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToMetadataResponsesGetFieldMultiPtr(out), nil
}

// GetMetadataFields get the tags names of the metadata fields available for the item
func (c *MetadataClient) GetMetadataFields(in *MetadataRequestsGetMetadataFields) (*MetadataResponsesGetMetadataFields, error) {
	return c.GetMetadataFieldsWithContext(context.Background(), in)
}

// GetMetadataFieldsWithContext get the tags names of the metadata fields available for the item
func (c *MetadataClient) GetMetadataFieldsWithContext(ctx context.Context, in *MetadataRequestsGetMetadataFields) (*MetadataResponsesGetMetadataFields, error) {
	out, err := c.pb.GetMetadataFields(ctx, convertFromMetadataRequestsGetMetadataFieldsPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToMetadataResponsesGetMetadataFieldsPtr(out), nil
}

// GetMetadataPanels get the IDs of the metadata panels available for the item
func (c *MetadataClient) GetMetadataPanels(in *MetadataRequestsGetMetadataPanels) (*MetadataResponsesGetMetadataPanels, error) {
	return c.GetMetadataPanelsWithContext(context.Background(), in)
}

// GetMetadataPanelsWithContext get the IDs of the metadata panels available for the item
func (c *MetadataClient) GetMetadataPanelsWithContext(ctx context.Context, in *MetadataRequestsGetMetadataPanels) (*MetadataResponsesGetMetadataPanels, error) {
	out, err := c.pb.GetMetadataPanels(ctx, convertFromMetadataRequestsGetMetadataPanelsPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToMetadataResponsesGetMetadataPanelsPtr(out), nil
}

// GetMetadataValues get the details of the metadata for the tag for the item
func (c *MetadataClient) GetMetadataValues(in *MetadataRequestsGetMetadataValues) (MetadataValue, error) {
	return c.GetMetadataValuesWithContext(context.Background(), in)
}

// GetMetadataValuesWithContext get the details of the metadata for the tag for the item
func (c *MetadataClient) GetMetadataValuesWithContext(ctx context.Context, in *MetadataRequestsGetMetadataValues) (MetadataValue, error) {
	out, err := c.pb.GetMetadataValues(ctx, convertFromMetadataRequestsGetMetadataValuesPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToMetadataValue(out), nil
}

// GetPanel get the details of the metadata panel with the given ID
func (c *MetadataClient) GetPanel(in *MetadataRequestsGetPanel) (*Panel, error) {
	return c.GetPanelWithContext(context.Background(), in)
}

// GetPanelWithContext get the details of the metadata panel with the given ID
func (c *MetadataClient) GetPanelWithContext(ctx context.Context, in *MetadataRequestsGetPanel) (*Panel, error) {
	out, err := c.pb.GetPanel(ctx, convertFromMetadataRequestsGetPanelPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToPanelPtr(out), nil
}

// GetPanelMulti get the details of the metadata panel with the given ID
func (c *MetadataClient) GetPanelMulti(in *MetadataRequestsGetPanelMulti) (*MetadataResponsesGetPanelMulti, error) {
	return c.GetPanelMultiWithContext(context.Background(), in)
}

// GetPanelMultiWithContext get the details of the metadata panel with the given ID
func (c *MetadataClient) GetPanelMultiWithContext(ctx context.Context, in *MetadataRequestsGetPanelMulti) (*MetadataResponsesGetPanelMulti, error) {
	out, err := c.pb.GetPanelMulti(ctx, convertFromMetadataRequestsGetPanelMultiPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToMetadataResponsesGetPanelMultiPtr(out), nil
}

// GetProfileVocab get the details of the vocab for the metadata tag for the current profile
func (c *MetadataClient) GetProfileVocab(in *MetadataRequestsGetProfileVocab) (*Vocab, error) {
	return c.GetProfileVocabWithContext(context.Background(), in)
}

// GetProfileVocabWithContext get the details of the vocab for the metadata tag for the current profile
func (c *MetadataClient) GetProfileVocabWithContext(ctx context.Context, in *MetadataRequestsGetProfileVocab) (*Vocab, error) {
	out, err := c.pb.GetProfileVocab(ctx, convertFromMetadataRequestsGetProfileVocabPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToVocabPtr(out), nil
}

// GetSiteVocab get the details of the vocab for the metadata tag for the site
func (c *MetadataClient) GetSiteVocab(in *MetadataRequestsGetSiteVocab) (*Vocab, error) {
	return c.GetSiteVocabWithContext(context.Background(), in)
}

// GetSiteVocabWithContext get the details of the vocab for the metadata tag for the site
func (c *MetadataClient) GetSiteVocabWithContext(ctx context.Context, in *MetadataRequestsGetSiteVocab) (*Vocab, error) {
	out, err := c.pb.GetSiteVocab(ctx, convertFromMetadataRequestsGetSiteVocabPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToVocabPtr(out), nil
}

// GetSpaceVocab get the details of the vocab for the metadata tag for the space
func (c *MetadataClient) GetSpaceVocab(in *MetadataRequestsGetSpaceVocab) (*Vocab, error) {
	return c.GetSpaceVocabWithContext(context.Background(), in)
}

// GetSpaceVocabWithContext get the details of the vocab for the metadata tag for the space
func (c *MetadataClient) GetSpaceVocabWithContext(ctx context.Context, in *MetadataRequestsGetSpaceVocab) (*Vocab, error) {
	out, err := c.pb.GetSpaceVocab(ctx, convertFromMetadataRequestsGetSpaceVocabPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToVocabPtr(out), nil
}

// ReplaceMetadataValues replace the metadata for the metadata tag for the item
func (c *MetadataClient) ReplaceMetadataValues(in *MetadataRequestsReplaceMetadataValues) (MetadataValue, error) {
	return c.ReplaceMetadataValuesWithContext(context.Background(), in)
}

// ReplaceMetadataValuesWithContext replace the metadata for the metadata tag for the item
func (c *MetadataClient) ReplaceMetadataValuesWithContext(ctx context.Context, in *MetadataRequestsReplaceMetadataValues) (MetadataValue, error) {
	out, err := c.pb.ReplaceMetadataValues(ctx, convertFromMetadataRequestsReplaceMetadataValuesPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToMetadataValue(out), nil
}

// ReplaceProfileVocab amend the details of the vocab for user profile for the given metadata.
func (c *MetadataClient) ReplaceProfileVocab(in *MetadataRequestsReplaceProfileVocab) (*Vocab, error) {
	return c.ReplaceProfileVocabWithContext(context.Background(), in)
}

// ReplaceProfileVocabWithContext amend the details of the vocab for user profile for the given metadata.
func (c *MetadataClient) ReplaceProfileVocabWithContext(ctx context.Context, in *MetadataRequestsReplaceProfileVocab) (*Vocab, error) {
	out, err := c.pb.ReplaceProfileVocab(ctx, convertFromMetadataRequestsReplaceProfileVocabPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToVocabPtr(out), nil
}

// ReplaceSiteVocab amend the details of the vocab for the given metadata for the site.
func (c *MetadataClient) ReplaceSiteVocab(in *MetadataRequestsReplaceSiteVocab) (*Vocab, error) {
	return c.ReplaceSiteVocabWithContext(context.Background(), in)
}

// ReplaceSiteVocabWithContext amend the details of the vocab for the given metadata for the site.
func (c *MetadataClient) ReplaceSiteVocabWithContext(ctx context.Context, in *MetadataRequestsReplaceSiteVocab) (*Vocab, error) {
	out, err := c.pb.ReplaceSiteVocab(ctx, convertFromMetadataRequestsReplaceSiteVocabPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToVocabPtr(out), nil
}

// ReplaceSpaceVocab amend the details of the vocab for the given metadata.
func (c *MetadataClient) ReplaceSpaceVocab(in *MetadataRequestsReplaceSpaceVocab) (*Vocab, error) {
	return c.ReplaceSpaceVocabWithContext(context.Background(), in)
}

// ReplaceSpaceVocabWithContext amend the details of the vocab for the given metadata.
func (c *MetadataClient) ReplaceSpaceVocabWithContext(ctx context.Context, in *MetadataRequestsReplaceSpaceVocab) (*Vocab, error) {
	out, err := c.pb.ReplaceSpaceVocab(ctx, convertFromMetadataRequestsReplaceSpaceVocabPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToVocabPtr(out), nil
}

// UpdateMetadataValues update the details of the metadata for the tag for the item. If called on a single value field, it will
// replace the metadata, otherwise it will add it.
func (c *MetadataClient) UpdateMetadataValues(in *MetadataRequestsUpdateMetadataValues) (MetadataValue, error) {
	return c.UpdateMetadataValuesWithContext(context.Background(), in)
}

// UpdateMetadataValuesWithContext update the details of the metadata for the tag for the item. If called on a single value field, it will
// replace the metadata, otherwise it will add it.
func (c *MetadataClient) UpdateMetadataValuesWithContext(ctx context.Context, in *MetadataRequestsUpdateMetadataValues) (MetadataValue, error) {
	out, err := c.pb.UpdateMetadataValues(ctx, convertFromMetadataRequestsUpdateMetadataValuesPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToMetadataValue(out), nil
}

// UpdateProfileVocab amend the details of the vocab for the metadata in the profile of the current user. Existing vocab values
// which are not specified are not removed.
func (c *MetadataClient) UpdateProfileVocab(in *MetadataRequestsUpdateProfileVocab) (*Vocab, error) {
	return c.UpdateProfileVocabWithContext(context.Background(), in)
}

// UpdateProfileVocabWithContext amend the details of the vocab for the metadata in the profile of the current user. Existing vocab values
// which are not specified are not removed.
func (c *MetadataClient) UpdateProfileVocabWithContext(ctx context.Context, in *MetadataRequestsUpdateProfileVocab) (*Vocab, error) {
	out, err := c.pb.UpdateProfileVocab(ctx, convertFromMetadataRequestsUpdateProfileVocabPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToVocabPtr(out), nil
}

// UpdateSiteVocab amend the details of the vocab for the site metadata. Existing vocab values
// which are not specified are not removed.
func (c *MetadataClient) UpdateSiteVocab(in *MetadataRequestsUpdateSiteVocab) (*Vocab, error) {
	return c.UpdateSiteVocabWithContext(context.Background(), in)
}

// UpdateSiteVocabWithContext amend the details of the vocab for the site metadata. Existing vocab values
// which are not specified are not removed.
func (c *MetadataClient) UpdateSiteVocabWithContext(ctx context.Context, in *MetadataRequestsUpdateSiteVocab) (*Vocab, error) {
	out, err := c.pb.UpdateSiteVocab(ctx, convertFromMetadataRequestsUpdateSiteVocabPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToVocabPtr(out), nil
}

// UpdateSpaceVocab amend the details of the vocab for the given metadata. Existing vocab values
// which are not specified are not removed.
func (c *MetadataClient) UpdateSpaceVocab(in *MetadataRequestsUpdateSpaceVocab) (*Vocab, error) {
	return c.UpdateSpaceVocabWithContext(context.Background(), in)
}

// UpdateSpaceVocabWithContext amend the details of the vocab for the given metadata. Existing vocab values
// which are not specified are not removed.
func (c *MetadataClient) UpdateSpaceVocabWithContext(ctx context.Context, in *MetadataRequestsUpdateSpaceVocab) (*Vocab, error) {
	out, err := c.pb.UpdateSpaceVocab(ctx, convertFromMetadataRequestsUpdateSpaceVocabPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToVocabPtr(out), nil
}
