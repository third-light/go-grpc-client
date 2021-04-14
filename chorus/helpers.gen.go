// This is a generated file, please do not edit by hand
package chorus

import (
	"fmt"

	"github.com/third-light/go-grpc-client/chorus/authdetailstype"
	"github.com/third-light/go-grpc-client/chorus/contentitemtype"
	"github.com/third-light/go-grpc-client/chorus/contexttype"
	"github.com/third-light/go-grpc-client/chorus/directurlsettingsfileformat"
	"github.com/third-light/go-grpc-client/chorus/directurlsettingsfitmode"
	"github.com/third-light/go-grpc-client/chorus/filetype"
	"github.com/third-light/go-grpc-client/chorus/folderdetailstype"
	"github.com/third-light/go-grpc-client/chorus/uploaddetailsstatus"
	"github.com/third-light/go-grpc-client/chorus/vocabmode"
	pb "github.com/third-light/go-grpc-client/protobuf"
)

func convertToApiKeyDetails(in *pb.ApiKeyDetails) ApiKeyDetails {
	return ApiKeyDetails{
		Context:      in.Context,
		CreatedDate:  convertToDateTime(in.CreatedDate),
		Id:           in.Id,
		Label:        in.Label,
		ModifiedDate: convertToDateTime(in.ModifiedDate),
	}
}

func convertFromApiKeyDetails(in ApiKeyDetails) *pb.ApiKeyDetails {
	return &pb.ApiKeyDetails{
		Context:      in.Context,
		CreatedDate:  convertFromDateTime(in.CreatedDate),
		Id:           in.Id,
		Label:        in.Label,
		ModifiedDate: convertFromDateTime(in.ModifiedDate),
	}
}

func convertToApiKeyDetailsPtr(in *pb.ApiKeyDetails) *ApiKeyDetails {
	if in == nil {
		return nil
	}
	value := convertToApiKeyDetails(in)
	return &value
}

func convertToApiKeyDetailsSlice(in []*pb.ApiKeyDetails) []ApiKeyDetails {
	s := make([]ApiKeyDetails, len(in))
	for i := range in {
		s[i] = convertToApiKeyDetails(in[i])
	}
	return s
}

func convertFromApiKeyDetailsSlice(in []ApiKeyDetails) []*pb.ApiKeyDetails {
	s := make([]*pb.ApiKeyDetails, len(in))
	for i := range in {
		s[i] = convertFromApiKeyDetails(in[i])
	}
	return s
}

func convertToApiKeyDetailsWithToken(in *pb.ApiKeyDetailsWithToken) ApiKeyDetailsWithToken {
	return ApiKeyDetailsWithToken{
		ApiKey:       in.ApiKey,
		Context:      in.Context,
		CreatedDate:  convertToDateTime(in.CreatedDate),
		Id:           in.Id,
		Label:        in.Label,
		ModifiedDate: convertToDateTime(in.ModifiedDate),
	}
}

func convertToApiKeyDetailsWithTokenPtr(in *pb.ApiKeyDetailsWithToken) *ApiKeyDetailsWithToken {
	if in == nil {
		return nil
	}
	value := convertToApiKeyDetailsWithToken(in)
	return &value
}

func convertFromApiKeysRequestsCreateProfileKey(in ApiKeysRequestsCreateProfileKey) *pb.ApiKeysRequests_CreateProfileKey {
	return &pb.ApiKeysRequests_CreateProfileKey{
		Details: convertFromCreateApiKeyDetails(in.Details),
	}
}

func convertFromApiKeysRequestsCreateProfileKeyPtr(in *ApiKeysRequestsCreateProfileKey) *pb.ApiKeysRequests_CreateProfileKey {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromApiKeysRequestsCreateProfileKey(value)
}

func convertFromApiKeysRequestsCreateSiteKey(in ApiKeysRequestsCreateSiteKey) *pb.ApiKeysRequests_CreateSiteKey {
	return &pb.ApiKeysRequests_CreateSiteKey{
		Details: convertFromCreateApiKeyDetails(in.Details),
	}
}

func convertFromApiKeysRequestsCreateSiteKeyPtr(in *ApiKeysRequestsCreateSiteKey) *pb.ApiKeysRequests_CreateSiteKey {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromApiKeysRequestsCreateSiteKey(value)
}

func convertFromApiKeysRequestsCreateSpacesKey(in ApiKeysRequestsCreateSpacesKey) *pb.ApiKeysRequests_CreateSpacesKey {
	return &pb.ApiKeysRequests_CreateSpacesKey{
		Details: convertFromCreateApiKeyDetails(in.Details),
		SpaceId: in.SpaceId,
	}
}

func convertFromApiKeysRequestsCreateSpacesKeyPtr(in *ApiKeysRequestsCreateSpacesKey) *pb.ApiKeysRequests_CreateSpacesKey {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromApiKeysRequestsCreateSpacesKey(value)
}

func convertFromApiKeysRequestsDelete(in ApiKeysRequestsDelete) *pb.ApiKeysRequests_Delete {
	return &pb.ApiKeysRequests_Delete{
		ApiKeyId: in.ApiKeyId,
	}
}

func convertFromApiKeysRequestsDeletePtr(in *ApiKeysRequestsDelete) *pb.ApiKeysRequests_Delete {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromApiKeysRequestsDelete(value)
}

func convertFromApiKeysRequestsGet(in ApiKeysRequestsGet) *pb.ApiKeysRequests_Get {
	return &pb.ApiKeysRequests_Get{
		ApiKeyId: in.ApiKeyId,
	}
}

func convertFromApiKeysRequestsGetPtr(in *ApiKeysRequestsGet) *pb.ApiKeysRequests_Get {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromApiKeysRequestsGet(value)
}

func convertFromApiKeysRequestsGetMulti(in ApiKeysRequestsGetMulti) *pb.ApiKeysRequests_GetMulti {
	return &pb.ApiKeysRequests_GetMulti{
		ApiKeyId: in.ApiKeyId,
	}
}

func convertFromApiKeysRequestsGetMultiPtr(in *ApiKeysRequestsGetMulti) *pb.ApiKeysRequests_GetMulti {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromApiKeysRequestsGetMulti(value)
}

func convertFromApiKeysRequestsGetProfileKeys(in ApiKeysRequestsGetProfileKeys) *pb.ApiKeysRequests_GetProfileKeys {
	return &pb.ApiKeysRequests_GetProfileKeys{}
}

func convertFromApiKeysRequestsGetProfileKeysPtr(in *ApiKeysRequestsGetProfileKeys) *pb.ApiKeysRequests_GetProfileKeys {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromApiKeysRequestsGetProfileKeys(value)
}

func convertFromApiKeysRequestsGetSiteKeys(in ApiKeysRequestsGetSiteKeys) *pb.ApiKeysRequests_GetSiteKeys {
	return &pb.ApiKeysRequests_GetSiteKeys{}
}

func convertFromApiKeysRequestsGetSiteKeysPtr(in *ApiKeysRequestsGetSiteKeys) *pb.ApiKeysRequests_GetSiteKeys {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromApiKeysRequestsGetSiteKeys(value)
}

func convertFromApiKeysRequestsGetSpacesKeys(in ApiKeysRequestsGetSpacesKeys) *pb.ApiKeysRequests_GetSpacesKeys {
	return &pb.ApiKeysRequests_GetSpacesKeys{
		SpaceId: in.SpaceId,
	}
}

func convertFromApiKeysRequestsGetSpacesKeysPtr(in *ApiKeysRequestsGetSpacesKeys) *pb.ApiKeysRequests_GetSpacesKeys {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromApiKeysRequestsGetSpacesKeys(value)
}

func convertFromApiKeysRequestsSet(in ApiKeysRequestsSet) *pb.ApiKeysRequests_Set {
	return &pb.ApiKeysRequests_Set{
		ApiKeyId:   in.ApiKeyId,
		Details:    convertFromSettableApiKeyDetails(in.Details),
		UpdateMask: getFieldMask(in),
	}
}

func convertFromApiKeysRequestsSetPtr(in *ApiKeysRequestsSet) *pb.ApiKeysRequests_Set {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromApiKeysRequestsSet(value)
}

func convertToApiKeysResponsesDelete(in *pb.ApiKeysResponses_Delete) ApiKeysResponsesDelete {
	return ApiKeysResponsesDelete{}
}

func convertToApiKeysResponsesDeletePtr(in *pb.ApiKeysResponses_Delete) *ApiKeysResponsesDelete {
	if in == nil {
		return nil
	}
	value := convertToApiKeysResponsesDelete(in)
	return &value
}

func convertToApiKeysResponsesGetMulti(in *pb.ApiKeysResponses_GetMulti) ApiKeysResponsesGetMulti {
	return ApiKeysResponsesGetMulti{
		Response: convertToApiKeyDetailsSlice(in.Response),
	}
}

func convertToApiKeysResponsesGetMultiPtr(in *pb.ApiKeysResponses_GetMulti) *ApiKeysResponsesGetMulti {
	if in == nil {
		return nil
	}
	value := convertToApiKeysResponsesGetMulti(in)
	return &value
}

func convertToApiKeysResponsesGetProfileKeys(in *pb.ApiKeysResponses_GetProfileKeys) ApiKeysResponsesGetProfileKeys {
	return ApiKeysResponsesGetProfileKeys{
		Response: in.Response,
	}
}

func convertToApiKeysResponsesGetProfileKeysPtr(in *pb.ApiKeysResponses_GetProfileKeys) *ApiKeysResponsesGetProfileKeys {
	if in == nil {
		return nil
	}
	value := convertToApiKeysResponsesGetProfileKeys(in)
	return &value
}

func convertToApiKeysResponsesGetSiteKeys(in *pb.ApiKeysResponses_GetSiteKeys) ApiKeysResponsesGetSiteKeys {
	return ApiKeysResponsesGetSiteKeys{
		Response: in.Response,
	}
}

func convertToApiKeysResponsesGetSiteKeysPtr(in *pb.ApiKeysResponses_GetSiteKeys) *ApiKeysResponsesGetSiteKeys {
	if in == nil {
		return nil
	}
	value := convertToApiKeysResponsesGetSiteKeys(in)
	return &value
}

func convertToApiKeysResponsesGetSpacesKeys(in *pb.ApiKeysResponses_GetSpacesKeys) ApiKeysResponsesGetSpacesKeys {
	return ApiKeysResponsesGetSpacesKeys{
		Response: in.Response,
	}
}

func convertToApiKeysResponsesGetSpacesKeysPtr(in *pb.ApiKeysResponses_GetSpacesKeys) *ApiKeysResponsesGetSpacesKeys {
	if in == nil {
		return nil
	}
	value := convertToApiKeysResponsesGetSpacesKeys(in)
	return &value
}

func convertToCreateApiKeyDetails(in *pb.CreateApiKeyDetails) CreateApiKeyDetails {
	return CreateApiKeyDetails{
		Label: in.Label,
	}
}

func convertFromCreateApiKeyDetails(in CreateApiKeyDetails) *pb.CreateApiKeyDetails {
	return &pb.CreateApiKeyDetails{
		Label: in.Label,
	}
}

func convertToSettableApiKeyDetails(in *pb.SettableApiKeyDetails) SettableApiKeyDetails {
	return SettableApiKeyDetails{
		Label: convertToStringPtr(in.Label),
	}
}

