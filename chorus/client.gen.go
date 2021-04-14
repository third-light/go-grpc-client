// This is a generated file, please do not edit by hand
package chorus

import (
	"google.golang.org/grpc"
)

// Client is the Chorus super-client, simplifying setup and access to individual module clients
type Client struct {
	cc                *grpc.ClientConn
	apiKeysClient     *ApiKeysClient
	authClient        *AuthClient
	binClient         *BinClient
	contentClient     *ContentClient
	contextsClient    *ContextsClient
	coreClient        *CoreClient
	filesClient       *FilesClient
	foldersClient     *FoldersClient
	metadataClient    *MetadataClient
	sharedLinksClient *SharedLinksClient
	spacesClient      *SpacesClient
	uploadsClient     *UploadsClient
	usersClient       *UsersClient
}

// NewClient returns a new Chorus client
func NewClient(cc *grpc.ClientConn) *Client {
	return &Client{
		cc:                cc,
		apiKeysClient:     NewApiKeysClient(cc),
		authClient:        NewAuthClient(cc),
		binClient:         NewBinClient(cc),
		contentClient:     NewContentClient(cc),
		contextsClient:    NewContextsClient(cc),
		coreClient:        NewCoreClient(cc),
		filesClient:       NewFilesClient(cc),
		foldersClient:     NewFoldersClient(cc),
		metadataClient:    NewMetadataClient(cc),
		sharedLinksClient: NewSharedLinksClient(cc),
		spacesClient:      NewSpacesClient(cc),
		uploadsClient:     NewUploadsClient(cc),
		usersClient:       NewUsersClient(cc),
	}
}

// Conn provides access to the underlying gRPC ClientConn
func (c *Client) Conn() *grpc.ClientConn {
	return c.cc
}

// ApiKeysClient returns the ApiKeys module client
func (c *Client) ApiKeys() *ApiKeysClient {
	return c.apiKeysClient
}

// AuthClient returns the Auth module client
func (c *Client) Auth() *AuthClient {
	return c.authClient
}

// BinClient returns the Bin module client
func (c *Client) Bin() *BinClient {
	return c.binClient
}

// ContentClient returns the Content module client
func (c *Client) Content() *ContentClient {
	return c.contentClient
}

// ContextsClient returns the Contexts module client
func (c *Client) Contexts() *ContextsClient {
	return c.contextsClient
}

// CoreClient returns the Core module client
func (c *Client) Core() *CoreClient {
	return c.coreClient
}

// FilesClient returns the Files module client
func (c *Client) Files() *FilesClient {
	return c.filesClient
}

// FoldersClient returns the Folders module client
func (c *Client) Folders() *FoldersClient {
	return c.foldersClient
}

// MetadataClient returns the Metadata module client
func (c *Client) Metadata() *MetadataClient {
	return c.metadataClient
}

// SharedLinksClient returns the SharedLinks module client
func (c *Client) SharedLinks() *SharedLinksClient {
	return c.sharedLinksClient
}

// SpacesClient returns the Spaces module client
func (c *Client) Spaces() *SpacesClient {
	return c.spacesClient
}

// UploadsClient returns the Uploads module client
func (c *Client) Uploads() *UploadsClient {
	return c.uploadsClient
}

// UsersClient returns the Users module client
func (c *Client) Users() *UsersClient {
	return c.usersClient
}
