// This is a generated file, please do not edit by hand
package chorus

import (
	"context"

	pb "github.com/third-light/go-grpc-client/protobuf"
	"google.golang.org/grpc"
)

// CreateUserDetails settings to use when creating a user
type CreateUserDetails struct {
	// Description a description about the user, e.g. their job title or department.
	Description string
	// Email the email address of the user.
	Email string
	// HideHomeSpaceShortcut flag used to decide whether to show the user's 'Home Space' in the UI. This should be used if presenting
	// users with a similar layout to that seen in Chorus.
	HideHomeSpaceShortcut bool
	// Name the user's name.
	Name string
	// Username the "username" of the user.
	Username string
}

// SettableUserDetails settings to use for updating a user's details
type SettableUserDetails struct {
	// Description a description about the user, e.g. their job title or department.
	Description *string `fieldMask:"description"`
	// Email the email address of the user.
	Email *string `fieldMask:"email"`
	// HideHomeSpaceShortcut flag used to decide whether to show the user's 'Home Space' in the UI. This should be used if presenting
	// users with a similar layout to that seen in Chorus.
	HideHomeSpaceShortcut *bool `fieldMask:"hideHomeSpaceShortcut"`
	// Name the user's name.
	Name *string `fieldMask:"name"`
}

// UserDetails represents the details of a normal Chorus User.
type UserDetails struct {
	// AvatarUrl a URL to the current avatar
	AvatarUrl string
	// BackingFolderId ID of the folder that contains this user's contents. This will be null if requesting details for a
	// different user to the one currently logged in.
	BackingFolderId *string
	// CreatedDate the date/time that this object was created
	CreatedDate DateTime
	// Description a description about the user, e.g. their job title or department.
	Description string
	// Email the email address of the user. Will be null if the authenticated user cannot see this detail.
	Email *string
	// HideHomeSpaceShortcut flag used to decide whether to show the user's 'Home Space' in the UI. This should be used if presenting
	// users with a similar layout to that seen in Chorus. Will be null if requesting the details of another user.
	HideHomeSpaceShortcut *bool
	// HomeSpaceId ID of the user's Home Space. Null if the user does not have a Home Space, or if the current user
	// cannot see the Home Space.
	HomeSpaceId *string
	// Id unique ID of this object.
	Id string
	// ModifiedDate the date/time that this object was last changed
	ModifiedDate DateTime
	// Name the user's name.
	Name string
	// Username the "username" of the user.
	Username string
}

// UsersRequests
type UsersRequests struct {
}

// UsersRequestsCreateSiteUser
type UsersRequestsCreateSiteUser struct {
	// Details settings to use when creating the new user
	Details CreateUserDetails
}

// UsersRequestsCreateSpacesUser
type UsersRequestsCreateSpacesUser struct {
	// Details settings to use when creating the new user
	Details CreateUserDetails
	// SpaceId the space for which to create the API key
	SpaceId string
}

// UsersRequestsCurrent
type UsersRequestsCurrent struct {
}

// UsersRequestsDelete
type UsersRequestsDelete struct {
	// UserId ID of the user to delete
	UserId string
}

// UsersRequestsGet
type UsersRequestsGet struct {
	// UserId the ID of the user
	UserId string
}

// UsersRequestsGetMulti
type UsersRequestsGetMulti struct {
	// UserIds the IDs of the users
	UserIds []string
}

// UsersRequestsGetSiteUsers
type UsersRequestsGetSiteUsers struct {
}

// UsersRequestsGetSpacesUsers
type UsersRequestsGetSpacesUsers struct {
	// SpaceId the space for which to get the users
	SpaceId string
}

// UsersRequestsSet
type UsersRequestsSet struct {
	// Details user details to set
	Details SettableUserDetails `fieldMask:"details"`
	// UserId ID of the user to update
	UserId string
}

// UsersResponses
type UsersResponses struct {
}

// UsersResponsesDelete
type UsersResponsesDelete struct {
}

// UsersResponsesGetMulti
type UsersResponsesGetMulti struct {
	Response []UserDetails
}

// UsersResponsesGetSiteUsers
type UsersResponsesGetSiteUsers struct {
	Response []string
}

// UsersResponsesGetSpacesUsers
type UsersResponsesGetSpacesUsers struct {
	Response []string
}

// UsersClient is the Users client
type UsersClient struct {
	pb pb.UsersClient
}

// NewUsersClient returns a new UsersClient created from supplied gRPC ClientConn
func NewUsersClient(cc *grpc.ClientConn) *UsersClient {
	return &UsersClient{
		pb: pb.NewUsersClient(cc),
	}
}

// CreateSiteUser create a new site user
func (c *UsersClient) CreateSiteUser(in *UsersRequestsCreateSiteUser) (*UserDetails, error) {
	return c.CreateSiteUserWithContext(context.Background(), in)
}