func convertFromSettableApiKeyDetails(in SettableApiKeyDetails) *pb.SettableApiKeyDetails {
	return &pb.SettableApiKeyDetails{
		Label: convertFromStringPtr(in.Label),
	}
}

func convertToAuthDetails(in *pb.AuthDetails) AuthDetails {
	return AuthDetails{
		UserId:   in.UserId,
		UserType: authdetailstype.AuthDetailsType(in.UserType),
	}
}

func convertToAuthDetailsPtr(in *pb.AuthDetails) *AuthDetails {
	if in == nil {
		return nil
	}
	value := convertToAuthDetails(in)
	return &value
}

func convertFromAuthRequestsGetAuthDetails(in AuthRequestsGetAuthDetails) *pb.AuthRequests_GetAuthDetails {
	return &pb.AuthRequests_GetAuthDetails{}
}

func convertFromAuthRequestsGetAuthDetailsPtr(in *AuthRequestsGetAuthDetails) *pb.AuthRequests_GetAuthDetails {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromAuthRequestsGetAuthDetails(value)
}

func convertFromAuthRequestsLogin(in AuthRequestsLogin) *pb.AuthRequests_Login {
	return &pb.AuthRequests_Login{
		Password: in.Password,
		Username: in.Username,
	}
}

func convertFromAuthRequestsLoginPtr(in *AuthRequestsLogin) *pb.AuthRequests_Login {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromAuthRequestsLogin(value)
}

func convertFromAuthRequestsLoginWithKey(in AuthRequestsLoginWithKey) *pb.AuthRequests_LoginWithKey {
	return &pb.AuthRequests_LoginWithKey{
		ApiKey: in.ApiKey,
	}
}

func convertFromAuthRequestsLoginWithKeyPtr(in *AuthRequestsLoginWithKey) *pb.AuthRequests_LoginWithKey {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromAuthRequestsLoginWithKey(value)
}

func convertFromAuthRequestsLogout(in AuthRequestsLogout) *pb.AuthRequests_Logout {
	return &pb.AuthRequests_Logout{}
}

func convertFromAuthRequestsLogoutPtr(in *AuthRequestsLogout) *pb.AuthRequests_Logout {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromAuthRequestsLogout(value)
}

func convertToAuthResponsesLogout(in *pb.AuthResponses_Logout) AuthResponsesLogout {
	return AuthResponsesLogout{}
}

func convertToAuthResponsesLogoutPtr(in *pb.AuthResponses_Logout) *AuthResponsesLogout {
	if in == nil {
		return nil
	}
	value := convertToAuthResponsesLogout(in)
	return &value
}

func convertToLoginDetails(in *pb.LoginDetails) LoginDetails {
	return LoginDetails{
		SessionId:   in.SessionId,
		UserDetails: convertToUserDetails(in.UserDetails),
	}
}

func convertToLoginDetailsPtr(in *pb.LoginDetails) *LoginDetails {
	if in == nil {
		return nil
	}
	value := convertToLoginDetails(in)
	return &value
}

func convertFromBinRequestsEmpty(in BinRequestsEmpty) *pb.BinRequests_Empty {
	return &pb.BinRequests_Empty{}
}

func convertFromBinRequestsEmptyPtr(in *BinRequestsEmpty) *pb.BinRequests_Empty {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromBinRequestsEmpty(value)
}

func convertFromBinRequestsGet(in BinRequestsGet) *pb.BinRequests_Get {
	return &pb.BinRequests_Get{}
}

func convertFromBinRequestsGetPtr(in *BinRequestsGet) *pb.BinRequests_Get {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromBinRequestsGet(value)
}

func convertToBinResponsesEmpty(in *pb.BinResponses_Empty) BinResponsesEmpty {
	return BinResponsesEmpty{}
}

func convertToBinResponsesEmptyPtr(in *pb.BinResponses_Empty) *BinResponsesEmpty {
	if in == nil {
		return nil
	}
	value := convertToBinResponsesEmpty(in)
	return &value
}

func convertToFolderDetails(in *pb.FolderDetails) FolderDetails {
	return FolderDetails{
		Context:      in.Context,
		CreatedDate:  convertToDateTime(in.CreatedDate),
		Description:  in.Description,
		Id:           in.Id,
		ModifiedDate: convertToDateTime(in.ModifiedDate),
		Name:         in.Name,
		ParentId:     convertToStringPtr(in.ParentId),
		Type:         folderdetailstype.FolderDetailsType(in.Type),
	}
}

func convertFromFolderDetails(in FolderDetails) *pb.FolderDetails {
	return &pb.FolderDetails{
		Context:      in.Context,
		CreatedDate:  convertFromDateTime(in.CreatedDate),
		Description:  in.Description,
		Id:           in.Id,
		ModifiedDate: convertFromDateTime(in.ModifiedDate),
		Name:         in.Name,
		ParentId:     convertFromStringPtr(in.ParentId),
		Type:         pb.FolderDetails_Type(in.Type),
	}
}

func convertToFolderDetailsPtr(in *pb.FolderDetails) *FolderDetails {
	if in == nil {
		return nil
	}
	value := convertToFolderDetails(in)
	return &value
}

func convertToFolderDetailsSlice(in []*pb.FolderDetails) []FolderDetails {
	s := make([]FolderDetails, len(in))
	for i := range in {
		s[i] = convertToFolderDetails(in[i])
	}
	return s
}

func convertFromFolderDetailsSlice(in []FolderDetails) []*pb.FolderDetails {
	s := make([]*pb.FolderDetails, len(in))
	for i := range in {
		s[i] = convertFromFolderDetails(in[i])
	}
	return s
}

func convertToFileDetails(in *pb.FileDetails) FileDetails {
	return FileDetails{
		CanViewOriginal: in.CanViewOriginal,
		CreatedDate:     convertToDateTime(in.CreatedDate),
		Filename:        in.Filename,
		FileSizeBytes:   in.FileSizeBytes,
		FileType:        filetype.FileType(in.FileType),
		Id:              in.Id,
		IsDerivative:    in.IsDerivative,
		Media:           convertToMediaDetails(in.Media),
		ModifiedDate:    convertToDateTime(in.ModifiedDate),
		OriginalAssetId: convertToStringPtr(in.OriginalAssetId),
		OwnerId:         convertToStringPtr(in.OwnerId),
		ParentId:        in.ParentId,
		RevisionNumber:  in.RevisionNumber,
		Thumbnails:      convertToThumbnails(in.Thumbnails),
	}
}

func convertFromFileDetails(in FileDetails) *pb.FileDetails {
	return &pb.FileDetails{
		CanViewOriginal: in.CanViewOriginal,
		CreatedDate:     convertFromDateTime(in.CreatedDate),
		Filename:        in.Filename,
		FileSizeBytes:   in.FileSizeBytes,
		FileType:        pb.FileType(in.FileType),
		Id:              in.Id,
		IsDerivative:    in.IsDerivative,
		Media:           in.Media.getProto(),
		ModifiedDate:    convertFromDateTime(in.ModifiedDate),
		OriginalAssetId: convertFromStringPtr(in.OriginalAssetId),
		OwnerId:         convertFromStringPtr(in.OwnerId),
		ParentId:        in.ParentId,
		RevisionNumber:  in.RevisionNumber,
		Thumbnails:      convertFromThumbnails(in.Thumbnails),
	}
}

func convertToFileDetailsPtr(in *pb.FileDetails) *FileDetails {
	if in == nil {
		return nil
	}
	value := convertToFileDetails(in)
	return &value
}

func convertToFileDetailsSlice(in []*pb.FileDetails) []FileDetails {
	s := make([]FileDetails, len(in))
	for i := range in {
		s[i] = convertToFileDetails(in[i])
	}
	return s
}

func convertFromFileDetailsSlice(in []FileDetails) []*pb.FileDetails {
	s := make([]*pb.FileDetails, len(in))
	for i := range in {
		s[i] = convertFromFileDetails(in[i])
	}
	return s
}

func convertToFileLinkDetails(in *pb.FileLinkDetails) FileLinkDetails {
	return FileLinkDetails{
		CreatedDate:      convertToDateTime(in.CreatedDate),
		Id:               in.Id,
		LinkName:         in.LinkName,
		ModifiedDate:     convertToDateTime(in.ModifiedDate),
		OriginalFileId:   in.OriginalFileId,
		OriginalFileType: filetype.FileType(in.OriginalFileType),
	}
}

func convertFromFileLinkDetails(in FileLinkDetails) *pb.FileLinkDetails {
	return &pb.FileLinkDetails{
		CreatedDate:      convertFromDateTime(in.CreatedDate),
		Id:               in.Id,
		LinkName:         in.LinkName,
		ModifiedDate:     convertFromDateTime(in.ModifiedDate),
		OriginalFileId:   in.OriginalFileId,
		OriginalFileType: pb.FileType(in.OriginalFileType),
	}
}

func convertToFileLinkDetailsPtr(in *pb.FileLinkDetails) *FileLinkDetails {
	if in == nil {
		return nil
	}
	value := convertToFileLinkDetails(in)
	return &value
}

func convertToFileLinkDetailsSlice(in []*pb.FileLinkDetails) []FileLinkDetails {
	s := make([]FileLinkDetails, len(in))
	for i := range in {
		s[i] = convertToFileLinkDetails(in[i])
	}
	return s
}

func convertFromFileLinkDetailsSlice(in []FileLinkDetails) []*pb.FileLinkDetails {
	s := make([]*pb.FileLinkDetails, len(in))
	for i := range in {
		s[i] = convertFromFileLinkDetails(in[i])
	}
	return s
}

func convertToFolderLinkDetails(in *pb.FolderLinkDetails) FolderLinkDetails {
	return FolderLinkDetails{
		CreatedDate:      convertToDateTime(in.CreatedDate),
		Id:               in.Id,
		LinkName:         in.LinkName,
		ModifiedDate:     convertToDateTime(in.ModifiedDate),
		OriginalFolderId: in.OriginalFolderId,
	}
}

func convertFromFolderLinkDetails(in FolderLinkDetails) *pb.FolderLinkDetails {
	return &pb.FolderLinkDetails{
		CreatedDate:      convertFromDateTime(in.CreatedDate),
		Id:               in.Id,
		LinkName:         in.LinkName,
		ModifiedDate:     convertFromDateTime(in.ModifiedDate),
		OriginalFolderId: in.OriginalFolderId,
	}
}

func convertToFolderLinkDetailsPtr(in *pb.FolderLinkDetails) *FolderLinkDetails {
	if in == nil {
		return nil
	}
	value := convertToFolderLinkDetails(in)
	return &value
}

func convertToFolderLinkDetailsSlice(in []*pb.FolderLinkDetails) []FolderLinkDetails {
	s := make([]FolderLinkDetails, len(in))
	for i := range in {
		s[i] = convertToFolderLinkDetails(in[i])
	}
	return s
}

func convertFromFolderLinkDetailsSlice(in []FolderLinkDetails) []*pb.FolderLinkDetails {
	s := make([]*pb.FolderLinkDetails, len(in))
	for i := range in {
		s[i] = convertFromFolderLinkDetails(in[i])
	}
	return s
}

func convertFromContentItemDetails(in ContentItemDetails) *pb.ContentItemDetails {
	return in.getProto()
}

