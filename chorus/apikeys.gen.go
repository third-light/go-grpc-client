// This is a generated file, please do not edit by hand
package chorus

import (
	"context"

	pb "github.com/third-light/go-grpc-client/protobuf"
	"google.golang.org/grpc"
)

// ApiKeyDetails chorus API key details
type ApiKeyDetails struct {
	// Context the context of the API key
	Context string
	// CreatedDate the date/time that this object was created
	CreatedDate DateTime
	// Id unique ID of this object.
	Id string
	// Label an arbitrary chosen identifier for the API key
	Label string
	// ModifiedDate the date/time that this object was last changed
	ModifiedDate DateTime
}

// ApiKeyDetailsWithToken chorus API key details containing the token string, only returned
// on create operations
type ApiKeyDetailsWithToken struct {
	// ApiKey the key token
	ApiKey string
	// Context the context of the API key
	Context string
	// CreatedDate the date/time that this object was created
	CreatedDate DateTime
	// Id unique ID of this object.
	Id string
	// Label an arbitrary chosen identifier for the API key
	Label string
	// ModifiedDate the date/time that this object was last changed
	ModifiedDate DateTime
}

// ApiKeysRequests
type ApiKeysRequests struct {
}

// ApiKeysRequestsCreateProfileKey
type ApiKeysRequestsCreateProfileKey struct {
	// Details settings to use when creating the new API key
	Details CreateApiKeyDetails
}

// ApiKeysRequestsCreateSiteKey
type ApiKeysRequestsCreateSiteKey struct {
	// Details settings to use when creating the new API key
	Details CreateApiKeyDetails
}

// ApiKeysRequestsCreateSpacesKey
type ApiKeysRequestsCreateSpacesKey struct {
	// Details settings to use when creating the new API key
	Details CreateApiKeyDetails
	// SpaceId the space for which to create the API key
	SpaceId string
}

// ApiKeysRequestsDelete
type ApiKeysRequestsDelete struct {
	// ApiKeyId the ID of the API key to delete
	ApiKeyId string
}

// ApiKeysRequestsGet
type ApiKeysRequestsGet struct {
	// ApiKeyId iDs of the API keys to fetch details for
	ApiKeyId string
}

// ApiKeysRequestsGetMulti
type ApiKeysRequestsGetMulti struct {
	// ApiKeyId iDs of the API keys to fetch details for
	ApiKeyId []string
}

// ApiKeysRequestsGetProfileKeys
type ApiKeysRequestsGetProfileKeys struct {
}

// ApiKeysRequestsGetSiteKeys
type ApiKeysRequestsGetSiteKeys struct {
}

// ApiKeysRequestsGetSpacesKeys
type ApiKeysRequestsGetSpacesKeys struct {
	// SpaceId the space for which to get the API keys
	SpaceId string
}

// ApiKeysRequestsSet
type ApiKeysRequestsSet struct {
	// ApiKeyId ID of the API key to update
	ApiKeyId string
	// Details API key details to set
	Details SettableApiKeyDetails `fieldMask:"details"`
}

// ApiKeysResponses
type ApiKeysResponses struct {
}

// ApiKeysResponsesDelete
type ApiKeysResponsesDelete struct {
}

// ApiKeysResponsesGetMulti
type ApiKeysResponsesGetMulti struct {
	Response []ApiKeyDetails
}

// ApiKeysResponsesGetProfileKeys
type ApiKeysResponsesGetProfileKeys struct {
	Response []string
}

// ApiKeysResponsesGetSiteKeys
type ApiKeysResponsesGetSiteKeys struct {
	Response []string
}

// ApiKeysResponsesGetSpacesKeys
type ApiKeysResponsesGetSpacesKeys struct {
	Response []string
}

// CreateApiKeyDetails settings to use when creating a new API key
type CreateApiKeyDetails struct {
	// Label an arbitrary chosen identifier for the API key
	Label string
}

// SettableApiKeyDetails settings to use for updating an API key's details
type SettableApiKeyDetails struct {
	// Label an arbitrary chosen identifier for the API key
	Label *string `fieldMask:"label"`
}

// ApiKeysClient is the ApiKeys client
type ApiKeysClient struct {
	pb pb.ApiKeysClient
}

// NewApiKeysClient returns a new ApiKeysClient created from supplied gRPC ClientConn
func NewApiKeysClient(cc *grpc.ClientConn) *ApiKeysClient {
	return &ApiKeysClient{
		pb: pb.NewApiKeysClient(cc),
	}
}

// CreateProfileKey create a new profile API key with the supplied label
func (c *ApiKeysClient) CreateProfileKey(in *ApiKeysRequestsCreateProfileKey) (*ApiKeyDetailsWithToken, error) {
	return c.CreateProfileKeyWithContext(context.Background(), in)
}