// CreateSiteUserWithContext create a new site user
func (c *UsersClient) CreateSiteUserWithContext(ctx context.Context, in *UsersRequestsCreateSiteUser) (*UserDetails, error) {
	out, err := c.pb.CreateSiteUser(ctx, convertFromUsersRequestsCreateSiteUserPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToUserDetailsPtr(out), nil
}

// CreateSpacesUser create a new spaces user
func (c *UsersClient) CreateSpacesUser(in *UsersRequestsCreateSpacesUser) (*UserDetails, error) {
	return c.CreateSpacesUserWithContext(context.Background(), in)
}

// CreateSpacesUserWithContext create a new spaces user
func (c *UsersClient) CreateSpacesUserWithContext(ctx context.Context, in *UsersRequestsCreateSpacesUser) (*UserDetails, error) {
	out, err := c.pb.CreateSpacesUser(ctx, convertFromUsersRequestsCreateSpacesUserPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToUserDetailsPtr(out), nil
}

// Current returns details of the current user.
// Note: this will return an error if the current user is not a Normal User.
func (c *UsersClient) Current(in *UsersRequestsCurrent) (*UserDetails, error) {
	return c.CurrentWithContext(context.Background(), in)
}

// CurrentWithContext returns details of the current user.
// Note: this will return an error if the current user is not a Normal User.
func (c *UsersClient) CurrentWithContext(ctx context.Context, in *UsersRequestsCurrent) (*UserDetails, error) {
	out, err := c.pb.Current(ctx, convertFromUsersRequestsCurrentPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToUserDetailsPtr(out), nil
}

// Delete the specified user
func (c *UsersClient) Delete(in *UsersRequestsDelete) (*UsersResponsesDelete, error) {
	return c.DeleteWithContext(context.Background(), in)
}

// DeleteWithContext the specified user
func (c *UsersClient) DeleteWithContext(ctx context.Context, in *UsersRequestsDelete) (*UsersResponsesDelete, error) {
	out, err := c.pb.Delete(ctx, convertFromUsersRequestsDeletePtr(in))
	if err != nil {
		return nil, err
	}
	return convertToUsersResponsesDeletePtr(out), nil
}

// Get details about users.
func (c *UsersClient) Get(in *UsersRequestsGet) (*UserDetails, error) {
	return c.GetWithContext(context.Background(), in)
}

// GetWithContext details about users.
func (c *UsersClient) GetWithContext(ctx context.Context, in *UsersRequestsGet) (*UserDetails, error) {
	out, err := c.pb.Get(ctx, convertFromUsersRequestsGetPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToUserDetailsPtr(out), nil
}

// GetMulti get details about users.
func (c *UsersClient) GetMulti(in *UsersRequestsGetMulti) (*UsersResponsesGetMulti, error) {
	return c.GetMultiWithContext(context.Background(), in)
}

// GetMultiWithContext get details about users.
func (c *UsersClient) GetMultiWithContext(ctx context.Context, in *UsersRequestsGetMulti) (*UsersResponsesGetMulti, error) {
	out, err := c.pb.GetMulti(ctx, convertFromUsersRequestsGetMultiPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToUsersResponsesGetMultiPtr(out), nil
}

// GetSiteUsers retrieve a list of all user IDs in the system
func (c *UsersClient) GetSiteUsers(in *UsersRequestsGetSiteUsers) (*UsersResponsesGetSiteUsers, error) {
	return c.GetSiteUsersWithContext(context.Background(), in)
}

// GetSiteUsersWithContext retrieve a list of all user IDs in the system
func (c *UsersClient) GetSiteUsersWithContext(ctx context.Context, in *UsersRequestsGetSiteUsers) (*UsersResponsesGetSiteUsers, error) {
	out, err := c.pb.GetSiteUsers(ctx, convertFromUsersRequestsGetSiteUsersPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToUsersResponsesGetSiteUsersPtr(out), nil
}

// GetSpacesUsers retrieve a list of all user IDs in the provided space
func (c *UsersClient) GetSpacesUsers(in *UsersRequestsGetSpacesUsers) (*UsersResponsesGetSpacesUsers, error) {
	return c.GetSpacesUsersWithContext(context.Background(), in)
}

// GetSpacesUsersWithContext retrieve a list of all user IDs in the provided space
func (c *UsersClient) GetSpacesUsersWithContext(ctx context.Context, in *UsersRequestsGetSpacesUsers) (*UsersResponsesGetSpacesUsers, error) {
	out, err := c.pb.GetSpacesUsers(ctx, convertFromUsersRequestsGetSpacesUsersPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToUsersResponsesGetSpacesUsersPtr(out), nil
}

// Set update the details of an API key
func (c *UsersClient) Set(in *UsersRequestsSet) (*UserDetails, error) {
	return c.SetWithContext(context.Background(), in)
}

// SetWithContext update the details of an API key
func (c *UsersClient) SetWithContext(ctx context.Context, in *UsersRequestsSet) (*UserDetails, error) {
	out, err := c.pb.Set(ctx, convertFromUsersRequestsSetPtr(in))
	if err != nil {
		return nil, err
	}
	return convertToUserDetailsPtr(out), nil
}