func convertToContentItemDetails(in *pb.ContentItemDetails) ContentItemDetails {
	switch in.Details.(type) {
	case *pb.ContentItemDetails_Derivative:
		v := in.Details.(*pb.ContentItemDetails_Derivative)
		return convertToFileDetails(v.Derivative)
	case *pb.ContentItemDetails_File:
		v := in.Details.(*pb.ContentItemDetails_File)
		return convertToFileDetails(v.File)
	case *pb.ContentItemDetails_FileLink:
		v := in.Details.(*pb.ContentItemDetails_FileLink)
		return convertToFileLinkDetails(v.FileLink)
	case *pb.ContentItemDetails_Folder:
		v := in.Details.(*pb.ContentItemDetails_Folder)
		return convertToFolderDetails(v.Folder)
	case *pb.ContentItemDetails_FolderLink:
		v := in.Details.(*pb.ContentItemDetails_FolderLink)
		return convertToFolderLinkDetails(v.FolderLink)
	default:
		panic(fmt.Errorf("Unknown type '%T' for pb.ContentItemDetails", in.Details))
	}
}

func convertToContentItemWithDetails(in *pb.ContentItemWithDetails) ContentItemWithDetails {
	return ContentItemWithDetails{
		Details: convertToContentItemDetails(in.Details),
		Id:      in.Id,
		Name:    in.Name,
		Type:    contentitemtype.ContentItemType(in.Type),
	}
}

func convertFromContentItemWithDetails(in ContentItemWithDetails) *pb.ContentItemWithDetails {
	return &pb.ContentItemWithDetails{
		Details: in.Details.getProto(),
		Id:      in.Id,
		Name:    in.Name,
		Type:    pb.ContentItemType(in.Type),
	}
}

func convertToContentItemWithDetailsPtr(in *pb.ContentItemWithDetails) *ContentItemWithDetails {
	if in == nil {
		return nil
	}
	value := convertToContentItemWithDetails(in)
	return &value
}

func convertToContentItemWithDetailsSlice(in []*pb.ContentItemWithDetails) []ContentItemWithDetails {
	s := make([]ContentItemWithDetails, len(in))
	for i := range in {
		s[i] = convertToContentItemWithDetails(in[i])
	}
	return s
}

func convertFromContentItemWithDetailsSlice(in []ContentItemWithDetails) []*pb.ContentItemWithDetails {
	s := make([]*pb.ContentItemWithDetails, len(in))
	for i := range in {
		s[i] = convertFromContentItemWithDetails(in[i])
	}
	return s
}

func convertFromContentRequestsDelete(in ContentRequestsDelete) *pb.ContentRequests_Delete {
	return &pb.ContentRequests_Delete{
		ItemId: in.ItemId,
	}
}

func convertFromContentRequestsDeletePtr(in *ContentRequestsDelete) *pb.ContentRequests_Delete {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromContentRequestsDelete(value)
}

func convertFromContentRequestsGet(in ContentRequestsGet) *pb.ContentRequests_Get {
	return &pb.ContentRequests_Get{
		ItemId: in.ItemId,
	}
}

func convertFromContentRequestsGetPtr(in *ContentRequestsGet) *pb.ContentRequests_Get {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromContentRequestsGet(value)
}

func convertFromContentRequestsGetMulti(in ContentRequestsGetMulti) *pb.ContentRequests_GetMulti {
	return &pb.ContentRequests_GetMulti{
		ItemIds: in.ItemIds,
	}
}

func convertFromContentRequestsGetMultiPtr(in *ContentRequestsGetMulti) *pb.ContentRequests_GetMulti {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromContentRequestsGetMulti(value)
}

func convertToContentResponsesDelete(in *pb.ContentResponses_Delete) ContentResponsesDelete {
	return ContentResponsesDelete{}
}

func convertToContentResponsesDeletePtr(in *pb.ContentResponses_Delete) *ContentResponsesDelete {
	if in == nil {
		return nil
	}
	value := convertToContentResponsesDelete(in)
	return &value
}

func convertToContentResponsesGetMulti(in *pb.ContentResponses_GetMulti) ContentResponsesGetMulti {
	return ContentResponsesGetMulti{
		Response: convertToContentItemWithDetailsSlice(in.Response),
	}
}

func convertToContentResponsesGetMultiPtr(in *pb.ContentResponses_GetMulti) *ContentResponsesGetMulti {
	if in == nil {
		return nil
	}
	value := convertToContentResponsesGetMulti(in)
	return &value
}

func convertToContextDetails(in *pb.ContextDetails) ContextDetails {
	return ContextDetails{
		AvatarUrl:       in.AvatarUrl,
		BackingFolderId: convertToStringPtr(in.BackingFolderId),
		Id:              in.Id,
		Name:            in.Name,
		OwnerId:         convertToStringPtr(in.OwnerId),
		ParentId:        in.ParentId,
		Type:            contexttype.ContextType(in.Type),
	}
}

func convertFromContextDetails(in ContextDetails) *pb.ContextDetails {
	return &pb.ContextDetails{
		AvatarUrl:       in.AvatarUrl,
		BackingFolderId: convertFromStringPtr(in.BackingFolderId),
		Id:              in.Id,
		Name:            in.Name,
		OwnerId:         convertFromStringPtr(in.OwnerId),
		ParentId:        in.ParentId,
		Type:            pb.ContextType(in.Type),
	}
}

func convertToContextDetailsPtr(in *pb.ContextDetails) *ContextDetails {
	if in == nil {
		return nil
	}
	value := convertToContextDetails(in)
	return &value
}

func convertToContextDetailsSlice(in []*pb.ContextDetails) []ContextDetails {
	s := make([]ContextDetails, len(in))
	for i := range in {
		s[i] = convertToContextDetails(in[i])
	}
	return s
}

func convertFromContextDetailsSlice(in []ContextDetails) []*pb.ContextDetails {
	s := make([]*pb.ContextDetails, len(in))
	for i := range in {
		s[i] = convertFromContextDetails(in[i])
	}
	return s
}

func convertFromContextsRequestsGet(in ContextsRequestsGet) *pb.ContextsRequests_Get {
	return &pb.ContextsRequests_Get{
		ContextId: in.ContextId,
	}
}

func convertFromContextsRequestsGetPtr(in *ContextsRequestsGet) *pb.ContextsRequests_Get {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromContextsRequestsGet(value)
}

func convertFromContextsRequestsGetChildren(in ContextsRequestsGetChildren) *pb.ContextsRequests_GetChildren {
	return &pb.ContextsRequests_GetChildren{
		ContextId: in.ContextId,
	}
}

func convertFromContextsRequestsGetChildrenPtr(in *ContextsRequestsGetChildren) *pb.ContextsRequests_GetChildren {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromContextsRequestsGetChildren(value)
}

func convertFromContextsRequestsGetMulti(in ContextsRequestsGetMulti) *pb.ContextsRequests_GetMulti {
	return &pb.ContextsRequests_GetMulti{
		ContextIds: in.ContextIds,
	}
}

func convertFromContextsRequestsGetMultiPtr(in *ContextsRequestsGetMulti) *pb.ContextsRequests_GetMulti {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromContextsRequestsGetMulti(value)
}

func convertFromContextsRequestsGetSharedLinks(in ContextsRequestsGetSharedLinks) *pb.ContextsRequests_GetSharedLinks {
	return &pb.ContextsRequests_GetSharedLinks{
		ContextId: in.ContextId,
	}
}

func convertFromContextsRequestsGetSharedLinksPtr(in *ContextsRequestsGetSharedLinks) *pb.ContextsRequests_GetSharedLinks {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromContextsRequestsGetSharedLinks(value)
}

func convertToContextsResponsesGetChildren(in *pb.ContextsResponses_GetChildren) ContextsResponsesGetChildren {
	return ContextsResponsesGetChildren{
		Response: in.Response,
	}
}

func convertToContextsResponsesGetChildrenPtr(in *pb.ContextsResponses_GetChildren) *ContextsResponsesGetChildren {
	if in == nil {
		return nil
	}
	value := convertToContextsResponsesGetChildren(in)
	return &value
}

func convertToContextsResponsesGetMulti(in *pb.ContextsResponses_GetMulti) ContextsResponsesGetMulti {
	return ContextsResponsesGetMulti{
		Response: convertToContextDetailsSlice(in.Response),
	}
}

func convertToContextsResponsesGetMultiPtr(in *pb.ContextsResponses_GetMulti) *ContextsResponsesGetMulti {
	if in == nil {
		return nil
	}
	value := convertToContextsResponsesGetMulti(in)
	return &value
}

func convertToContextsResponsesGetSharedLinks(in *pb.ContextsResponses_GetSharedLinks) ContextsResponsesGetSharedLinks {
	return ContextsResponsesGetSharedLinks{
		Response: convertToSharedLinkDetailsSlice(in.Response),
	}
}

func convertToContextsResponsesGetSharedLinksPtr(in *pb.ContextsResponses_GetSharedLinks) *ContextsResponsesGetSharedLinks {
	if in == nil {
		return nil
	}
	value := convertToContextsResponsesGetSharedLinks(in)
	return &value
}

func convertFromCoreRequestsGetEnvironment(in CoreRequestsGetEnvironment) *pb.CoreRequests_GetEnvironment {
	return &pb.CoreRequests_GetEnvironment{}
}

func convertFromCoreRequestsGetEnvironmentPtr(in *CoreRequestsGetEnvironment) *pb.CoreRequests_GetEnvironment {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromCoreRequestsGetEnvironment(value)
}

func convertToDateTime(in *pb.DateTime) DateTime {
	return DateTime{
		Rfc3339:   in.Rfc3339,
		Timestamp: in.Timestamp,
	}
}

func convertFromDateTime(in DateTime) *pb.DateTime {
	return &pb.DateTime{
		Rfc3339:   in.Rfc3339,
		Timestamp: in.Timestamp,
	}
}

func convertToEnvironmentDetails(in *pb.EnvironmentDetails) EnvironmentDetails {
	return EnvironmentDetails{
		BrowserTitle: in.BrowserTitle,
		Title:        in.Title,
		Version:      in.Version,
	}
}

func convertToEnvironmentDetailsPtr(in *pb.EnvironmentDetails) *EnvironmentDetails {
	if in == nil {
		return nil
	}
	value := convertToEnvironmentDetails(in)
	return &value
}

func convertToDirectUrlSettings(in *pb.DirectUrlSettings) DirectUrlSettings {
	return DirectUrlSettings{
		Blur:     convertToInt64Ptr(in.Blur),
		Crop:     convertToDirectUrlSettingsCropSettingsPtr(in.Crop),
		Download: convertToBoolPtr(in.Download),
		Dpi:      convertToInt64Ptr(in.Dpi),
		Filename: convertToStringPtr(in.Filename),
		Fit:      directurlsettingsfitmode.DirectUrlSettingsFitModePtr(directurlsettingsfitmode.DirectUrlSettingsFitMode(in.Fit)),
		Format:   directurlsettingsfileformat.DirectUrlSettingsFileFormatPtr(directurlsettingsfileformat.DirectUrlSettingsFileFormat(in.Format)),
		Height:   convertToInt64Ptr(in.Height),
		Page:     convertToInt64Ptr(in.Page),
		Quality:  convertToInt64Ptr(in.Quality),
		Rotate:   convertToInt64Ptr(in.Rotate),
		Width:    convertToInt64Ptr(in.Width),
	}
}

