package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetIGUserAgenciesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserAgenciesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGUserAgenciesBatchCall(id string, params GetIGUserAgenciesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "agencies"), params.ToParams(), options...)
}

func NewGetIGUserAgenciesBatchRequest(id string, params GetIGUserAgenciesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Business]] {
	return core.NewBatchRequest[core.Cursor[objects.Business]](GetIGUserAgenciesBatchCall(id, params, options...))
}

func DecodeGetIGUserAgenciesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Business], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Business]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserAgenciesWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserAgenciesParams) (*core.Cursor[objects.Business], *core.Response, error) {
	var out core.Cursor[objects.Business]
	call := GetIGUserAgenciesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserAgencies(ctx context.Context, client *core.Client, id string, params GetIGUserAgenciesParams) (*core.Cursor[objects.Business], error) {
	out, _, err := GetIGUserAgenciesWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserAuthorizedAdaccountsParams struct {
	Business string      `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params GetIGUserAuthorizedAdaccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	return out
}

func GetIGUserAuthorizedAdaccountsBatchCall(id string, params GetIGUserAuthorizedAdaccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "authorized_adaccounts"), params.ToParams(), options...)
}

func NewGetIGUserAuthorizedAdaccountsBatchRequest(id string, params GetIGUserAuthorizedAdaccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccount]](GetIGUserAuthorizedAdaccountsBatchCall(id, params, options...))
}

func DecodeGetIGUserAuthorizedAdaccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccount], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccount]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserAuthorizedAdaccountsWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserAuthorizedAdaccountsParams) (*core.Cursor[objects.AdAccount], *core.Response, error) {
	var out core.Cursor[objects.AdAccount]
	call := GetIGUserAuthorizedAdaccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserAuthorizedAdaccounts(ctx context.Context, client *core.Client, id string, params GetIGUserAuthorizedAdaccountsParams) (*core.Cursor[objects.AdAccount], error) {
	out, _, err := GetIGUserAuthorizedAdaccountsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateIGUserAuthorizedAdaccountsParams struct {
	AccountID core.ID     `facebook:"account_id"`
	Business  string      `facebook:"business"`
	Extra     core.Params `facebook:"-"`
}

func (params CreateIGUserAuthorizedAdaccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["account_id"] = params.AccountID
	out["business"] = params.Business
	return out
}

func CreateIGUserAuthorizedAdaccountsBatchCall(id string, params CreateIGUserAuthorizedAdaccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "authorized_adaccounts"), params.ToParams(), options...)
}

func NewCreateIGUserAuthorizedAdaccountsBatchRequest(id string, params CreateIGUserAuthorizedAdaccountsParams, options ...core.BatchOption) *core.BatchRequest[objects.IGUser] {
	return core.NewBatchRequest[objects.IGUser](CreateIGUserAuthorizedAdaccountsBatchCall(id, params, options...))
}

func DecodeCreateIGUserAuthorizedAdaccountsBatchResponse(response *core.BatchResponse) (*objects.IGUser, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.IGUser
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateIGUserAuthorizedAdaccountsWithResponse(ctx context.Context, client *core.Client, id string, params CreateIGUserAuthorizedAdaccountsParams) (*objects.IGUser, *core.Response, error) {
	var out objects.IGUser
	call := CreateIGUserAuthorizedAdaccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateIGUserAuthorizedAdaccounts(ctx context.Context, client *core.Client, id string, params CreateIGUserAuthorizedAdaccountsParams) (*objects.IGUser, error) {
	out, _, err := CreateIGUserAuthorizedAdaccountsWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserAvailableCatalogsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserAvailableCatalogsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGUserAvailableCatalogsBatchCall(id string, params GetIGUserAvailableCatalogsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "available_catalogs"), params.ToParams(), options...)
}

func NewGetIGUserAvailableCatalogsBatchRequest(id string, params GetIGUserAvailableCatalogsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.UserAvailableCatalogs]] {
	return core.NewBatchRequest[core.Cursor[objects.UserAvailableCatalogs]](GetIGUserAvailableCatalogsBatchCall(id, params, options...))
}

func DecodeGetIGUserAvailableCatalogsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.UserAvailableCatalogs], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.UserAvailableCatalogs]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserAvailableCatalogsWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserAvailableCatalogsParams) (*core.Cursor[objects.UserAvailableCatalogs], *core.Response, error) {
	var out core.Cursor[objects.UserAvailableCatalogs]
	call := GetIGUserAvailableCatalogsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserAvailableCatalogs(ctx context.Context, client *core.Client, id string, params GetIGUserAvailableCatalogsParams) (*core.Cursor[objects.UserAvailableCatalogs], error) {
	out, _, err := GetIGUserAvailableCatalogsWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserBrandedContentAdPermissionsParams struct {
	CreatorUsername *string     `facebook:"creator_username"`
	Extra           core.Params `facebook:"-"`
}

func (params GetIGUserBrandedContentAdPermissionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CreatorUsername != nil {
		out["creator_username"] = *params.CreatorUsername
	}
	return out
}

func GetIGUserBrandedContentAdPermissionsBatchCall(id string, params GetIGUserBrandedContentAdPermissionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "branded_content_ad_permissions"), params.ToParams(), options...)
}

func NewGetIGUserBrandedContentAdPermissionsBatchRequest(id string, params GetIGUserBrandedContentAdPermissionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGBCAdsPermission]] {
	return core.NewBatchRequest[core.Cursor[objects.IGBCAdsPermission]](GetIGUserBrandedContentAdPermissionsBatchCall(id, params, options...))
}

func DecodeGetIGUserBrandedContentAdPermissionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGBCAdsPermission], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGBCAdsPermission]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserBrandedContentAdPermissionsWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserBrandedContentAdPermissionsParams) (*core.Cursor[objects.IGBCAdsPermission], *core.Response, error) {
	var out core.Cursor[objects.IGBCAdsPermission]
	call := GetIGUserBrandedContentAdPermissionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserBrandedContentAdPermissions(ctx context.Context, client *core.Client, id string, params GetIGUserBrandedContentAdPermissionsParams) (*core.Cursor[objects.IGBCAdsPermission], error) {
	out, _, err := GetIGUserBrandedContentAdPermissionsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateIGUserBrandedContentAdPermissionsParams struct {
	CreatorInstagramAccount  *string     `facebook:"creator_instagram_account"`
	CreatorInstagramUsername *string     `facebook:"creator_instagram_username"`
	Revoke                   *bool       `facebook:"revoke"`
	Extra                    core.Params `facebook:"-"`
}

func (params CreateIGUserBrandedContentAdPermissionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CreatorInstagramAccount != nil {
		out["creator_instagram_account"] = *params.CreatorInstagramAccount
	}
	if params.CreatorInstagramUsername != nil {
		out["creator_instagram_username"] = *params.CreatorInstagramUsername
	}
	if params.Revoke != nil {
		out["revoke"] = *params.Revoke
	}
	return out
}

func CreateIGUserBrandedContentAdPermissionsBatchCall(id string, params CreateIGUserBrandedContentAdPermissionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "branded_content_ad_permissions"), params.ToParams(), options...)
}

func NewCreateIGUserBrandedContentAdPermissionsBatchRequest(id string, params CreateIGUserBrandedContentAdPermissionsParams, options ...core.BatchOption) *core.BatchRequest[objects.IGBCAdsPermission] {
	return core.NewBatchRequest[objects.IGBCAdsPermission](CreateIGUserBrandedContentAdPermissionsBatchCall(id, params, options...))
}

func DecodeCreateIGUserBrandedContentAdPermissionsBatchResponse(response *core.BatchResponse) (*objects.IGBCAdsPermission, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.IGBCAdsPermission
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateIGUserBrandedContentAdPermissionsWithResponse(ctx context.Context, client *core.Client, id string, params CreateIGUserBrandedContentAdPermissionsParams) (*objects.IGBCAdsPermission, *core.Response, error) {
	var out objects.IGBCAdsPermission
	call := CreateIGUserBrandedContentAdPermissionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateIGUserBrandedContentAdPermissions(ctx context.Context, client *core.Client, id string, params CreateIGUserBrandedContentAdPermissionsParams) (*objects.IGBCAdsPermission, error) {
	out, _, err := CreateIGUserBrandedContentAdPermissionsWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserBrandedContentAdvertisableMediasParams struct {
	AdCode                      *string                                                                         `facebook:"ad_code"`
	CreatorUsername             *string                                                                         `facebook:"creator_username"`
	MediaRelationship           *[]enums.ShadowiguserbrandedContentAdvertisableMediasMediaRelationshipEnumParam `facebook:"media_relationship"`
	OnlyFetchAllowlisted        *bool                                                                           `facebook:"only_fetch_allowlisted"`
	OnlyFetchRecommendedContent *bool                                                                           `facebook:"only_fetch_recommended_content"`
	Permalinks                  *[]string                                                                       `facebook:"permalinks"`
	Extra                       core.Params                                                                     `facebook:"-"`
}

func (params GetIGUserBrandedContentAdvertisableMediasParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdCode != nil {
		out["ad_code"] = *params.AdCode
	}
	if params.CreatorUsername != nil {
		out["creator_username"] = *params.CreatorUsername
	}
	if params.MediaRelationship != nil {
		out["media_relationship"] = *params.MediaRelationship
	}
	if params.OnlyFetchAllowlisted != nil {
		out["only_fetch_allowlisted"] = *params.OnlyFetchAllowlisted
	}
	if params.OnlyFetchRecommendedContent != nil {
		out["only_fetch_recommended_content"] = *params.OnlyFetchRecommendedContent
	}
	if params.Permalinks != nil {
		out["permalinks"] = *params.Permalinks
	}
	return out
}

func GetIGUserBrandedContentAdvertisableMediasBatchCall(id string, params GetIGUserBrandedContentAdvertisableMediasParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "branded_content_advertisable_medias"), params.ToParams(), options...)
}

func NewGetIGUserBrandedContentAdvertisableMediasBatchRequest(id string, params GetIGUserBrandedContentAdvertisableMediasParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BrandedContentShadowIGMediaID]] {
	return core.NewBatchRequest[core.Cursor[objects.BrandedContentShadowIGMediaID]](GetIGUserBrandedContentAdvertisableMediasBatchCall(id, params, options...))
}

func DecodeGetIGUserBrandedContentAdvertisableMediasBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BrandedContentShadowIGMediaID], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.BrandedContentShadowIGMediaID]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserBrandedContentAdvertisableMediasWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserBrandedContentAdvertisableMediasParams) (*core.Cursor[objects.BrandedContentShadowIGMediaID], *core.Response, error) {
	var out core.Cursor[objects.BrandedContentShadowIGMediaID]
	call := GetIGUserBrandedContentAdvertisableMediasBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserBrandedContentAdvertisableMedias(ctx context.Context, client *core.Client, id string, params GetIGUserBrandedContentAdvertisableMediasParams) (*core.Cursor[objects.BrandedContentShadowIGMediaID], error) {
	out, _, err := GetIGUserBrandedContentAdvertisableMediasWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteIGUserBrandedContentTagApprovalParams struct {
	UserIds []core.ID   `facebook:"user_ids"`
	Extra   core.Params `facebook:"-"`
}

func (params DeleteIGUserBrandedContentTagApprovalParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["user_ids"] = params.UserIds
	return out
}

func DeleteIGUserBrandedContentTagApprovalBatchCall(id string, params DeleteIGUserBrandedContentTagApprovalParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "branded_content_tag_approval"), params.ToParams(), options...)
}

func NewDeleteIGUserBrandedContentTagApprovalBatchRequest(id string, params DeleteIGUserBrandedContentTagApprovalParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteIGUserBrandedContentTagApprovalBatchCall(id, params, options...))
}

func DecodeDeleteIGUserBrandedContentTagApprovalBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func DeleteIGUserBrandedContentTagApprovalWithResponse(ctx context.Context, client *core.Client, id string, params DeleteIGUserBrandedContentTagApprovalParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteIGUserBrandedContentTagApprovalBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteIGUserBrandedContentTagApproval(ctx context.Context, client *core.Client, id string, params DeleteIGUserBrandedContentTagApprovalParams) (*map[string]interface{}, error) {
	out, _, err := DeleteIGUserBrandedContentTagApprovalWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserBrandedContentTagApprovalParams struct {
	UserIds []core.ID   `facebook:"user_ids"`
	Extra   core.Params `facebook:"-"`
}

func (params GetIGUserBrandedContentTagApprovalParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["user_ids"] = params.UserIds
	return out
}

func GetIGUserBrandedContentTagApprovalBatchCall(id string, params GetIGUserBrandedContentTagApprovalParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "branded_content_tag_approval"), params.ToParams(), options...)
}

func NewGetIGUserBrandedContentTagApprovalBatchRequest(id string, params GetIGUserBrandedContentTagApprovalParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BrandedContentShadowIGUserID]] {
	return core.NewBatchRequest[core.Cursor[objects.BrandedContentShadowIGUserID]](GetIGUserBrandedContentTagApprovalBatchCall(id, params, options...))
}

func DecodeGetIGUserBrandedContentTagApprovalBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BrandedContentShadowIGUserID], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.BrandedContentShadowIGUserID]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserBrandedContentTagApprovalWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserBrandedContentTagApprovalParams) (*core.Cursor[objects.BrandedContentShadowIGUserID], *core.Response, error) {
	var out core.Cursor[objects.BrandedContentShadowIGUserID]
	call := GetIGUserBrandedContentTagApprovalBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserBrandedContentTagApproval(ctx context.Context, client *core.Client, id string, params GetIGUserBrandedContentTagApprovalParams) (*core.Cursor[objects.BrandedContentShadowIGUserID], error) {
	out, _, err := GetIGUserBrandedContentTagApprovalWithResponse(ctx, client, id, params)
	return out, err
}

type CreateIGUserBrandedContentTagApprovalParams struct {
	UserIds []core.ID   `facebook:"user_ids"`
	Extra   core.Params `facebook:"-"`
}

func (params CreateIGUserBrandedContentTagApprovalParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["user_ids"] = params.UserIds
	return out
}

func CreateIGUserBrandedContentTagApprovalBatchCall(id string, params CreateIGUserBrandedContentTagApprovalParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "branded_content_tag_approval"), params.ToParams(), options...)
}

func NewCreateIGUserBrandedContentTagApprovalBatchRequest(id string, params CreateIGUserBrandedContentTagApprovalParams, options ...core.BatchOption) *core.BatchRequest[objects.BrandedContentShadowIGUserID] {
	return core.NewBatchRequest[objects.BrandedContentShadowIGUserID](CreateIGUserBrandedContentTagApprovalBatchCall(id, params, options...))
}

func DecodeCreateIGUserBrandedContentTagApprovalBatchResponse(response *core.BatchResponse) (*objects.BrandedContentShadowIGUserID, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BrandedContentShadowIGUserID
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateIGUserBrandedContentTagApprovalWithResponse(ctx context.Context, client *core.Client, id string, params CreateIGUserBrandedContentTagApprovalParams) (*objects.BrandedContentShadowIGUserID, *core.Response, error) {
	var out objects.BrandedContentShadowIGUserID
	call := CreateIGUserBrandedContentTagApprovalBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateIGUserBrandedContentTagApproval(ctx context.Context, client *core.Client, id string, params CreateIGUserBrandedContentTagApprovalParams) (*objects.BrandedContentShadowIGUserID, error) {
	out, _, err := CreateIGUserBrandedContentTagApprovalWithResponse(ctx, client, id, params)
	return out, err
}

type CreateIGUserBusinessMessagingFeatureStatusParams struct {
	Features []map[string]interface{} `facebook:"features"`
	Extra    core.Params              `facebook:"-"`
}

func (params CreateIGUserBusinessMessagingFeatureStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["features"] = params.Features
	return out
}

func CreateIGUserBusinessMessagingFeatureStatusBatchCall(id string, params CreateIGUserBusinessMessagingFeatureStatusParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "business_messaging_feature_status"), params.ToParams(), options...)
}

func NewCreateIGUserBusinessMessagingFeatureStatusBatchRequest(id string, params CreateIGUserBusinessMessagingFeatureStatusParams, options ...core.BatchOption) *core.BatchRequest[objects.IGUser] {
	return core.NewBatchRequest[objects.IGUser](CreateIGUserBusinessMessagingFeatureStatusBatchCall(id, params, options...))
}

func DecodeCreateIGUserBusinessMessagingFeatureStatusBatchResponse(response *core.BatchResponse) (*objects.IGUser, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.IGUser
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateIGUserBusinessMessagingFeatureStatusWithResponse(ctx context.Context, client *core.Client, id string, params CreateIGUserBusinessMessagingFeatureStatusParams) (*objects.IGUser, *core.Response, error) {
	var out objects.IGUser
	call := CreateIGUserBusinessMessagingFeatureStatusBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateIGUserBusinessMessagingFeatureStatus(ctx context.Context, client *core.Client, id string, params CreateIGUserBusinessMessagingFeatureStatusParams) (*objects.IGUser, error) {
	out, _, err := CreateIGUserBusinessMessagingFeatureStatusWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserCatalogProductSearchParams struct {
	CatalogID core.ID     `facebook:"catalog_id"`
	Q         *string     `facebook:"q"`
	Extra     core.Params `facebook:"-"`
}

func (params GetIGUserCatalogProductSearchParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["catalog_id"] = params.CatalogID
	if params.Q != nil {
		out["q"] = *params.Q
	}
	return out
}

func GetIGUserCatalogProductSearchBatchCall(id string, params GetIGUserCatalogProductSearchParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "catalog_product_search"), params.ToParams(), options...)
}

func NewGetIGUserCatalogProductSearchBatchRequest(id string, params GetIGUserCatalogProductSearchParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ShadowIGUserCatalogProductSearch]] {
	return core.NewBatchRequest[core.Cursor[objects.ShadowIGUserCatalogProductSearch]](GetIGUserCatalogProductSearchBatchCall(id, params, options...))
}

func DecodeGetIGUserCatalogProductSearchBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ShadowIGUserCatalogProductSearch], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ShadowIGUserCatalogProductSearch]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserCatalogProductSearchWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserCatalogProductSearchParams) (*core.Cursor[objects.ShadowIGUserCatalogProductSearch], *core.Response, error) {
	var out core.Cursor[objects.ShadowIGUserCatalogProductSearch]
	call := GetIGUserCatalogProductSearchBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserCatalogProductSearch(ctx context.Context, client *core.Client, id string, params GetIGUserCatalogProductSearchParams) (*core.Cursor[objects.ShadowIGUserCatalogProductSearch], error) {
	out, _, err := GetIGUserCatalogProductSearchWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserCollaborationInvitesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserCollaborationInvitesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGUserCollaborationInvitesBatchCall(id string, params GetIGUserCollaborationInvitesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "collaboration_invites"), params.ToParams(), options...)
}

func NewGetIGUserCollaborationInvitesBatchRequest(id string, params GetIGUserCollaborationInvitesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ShadowIGUserCollaborationInvites]] {
	return core.NewBatchRequest[core.Cursor[objects.ShadowIGUserCollaborationInvites]](GetIGUserCollaborationInvitesBatchCall(id, params, options...))
}

func DecodeGetIGUserCollaborationInvitesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ShadowIGUserCollaborationInvites], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ShadowIGUserCollaborationInvites]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserCollaborationInvitesWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserCollaborationInvitesParams) (*core.Cursor[objects.ShadowIGUserCollaborationInvites], *core.Response, error) {
	var out core.Cursor[objects.ShadowIGUserCollaborationInvites]
	call := GetIGUserCollaborationInvitesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserCollaborationInvites(ctx context.Context, client *core.Client, id string, params GetIGUserCollaborationInvitesParams) (*core.Cursor[objects.ShadowIGUserCollaborationInvites], error) {
	out, _, err := GetIGUserCollaborationInvitesWithResponse(ctx, client, id, params)
	return out, err
}

type CreateIGUserCollaborationInvitesParams struct {
	Accept  bool        `facebook:"accept"`
	MediaID core.ID     `facebook:"media_id"`
	Extra   core.Params `facebook:"-"`
}

func (params CreateIGUserCollaborationInvitesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["accept"] = params.Accept
	out["media_id"] = params.MediaID
	return out
}

func CreateIGUserCollaborationInvitesBatchCall(id string, params CreateIGUserCollaborationInvitesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "collaboration_invites"), params.ToParams(), options...)
}

func NewCreateIGUserCollaborationInvitesBatchRequest(id string, params CreateIGUserCollaborationInvitesParams, options ...core.BatchOption) *core.BatchRequest[objects.ShadowIGUserCollaborationInvites] {
	return core.NewBatchRequest[objects.ShadowIGUserCollaborationInvites](CreateIGUserCollaborationInvitesBatchCall(id, params, options...))
}

func DecodeCreateIGUserCollaborationInvitesBatchResponse(response *core.BatchResponse) (*objects.ShadowIGUserCollaborationInvites, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ShadowIGUserCollaborationInvites
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateIGUserCollaborationInvitesWithResponse(ctx context.Context, client *core.Client, id string, params CreateIGUserCollaborationInvitesParams) (*objects.ShadowIGUserCollaborationInvites, *core.Response, error) {
	var out objects.ShadowIGUserCollaborationInvites
	call := CreateIGUserCollaborationInvitesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateIGUserCollaborationInvites(ctx context.Context, client *core.Client, id string, params CreateIGUserCollaborationInvitesParams) (*objects.ShadowIGUserCollaborationInvites, error) {
	out, _, err := CreateIGUserCollaborationInvitesWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserCollaborativeMediaParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserCollaborativeMediaParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGUserCollaborativeMediaBatchCall(id string, params GetIGUserCollaborativeMediaParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "collaborative_media"), params.ToParams(), options...)
}

func NewGetIGUserCollaborativeMediaBatchRequest(id string, params GetIGUserCollaborativeMediaParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ShadowIGUserCollaborativeMedia]] {
	return core.NewBatchRequest[core.Cursor[objects.ShadowIGUserCollaborativeMedia]](GetIGUserCollaborativeMediaBatchCall(id, params, options...))
}

func DecodeGetIGUserCollaborativeMediaBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ShadowIGUserCollaborativeMedia], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ShadowIGUserCollaborativeMedia]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserCollaborativeMediaWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserCollaborativeMediaParams) (*core.Cursor[objects.ShadowIGUserCollaborativeMedia], *core.Response, error) {
	var out core.Cursor[objects.ShadowIGUserCollaborativeMedia]
	call := GetIGUserCollaborativeMediaBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserCollaborativeMedia(ctx context.Context, client *core.Client, id string, params GetIGUserCollaborativeMediaParams) (*core.Cursor[objects.ShadowIGUserCollaborativeMedia], error) {
	out, _, err := GetIGUserCollaborativeMediaWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserConnectedThreadsUserParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserConnectedThreadsUserParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGUserConnectedThreadsUserBatchCall(id string, params GetIGUserConnectedThreadsUserParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "connected_threads_user"), params.ToParams(), options...)
}

func NewGetIGUserConnectedThreadsUserBatchRequest(id string, params GetIGUserConnectedThreadsUserParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ThreadsUser]] {
	return core.NewBatchRequest[core.Cursor[objects.ThreadsUser]](GetIGUserConnectedThreadsUserBatchCall(id, params, options...))
}

func DecodeGetIGUserConnectedThreadsUserBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ThreadsUser], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ThreadsUser]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserConnectedThreadsUserWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserConnectedThreadsUserParams) (*core.Cursor[objects.ThreadsUser], *core.Response, error) {
	var out core.Cursor[objects.ThreadsUser]
	call := GetIGUserConnectedThreadsUserBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserConnectedThreadsUser(ctx context.Context, client *core.Client, id string, params GetIGUserConnectedThreadsUserParams) (*core.Cursor[objects.ThreadsUser], error) {
	out, _, err := GetIGUserConnectedThreadsUserWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserContentPublishingLimitParams struct {
	Since *core.Time  `facebook:"since"`
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserContentPublishingLimitParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	return out
}

func GetIGUserContentPublishingLimitBatchCall(id string, params GetIGUserContentPublishingLimitParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "content_publishing_limit"), params.ToParams(), options...)
}

func NewGetIGUserContentPublishingLimitBatchRequest(id string, params GetIGUserContentPublishingLimitParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ContentPublishingLimitResponse]] {
	return core.NewBatchRequest[core.Cursor[objects.ContentPublishingLimitResponse]](GetIGUserContentPublishingLimitBatchCall(id, params, options...))
}

func DecodeGetIGUserContentPublishingLimitBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ContentPublishingLimitResponse], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ContentPublishingLimitResponse]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserContentPublishingLimitWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserContentPublishingLimitParams) (*core.Cursor[objects.ContentPublishingLimitResponse], *core.Response, error) {
	var out core.Cursor[objects.ContentPublishingLimitResponse]
	call := GetIGUserContentPublishingLimitBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserContentPublishingLimit(ctx context.Context, client *core.Client, id string, params GetIGUserContentPublishingLimitParams) (*core.Cursor[objects.ContentPublishingLimitResponse], error) {
	out, _, err := GetIGUserContentPublishingLimitWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserCreatorMarketplaceCreatorsParams struct {
	CreatorAgeBucket                 *[]map[string]interface{}                                                      `facebook:"creator_age_bucket"`
	CreatorCountries                 *[]enums.ShadowigusercreatorMarketplaceCreatorsCreatorCountriesEnumParam       `facebook:"creator_countries"`
	CreatorFollowerGrowth            *map[string]interface{}                                                        `facebook:"creator_follower_growth"`
	CreatorGender                    *[]enums.ShadowigusercreatorMarketplaceCreatorsCreatorGenderEnumParam          `facebook:"creator_gender"`
	CreatorInterests                 *[]map[string]interface{}                                                      `facebook:"creator_interests"`
	CreatorLatestPostActivity        *map[string]interface{}                                                        `facebook:"creator_latest_post_activity"`
	CreatorMaxEngagedAccounts        *uint64                                                                        `facebook:"creator_max_engaged_accounts"`
	CreatorMaxFollowers              *uint64                                                                        `facebook:"creator_max_followers"`
	CreatorMinEngagedAccounts        *uint64                                                                        `facebook:"creator_min_engaged_accounts"`
	CreatorMinFollowers              *uint64                                                                        `facebook:"creator_min_followers"`
	CreatorStates                    *[]string                                                                      `facebook:"creator_states"`
	CustomAudienceID                 *core.ID                                                                       `facebook:"custom_audience_id"`
	HasPublicContactEmail            *bool                                                                          `facebook:"has_public_contact_email"`
	IsPaidPartnershipMessagesEnabled *bool                                                                          `facebook:"is_paid_partnership_messages_enabled"`
	MajorAudienceAgeBucket           *[]map[string]interface{}                                                      `facebook:"major_audience_age_bucket"`
	MajorAudienceCountries           *[]enums.ShadowigusercreatorMarketplaceCreatorsMajorAudienceCountriesEnumParam `facebook:"major_audience_countries"`
	MajorAudienceDeviceType          *[]map[string]interface{}                                                      `facebook:"major_audience_device_type"`
	MajorAudienceGender              *[]enums.ShadowigusercreatorMarketplaceCreatorsMajorAudienceGenderEnumParam    `facebook:"major_audience_gender"`
	MajorAudienceStates              *[]string                                                                      `facebook:"major_audience_states"`
	Query                            *string                                                                        `facebook:"query"`
	RecommendationType               *enums.ShadowigusercreatorMarketplaceCreatorsRecommendationTypeEnumParam       `facebook:"recommendation_type"`
	ReelsInteractionRate             *map[string]interface{}                                                        `facebook:"reels_interaction_rate"`
	ShowOnboardedCreatorsOnly        *bool                                                                          `facebook:"show_onboarded_creators_only"`
	SimilarToCreators                *[]string                                                                      `facebook:"similar_to_creators"`
	Username                         *string                                                                        `facebook:"username"`
	Usernames                        *[]string                                                                      `facebook:"usernames"`
	Extra                            core.Params                                                                    `facebook:"-"`
}

func (params GetIGUserCreatorMarketplaceCreatorsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CreatorAgeBucket != nil {
		out["creator_age_bucket"] = *params.CreatorAgeBucket
	}
	if params.CreatorCountries != nil {
		out["creator_countries"] = *params.CreatorCountries
	}
	if params.CreatorFollowerGrowth != nil {
		out["creator_follower_growth"] = *params.CreatorFollowerGrowth
	}
	if params.CreatorGender != nil {
		out["creator_gender"] = *params.CreatorGender
	}
	if params.CreatorInterests != nil {
		out["creator_interests"] = *params.CreatorInterests
	}
	if params.CreatorLatestPostActivity != nil {
		out["creator_latest_post_activity"] = *params.CreatorLatestPostActivity
	}
	if params.CreatorMaxEngagedAccounts != nil {
		out["creator_max_engaged_accounts"] = *params.CreatorMaxEngagedAccounts
	}
	if params.CreatorMaxFollowers != nil {
		out["creator_max_followers"] = *params.CreatorMaxFollowers
	}
	if params.CreatorMinEngagedAccounts != nil {
		out["creator_min_engaged_accounts"] = *params.CreatorMinEngagedAccounts
	}
	if params.CreatorMinFollowers != nil {
		out["creator_min_followers"] = *params.CreatorMinFollowers
	}
	if params.CreatorStates != nil {
		out["creator_states"] = *params.CreatorStates
	}
	if params.CustomAudienceID != nil {
		out["custom_audience_id"] = *params.CustomAudienceID
	}
	if params.HasPublicContactEmail != nil {
		out["has_public_contact_email"] = *params.HasPublicContactEmail
	}
	if params.IsPaidPartnershipMessagesEnabled != nil {
		out["is_paid_partnership_messages_enabled"] = *params.IsPaidPartnershipMessagesEnabled
	}
	if params.MajorAudienceAgeBucket != nil {
		out["major_audience_age_bucket"] = *params.MajorAudienceAgeBucket
	}
	if params.MajorAudienceCountries != nil {
		out["major_audience_countries"] = *params.MajorAudienceCountries
	}
	if params.MajorAudienceDeviceType != nil {
		out["major_audience_device_type"] = *params.MajorAudienceDeviceType
	}
	if params.MajorAudienceGender != nil {
		out["major_audience_gender"] = *params.MajorAudienceGender
	}
	if params.MajorAudienceStates != nil {
		out["major_audience_states"] = *params.MajorAudienceStates
	}
	if params.Query != nil {
		out["query"] = *params.Query
	}
	if params.RecommendationType != nil {
		out["recommendation_type"] = *params.RecommendationType
	}
	if params.ReelsInteractionRate != nil {
		out["reels_interaction_rate"] = *params.ReelsInteractionRate
	}
	if params.ShowOnboardedCreatorsOnly != nil {
		out["show_onboarded_creators_only"] = *params.ShowOnboardedCreatorsOnly
	}
	if params.SimilarToCreators != nil {
		out["similar_to_creators"] = *params.SimilarToCreators
	}
	if params.Username != nil {
		out["username"] = *params.Username
	}
	if params.Usernames != nil {
		out["usernames"] = *params.Usernames
	}
	return out
}

func GetIGUserCreatorMarketplaceCreatorsBatchCall(id string, params GetIGUserCreatorMarketplaceCreatorsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "creator_marketplace_creators"), params.ToParams(), options...)
}

func NewGetIGUserCreatorMarketplaceCreatorsBatchRequest(id string, params GetIGUserCreatorMarketplaceCreatorsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGUserExportForCAM]] {
	return core.NewBatchRequest[core.Cursor[objects.IGUserExportForCAM]](GetIGUserCreatorMarketplaceCreatorsBatchCall(id, params, options...))
}

func DecodeGetIGUserCreatorMarketplaceCreatorsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGUserExportForCAM], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGUserExportForCAM]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserCreatorMarketplaceCreatorsWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserCreatorMarketplaceCreatorsParams) (*core.Cursor[objects.IGUserExportForCAM], *core.Response, error) {
	var out core.Cursor[objects.IGUserExportForCAM]
	call := GetIGUserCreatorMarketplaceCreatorsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserCreatorMarketplaceCreators(ctx context.Context, client *core.Client, id string, params GetIGUserCreatorMarketplaceCreatorsParams) (*core.Cursor[objects.IGUserExportForCAM], error) {
	out, _, err := GetIGUserCreatorMarketplaceCreatorsWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserDatasetParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserDatasetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGUserDatasetBatchCall(id string, params GetIGUserDatasetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "dataset"), params.ToParams(), options...)
}

func NewGetIGUserDatasetBatchRequest(id string, params GetIGUserDatasetParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Dataset]] {
	return core.NewBatchRequest[core.Cursor[objects.Dataset]](GetIGUserDatasetBatchCall(id, params, options...))
}

func DecodeGetIGUserDatasetBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Dataset], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Dataset]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserDatasetWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserDatasetParams) (*core.Cursor[objects.Dataset], *core.Response, error) {
	var out core.Cursor[objects.Dataset]
	call := GetIGUserDatasetBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserDataset(ctx context.Context, client *core.Client, id string, params GetIGUserDatasetParams) (*core.Cursor[objects.Dataset], error) {
	out, _, err := GetIGUserDatasetWithResponse(ctx, client, id, params)
	return out, err
}

type CreateIGUserDatasetParams struct {
	DatasetName *string     `facebook:"dataset_name"`
	Extra       core.Params `facebook:"-"`
}

func (params CreateIGUserDatasetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.DatasetName != nil {
		out["dataset_name"] = *params.DatasetName
	}
	return out
}

func CreateIGUserDatasetBatchCall(id string, params CreateIGUserDatasetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "dataset"), params.ToParams(), options...)
}

func NewCreateIGUserDatasetBatchRequest(id string, params CreateIGUserDatasetParams, options ...core.BatchOption) *core.BatchRequest[objects.Dataset] {
	return core.NewBatchRequest[objects.Dataset](CreateIGUserDatasetBatchCall(id, params, options...))
}

func DecodeCreateIGUserDatasetBatchResponse(response *core.BatchResponse) (*objects.Dataset, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Dataset
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateIGUserDatasetWithResponse(ctx context.Context, client *core.Client, id string, params CreateIGUserDatasetParams) (*objects.Dataset, *core.Response, error) {
	var out objects.Dataset
	call := CreateIGUserDatasetBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateIGUserDataset(ctx context.Context, client *core.Client, id string, params CreateIGUserDatasetParams) (*objects.Dataset, error) {
	out, _, err := CreateIGUserDatasetWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserInsightsParams struct {
	Breakdown  *[]enums.ShadowiguserinsightsBreakdownEnumParam `facebook:"breakdown"`
	Metric     []enums.ShadowiguserinsightsMetricEnumParam     `facebook:"metric"`
	MetricType *enums.ShadowiguserinsightsMetricTypeEnumParam  `facebook:"metric_type"`
	Period     []enums.ShadowiguserinsightsPeriodEnumParam     `facebook:"period"`
	Since      *core.Time                                      `facebook:"since"`
	Timeframe  *enums.ShadowiguserinsightsTimeframeEnumParam   `facebook:"timeframe"`
	Until      *core.Time                                      `facebook:"until"`
	Extra      core.Params                                     `facebook:"-"`
}

func (params GetIGUserInsightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Breakdown != nil {
		out["breakdown"] = *params.Breakdown
	}
	out["metric"] = params.Metric
	if params.MetricType != nil {
		out["metric_type"] = *params.MetricType
	}
	out["period"] = params.Period
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.Timeframe != nil {
		out["timeframe"] = *params.Timeframe
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	return out
}

func GetIGUserInsightsBatchCall(id string, params GetIGUserInsightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "insights"), params.ToParams(), options...)
}

func NewGetIGUserInsightsBatchRequest(id string, params GetIGUserInsightsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.InstagramInsightsResult]] {
	return core.NewBatchRequest[core.Cursor[objects.InstagramInsightsResult]](GetIGUserInsightsBatchCall(id, params, options...))
}

func DecodeGetIGUserInsightsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.InstagramInsightsResult], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.InstagramInsightsResult]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserInsightsWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserInsightsParams) (*core.Cursor[objects.InstagramInsightsResult], *core.Response, error) {
	var out core.Cursor[objects.InstagramInsightsResult]
	call := GetIGUserInsightsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserInsights(ctx context.Context, client *core.Client, id string, params GetIGUserInsightsParams) (*core.Cursor[objects.InstagramInsightsResult], error) {
	out, _, err := GetIGUserInsightsWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserInstagramBackedThreadsUserParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserInstagramBackedThreadsUserParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGUserInstagramBackedThreadsUserBatchCall(id string, params GetIGUserInstagramBackedThreadsUserParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "instagram_backed_threads_user"), params.ToParams(), options...)
}

func NewGetIGUserInstagramBackedThreadsUserBatchRequest(id string, params GetIGUserInstagramBackedThreadsUserParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ThreadsUser]] {
	return core.NewBatchRequest[core.Cursor[objects.ThreadsUser]](GetIGUserInstagramBackedThreadsUserBatchCall(id, params, options...))
}

func DecodeGetIGUserInstagramBackedThreadsUserBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ThreadsUser], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ThreadsUser]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserInstagramBackedThreadsUserWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserInstagramBackedThreadsUserParams) (*core.Cursor[objects.ThreadsUser], *core.Response, error) {
	var out core.Cursor[objects.ThreadsUser]
	call := GetIGUserInstagramBackedThreadsUserBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserInstagramBackedThreadsUser(ctx context.Context, client *core.Client, id string, params GetIGUserInstagramBackedThreadsUserParams) (*core.Cursor[objects.ThreadsUser], error) {
	out, _, err := GetIGUserInstagramBackedThreadsUserWithResponse(ctx, client, id, params)
	return out, err
}

type CreateIGUserInstagramBackedThreadsUserParams struct {
	Extra core.Params `facebook:"-"`
}

func (params CreateIGUserInstagramBackedThreadsUserParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func CreateIGUserInstagramBackedThreadsUserBatchCall(id string, params CreateIGUserInstagramBackedThreadsUserParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "instagram_backed_threads_user"), params.ToParams(), options...)
}

func NewCreateIGUserInstagramBackedThreadsUserBatchRequest(id string, params CreateIGUserInstagramBackedThreadsUserParams, options ...core.BatchOption) *core.BatchRequest[objects.ThreadsUser] {
	return core.NewBatchRequest[objects.ThreadsUser](CreateIGUserInstagramBackedThreadsUserBatchCall(id, params, options...))
}

func DecodeCreateIGUserInstagramBackedThreadsUserBatchResponse(response *core.BatchResponse) (*objects.ThreadsUser, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ThreadsUser
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateIGUserInstagramBackedThreadsUserWithResponse(ctx context.Context, client *core.Client, id string, params CreateIGUserInstagramBackedThreadsUserParams) (*objects.ThreadsUser, *core.Response, error) {
	var out objects.ThreadsUser
	call := CreateIGUserInstagramBackedThreadsUserBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateIGUserInstagramBackedThreadsUser(ctx context.Context, client *core.Client, id string, params CreateIGUserInstagramBackedThreadsUserParams) (*objects.ThreadsUser, error) {
	out, _, err := CreateIGUserInstagramBackedThreadsUserWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserLiveMediaParams struct {
	Since *core.Time  `facebook:"since"`
	Until *core.Time  `facebook:"until"`
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserLiveMediaParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	return out
}

func GetIGUserLiveMediaBatchCall(id string, params GetIGUserLiveMediaParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "live_media"), params.ToParams(), options...)
}

func NewGetIGUserLiveMediaBatchRequest(id string, params GetIGUserLiveMediaParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGMedia]] {
	return core.NewBatchRequest[core.Cursor[objects.IGMedia]](GetIGUserLiveMediaBatchCall(id, params, options...))
}

func DecodeGetIGUserLiveMediaBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGMedia], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGMedia]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserLiveMediaWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserLiveMediaParams) (*core.Cursor[objects.IGMedia], *core.Response, error) {
	var out core.Cursor[objects.IGMedia]
	call := GetIGUserLiveMediaBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserLiveMedia(ctx context.Context, client *core.Client, id string, params GetIGUserLiveMediaParams) (*core.Cursor[objects.IGMedia], error) {
	out, _, err := GetIGUserLiveMediaWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserMediaParams struct {
	Since *core.Time  `facebook:"since"`
	Until *core.Time  `facebook:"until"`
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserMediaParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	return out
}

func GetIGUserMediaBatchCall(id string, params GetIGUserMediaParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "media"), params.ToParams(), options...)
}

func NewGetIGUserMediaBatchRequest(id string, params GetIGUserMediaParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGMedia]] {
	return core.NewBatchRequest[core.Cursor[objects.IGMedia]](GetIGUserMediaBatchCall(id, params, options...))
}

func DecodeGetIGUserMediaBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGMedia], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGMedia]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserMediaWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserMediaParams) (*core.Cursor[objects.IGMedia], *core.Response, error) {
	var out core.Cursor[objects.IGMedia]
	call := GetIGUserMediaBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserMedia(ctx context.Context, client *core.Client, id string, params GetIGUserMediaParams) (*core.Cursor[objects.IGMedia], error) {
	out, _, err := GetIGUserMediaWithResponse(ctx, client, id, params)
	return out, err
}

type CreateIGUserMediaParams struct {
	AltText                  *string                   `facebook:"alt_text"`
	AudioConfiguration       *string                   `facebook:"audio_configuration"`
	AudioName                *string                   `facebook:"audio_name"`
	BrandedContentSponsorIds *[]core.ID                `facebook:"branded_content_sponsor_ids"`
	Caption                  *string                   `facebook:"caption"`
	Children                 *[]string                 `facebook:"children"`
	Collaborators            *[]string                 `facebook:"collaborators"`
	CoverURL                 *string                   `facebook:"cover_url"`
	ImageURL                 *string                   `facebook:"image_url"`
	IsCarouselItem           *bool                     `facebook:"is_carousel_item"`
	IsPaidPartnership        *bool                     `facebook:"is_paid_partnership"`
	LocationID               *core.ID                  `facebook:"location_id"`
	MediaType                *string                   `facebook:"media_type"`
	ProductTags              *[]map[string]interface{} `facebook:"product_tags"`
	ShareToFeed              *bool                     `facebook:"share_to_feed"`
	ThumbOffset              *string                   `facebook:"thumb_offset"`
	TrialParams              *map[string]interface{}   `facebook:"trial_params"`
	UploadType               *string                   `facebook:"upload_type"`
	UserTags                 *[]map[string]interface{} `facebook:"user_tags"`
	VideoURL                 *string                   `facebook:"video_url"`
	Extra                    core.Params               `facebook:"-"`
}

func (params CreateIGUserMediaParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AltText != nil {
		out["alt_text"] = *params.AltText
	}
	if params.AudioConfiguration != nil {
		out["audio_configuration"] = *params.AudioConfiguration
	}
	if params.AudioName != nil {
		out["audio_name"] = *params.AudioName
	}
	if params.BrandedContentSponsorIds != nil {
		out["branded_content_sponsor_ids"] = *params.BrandedContentSponsorIds
	}
	if params.Caption != nil {
		out["caption"] = *params.Caption
	}
	if params.Children != nil {
		out["children"] = *params.Children
	}
	if params.Collaborators != nil {
		out["collaborators"] = *params.Collaborators
	}
	if params.CoverURL != nil {
		out["cover_url"] = *params.CoverURL
	}
	if params.ImageURL != nil {
		out["image_url"] = *params.ImageURL
	}
	if params.IsCarouselItem != nil {
		out["is_carousel_item"] = *params.IsCarouselItem
	}
	if params.IsPaidPartnership != nil {
		out["is_paid_partnership"] = *params.IsPaidPartnership
	}
	if params.LocationID != nil {
		out["location_id"] = *params.LocationID
	}
	if params.MediaType != nil {
		out["media_type"] = *params.MediaType
	}
	if params.ProductTags != nil {
		out["product_tags"] = *params.ProductTags
	}
	if params.ShareToFeed != nil {
		out["share_to_feed"] = *params.ShareToFeed
	}
	if params.ThumbOffset != nil {
		out["thumb_offset"] = *params.ThumbOffset
	}
	if params.TrialParams != nil {
		out["trial_params"] = *params.TrialParams
	}
	if params.UploadType != nil {
		out["upload_type"] = *params.UploadType
	}
	if params.UserTags != nil {
		out["user_tags"] = *params.UserTags
	}
	if params.VideoURL != nil {
		out["video_url"] = *params.VideoURL
	}
	return out
}

func CreateIGUserMediaBatchCall(id string, params CreateIGUserMediaParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "media"), params.ToParams(), options...)
}

func NewCreateIGUserMediaBatchRequest(id string, params CreateIGUserMediaParams, options ...core.BatchOption) *core.BatchRequest[objects.IGMedia] {
	return core.NewBatchRequest[objects.IGMedia](CreateIGUserMediaBatchCall(id, params, options...))
}

func DecodeCreateIGUserMediaBatchResponse(response *core.BatchResponse) (*objects.IGMedia, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.IGMedia
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateIGUserMediaWithResponse(ctx context.Context, client *core.Client, id string, params CreateIGUserMediaParams) (*objects.IGMedia, *core.Response, error) {
	var out objects.IGMedia
	call := CreateIGUserMediaBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateIGUserMedia(ctx context.Context, client *core.Client, id string, params CreateIGUserMediaParams) (*objects.IGMedia, error) {
	out, _, err := CreateIGUserMediaWithResponse(ctx, client, id, params)
	return out, err
}

type CreateIGUserMediaPublishParams struct {
	CreationID core.ID     `facebook:"creation_id"`
	Extra      core.Params `facebook:"-"`
}

func (params CreateIGUserMediaPublishParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["creation_id"] = params.CreationID
	return out
}

func CreateIGUserMediaPublishBatchCall(id string, params CreateIGUserMediaPublishParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "media_publish"), params.ToParams(), options...)
}

func NewCreateIGUserMediaPublishBatchRequest(id string, params CreateIGUserMediaPublishParams, options ...core.BatchOption) *core.BatchRequest[objects.IGMedia] {
	return core.NewBatchRequest[objects.IGMedia](CreateIGUserMediaPublishBatchCall(id, params, options...))
}

func DecodeCreateIGUserMediaPublishBatchResponse(response *core.BatchResponse) (*objects.IGMedia, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.IGMedia
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateIGUserMediaPublishWithResponse(ctx context.Context, client *core.Client, id string, params CreateIGUserMediaPublishParams) (*objects.IGMedia, *core.Response, error) {
	var out objects.IGMedia
	call := CreateIGUserMediaPublishBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateIGUserMediaPublish(ctx context.Context, client *core.Client, id string, params CreateIGUserMediaPublishParams) (*objects.IGMedia, error) {
	out, _, err := CreateIGUserMediaPublishWithResponse(ctx, client, id, params)
	return out, err
}

type CreateIGUserMentionsParams struct {
	CommentID *core.ID    `facebook:"comment_id"`
	MediaID   core.ID     `facebook:"media_id"`
	Message   string      `facebook:"message"`
	Extra     core.Params `facebook:"-"`
}

func (params CreateIGUserMentionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CommentID != nil {
		out["comment_id"] = *params.CommentID
	}
	out["media_id"] = params.MediaID
	out["message"] = params.Message
	return out
}

func CreateIGUserMentionsBatchCall(id string, params CreateIGUserMentionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "mentions"), params.ToParams(), options...)
}

func NewCreateIGUserMentionsBatchRequest(id string, params CreateIGUserMentionsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateIGUserMentionsBatchCall(id, params, options...))
}

func DecodeCreateIGUserMentionsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateIGUserMentionsWithResponse(ctx context.Context, client *core.Client, id string, params CreateIGUserMentionsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreateIGUserMentionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateIGUserMentions(ctx context.Context, client *core.Client, id string, params CreateIGUserMentionsParams) (*map[string]interface{}, error) {
	out, _, err := CreateIGUserMentionsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateIGUserModerateConversationsParams struct {
	Actions []enums.ShadowigusermoderateConversationsActionsEnumParam `facebook:"actions"`
	UserIds []map[string]interface{}                                  `facebook:"user_ids"`
	Extra   core.Params                                               `facebook:"-"`
}

func (params CreateIGUserModerateConversationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["actions"] = params.Actions
	out["user_ids"] = params.UserIds
	return out
}

func CreateIGUserModerateConversationsBatchCall(id string, params CreateIGUserModerateConversationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "moderate_conversations"), params.ToParams(), options...)
}

func NewCreateIGUserModerateConversationsBatchRequest(id string, params CreateIGUserModerateConversationsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateIGUserModerateConversationsBatchCall(id, params, options...))
}

func DecodeCreateIGUserModerateConversationsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateIGUserModerateConversationsWithResponse(ctx context.Context, client *core.Client, id string, params CreateIGUserModerateConversationsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreateIGUserModerateConversationsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateIGUserModerateConversations(ctx context.Context, client *core.Client, id string, params CreateIGUserModerateConversationsParams) (*map[string]interface{}, error) {
	out, _, err := CreateIGUserModerateConversationsWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserNotificationMessageTokensParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserNotificationMessageTokensParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGUserNotificationMessageTokensBatchCall(id string, params GetIGUserNotificationMessageTokensParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "notification_message_tokens"), params.ToParams(), options...)
}

func NewGetIGUserNotificationMessageTokensBatchRequest(id string, params GetIGUserNotificationMessageTokensParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.UserPageOneTimeOptInTokenSettings]] {
	return core.NewBatchRequest[core.Cursor[objects.UserPageOneTimeOptInTokenSettings]](GetIGUserNotificationMessageTokensBatchCall(id, params, options...))
}

func DecodeGetIGUserNotificationMessageTokensBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.UserPageOneTimeOptInTokenSettings], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.UserPageOneTimeOptInTokenSettings]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserNotificationMessageTokensWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserNotificationMessageTokensParams) (*core.Cursor[objects.UserPageOneTimeOptInTokenSettings], *core.Response, error) {
	var out core.Cursor[objects.UserPageOneTimeOptInTokenSettings]
	call := GetIGUserNotificationMessageTokensBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserNotificationMessageTokens(ctx context.Context, client *core.Client, id string, params GetIGUserNotificationMessageTokensParams) (*core.Cursor[objects.UserPageOneTimeOptInTokenSettings], error) {
	out, _, err := GetIGUserNotificationMessageTokensWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserProductAppealParams struct {
	ProductID core.ID     `facebook:"product_id"`
	Extra     core.Params `facebook:"-"`
}

func (params GetIGUserProductAppealParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["product_id"] = params.ProductID
	return out
}

func GetIGUserProductAppealBatchCall(id string, params GetIGUserProductAppealParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "product_appeal"), params.ToParams(), options...)
}

func NewGetIGUserProductAppealBatchRequest(id string, params GetIGUserProductAppealParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGShoppingProductAppeal]] {
	return core.NewBatchRequest[core.Cursor[objects.IGShoppingProductAppeal]](GetIGUserProductAppealBatchCall(id, params, options...))
}

func DecodeGetIGUserProductAppealBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGShoppingProductAppeal], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGShoppingProductAppeal]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserProductAppealWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserProductAppealParams) (*core.Cursor[objects.IGShoppingProductAppeal], *core.Response, error) {
	var out core.Cursor[objects.IGShoppingProductAppeal]
	call := GetIGUserProductAppealBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserProductAppeal(ctx context.Context, client *core.Client, id string, params GetIGUserProductAppealParams) (*core.Cursor[objects.IGShoppingProductAppeal], error) {
	out, _, err := GetIGUserProductAppealWithResponse(ctx, client, id, params)
	return out, err
}

type CreateIGUserProductAppealParams struct {
	AppealReason string      `facebook:"appeal_reason"`
	ProductID    core.ID     `facebook:"product_id"`
	Extra        core.Params `facebook:"-"`
}

func (params CreateIGUserProductAppealParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["appeal_reason"] = params.AppealReason
	out["product_id"] = params.ProductID
	return out
}

func CreateIGUserProductAppealBatchCall(id string, params CreateIGUserProductAppealParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "product_appeal"), params.ToParams(), options...)
}

func NewCreateIGUserProductAppealBatchRequest(id string, params CreateIGUserProductAppealParams, options ...core.BatchOption) *core.BatchRequest[objects.IGShoppingProductAppeal] {
	return core.NewBatchRequest[objects.IGShoppingProductAppeal](CreateIGUserProductAppealBatchCall(id, params, options...))
}

func DecodeCreateIGUserProductAppealBatchResponse(response *core.BatchResponse) (*objects.IGShoppingProductAppeal, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.IGShoppingProductAppeal
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateIGUserProductAppealWithResponse(ctx context.Context, client *core.Client, id string, params CreateIGUserProductAppealParams) (*objects.IGShoppingProductAppeal, *core.Response, error) {
	var out objects.IGShoppingProductAppeal
	call := CreateIGUserProductAppealBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateIGUserProductAppeal(ctx context.Context, client *core.Client, id string, params CreateIGUserProductAppealParams) (*objects.IGShoppingProductAppeal, error) {
	out, _, err := CreateIGUserProductAppealWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserRecentlySearchedHashtagsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserRecentlySearchedHashtagsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGUserRecentlySearchedHashtagsBatchCall(id string, params GetIGUserRecentlySearchedHashtagsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "recently_searched_hashtags"), params.ToParams(), options...)
}

func NewGetIGUserRecentlySearchedHashtagsBatchRequest(id string, params GetIGUserRecentlySearchedHashtagsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ShadowIGHashtag]] {
	return core.NewBatchRequest[core.Cursor[objects.ShadowIGHashtag]](GetIGUserRecentlySearchedHashtagsBatchCall(id, params, options...))
}

func DecodeGetIGUserRecentlySearchedHashtagsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ShadowIGHashtag], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ShadowIGHashtag]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserRecentlySearchedHashtagsWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserRecentlySearchedHashtagsParams) (*core.Cursor[objects.ShadowIGHashtag], *core.Response, error) {
	var out core.Cursor[objects.ShadowIGHashtag]
	call := GetIGUserRecentlySearchedHashtagsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserRecentlySearchedHashtags(ctx context.Context, client *core.Client, id string, params GetIGUserRecentlySearchedHashtagsParams) (*core.Cursor[objects.ShadowIGHashtag], error) {
	out, _, err := GetIGUserRecentlySearchedHashtagsWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserScheduledMediaParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserScheduledMediaParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGUserScheduledMediaBatchCall(id string, params GetIGUserScheduledMediaParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "scheduled_media"), params.ToParams(), options...)
}

func NewGetIGUserScheduledMediaBatchRequest(id string, params GetIGUserScheduledMediaParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ShadowIGScheduledMedia]] {
	return core.NewBatchRequest[core.Cursor[objects.ShadowIGScheduledMedia]](GetIGUserScheduledMediaBatchCall(id, params, options...))
}

func DecodeGetIGUserScheduledMediaBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ShadowIGScheduledMedia], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ShadowIGScheduledMedia]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserScheduledMediaWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserScheduledMediaParams) (*core.Cursor[objects.ShadowIGScheduledMedia], *core.Response, error) {
	var out core.Cursor[objects.ShadowIGScheduledMedia]
	call := GetIGUserScheduledMediaBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserScheduledMedia(ctx context.Context, client *core.Client, id string, params GetIGUserScheduledMediaParams) (*core.Cursor[objects.ShadowIGScheduledMedia], error) {
	out, _, err := GetIGUserScheduledMediaWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserStoriesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserStoriesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGUserStoriesBatchCall(id string, params GetIGUserStoriesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "stories"), params.ToParams(), options...)
}

func NewGetIGUserStoriesBatchRequest(id string, params GetIGUserStoriesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGMedia]] {
	return core.NewBatchRequest[core.Cursor[objects.IGMedia]](GetIGUserStoriesBatchCall(id, params, options...))
}

func DecodeGetIGUserStoriesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGMedia], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGMedia]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserStoriesWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserStoriesParams) (*core.Cursor[objects.IGMedia], *core.Response, error) {
	var out core.Cursor[objects.IGMedia]
	call := GetIGUserStoriesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserStories(ctx context.Context, client *core.Client, id string, params GetIGUserStoriesParams) (*core.Cursor[objects.IGMedia], error) {
	out, _, err := GetIGUserStoriesWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserTagsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserTagsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGUserTagsBatchCall(id string, params GetIGUserTagsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "tags"), params.ToParams(), options...)
}

func NewGetIGUserTagsBatchRequest(id string, params GetIGUserTagsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGMedia]] {
	return core.NewBatchRequest[core.Cursor[objects.IGMedia]](GetIGUserTagsBatchCall(id, params, options...))
}

func DecodeGetIGUserTagsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGMedia], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGMedia]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserTagsWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserTagsParams) (*core.Cursor[objects.IGMedia], *core.Response, error) {
	var out core.Cursor[objects.IGMedia]
	call := GetIGUserTagsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserTags(ctx context.Context, client *core.Client, id string, params GetIGUserTagsParams) (*core.Cursor[objects.IGMedia], error) {
	out, _, err := GetIGUserTagsWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserUpcomingEventsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserUpcomingEventsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGUserUpcomingEventsBatchCall(id string, params GetIGUserUpcomingEventsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "upcoming_events"), params.ToParams(), options...)
}

func NewGetIGUserUpcomingEventsBatchRequest(id string, params GetIGUserUpcomingEventsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGUpcomingEvent]] {
	return core.NewBatchRequest[core.Cursor[objects.IGUpcomingEvent]](GetIGUserUpcomingEventsBatchCall(id, params, options...))
}

func DecodeGetIGUserUpcomingEventsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGUpcomingEvent], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGUpcomingEvent]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserUpcomingEventsWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserUpcomingEventsParams) (*core.Cursor[objects.IGUpcomingEvent], *core.Response, error) {
	var out core.Cursor[objects.IGUpcomingEvent]
	call := GetIGUserUpcomingEventsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserUpcomingEvents(ctx context.Context, client *core.Client, id string, params GetIGUserUpcomingEventsParams) (*core.Cursor[objects.IGUpcomingEvent], error) {
	out, _, err := GetIGUserUpcomingEventsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateIGUserUpcomingEventsParams struct {
	EndTime              *core.Time                                                       `facebook:"end_time"`
	NotificationSubtypes *[]enums.ShadowiguserupcomingEventsNotificationSubtypesEnumParam `facebook:"notification_subtypes"`
	StartTime            core.Time                                                        `facebook:"start_time"`
	Title                string                                                           `facebook:"title"`
	Extra                core.Params                                                      `facebook:"-"`
}

func (params CreateIGUserUpcomingEventsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.EndTime != nil {
		out["end_time"] = *params.EndTime
	}
	if params.NotificationSubtypes != nil {
		out["notification_subtypes"] = *params.NotificationSubtypes
	}
	out["start_time"] = params.StartTime
	out["title"] = params.Title
	return out
}

func CreateIGUserUpcomingEventsBatchCall(id string, params CreateIGUserUpcomingEventsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "upcoming_events"), params.ToParams(), options...)
}

func NewCreateIGUserUpcomingEventsBatchRequest(id string, params CreateIGUserUpcomingEventsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateIGUserUpcomingEventsBatchCall(id, params, options...))
}

func DecodeCreateIGUserUpcomingEventsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateIGUserUpcomingEventsWithResponse(ctx context.Context, client *core.Client, id string, params CreateIGUserUpcomingEventsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreateIGUserUpcomingEventsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateIGUserUpcomingEvents(ctx context.Context, client *core.Client, id string, params CreateIGUserUpcomingEventsParams) (*map[string]interface{}, error) {
	out, _, err := CreateIGUserUpcomingEventsWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserWelcomeMessageFlowsParams struct {
	AppID  *core.ID    `facebook:"app_id"`
	FlowID *core.ID    `facebook:"flow_id"`
	Extra  core.Params `facebook:"-"`
}

func (params GetIGUserWelcomeMessageFlowsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AppID != nil {
		out["app_id"] = *params.AppID
	}
	if params.FlowID != nil {
		out["flow_id"] = *params.FlowID
	}
	return out
}

func GetIGUserWelcomeMessageFlowsBatchCall(id string, params GetIGUserWelcomeMessageFlowsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "welcome_message_flows"), params.ToParams(), options...)
}

func NewGetIGUserWelcomeMessageFlowsBatchRequest(id string, params GetIGUserWelcomeMessageFlowsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ShadowIGUserCTXPartnerAppWelcomeMessageFlow]] {
	return core.NewBatchRequest[core.Cursor[objects.ShadowIGUserCTXPartnerAppWelcomeMessageFlow]](GetIGUserWelcomeMessageFlowsBatchCall(id, params, options...))
}

func DecodeGetIGUserWelcomeMessageFlowsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ShadowIGUserCTXPartnerAppWelcomeMessageFlow], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ShadowIGUserCTXPartnerAppWelcomeMessageFlow]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserWelcomeMessageFlowsWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserWelcomeMessageFlowsParams) (*core.Cursor[objects.ShadowIGUserCTXPartnerAppWelcomeMessageFlow], *core.Response, error) {
	var out core.Cursor[objects.ShadowIGUserCTXPartnerAppWelcomeMessageFlow]
	call := GetIGUserWelcomeMessageFlowsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserWelcomeMessageFlows(ctx context.Context, client *core.Client, id string, params GetIGUserWelcomeMessageFlowsParams) (*core.Cursor[objects.ShadowIGUserCTXPartnerAppWelcomeMessageFlow], error) {
	out, _, err := GetIGUserWelcomeMessageFlowsWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserParams struct {
	AdgroupID *core.ID    `facebook:"adgroup_id"`
	Extra     core.Params `facebook:"-"`
}

func (params GetIGUserParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdgroupID != nil {
		out["adgroup_id"] = *params.AdgroupID
	}
	return out
}

func GetIGUserBatchCall(id string, params GetIGUserParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetIGUserBatchRequest(id string, params GetIGUserParams, options ...core.BatchOption) *core.BatchRequest[objects.IGUser] {
	return core.NewBatchRequest[objects.IGUser](GetIGUserBatchCall(id, params, options...))
}

func DecodeGetIGUserBatchResponse(response *core.BatchResponse) (*objects.IGUser, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.IGUser
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserParams) (*objects.IGUser, *core.Response, error) {
	var out objects.IGUser
	call := GetIGUserBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUser(ctx context.Context, client *core.Client, id string, params GetIGUserParams) (*objects.IGUser, error) {
	out, _, err := GetIGUserWithResponse(ctx, client, id, params)
	return out, err
}