// CreateProfileKeyWithContext create a new profile API key with the supplied label
func (c *ApiKeysClient) CreateProfileKeyWithContext(ctx context.Context, in *ApiKeysRequestsCreateProfileKey) (*ApiKeyDetailsWithToken, error) {
	out, err := c.pb.CreateProfileKey(ctx, convertFromApiKeysRequestsCreateProfileKeyPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToApiKeyDetailsWithTokenPtr(out), nil
}

// CreateSiteKey create a new site API key with the supplied label
func (c *ApiKeysClient) CreateSiteKey(in *ApiKeysRequestsCreateSiteKey) (*ApiKeyDetailsWithToken, error) {
	return c.CreateSiteKeyWithContext(context.Background(), in)
}

// CreateSiteKeyWithContext create a new site API key with the supplied label
func (c *ApiKeysClient) CreateSiteKeyWithContext(ctx context.Context, in *ApiKeysRequestsCreateSiteKey) (*ApiKeyDetailsWithToken, error) {
	out, err := c.pb.CreateSiteKey(ctx, convertFromApiKeysRequestsCreateSiteKeyPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToApiKeyDetailsWithTokenPtr(out), nil
}

// CreateSpacesKey create a new space API key with the supplied label
func (c *ApiKeysClient) CreateSpacesKey(in *ApiKeysRequestsCreateSpacesKey) (*ApiKeyDetailsWithToken, error) {
	return c.CreateSpacesKeyWithContext(context.Background(), in)
}

// CreateSpacesKeyWithContext create a new space API key with the supplied label
func (c *ApiKeysClient) CreateSpacesKeyWithContext(ctx context.Context, in *ApiKeysRequestsCreateSpacesKey) (*ApiKeyDetailsWithToken, error) {
	out, err := c.pb.CreateSpacesKey(ctx, convertFromApiKeysRequestsCreateSpacesKeyPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToApiKeyDetailsWithTokenPtr(out), nil
}

// Delete the specified API key
func (c *ApiKeysClient) Delete(in *ApiKeysRequestsDelete) (*ApiKeysResponsesDelete, error) {
	return c.DeleteWithContext(context.Background(), in)
}

// DeleteWithContext the specified API key
func (c *ApiKeysClient) DeleteWithContext(ctx context.Context, in *ApiKeysRequestsDelete) (*ApiKeysResponsesDelete, error) {
	out, err := c.pb.Delete(ctx, convertFromApiKeysRequestsDeletePtr(in))
	if err != nil {
		return nil, err
	}
	return convertToApiKeysResponsesDeletePtr(out), nil
}

// Get returns details about API keys
func (c *ApiKeysClient) Get(in *ApiKeysRequestsGet) (*ApiKeyDetails, error) {
	return c.GetWithContext(context.Background(), in)
}

// GetWithContext returns details about API keys
func (c *ApiKeysClient) GetWithContext(ctx context.Context, in *ApiKeysRequestsGet) (*ApiKeyDetails, error) {
	out, err := c.pb.Get(ctx, convertFromApiKeysRequestsGetPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToApiKeyDetailsPtr(out), nil
}

// GetMulti returns details about API keys
func (c *ApiKeysClient) GetMulti(in *ApiKeysRequestsGetMulti) (*ApiKeysResponsesGetMulti, error) {
	return c.GetMultiWithContext(context.Background(), in)
}

// GetMultiWithContext returns details about API keys
func (c *ApiKeysClient) GetMultiWithContext(ctx context.Context, in *ApiKeysRequestsGetMulti) (*ApiKeysResponsesGetMulti, error) {
	out, err := c.pb.GetMulti(ctx, convertFromApiKeysRequestsGetMultiPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToApiKeysResponsesGetMultiPtr(out), nil
}

// GetProfileKeys retrieve a list of API key IDs configured for this user
func (c *ApiKeysClient) GetProfileKeys(in *ApiKeysRequestsGetProfileKeys) (*ApiKeysResponsesGetProfileKeys, error) {
	return c.GetProfileKeysWithContext(context.Background(), in)
}

// GetProfileKeysWithContext retrieve a list of API key IDs configured for this user
func (c *ApiKeysClient) GetProfileKeysWithContext(ctx context.Context, in *ApiKeysRequestsGetProfileKeys) (*ApiKeysResponsesGetProfileKeys, error) {
	out, err := c.pb.GetProfileKeys(ctx, convertFromApiKeysRequestsGetProfileKeysPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToApiKeysResponsesGetProfileKeysPtr(out), nil
}

// GetSiteKeys retrieve a list of API key IDs configured for the site
func (c *ApiKeysClient) GetSiteKeys(in *ApiKeysRequestsGetSiteKeys) (*ApiKeysResponsesGetSiteKeys, error) {
	return c.GetSiteKeysWithContext(context.Background(), in)
}

// GetSiteKeysWithContext retrieve a list of API key IDs configured for the site
func (c *ApiKeysClient) GetSiteKeysWithContext(ctx context.Context, in *ApiKeysRequestsGetSiteKeys) (*ApiKeysResponsesGetSiteKeys, error) {
	out, err := c.pb.GetSiteKeys(ctx, convertFromApiKeysRequestsGetSiteKeysPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToApiKeysResponsesGetSiteKeysPtr(out), nil
}

// GetSpacesKeys retrieve a list of API key IDs configured for this user
func (c *ApiKeysClient) GetSpacesKeys(in *ApiKeysRequestsGetSpacesKeys) (*ApiKeysResponsesGetSpacesKeys, error) {
	return c.GetSpacesKeysWithContext(context.Background(), in)
}

// GetSpacesKeysWithContext retrieve a list of API key IDs configured for this user
func (c *ApiKeysClient) GetSpacesKeysWithContext(ctx context.Context, in *ApiKeysRequestsGetSpacesKeys) (*ApiKeysResponsesGetSpacesKeys, error) {
	out, err := c.pb.GetSpacesKeys(ctx, convertFromApiKeysRequestsGetSpacesKeysPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToApiKeysResponsesGetSpacesKeysPtr(out), nil
}

// Set update the details of an API key
func (c *ApiKeysClient) Set(in *ApiKeysRequestsSet) (*ApiKeyDetails, error) {
	return c.SetWithContext(context.Background(), in)
}

// SetWithContext update the details of an API key
func (c *ApiKeysClient) SetWithContext(ctx context.Context, in *ApiKeysRequestsSet) (*ApiKeyDetails, error) {
	out, err := c.pb.Set(ctx, convertFromApiKeysRequestsSetPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToApiKeyDetailsPtr(out), nil
}