func convertFromDirectUrlSettings(in DirectUrlSettings) *pb.DirectUrlSettings {
	return &pb.DirectUrlSettings{
		Blur:     convertFromInt64Ptr(in.Blur),
		Crop:     convertFromDirectUrlSettingsCropSettingsPtr(in.Crop),
		Download: convertFromBoolPtr(in.Download),
		Dpi:      convertFromInt64Ptr(in.Dpi),
		Filename: convertFromStringPtr(in.Filename),
		Fit:      pb.DirectUrlSettings_FitMode(directurlsettingsfitmode.DirectUrlSettingsFitModeFromPtr(in.Fit)),
		Format:   pb.DirectUrlSettings_FileFormat(directurlsettingsfileformat.DirectUrlSettingsFileFormatFromPtr(in.Format)),
		Height:   convertFromInt64Ptr(in.Height),
		Page:     convertFromInt64Ptr(in.Page),
		Quality:  convertFromInt64Ptr(in.Quality),
		Rotate:   convertFromInt64Ptr(in.Rotate),
		Width:    convertFromInt64Ptr(in.Width),
	}
}

func convertToDirectUrlSettingsCropSettings(in *pb.DirectUrlSettings_CropSettings) DirectUrlSettingsCropSettings {
	return DirectUrlSettingsCropSettings{
		Height: in.Height,
		Width:  in.Width,
		X:      in.X,
		Y:      in.Y,
	}
}

func convertFromDirectUrlSettingsCropSettings(in DirectUrlSettingsCropSettings) *pb.DirectUrlSettings_CropSettings {
	return &pb.DirectUrlSettings_CropSettings{
		Height: in.Height,
		Width:  in.Width,
		X:      in.X,
		Y:      in.Y,
	}
}

func convertToDirectUrlSettingsCropSettingsPtr(in *pb.DirectUrlSettings_CropSettings) *DirectUrlSettingsCropSettings {
	if in == nil {
		return nil
	}
	value := convertToDirectUrlSettingsCropSettings(in)
	return &value
}

func convertFromDirectUrlSettingsCropSettingsPtr(in *DirectUrlSettingsCropSettings) *pb.DirectUrlSettings_CropSettings {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromDirectUrlSettingsCropSettings(value)
}

func convertToAudioDetails(in *pb.AudioDetails) AudioDetails {
	return AudioDetails{
		DurationSecs: in.DurationSecs,
	}
}

func convertFromAudioDetails(in AudioDetails) *pb.AudioDetails {
	return &pb.AudioDetails{
		DurationSecs: in.DurationSecs,
	}
}

func convertToDocumentDetails(in *pb.DocumentDetails) DocumentDetails {
	return DocumentDetails{
		Height:    in.Height,
		PageCount: in.PageCount,
		Width:     in.Width,
	}
}

func convertFromDocumentDetails(in DocumentDetails) *pb.DocumentDetails {
	return &pb.DocumentDetails{
		Height:    in.Height,
		PageCount: in.PageCount,
		Width:     in.Width,
	}
}

func convertFromFilesRequestsAttach(in FilesRequestsAttach) *pb.FilesRequests_Attach {
	return &pb.FilesRequests_Attach{
		AttachedFileIds: in.AttachedFileIds,
		FileId:          in.FileId,
	}
}

func convertFromFilesRequestsAttachPtr(in *FilesRequestsAttach) *pb.FilesRequests_Attach {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromFilesRequestsAttach(value)
}

func convertFromFilesRequestsDeleteDirectUrl(in FilesRequestsDeleteDirectUrl) *pb.FilesRequests_DeleteDirectUrl {
	return &pb.FilesRequests_DeleteDirectUrl{
		FileId: in.FileId,
	}
}

func convertFromFilesRequestsDeleteDirectUrlPtr(in *FilesRequestsDeleteDirectUrl) *pb.FilesRequests_DeleteDirectUrl {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromFilesRequestsDeleteDirectUrl(value)
}

func convertFromFilesRequestsDetach(in FilesRequestsDetach) *pb.FilesRequests_Detach {
	return &pb.FilesRequests_Detach{
		AttachedFileId: in.AttachedFileId,
		FileId:         in.FileId,
	}
}

func convertFromFilesRequestsDetachPtr(in *FilesRequestsDetach) *pb.FilesRequests_Detach {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromFilesRequestsDetach(value)
}

func convertFromFilesRequestsDetachAll(in FilesRequestsDetachAll) *pb.FilesRequests_DetachAll {
	return &pb.FilesRequests_DetachAll{
		FileId: in.FileId,
	}
}

func convertFromFilesRequestsDetachAllPtr(in *FilesRequestsDetachAll) *pb.FilesRequests_DetachAll {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromFilesRequestsDetachAll(value)
}

func convertFromFilesRequestsGet(in FilesRequestsGet) *pb.FilesRequests_Get {
	return &pb.FilesRequests_Get{
		FileId: in.FileId,
	}
}

func convertFromFilesRequestsGetPtr(in *FilesRequestsGet) *pb.FilesRequests_Get {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromFilesRequestsGet(value)
}

func convertFromFilesRequestsGetAttachments(in FilesRequestsGetAttachments) *pb.FilesRequests_GetAttachments {
	return &pb.FilesRequests_GetAttachments{
		FileId: in.FileId,
	}
}

func convertFromFilesRequestsGetAttachmentsPtr(in *FilesRequestsGetAttachments) *pb.FilesRequests_GetAttachments {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromFilesRequestsGetAttachments(value)
}

func convertFromFilesRequestsGetDirectUrl(in FilesRequestsGetDirectUrl) *pb.FilesRequests_GetDirectUrl {
	return &pb.FilesRequests_GetDirectUrl{
		FileId:   in.FileId,
		Settings: convertFromDirectUrlSettings(in.Settings),
	}
}

func convertFromFilesRequestsGetDirectUrlPtr(in *FilesRequestsGetDirectUrl) *pb.FilesRequests_GetDirectUrl {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromFilesRequestsGetDirectUrl(value)
}

func convertFromFilesRequestsGetLink(in FilesRequestsGetLink) *pb.FilesRequests_GetLink {
	return &pb.FilesRequests_GetLink{
		LinkId: in.LinkId,
	}
}

func convertFromFilesRequestsGetLinkPtr(in *FilesRequestsGetLink) *pb.FilesRequests_GetLink {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromFilesRequestsGetLink(value)
}

func convertFromFilesRequestsGetLinkMulti(in FilesRequestsGetLinkMulti) *pb.FilesRequests_GetLinkMulti {
	return &pb.FilesRequests_GetLinkMulti{
		LinkIds: in.LinkIds,
	}
}

func convertFromFilesRequestsGetLinkMultiPtr(in *FilesRequestsGetLinkMulti) *pb.FilesRequests_GetLinkMulti {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromFilesRequestsGetLinkMulti(value)
}

func convertFromFilesRequestsGetMetadata(in FilesRequestsGetMetadata) *pb.FilesRequests_GetMetadata {
	return &pb.FilesRequests_GetMetadata{
		FileId: in.FileId,
	}
}

func convertFromFilesRequestsGetMetadataPtr(in *FilesRequestsGetMetadata) *pb.FilesRequests_GetMetadata {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromFilesRequestsGetMetadata(value)
}

func convertFromFilesRequestsGetMulti(in FilesRequestsGetMulti) *pb.FilesRequests_GetMulti {
	return &pb.FilesRequests_GetMulti{
		FileIds: in.FileIds,
	}
}

func convertFromFilesRequestsGetMultiPtr(in *FilesRequestsGetMulti) *pb.FilesRequests_GetMulti {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromFilesRequestsGetMulti(value)
}

func convertFromFilesRequestsGetTemporaryDirectUrl(in FilesRequestsGetTemporaryDirectUrl) *pb.FilesRequests_GetTemporaryDirectUrl {
	return &pb.FilesRequests_GetTemporaryDirectUrl{
		FileId:   in.FileId,
		Settings: convertFromDirectUrlSettings(in.Settings),
	}
}

func convertFromFilesRequestsGetTemporaryDirectUrlPtr(in *FilesRequestsGetTemporaryDirectUrl) *pb.FilesRequests_GetTemporaryDirectUrl {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromFilesRequestsGetTemporaryDirectUrl(value)
}

func convertToFilesResponsesAttach(in *pb.FilesResponses_Attach) FilesResponsesAttach {
	return FilesResponsesAttach{
		Response: convertToFileDetailsSlice(in.Response),
	}
}

func convertToFilesResponsesAttachPtr(in *pb.FilesResponses_Attach) *FilesResponsesAttach {
	if in == nil {
		return nil
	}
	value := convertToFilesResponsesAttach(in)
	return &value
}

func convertToFilesResponsesDeleteDirectUrl(in *pb.FilesResponses_DeleteDirectUrl) FilesResponsesDeleteDirectUrl {
	return FilesResponsesDeleteDirectUrl{}
}

func convertToFilesResponsesDeleteDirectUrlPtr(in *pb.FilesResponses_DeleteDirectUrl) *FilesResponsesDeleteDirectUrl {
	if in == nil {
		return nil
	}
	value := convertToFilesResponsesDeleteDirectUrl(in)
	return &value
}

func convertToFilesResponsesDetach(in *pb.FilesResponses_Detach) FilesResponsesDetach {
	return FilesResponsesDetach{
		Response: convertToFileDetailsSlice(in.Response),
	}
}

func convertToFilesResponsesDetachPtr(in *pb.FilesResponses_Detach) *FilesResponsesDetach {
	if in == nil {
		return nil
	}
	value := convertToFilesResponsesDetach(in)
	return &value
}

func convertToFilesResponsesDetachAll(in *pb.FilesResponses_DetachAll) FilesResponsesDetachAll {
	return FilesResponsesDetachAll{}
}

func convertToFilesResponsesDetachAllPtr(in *pb.FilesResponses_DetachAll) *FilesResponsesDetachAll {
	if in == nil {
		return nil
	}
	value := convertToFilesResponsesDetachAll(in)
	return &value
}

func convertToFilesResponsesGetAttachments(in *pb.FilesResponses_GetAttachments) FilesResponsesGetAttachments {
	return FilesResponsesGetAttachments{
		Response: convertToFileDetailsSlice(in.Response),
	}
}

func convertToFilesResponsesGetAttachmentsPtr(in *pb.FilesResponses_GetAttachments) *FilesResponsesGetAttachments {
	if in == nil {
		return nil
	}
	value := convertToFilesResponsesGetAttachments(in)
	return &value
}

func convertToFilesResponsesGetDirectUrl(in *pb.FilesResponses_GetDirectUrl) FilesResponsesGetDirectUrl {
	return FilesResponsesGetDirectUrl{
		Response: in.Response,
	}
}

func convertToFilesResponsesGetDirectUrlPtr(in *pb.FilesResponses_GetDirectUrl) *FilesResponsesGetDirectUrl {
	if in == nil {
		return nil
	}
	value := convertToFilesResponsesGetDirectUrl(in)
	return &value
}

func convertToFilesResponsesGetLinkMulti(in *pb.FilesResponses_GetLinkMulti) FilesResponsesGetLinkMulti {
	return FilesResponsesGetLinkMulti{
		Response: convertToFileLinkDetailsSlice(in.Response),
	}
}

func convertToFilesResponsesGetLinkMultiPtr(in *pb.FilesResponses_GetLinkMulti) *FilesResponsesGetLinkMulti {
	if in == nil {
		return nil
	}
	value := convertToFilesResponsesGetLinkMulti(in)
	return &value
}

func convertToFilesResponsesGetMulti(in *pb.FilesResponses_GetMulti) FilesResponsesGetMulti {
	return FilesResponsesGetMulti{
		Response: convertToFileDetailsSlice(in.Response),
	}
}

func convertToFilesResponsesGetMultiPtr(in *pb.FilesResponses_GetMulti) *FilesResponsesGetMulti {
	if in == nil {
		return nil
	}
	value := convertToFilesResponsesGetMulti(in)
	return &value
}

func convertToFilesResponsesGetTemporaryDirectUrl(in *pb.FilesResponses_GetTemporaryDirectUrl) FilesResponsesGetTemporaryDirectUrl {
	return FilesResponsesGetTemporaryDirectUrl{
		Response: in.Response,
	}
}

func convertToFilesResponsesGetTemporaryDirectUrlPtr(in *pb.FilesResponses_GetTemporaryDirectUrl) *FilesResponsesGetTemporaryDirectUrl {
	if in == nil {
		return nil
	}
	value := convertToFilesResponsesGetTemporaryDirectUrl(in)
	return &value
}

func convertToGenericDetails(in *pb.GenericDetails) GenericDetails {
	return GenericDetails{}
}

func convertFromGenericDetails(in GenericDetails) *pb.GenericDetails {
	return &pb.GenericDetails{}
}

func convertToImageDetails(in *pb.ImageDetails) ImageDetails {
	return ImageDetails{
		Height: in.Height,
		Width:  in.Width,
	}
}

func convertFromImageDetails(in ImageDetails) *pb.ImageDetails {
	return &pb.ImageDetails{
		Height: in.Height,
		Width:  in.Width,
	}
}

func convertToVideoDetails(in *pb.VideoDetails) VideoDetails {
	return VideoDetails{
		DurationSecs: in.DurationSecs,
		Framerate:    in.Framerate,
		Height:       in.Height,
		Width:        in.Width,
	}
}

func convertFromVideoDetails(in VideoDetails) *pb.VideoDetails {
	return &pb.VideoDetails{
		DurationSecs: in.DurationSecs,
		Framerate:    in.Framerate,
		Height:       in.Height,
		Width:        in.Width,
	}
}

func convertFromMediaDetails(in MediaDetails) *pb.MediaDetails {
	return in.getProto()
}

func convertToMediaDetails(in *pb.MediaDetails) MediaDetails {
	switch in.Details.(type) {
	case *pb.MediaDetails_Audio:
		v := in.Details.(*pb.MediaDetails_Audio)
		return convertToAudioDetails(v.Audio)
	case *pb.MediaDetails_Document:
		v := in.Details.(*pb.MediaDetails_Document)
		return convertToDocumentDetails(v.Document)
	case *pb.MediaDetails_Generic:
		v := in.Details.(*pb.MediaDetails_Generic)
		return convertToGenericDetails(v.Generic)
	case *pb.MediaDetails_Image:
		v := in.Details.(*pb.MediaDetails_Image)
		return convertToImageDetails(v.Image)
	case *pb.MediaDetails_Video:
		v := in.Details.(*pb.MediaDetails_Video)
		return convertToVideoDetails(v.Video)
	default:
		panic(fmt.Errorf("Unknown type '%T' for pb.MediaDetails", in.Details))
	}
}

func convertToThumbDetails(in *pb.ThumbDetails) ThumbDetails {
	return ThumbDetails{
		Height: in.Height,
		Url:    in.Url,
		Width:  in.Width,
	}
}

func convertFromThumbDetails(in ThumbDetails) *pb.ThumbDetails {
	return &pb.ThumbDetails{
		Height: in.Height,
		Url:    in.Url,
		Width:  in.Width,
	}
}

func convertToThumbDetailsPtr(in *pb.ThumbDetails) *ThumbDetails {
	if in == nil {
		return nil
	}
	value := convertToThumbDetails(in)
	return &value
}

func convertFromThumbDetailsPtr(in *ThumbDetails) *pb.ThumbDetails {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromThumbDetails(value)
}

func convertToThumbnails(in *pb.Thumbnails) Thumbnails {
	return Thumbnails{
		Large:  convertToThumbDetailsPtr(in.Large),
		Medium: convertToThumbDetailsPtr(in.Medium),
		Small:  convertToThumbDetailsPtr(in.Small),
	}
}

func convertFromThumbnails(in Thumbnails) *pb.Thumbnails {
	return &pb.Thumbnails{
		Large:  convertFromThumbDetailsPtr(in.Large),
		Medium: convertFromThumbDetailsPtr(in.Medium),
		Small:  convertFromThumbDetailsPtr(in.Small),
	}
}

func convertToMapOfstringToMetadataValue(in map[string]*pb.MetadataValue) map[string]MetadataValue {
	m := make(map[string]MetadataValue)
	for key, value := range in {
		m[key] = convertToMetadataValue(value)
	}
	return m
}

func convertFromMapOfstringToMetadataValue(in map[string]MetadataValue) map[string]*pb.MetadataValue {
	m := make(map[string]*pb.MetadataValue)
	for key, value := range in {
		m[key] = convertFromMetadataValue(value)
	}
	return m
}

func convertToMetadataValueMap(in *pb.MetadataValueMap) MetadataValueMap {
	return MetadataValueMap{
		Values: convertToMapOfstringToMetadataValue(in.Values),
	}
}

func convertToMetadataValueMapPtr(in *pb.MetadataValueMap) *MetadataValueMap {
	if in == nil {
		return nil
	}
	value := convertToMetadataValueMap(in)
	return &value
}

func convertFromFoldersRequestsGet(in FoldersRequestsGet) *pb.FoldersRequests_Get {
	return &pb.FoldersRequests_Get{
		FolderId: in.FolderId,
	}
}

func convertFromFoldersRequestsGetPtr(in *FoldersRequestsGet) *pb.FoldersRequests_Get {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromFoldersRequestsGet(value)
}

func convertFromFoldersRequestsGetChildFiles(in FoldersRequestsGetChildFiles) *pb.FoldersRequests_GetChildFiles {
	return &pb.FoldersRequests_GetChildFiles{
		FolderId: in.FolderId,
	}
}

func convertFromFoldersRequestsGetChildFilesPtr(in *FoldersRequestsGetChildFiles) *pb.FoldersRequests_GetChildFiles {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromFoldersRequestsGetChildFiles(value)
}

func convertFromFoldersRequestsGetChildFolders(in FoldersRequestsGetChildFolders) *pb.FoldersRequests_GetChildFolders {
	return &pb.FoldersRequests_GetChildFolders{
		FolderId: in.FolderId,
	}
}

func convertFromFoldersRequestsGetChildFoldersPtr(in *FoldersRequestsGetChildFolders) *pb.FoldersRequests_GetChildFolders {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromFoldersRequestsGetChildFolders(value)
}

func convertFromFoldersRequestsGetLink(in FoldersRequestsGetLink) *pb.FoldersRequests_GetLink {
	return &pb.FoldersRequests_GetLink{
		LinkId: in.LinkId,
	}
}

func convertFromFoldersRequestsGetLinkPtr(in *FoldersRequestsGetLink) *pb.FoldersRequests_GetLink {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromFoldersRequestsGetLink(value)
}

func convertFromFoldersRequestsGetLinkMulti(in FoldersRequestsGetLinkMulti) *pb.FoldersRequests_GetLinkMulti {
	return &pb.FoldersRequests_GetLinkMulti{
		LinkIds: in.LinkIds,
	}
}

func convertFromFoldersRequestsGetLinkMultiPtr(in *FoldersRequestsGetLinkMulti) *pb.FoldersRequests_GetLinkMulti {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromFoldersRequestsGetLinkMulti(value)
}

func convertFromFoldersRequestsGetMulti(in FoldersRequestsGetMulti) *pb.FoldersRequests_GetMulti {
	return &pb.FoldersRequests_GetMulti{
		FolderIds: in.FolderIds,
	}
}

func convertFromFoldersRequestsGetMultiPtr(in *FoldersRequestsGetMulti) *pb.FoldersRequests_GetMulti {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromFoldersRequestsGetMulti(value)
}

func convertFromFoldersRequestsSet(in FoldersRequestsSet) *pb.FoldersRequests_Set {
	return &pb.FoldersRequests_Set{
		Details:    convertFromSettableFolderDetails(in.Details),
		FolderId:   in.FolderId,
		UpdateMask: getFieldMask(in),
	}
}

func convertFromFoldersRequestsSetPtr(in *FoldersRequestsSet) *pb.FoldersRequests_Set {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromFoldersRequestsSet(value)
}

func convertToFoldersResponsesGetChildFiles(in *pb.FoldersResponses_GetChildFiles) FoldersResponsesGetChildFiles {
	return FoldersResponsesGetChildFiles{
		Response: convertToFileDetailsSlice(in.Response),
	}
}

func convertToFoldersResponsesGetChildFilesPtr(in *pb.FoldersResponses_GetChildFiles) *FoldersResponsesGetChildFiles {
	if in == nil {
		return nil
	}
	value := convertToFoldersResponsesGetChildFiles(in)
	return &value
}

func convertToFoldersResponsesGetChildFolders(in *pb.FoldersResponses_GetChildFolders) FoldersResponsesGetChildFolders {
	return FoldersResponsesGetChildFolders{
		Response: convertToFolderDetailsSlice(in.Response),
	}
}

func convertToFoldersResponsesGetChildFoldersPtr(in *pb.FoldersResponses_GetChildFolders) *FoldersResponsesGetChildFolders {
	if in == nil {
		return nil
	}
	value := convertToFoldersResponsesGetChildFolders(in)
	return &value
}

func convertToFoldersResponsesGetLinkMulti(in *pb.FoldersResponses_GetLinkMulti) FoldersResponsesGetLinkMulti {
	return FoldersResponsesGetLinkMulti{
		Response: convertToFolderLinkDetailsSlice(in.Response),
	}
}

func convertToFoldersResponsesGetLinkMultiPtr(in *pb.FoldersResponses_GetLinkMulti) *FoldersResponsesGetLinkMulti {
	if in == nil {
		return nil
	}
	value := convertToFoldersResponsesGetLinkMulti(in)
	return &value
}

func convertToFoldersResponsesGetMulti(in *pb.FoldersResponses_GetMulti) FoldersResponsesGetMulti {
	return FoldersResponsesGetMulti{
		Response: convertToFolderDetailsSlice(in.Response),
	}
}

func convertToFoldersResponsesGetMultiPtr(in *pb.FoldersResponses_GetMulti) *FoldersResponsesGetMulti {
	if in == nil {
		return nil
	}
	value := convertToFoldersResponsesGetMulti(in)
	return &value
}

func convertToSettableFolderDetails(in *pb.SettableFolderDetails) SettableFolderDetails {
	return SettableFolderDetails{
		Description: convertToStringPtr(in.Description),
		Name:        convertToStringPtr(in.Name),
	}
}

func convertFromSettableFolderDetails(in SettableFolderDetails) *pb.SettableFolderDetails {
	return &pb.SettableFolderDetails{
		Description: convertFromStringPtr(in.Description),
		Name:        convertFromStringPtr(in.Name),
	}
}

func convertToField(in *pb.Field) Field {
	return Field{
		DataType:           in.DataType,
		Description:        in.Description,
		DisplayType:        in.DisplayType,
		ReadOnly:           in.ReadOnly,
		SearchSettings:     convertToFieldSearchSettings(in.SearchSettings),
		SupportedFileTypes: in.SupportedFileTypes,
		TagName:            in.TagName,
	}
}

func convertFromField(in Field) *pb.Field {
	return &pb.Field{
		DataType:           in.DataType,
		Description:        in.Description,
		DisplayType:        in.DisplayType,
		ReadOnly:           in.ReadOnly,
		SearchSettings:     convertFromFieldSearchSettings(in.SearchSettings),
		SupportedFileTypes: in.SupportedFileTypes,
		TagName:            in.TagName,
	}
}

func convertToFieldPtr(in *pb.Field) *Field {
	if in == nil {
		return nil
	}
	value := convertToField(in)
	return &value
}

func convertToFieldSlice(in []*pb.Field) []Field {
	s := make([]Field, len(in))
	for i := range in {
		s[i] = convertToField(in[i])
	}
	return s
}

func convertFromFieldSlice(in []Field) []*pb.Field {
	s := make([]*pb.Field, len(in))
	for i := range in {
		s[i] = convertFromField(in[i])
	}
	return s
}

func convertToFieldSearchSettings(in *pb.FieldSearchSettings) FieldSearchSettings {
	return FieldSearchSettings{
		AdvancedSearch: in.AdvancedSearch,
		GeneralSearch:  in.GeneralSearch,
	}
}

func convertFromFieldSearchSettings(in FieldSearchSettings) *pb.FieldSearchSettings {
	return &pb.FieldSearchSettings{
		AdvancedSearch: in.AdvancedSearch,
		GeneralSearch:  in.GeneralSearch,
	}
}

func convertFromMetadataRequestsDeleteMetadataValues(in MetadataRequestsDeleteMetadataValues) *pb.MetadataRequests_DeleteMetadataValues {
	return &pb.MetadataRequests_DeleteMetadataValues{
		Details: in.Details.getProto(),
		ItemId:  in.ItemId,
		TagName: in.TagName,
	}
}

func convertFromMetadataRequestsDeleteMetadataValuesPtr(in *MetadataRequestsDeleteMetadataValues) *pb.MetadataRequests_DeleteMetadataValues {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromMetadataRequestsDeleteMetadataValues(value)
}

func convertFromMetadataRequestsGetField(in MetadataRequestsGetField) *pb.MetadataRequests_GetField {
	return &pb.MetadataRequests_GetField{
		TagName: in.TagName,
	}
}

func convertFromMetadataRequestsGetFieldPtr(in *MetadataRequestsGetField) *pb.MetadataRequests_GetField {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromMetadataRequestsGetField(value)
}

func convertFromMetadataRequestsGetFieldMulti(in MetadataRequestsGetFieldMulti) *pb.MetadataRequests_GetFieldMulti {
	return &pb.MetadataRequests_GetFieldMulti{
		TagNames: in.TagNames,
	}
}

func convertFromMetadataRequestsGetFieldMultiPtr(in *MetadataRequestsGetFieldMulti) *pb.MetadataRequests_GetFieldMulti {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromMetadataRequestsGetFieldMulti(value)
}

func convertFromMetadataRequestsGetMetadataFields(in MetadataRequestsGetMetadataFields) *pb.MetadataRequests_GetMetadataFields {
	return &pb.MetadataRequests_GetMetadataFields{
		ItemId: in.ItemId,
	}
}

func convertFromMetadataRequestsGetMetadataFieldsPtr(in *MetadataRequestsGetMetadataFields) *pb.MetadataRequests_GetMetadataFields {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromMetadataRequestsGetMetadataFields(value)
}

func convertFromMetadataRequestsGetMetadataPanels(in MetadataRequestsGetMetadataPanels) *pb.MetadataRequests_GetMetadataPanels {
	return &pb.MetadataRequests_GetMetadataPanels{
		ItemId: in.ItemId,
	}
}

func convertFromMetadataRequestsGetMetadataPanelsPtr(in *MetadataRequestsGetMetadataPanels) *pb.MetadataRequests_GetMetadataPanels {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromMetadataRequestsGetMetadataPanels(value)
}

func convertFromMetadataRequestsGetMetadataValues(in MetadataRequestsGetMetadataValues) *pb.MetadataRequests_GetMetadataValues {
	return &pb.MetadataRequests_GetMetadataValues{
		ItemId:  in.ItemId,
		TagName: in.TagName,
	}
}

func convertFromMetadataRequestsGetMetadataValuesPtr(in *MetadataRequestsGetMetadataValues) *pb.MetadataRequests_GetMetadataValues {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromMetadataRequestsGetMetadataValues(value)
}

func convertFromMetadataRequestsGetPanel(in MetadataRequestsGetPanel) *pb.MetadataRequests_GetPanel {
	return &pb.MetadataRequests_GetPanel{
		PanelId: in.PanelId,
	}
}

func convertFromMetadataRequestsGetPanelPtr(in *MetadataRequestsGetPanel) *pb.MetadataRequests_GetPanel {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromMetadataRequestsGetPanel(value)
}

func convertFromMetadataRequestsGetPanelMulti(in MetadataRequestsGetPanelMulti) *pb.MetadataRequests_GetPanelMulti {
	return &pb.MetadataRequests_GetPanelMulti{
		PanelIds: in.PanelIds,
	}
}

func convertFromMetadataRequestsGetPanelMultiPtr(in *MetadataRequestsGetPanelMulti) *pb.MetadataRequests_GetPanelMulti {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromMetadataRequestsGetPanelMulti(value)
}

func convertFromMetadataRequestsGetProfileVocab(in MetadataRequestsGetProfileVocab) *pb.MetadataRequests_GetProfileVocab {
	return &pb.MetadataRequests_GetProfileVocab{
		TagName: in.TagName,
	}
}

func convertFromMetadataRequestsGetProfileVocabPtr(in *MetadataRequestsGetProfileVocab) *pb.MetadataRequests_GetProfileVocab {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromMetadataRequestsGetProfileVocab(value)
}

func convertFromMetadataRequestsGetSiteVocab(in MetadataRequestsGetSiteVocab) *pb.MetadataRequests_GetSiteVocab {
	return &pb.MetadataRequests_GetSiteVocab{
		TagName: in.TagName,
	}
}

func convertFromMetadataRequestsGetSiteVocabPtr(in *MetadataRequestsGetSiteVocab) *pb.MetadataRequests_GetSiteVocab {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromMetadataRequestsGetSiteVocab(value)
}

func convertFromMetadataRequestsGetSpaceVocab(in MetadataRequestsGetSpaceVocab) *pb.MetadataRequests_GetSpaceVocab {
	return &pb.MetadataRequests_GetSpaceVocab{
		SpaceId: in.SpaceId,
		TagName: in.TagName,
	}
}

func convertFromMetadataRequestsGetSpaceVocabPtr(in *MetadataRequestsGetSpaceVocab) *pb.MetadataRequests_GetSpaceVocab {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromMetadataRequestsGetSpaceVocab(value)
}

func convertFromMetadataRequestsReplaceMetadataValues(in MetadataRequestsReplaceMetadataValues) *pb.MetadataRequests_ReplaceMetadataValues {
	return &pb.MetadataRequests_ReplaceMetadataValues{
		Details: in.Details.getProto(),
		ItemId:  in.ItemId,
		TagName: in.TagName,
	}
}

func convertFromMetadataRequestsReplaceMetadataValuesPtr(in *MetadataRequestsReplaceMetadataValues) *pb.MetadataRequests_ReplaceMetadataValues {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromMetadataRequestsReplaceMetadataValues(value)
}

func convertFromMetadataRequestsReplaceProfileVocab(in MetadataRequestsReplaceProfileVocab) *pb.MetadataRequests_ReplaceProfileVocab {
	return &pb.MetadataRequests_ReplaceProfileVocab{
		Details: convertFromVocab(in.Details),
		TagName: in.TagName,
	}
}

func convertFromMetadataRequestsReplaceProfileVocabPtr(in *MetadataRequestsReplaceProfileVocab) *pb.MetadataRequests_ReplaceProfileVocab {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromMetadataRequestsReplaceProfileVocab(value)
}

func convertFromMetadataRequestsReplaceSiteVocab(in MetadataRequestsReplaceSiteVocab) *pb.MetadataRequests_ReplaceSiteVocab {
	return &pb.MetadataRequests_ReplaceSiteVocab{
		Details: convertFromVocab(in.Details),
		TagName: in.TagName,
	}
}

func convertFromMetadataRequestsReplaceSiteVocabPtr(in *MetadataRequestsReplaceSiteVocab) *pb.MetadataRequests_ReplaceSiteVocab {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromMetadataRequestsReplaceSiteVocab(value)
}

func convertFromMetadataRequestsReplaceSpaceVocab(in MetadataRequestsReplaceSpaceVocab) *pb.MetadataRequests_ReplaceSpaceVocab {
	return &pb.MetadataRequests_ReplaceSpaceVocab{
		Details: convertFromVocab(in.Details),
		SpaceId: in.SpaceId,
		TagName: in.TagName,
	}
}

func convertFromMetadataRequestsReplaceSpaceVocabPtr(in *MetadataRequestsReplaceSpaceVocab) *pb.MetadataRequests_ReplaceSpaceVocab {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromMetadataRequestsReplaceSpaceVocab(value)
}

func convertFromMetadataRequestsUpdateMetadataValues(in MetadataRequestsUpdateMetadataValues) *pb.MetadataRequests_UpdateMetadataValues {
	return &pb.MetadataRequests_UpdateMetadataValues{
		Details: in.Details.getProto(),
		ItemId:  in.ItemId,
		TagName: in.TagName,
	}
}

func convertFromMetadataRequestsUpdateMetadataValuesPtr(in *MetadataRequestsUpdateMetadataValues) *pb.MetadataRequests_UpdateMetadataValues {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromMetadataRequestsUpdateMetadataValues(value)
}

func convertFromMetadataRequestsUpdateProfileVocab(in MetadataRequestsUpdateProfileVocab) *pb.MetadataRequests_UpdateProfileVocab {
	return &pb.MetadataRequests_UpdateProfileVocab{
		Details: convertFromVocab(in.Details),
		TagName: in.TagName,
	}
}

func convertFromMetadataRequestsUpdateProfileVocabPtr(in *MetadataRequestsUpdateProfileVocab) *pb.MetadataRequests_UpdateProfileVocab {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromMetadataRequestsUpdateProfileVocab(value)
}

func convertFromMetadataRequestsUpdateSiteVocab(in MetadataRequestsUpdateSiteVocab) *pb.MetadataRequests_UpdateSiteVocab {
	return &pb.MetadataRequests_UpdateSiteVocab{
		Details: convertFromVocab(in.Details),
		TagName: in.TagName,
	}
}

func convertFromMetadataRequestsUpdateSiteVocabPtr(in *MetadataRequestsUpdateSiteVocab) *pb.MetadataRequests_UpdateSiteVocab {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromMetadataRequestsUpdateSiteVocab(value)
}

func convertFromMetadataRequestsUpdateSpaceVocab(in MetadataRequestsUpdateSpaceVocab) *pb.MetadataRequests_UpdateSpaceVocab {
	return &pb.MetadataRequests_UpdateSpaceVocab{
		Details: convertFromVocab(in.Details),
		SpaceId: in.SpaceId,
		TagName: in.TagName,
	}
}

func convertFromMetadataRequestsUpdateSpaceVocabPtr(in *MetadataRequestsUpdateSpaceVocab) *pb.MetadataRequests_UpdateSpaceVocab {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromMetadataRequestsUpdateSpaceVocab(value)
}

func convertToMetadataResponsesGetFieldMulti(in *pb.MetadataResponses_GetFieldMulti) MetadataResponsesGetFieldMulti {
	return MetadataResponsesGetFieldMulti{
		Response: convertToFieldSlice(in.Response),
	}
}

func convertToMetadataResponsesGetFieldMultiPtr(in *pb.MetadataResponses_GetFieldMulti) *MetadataResponsesGetFieldMulti {
	if in == nil {
		return nil
	}
	value := convertToMetadataResponsesGetFieldMulti(in)
	return &value
}

func convertToMetadataResponsesGetMetadataFields(in *pb.MetadataResponses_GetMetadataFields) MetadataResponsesGetMetadataFields {
	return MetadataResponsesGetMetadataFields{
		Response: in.Response,
	}
}

func convertToMetadataResponsesGetMetadataFieldsPtr(in *pb.MetadataResponses_GetMetadataFields) *MetadataResponsesGetMetadataFields {
	if in == nil {
		return nil
	}
	value := convertToMetadataResponsesGetMetadataFields(in)
	return &value
}

func convertToMetadataResponsesGetMetadataPanels(in *pb.MetadataResponses_GetMetadataPanels) MetadataResponsesGetMetadataPanels {
	return MetadataResponsesGetMetadataPanels{
		Response: in.Response,
	}
}

func convertToMetadataResponsesGetMetadataPanelsPtr(in *pb.MetadataResponses_GetMetadataPanels) *MetadataResponsesGetMetadataPanels {
	if in == nil {
		return nil
	}
	value := convertToMetadataResponsesGetMetadataPanels(in)
	return &value
}

func convertToMetadataResponsesGetPanelMulti(in *pb.MetadataResponses_GetPanelMulti) MetadataResponsesGetPanelMulti {
	return MetadataResponsesGetPanelMulti{
		Response: convertToPanelSlice(in.Response),
	}
}

func convertToMetadataResponsesGetPanelMultiPtr(in *pb.MetadataResponses_GetPanelMulti) *MetadataResponsesGetPanelMulti {
	if in == nil {
		return nil
	}
	value := convertToMetadataResponsesGetPanelMulti(in)
	return &value
}

func convertFromMetadataValue(in MetadataValue) *pb.MetadataValue {
	return in.getProto()
}

func convertToMetadataValue(in *pb.MetadataValue) MetadataValue {
	switch in.Values.(type) {
	case *pb.MetadataValue_Numeric:
		v := in.Values.(*pb.MetadataValue_Numeric)
		return MetadataValueNumericType(v.Numeric.Values)
	case *pb.MetadataValue_Text:
		v := in.Values.(*pb.MetadataValue_Text)
		return MetadataValueTextType(v.Text.Values)
	default:
		panic(fmt.Errorf("Unknown type '%T' for pb.MetadataValue", in.Values))
	}
}

func convertToPanel(in *pb.Panel) Panel {
	return Panel{
		Description: in.Description,
		PanelId:     in.PanelId,
		TagNames:    in.TagNames,
	}
}

func convertFromPanel(in Panel) *pb.Panel {
	return &pb.Panel{
		Description: in.Description,
		PanelId:     in.PanelId,
		TagNames:    in.TagNames,
	}
}

func convertToPanelPtr(in *pb.Panel) *Panel {
	if in == nil {
		return nil
	}
	value := convertToPanel(in)
	return &value
}

func convertToPanelSlice(in []*pb.Panel) []Panel {
	s := make([]Panel, len(in))
	for i := range in {
		s[i] = convertToPanel(in[i])
	}
	return s
}

func convertFromPanelSlice(in []Panel) []*pb.Panel {
	s := make([]*pb.Panel, len(in))
	for i := range in {
		s[i] = convertFromPanel(in[i])
	}
	return s
}

func convertToVocab(in *pb.Vocab) Vocab {
	return Vocab{
		Mode:    vocabmode.VocabMode(in.Mode),
		TagName: in.TagName,
		Values:  in.Values,
	}
}

func convertFromVocab(in Vocab) *pb.Vocab {
	return &pb.Vocab{
		Mode:    pb.Vocab_Mode(in.Mode),
		TagName: in.TagName,
		Values:  in.Values,
	}
}

func convertToVocabPtr(in *pb.Vocab) *Vocab {
	if in == nil {
		return nil
	}
	value := convertToVocab(in)
	return &value
}

func convertToSharedLinkDetails(in *pb.SharedLinkDetails) SharedLinkDetails {
	return SharedLinkDetails{
		CreatedDate:  convertToDateTime(in.CreatedDate),
		Description:  in.Description,
		FolderId:     in.FolderId,
		Id:           in.Id,
		ModifiedDate: convertToDateTime(in.ModifiedDate),
		Name:         in.Name,
		Urls:         in.Urls,
	}
}

func convertFromSharedLinkDetails(in SharedLinkDetails) *pb.SharedLinkDetails {
	return &pb.SharedLinkDetails{
		CreatedDate:  convertFromDateTime(in.CreatedDate),
		Description:  in.Description,
		FolderId:     in.FolderId,
		Id:           in.Id,
		ModifiedDate: convertFromDateTime(in.ModifiedDate),
		Name:         in.Name,
		Urls:         in.Urls,
	}
}

func convertToSharedLinkDetailsPtr(in *pb.SharedLinkDetails) *SharedLinkDetails {
	if in == nil {
		return nil
	}
	value := convertToSharedLinkDetails(in)
	return &value
}

func convertToSharedLinkDetailsSlice(in []*pb.SharedLinkDetails) []SharedLinkDetails {
	s := make([]SharedLinkDetails, len(in))
	for i := range in {
		s[i] = convertToSharedLinkDetails(in[i])
	}
	return s
}

func convertFromSharedLinkDetailsSlice(in []SharedLinkDetails) []*pb.SharedLinkDetails {
	s := make([]*pb.SharedLinkDetails, len(in))
	for i := range in {
		s[i] = convertFromSharedLinkDetails(in[i])
	}
	return s
}

func convertFromSharedLinksRequestsGet(in SharedLinksRequestsGet) *pb.SharedLinksRequests_Get {
	return &pb.SharedLinksRequests_Get{
		SharedLinkId: in.SharedLinkId,
	}
}

func convertFromSharedLinksRequestsGetPtr(in *SharedLinksRequestsGet) *pb.SharedLinksRequests_Get {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromSharedLinksRequestsGet(value)
}

func convertFromSharedLinksRequestsGetMulti(in SharedLinksRequestsGetMulti) *pb.SharedLinksRequests_GetMulti {
	return &pb.SharedLinksRequests_GetMulti{
		SharedLinkIds: in.SharedLinkIds,
	}
}

func convertFromSharedLinksRequestsGetMultiPtr(in *SharedLinksRequestsGetMulti) *pb.SharedLinksRequests_GetMulti {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromSharedLinksRequestsGetMulti(value)
}

func convertToSharedLinksResponsesGetMulti(in *pb.SharedLinksResponses_GetMulti) SharedLinksResponsesGetMulti {
	return SharedLinksResponsesGetMulti{
		Response: convertToSharedLinkDetailsSlice(in.Response),
	}
}

func convertToSharedLinksResponsesGetMultiPtr(in *pb.SharedLinksResponses_GetMulti) *SharedLinksResponsesGetMulti {
	if in == nil {
		return nil
	}
	value := convertToSharedLinksResponsesGetMulti(in)
	return &value
}

func convertToSpaceDetails(in *pb.SpaceDetails) SpaceDetails {
	return SpaceDetails{
		AvatarUrl:    in.AvatarUrl,
		CreatedDate:  convertToDateTime(in.CreatedDate),
		Description:  in.Description,
		Id:           in.Id,
		ModifiedDate: convertToDateTime(in.ModifiedDate),
		Name:         in.Name,
		ParentId:     convertToStringPtr(in.ParentId),
	}
}

func convertFromSpaceDetails(in SpaceDetails) *pb.SpaceDetails {
	return &pb.SpaceDetails{
		AvatarUrl:    in.AvatarUrl,
		CreatedDate:  convertFromDateTime(in.CreatedDate),
		Description:  in.Description,
		Id:           in.Id,
		ModifiedDate: convertFromDateTime(in.ModifiedDate),
		Name:         in.Name,
		ParentId:     convertFromStringPtr(in.ParentId),
	}
}

func convertToSpaceDetailsPtr(in *pb.SpaceDetails) *SpaceDetails {
	if in == nil {
		return nil
	}
	value := convertToSpaceDetails(in)
	return &value
}

func convertToSpaceDetailsSlice(in []*pb.SpaceDetails) []SpaceDetails {
	s := make([]SpaceDetails, len(in))
	for i := range in {
		s[i] = convertToSpaceDetails(in[i])
	}
	return s
}

func convertFromSpaceDetailsSlice(in []SpaceDetails) []*pb.SpaceDetails {
	s := make([]*pb.SpaceDetails, len(in))
	for i := range in {
		s[i] = convertFromSpaceDetails(in[i])
	}
	return s
}

func convertFromSpacesRequestsGet(in SpacesRequestsGet) *pb.SpacesRequests_Get {
	return &pb.SpacesRequests_Get{
		SpaceId: in.SpaceId,
	}
}

func convertFromSpacesRequestsGetPtr(in *SpacesRequestsGet) *pb.SpacesRequests_Get {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromSpacesRequestsGet(value)
}

func convertFromSpacesRequestsGetMulti(in SpacesRequestsGetMulti) *pb.SpacesRequests_GetMulti {
	return &pb.SpacesRequests_GetMulti{
		SpaceIds: in.SpaceIds,
	}
}

func convertFromSpacesRequestsGetMultiPtr(in *SpacesRequestsGetMulti) *pb.SpacesRequests_GetMulti {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromSpacesRequestsGetMulti(value)
}

func convertToSpacesResponsesGetMulti(in *pb.SpacesResponses_GetMulti) SpacesResponsesGetMulti {
	return SpacesResponsesGetMulti{
		Response: convertToSpaceDetailsSlice(in.Response),
	}
}

func convertToSpacesResponsesGetMultiPtr(in *pb.SpacesResponses_GetMulti) *SpacesResponsesGetMulti {
	if in == nil {
		return nil
	}
	value := convertToSpacesResponsesGetMulti(in)
	return &value
}

func convertToUploadDetails(in *pb.UploadDetails) UploadDetails {
	return UploadDetails{
		CancelledCount:      in.CancelledCount,
		CreatedDate:         convertToDateTime(in.CreatedDate),
		DestinationFolderId: in.DestinationFolderId,
		FailedCount:         in.FailedCount,
		Id:                  in.Id,
		IsAcceptingFiles:    in.IsAcceptingFiles,
		ModifiedDate:        convertToDateTime(in.ModifiedDate),
		PostUrl:             in.PostUrl,
		Status:              uploaddetailsstatus.UploadDetailsStatus(in.Status),
		SuccessCount:        in.SuccessCount,
		TotalCount:          in.TotalCount,
		UploadedIds:         in.UploadedIds,
	}
}

func convertToUploadDetailsPtr(in *pb.UploadDetails) *UploadDetails {
	if in == nil {
		return nil
	}
	value := convertToUploadDetails(in)
	return &value
}

func convertFromUploadsRequestsCreate(in UploadsRequestsCreate) *pb.UploadsRequests_Create {
	return &pb.UploadsRequests_Create{
		DestinationFolderId: in.DestinationFolderId,
	}
}

func convertFromUploadsRequestsCreatePtr(in *UploadsRequestsCreate) *pb.UploadsRequests_Create {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromUploadsRequestsCreate(value)
}

func convertFromUploadsRequestsFinish(in UploadsRequestsFinish) *pb.UploadsRequests_Finish {
	return &pb.UploadsRequests_Finish{
		UploadId: in.UploadId,
	}
}

func convertFromUploadsRequestsFinishPtr(in *UploadsRequestsFinish) *pb.UploadsRequests_Finish {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromUploadsRequestsFinish(value)
}

func convertFromUploadsRequestsGet(in UploadsRequestsGet) *pb.UploadsRequests_Get {
	return &pb.UploadsRequests_Get{
		UploadId: in.UploadId,
	}
}

func convertFromUploadsRequestsGetPtr(in *UploadsRequestsGet) *pb.UploadsRequests_Get {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromUploadsRequestsGet(value)
}

func convertToCreateUserDetails(in *pb.CreateUserDetails) CreateUserDetails {
	return CreateUserDetails{
		Description:           in.Description,
		Email:                 in.Email,
		HideHomeSpaceShortcut: in.HideHomeSpaceShortcut,
		Name:                  in.Name,
		Username:              in.Username,
	}
}

func convertFromCreateUserDetails(in CreateUserDetails) *pb.CreateUserDetails {
	return &pb.CreateUserDetails{
		Description:           in.Description,
		Email:                 in.Email,
		HideHomeSpaceShortcut: in.HideHomeSpaceShortcut,
		Name:                  in.Name,
		Username:              in.Username,
	}
}

func convertToSettableUserDetails(in *pb.SettableUserDetails) SettableUserDetails {
	return SettableUserDetails{
		Description:           convertToStringPtr(in.Description),
		Email:                 convertToStringPtr(in.Email),
		HideHomeSpaceShortcut: convertToBoolPtr(in.HideHomeSpaceShortcut),
		Name:                  convertToStringPtr(in.Name),
	}
}

func convertFromSettableUserDetails(in SettableUserDetails) *pb.SettableUserDetails {
	return &pb.SettableUserDetails{
		Description:           convertFromStringPtr(in.Description),
		Email:                 convertFromStringPtr(in.Email),
		HideHomeSpaceShortcut: convertFromBoolPtr(in.HideHomeSpaceShortcut),
		Name:                  convertFromStringPtr(in.Name),
	}
}

func convertToUserDetails(in *pb.UserDetails) UserDetails {
	return UserDetails{
		AvatarUrl:             in.AvatarUrl,
		BackingFolderId:       convertToStringPtr(in.BackingFolderId),
		CreatedDate:           convertToDateTime(in.CreatedDate),
		Description:           in.Description,
		Email:                 convertToStringPtr(in.Email),
		HideHomeSpaceShortcut: convertToBoolPtr(in.HideHomeSpaceShortcut),
		HomeSpaceId:           convertToStringPtr(in.HomeSpaceId),
		Id:                    in.Id,
		ModifiedDate:          convertToDateTime(in.ModifiedDate),
		Name:                  in.Name,
		Username:              in.Username,
	}
}

func convertFromUserDetails(in UserDetails) *pb.UserDetails {
	return &pb.UserDetails{
		AvatarUrl:             in.AvatarUrl,
		BackingFolderId:       convertFromStringPtr(in.BackingFolderId),
		CreatedDate:           convertFromDateTime(in.CreatedDate),
		Description:           in.Description,
		Email:                 convertFromStringPtr(in.Email),
		HideHomeSpaceShortcut: convertFromBoolPtr(in.HideHomeSpaceShortcut),
		HomeSpaceId:           convertFromStringPtr(in.HomeSpaceId),
		Id:                    in.Id,
		ModifiedDate:          convertFromDateTime(in.ModifiedDate),
		Name:                  in.Name,
		Username:              in.Username,
	}
}

func convertToUserDetailsPtr(in *pb.UserDetails) *UserDetails {
	if in == nil {
		return nil
	}
	value := convertToUserDetails(in)
	return &value
}

func convertToUserDetailsSlice(in []*pb.UserDetails) []UserDetails {
	s := make([]UserDetails, len(in))
	for i := range in {
		s[i] = convertToUserDetails(in[i])
	}
	return s
}

func convertFromUserDetailsSlice(in []UserDetails) []*pb.UserDetails {
	s := make([]*pb.UserDetails, len(in))
	for i := range in {
		s[i] = convertFromUserDetails(in[i])
	}
	return s
}

func convertFromUsersRequestsCreateSiteUser(in UsersRequestsCreateSiteUser) *pb.UsersRequests_CreateSiteUser {
	return &pb.UsersRequests_CreateSiteUser{
		Details: convertFromCreateUserDetails(in.Details),
	}
}

func convertFromUsersRequestsCreateSiteUserPtr(in *UsersRequestsCreateSiteUser) *pb.UsersRequests_CreateSiteUser {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromUsersRequestsCreateSiteUser(value)
}

func convertFromUsersRequestsCreateSpacesUser(in UsersRequestsCreateSpacesUser) *pb.UsersRequests_CreateSpacesUser {
	return &pb.UsersRequests_CreateSpacesUser{
		Details: convertFromCreateUserDetails(in.Details),
		SpaceId: in.SpaceId,
	}
}

func convertFromUsersRequestsCreateSpacesUserPtr(in *UsersRequestsCreateSpacesUser) *pb.UsersRequests_CreateSpacesUser {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromUsersRequestsCreateSpacesUser(value)
}

func convertFromUsersRequestsCurrent(in UsersRequestsCurrent) *pb.UsersRequests_Current {
	return &pb.UsersRequests_Current{}
}

func convertFromUsersRequestsCurrentPtr(in *UsersRequestsCurrent) *pb.UsersRequests_Current {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromUsersRequestsCurrent(value)
}

func convertFromUsersRequestsDelete(in UsersRequestsDelete) *pb.UsersRequests_Delete {
	return &pb.UsersRequests_Delete{
		UserId: in.UserId,
	}
}

func convertFromUsersRequestsDeletePtr(in *UsersRequestsDelete) *pb.UsersRequests_Delete {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromUsersRequestsDelete(value)
}

func convertFromUsersRequestsGet(in UsersRequestsGet) *pb.UsersRequests_Get {
	return &pb.UsersRequests_Get{
		UserId: in.UserId,
	}
}

func convertFromUsersRequestsGetPtr(in *UsersRequestsGet) *pb.UsersRequests_Get {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromUsersRequestsGet(value)
}

func convertFromUsersRequestsGetMulti(in UsersRequestsGetMulti) *pb.UsersRequests_GetMulti {
	return &pb.UsersRequests_GetMulti{
		UserIds: in.UserIds,
	}
}

func convertFromUsersRequestsGetMultiPtr(in *UsersRequestsGetMulti) *pb.UsersRequests_GetMulti {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromUsersRequestsGetMulti(value)
}

func convertFromUsersRequestsGetSiteUsers(in UsersRequestsGetSiteUsers) *pb.UsersRequests_GetSiteUsers {
	return &pb.UsersRequests_GetSiteUsers{}
}

func convertFromUsersRequestsGetSiteUsersPtr(in *UsersRequestsGetSiteUsers) *pb.UsersRequests_GetSiteUsers {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromUsersRequestsGetSiteUsers(value)
}

func convertFromUsersRequestsGetSpacesUsers(in UsersRequestsGetSpacesUsers) *pb.UsersRequests_GetSpacesUsers {
	return &pb.UsersRequests_GetSpacesUsers{
		SpaceId: in.SpaceId,
	}
}

func convertFromUsersRequestsGetSpacesUsersPtr(in *UsersRequestsGetSpacesUsers) *pb.UsersRequests_GetSpacesUsers {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromUsersRequestsGetSpacesUsers(value)
}

func convertFromUsersRequestsSet(in UsersRequestsSet) *pb.UsersRequests_Set {
	return &pb.UsersRequests_Set{
		Details:    convertFromSettableUserDetails(in.Details),
		UpdateMask: getFieldMask(in),
		UserId:     in.UserId,
	}
}

func convertFromUsersRequestsSetPtr(in *UsersRequestsSet) *pb.UsersRequests_Set {
	if in == nil {
		return nil
	}
	value := *in
	return convertFromUsersRequestsSet(value)
}

func convertToUsersResponsesDelete(in *pb.UsersResponses_Delete) UsersResponsesDelete {
	return UsersResponsesDelete{}
}

func convertToUsersResponsesDeletePtr(in *pb.UsersResponses_Delete) *UsersResponsesDelete {
	if in == nil {
		return nil
	}
	value := convertToUsersResponsesDelete(in)
	return &value
}

func convertToUsersResponsesGetMulti(in *pb.UsersResponses_GetMulti) UsersResponsesGetMulti {
	return UsersResponsesGetMulti{
		Response: convertToUserDetailsSlice(in.Response),
	}
}

func convertToUsersResponsesGetMultiPtr(in *pb.UsersResponses_GetMulti) *UsersResponsesGetMulti {
	if in == nil {
		return nil
	}
	value := convertToUsersResponsesGetMulti(in)
	return &value
}

func convertToUsersResponsesGetSiteUsers(in *pb.UsersResponses_GetSiteUsers) UsersResponsesGetSiteUsers {
	return UsersResponsesGetSiteUsers{
		Response: in.Response,
	}
}

func convertToUsersResponsesGetSiteUsersPtr(in *pb.UsersResponses_GetSiteUsers) *UsersResponsesGetSiteUsers {
	if in == nil {
		return nil
	}
	value := convertToUsersResponsesGetSiteUsers(in)
	return &value
}

func convertToUsersResponsesGetSpacesUsers(in *pb.UsersResponses_GetSpacesUsers) UsersResponsesGetSpacesUsers {
	return UsersResponsesGetSpacesUsers{
		Response: in.Response,
	}
}

func convertToUsersResponsesGetSpacesUsersPtr(in *pb.UsersResponses_GetSpacesUsers) *UsersResponsesGetSpacesUsers {
	if in == nil {
		return nil
	}
	value := convertToUsersResponsesGetSpacesUsers(in)
	return &value
}
