package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type DeleteUserAccessTokensParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteUserAccessTokensParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteUserAccessTokensBatchCall(id string, params DeleteUserAccessTokensParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "access_tokens"), params.ToParams(), options...)
}

func NewDeleteUserAccessTokensBatchRequest(id string, params DeleteUserAccessTokensParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteUserAccessTokensBatchCall(id, params, options...))
}

func DecodeDeleteUserAccessTokensBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteUserAccessTokensWithResponse(ctx context.Context, client *core.Client, id string, params DeleteUserAccessTokensParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteUserAccessTokensBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteUserAccessTokens(ctx context.Context, client *core.Client, id string, params DeleteUserAccessTokensParams) (*map[string]interface{}, error) {
	out, _, err := DeleteUserAccessTokensWithResponse(ctx, client, id, params)
	return out, err
}

type CreateUserAccessTokensParams struct {
	BusinessApp              string      `facebook:"business_app"`
	PageID                   *core.ID    `facebook:"page_id"`
	Scope                    *[]string   `facebook:"scope"`
	SetTokenExpiresInX60Days *bool       `facebook:"set_token_expires_in_60_days"`
	Extra                    core.Params `facebook:"-"`
}

func (params CreateUserAccessTokensParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business_app"] = params.BusinessApp
	if params.PageID != nil {
		out["page_id"] = *params.PageID
	}
	if params.Scope != nil {
		out["scope"] = *params.Scope
	}
	if params.SetTokenExpiresInX60Days != nil {
		out["set_token_expires_in_60_days"] = *params.SetTokenExpiresInX60Days
	}
	return out
}

func CreateUserAccessTokensBatchCall(id string, params CreateUserAccessTokensParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "access_tokens"), params.ToParams(), options...)
}

func NewCreateUserAccessTokensBatchRequest(id string, params CreateUserAccessTokensParams, options ...core.BatchOption) *core.BatchRequest[objects.User] {
	return core.NewBatchRequest[objects.User](CreateUserAccessTokensBatchCall(id, params, options...))
}

func DecodeCreateUserAccessTokensBatchResponse(response *core.BatchResponse) (*objects.User, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.User
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateUserAccessTokensWithResponse(ctx context.Context, client *core.Client, id string, params CreateUserAccessTokensParams) (*objects.User, *core.Response, error) {
	var out objects.User
	call := CreateUserAccessTokensBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateUserAccessTokens(ctx context.Context, client *core.Client, id string, params CreateUserAccessTokensParams) (*objects.User, error) {
	out, _, err := CreateUserAccessTokensWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserAccountsParams struct {
	AdID         *core.ID    `facebook:"ad_id"`
	IsPlace      *bool       `facebook:"is_place"`
	IsPromotable *bool       `facebook:"is_promotable"`
	Extra        core.Params `facebook:"-"`
}

func (params GetUserAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdID != nil {
		out["ad_id"] = *params.AdID
	}
	if params.IsPlace != nil {
		out["is_place"] = *params.IsPlace
	}
	if params.IsPromotable != nil {
		out["is_promotable"] = *params.IsPromotable
	}
	return out
}

func GetUserAccountsBatchCall(id string, params GetUserAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "accounts"), params.ToParams(), options...)
}

func NewGetUserAccountsBatchRequest(id string, params GetUserAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Page]] {
	return core.NewBatchRequest[core.Cursor[objects.Page]](GetUserAccountsBatchCall(id, params, options...))
}

func DecodeGetUserAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Page], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Page]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserAccountsWithResponse(ctx context.Context, client *core.Client, id string, params GetUserAccountsParams) (*core.Cursor[objects.Page], *core.Response, error) {
	var out core.Cursor[objects.Page]
	call := GetUserAccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserAccounts(ctx context.Context, client *core.Client, id string, params GetUserAccountsParams) (*core.Cursor[objects.Page], error) {
	out, _, err := GetUserAccountsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateUserAccountsParams struct {
	About                    *string                 `facebook:"about"`
	Address                  *string                 `facebook:"address"`
	Category                 *int                    `facebook:"category"`
	CategoryEnum             *string                 `facebook:"category_enum"`
	CategoryList             *[]string               `facebook:"category_list"`
	CityID                   *core.ID                `facebook:"city_id"`
	Coordinates              *map[string]interface{} `facebook:"coordinates"`
	CoverPhoto               *map[string]interface{} `facebook:"cover_photo"`
	Description              *string                 `facebook:"description"`
	IgnoreCoordinateWarnings *bool                   `facebook:"ignore_coordinate_warnings"`
	Location                 *map[string]interface{} `facebook:"location"`
	Name                     string                  `facebook:"name"`
	Phone                    *string                 `facebook:"phone"`
	Picture                  *string                 `facebook:"picture"`
	Website                  *string                 `facebook:"website"`
	Zip                      *string                 `facebook:"zip"`
	Extra                    core.Params             `facebook:"-"`
}

func (params CreateUserAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.About != nil {
		out["about"] = *params.About
	}
	if params.Address != nil {
		out["address"] = *params.Address
	}
	if params.Category != nil {
		out["category"] = *params.Category
	}
	if params.CategoryEnum != nil {
		out["category_enum"] = *params.CategoryEnum
	}
	if params.CategoryList != nil {
		out["category_list"] = *params.CategoryList
	}
	if params.CityID != nil {
		out["city_id"] = *params.CityID
	}
	if params.Coordinates != nil {
		out["coordinates"] = *params.Coordinates
	}
	if params.CoverPhoto != nil {
		out["cover_photo"] = *params.CoverPhoto
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.IgnoreCoordinateWarnings != nil {
		out["ignore_coordinate_warnings"] = *params.IgnoreCoordinateWarnings
	}
	if params.Location != nil {
		out["location"] = *params.Location
	}
	out["name"] = params.Name
	if params.Phone != nil {
		out["phone"] = *params.Phone
	}
	if params.Picture != nil {
		out["picture"] = *params.Picture
	}
	if params.Website != nil {
		out["website"] = *params.Website
	}
	if params.Zip != nil {
		out["zip"] = *params.Zip
	}
	return out
}

func CreateUserAccountsBatchCall(id string, params CreateUserAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "accounts"), params.ToParams(), options...)
}

func NewCreateUserAccountsBatchRequest(id string, params CreateUserAccountsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateUserAccountsBatchCall(id, params, options...))
}

func DecodeCreateUserAccountsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreateUserAccountsWithResponse(ctx context.Context, client *core.Client, id string, params CreateUserAccountsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreateUserAccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateUserAccounts(ctx context.Context, client *core.Client, id string, params CreateUserAccountsParams) (*map[string]interface{}, error) {
	out, _, err := CreateUserAccountsWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserAdStudiesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserAdStudiesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserAdStudiesBatchCall(id string, params GetUserAdStudiesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ad_studies"), params.ToParams(), options...)
}

func NewGetUserAdStudiesBatchRequest(id string, params GetUserAdStudiesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdStudy]] {
	return core.NewBatchRequest[core.Cursor[objects.AdStudy]](GetUserAdStudiesBatchCall(id, params, options...))
}

func DecodeGetUserAdStudiesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdStudy], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdStudy]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserAdStudiesWithResponse(ctx context.Context, client *core.Client, id string, params GetUserAdStudiesParams) (*core.Cursor[objects.AdStudy], *core.Response, error) {
	var out core.Cursor[objects.AdStudy]
	call := GetUserAdStudiesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserAdStudies(ctx context.Context, client *core.Client, id string, params GetUserAdStudiesParams) (*core.Cursor[objects.AdStudy], error) {
	out, _, err := GetUserAdStudiesWithResponse(ctx, client, id, params)
	return out, err
}

type CreateUserAdStudiesParams struct {
	Cells              *[]map[string]interface{}         `facebook:"cells"`
	ClientBusiness     *string                           `facebook:"client_business"`
	ConfidenceLevel    *float64                          `facebook:"confidence_level"`
	CooldownStartTime  *int                              `facebook:"cooldown_start_time"`
	CreativeTestConfig *map[string]interface{}           `facebook:"creative_test_config"`
	Description        *string                           `facebook:"description"`
	EndTime            *int                              `facebook:"end_time"`
	Name               *string                           `facebook:"name"`
	Objectives         *[]map[string]interface{}         `facebook:"objectives"`
	ObservationEndTime *int                              `facebook:"observation_end_time"`
	StartTime          *int                              `facebook:"start_time"`
	Type               *enums.UseradStudiesTypeEnumParam `facebook:"type"`
	Viewers            *[]int                            `facebook:"viewers"`
	Extra              core.Params                       `facebook:"-"`
}

func (params CreateUserAdStudiesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Cells != nil {
		out["cells"] = *params.Cells
	}
	if params.ClientBusiness != nil {
		out["client_business"] = *params.ClientBusiness
	}
	if params.ConfidenceLevel != nil {
		out["confidence_level"] = *params.ConfidenceLevel
	}
	if params.CooldownStartTime != nil {
		out["cooldown_start_time"] = *params.CooldownStartTime
	}
	if params.CreativeTestConfig != nil {
		out["creative_test_config"] = *params.CreativeTestConfig
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.EndTime != nil {
		out["end_time"] = *params.EndTime
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.Objectives != nil {
		out["objectives"] = *params.Objectives
	}
	if params.ObservationEndTime != nil {
		out["observation_end_time"] = *params.ObservationEndTime
	}
	if params.StartTime != nil {
		out["start_time"] = *params.StartTime
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	if params.Viewers != nil {
		out["viewers"] = *params.Viewers
	}
	return out
}

func CreateUserAdStudiesBatchCall(id string, params CreateUserAdStudiesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "ad_studies"), params.ToParams(), options...)
}

func NewCreateUserAdStudiesBatchRequest(id string, params CreateUserAdStudiesParams, options ...core.BatchOption) *core.BatchRequest[objects.AdStudy] {
	return core.NewBatchRequest[objects.AdStudy](CreateUserAdStudiesBatchCall(id, params, options...))
}

func DecodeCreateUserAdStudiesBatchResponse(response *core.BatchResponse) (*objects.AdStudy, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdStudy
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateUserAdStudiesWithResponse(ctx context.Context, client *core.Client, id string, params CreateUserAdStudiesParams) (*objects.AdStudy, *core.Response, error) {
	var out objects.AdStudy
	call := CreateUserAdStudiesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateUserAdStudies(ctx context.Context, client *core.Client, id string, params CreateUserAdStudiesParams) (*objects.AdStudy, error) {
	out, _, err := CreateUserAdStudiesWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserAdaccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserAdaccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserAdaccountsBatchCall(id string, params GetUserAdaccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adaccounts"), params.ToParams(), options...)
}

func NewGetUserAdaccountsBatchRequest(id string, params GetUserAdaccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccount]](GetUserAdaccountsBatchCall(id, params, options...))
}

func DecodeGetUserAdaccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccount], error) {
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

func GetUserAdaccountsWithResponse(ctx context.Context, client *core.Client, id string, params GetUserAdaccountsParams) (*core.Cursor[objects.AdAccount], *core.Response, error) {
	var out core.Cursor[objects.AdAccount]
	call := GetUserAdaccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserAdaccounts(ctx context.Context, client *core.Client, id string, params GetUserAdaccountsParams) (*core.Cursor[objects.AdAccount], error) {
	out, _, err := GetUserAdaccountsWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserAlbumsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserAlbumsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserAlbumsBatchCall(id string, params GetUserAlbumsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "albums"), params.ToParams(), options...)
}

func NewGetUserAlbumsBatchRequest(id string, params GetUserAlbumsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Album]] {
	return core.NewBatchRequest[core.Cursor[objects.Album]](GetUserAlbumsBatchCall(id, params, options...))
}

func DecodeGetUserAlbumsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Album], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Album]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserAlbumsWithResponse(ctx context.Context, client *core.Client, id string, params GetUserAlbumsParams) (*core.Cursor[objects.Album], *core.Response, error) {
	var out core.Cursor[objects.Album]
	call := GetUserAlbumsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserAlbums(ctx context.Context, client *core.Client, id string, params GetUserAlbumsParams) (*core.Cursor[objects.Album], error) {
	out, _, err := GetUserAlbumsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateUserApplicationsParams struct {
	BusinessApp int         `facebook:"business_app"`
	Extra       core.Params `facebook:"-"`
}

func (params CreateUserApplicationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business_app"] = params.BusinessApp
	return out
}

func CreateUserApplicationsBatchCall(id string, params CreateUserApplicationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "applications"), params.ToParams(), options...)
}

func NewCreateUserApplicationsBatchRequest(id string, params CreateUserApplicationsParams, options ...core.BatchOption) *core.BatchRequest[objects.User] {
	return core.NewBatchRequest[objects.User](CreateUserApplicationsBatchCall(id, params, options...))
}

func DecodeCreateUserApplicationsBatchResponse(response *core.BatchResponse) (*objects.User, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.User
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateUserApplicationsWithResponse(ctx context.Context, client *core.Client, id string, params CreateUserApplicationsParams) (*objects.User, *core.Response, error) {
	var out objects.User
	call := CreateUserApplicationsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateUserApplications(ctx context.Context, client *core.Client, id string, params CreateUserApplicationsParams) (*objects.User, error) {
	out, _, err := CreateUserApplicationsWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserApprequestformerrecipientsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserApprequestformerrecipientsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserApprequestformerrecipientsBatchCall(id string, params GetUserApprequestformerrecipientsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "apprequestformerrecipients"), params.ToParams(), options...)
}

func NewGetUserApprequestformerrecipientsBatchRequest(id string, params GetUserApprequestformerrecipientsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AppRequestFormerRecipient]] {
	return core.NewBatchRequest[core.Cursor[objects.AppRequestFormerRecipient]](GetUserApprequestformerrecipientsBatchCall(id, params, options...))
}

func DecodeGetUserApprequestformerrecipientsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AppRequestFormerRecipient], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AppRequestFormerRecipient]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserApprequestformerrecipientsWithResponse(ctx context.Context, client *core.Client, id string, params GetUserApprequestformerrecipientsParams) (*core.Cursor[objects.AppRequestFormerRecipient], *core.Response, error) {
	var out core.Cursor[objects.AppRequestFormerRecipient]
	call := GetUserApprequestformerrecipientsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserApprequestformerrecipients(ctx context.Context, client *core.Client, id string, params GetUserApprequestformerrecipientsParams) (*core.Cursor[objects.AppRequestFormerRecipient], error) {
	out, _, err := GetUserApprequestformerrecipientsWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserApprequestsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserApprequestsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserApprequestsBatchCall(id string, params GetUserApprequestsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "apprequests"), params.ToParams(), options...)
}

func NewGetUserApprequestsBatchRequest(id string, params GetUserApprequestsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AppRequest]] {
	return core.NewBatchRequest[core.Cursor[objects.AppRequest]](GetUserApprequestsBatchCall(id, params, options...))
}

func DecodeGetUserApprequestsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AppRequest], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AppRequest]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserApprequestsWithResponse(ctx context.Context, client *core.Client, id string, params GetUserApprequestsParams) (*core.Cursor[objects.AppRequest], *core.Response, error) {
	var out core.Cursor[objects.AppRequest]
	call := GetUserApprequestsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserApprequests(ctx context.Context, client *core.Client, id string, params GetUserApprequestsParams) (*core.Cursor[objects.AppRequest], error) {
	out, _, err := GetUserApprequestsWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserAssignedAdAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserAssignedAdAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserAssignedAdAccountsBatchCall(id string, params GetUserAssignedAdAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "assigned_ad_accounts"), params.ToParams(), options...)
}

func NewGetUserAssignedAdAccountsBatchRequest(id string, params GetUserAssignedAdAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccount]](GetUserAssignedAdAccountsBatchCall(id, params, options...))
}

func DecodeGetUserAssignedAdAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccount], error) {
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

func GetUserAssignedAdAccountsWithResponse(ctx context.Context, client *core.Client, id string, params GetUserAssignedAdAccountsParams) (*core.Cursor[objects.AdAccount], *core.Response, error) {
	var out core.Cursor[objects.AdAccount]
	call := GetUserAssignedAdAccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserAssignedAdAccounts(ctx context.Context, client *core.Client, id string, params GetUserAssignedAdAccountsParams) (*core.Cursor[objects.AdAccount], error) {
	out, _, err := GetUserAssignedAdAccountsWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserAssignedApplicationsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserAssignedApplicationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserAssignedApplicationsBatchCall(id string, params GetUserAssignedApplicationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "assigned_applications"), params.ToParams(), options...)
}

func NewGetUserAssignedApplicationsBatchRequest(id string, params GetUserAssignedApplicationsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Application]] {
	return core.NewBatchRequest[core.Cursor[objects.Application]](GetUserAssignedApplicationsBatchCall(id, params, options...))
}

func DecodeGetUserAssignedApplicationsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Application], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Application]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserAssignedApplicationsWithResponse(ctx context.Context, client *core.Client, id string, params GetUserAssignedApplicationsParams) (*core.Cursor[objects.Application], *core.Response, error) {
	var out core.Cursor[objects.Application]
	call := GetUserAssignedApplicationsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserAssignedApplications(ctx context.Context, client *core.Client, id string, params GetUserAssignedApplicationsParams) (*core.Cursor[objects.Application], error) {
	out, _, err := GetUserAssignedApplicationsWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserAssignedBusinessAssetGroupsParams struct {
	ContainedAssetID *core.ID    `facebook:"contained_asset_id"`
	Extra            core.Params `facebook:"-"`
}

func (params GetUserAssignedBusinessAssetGroupsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ContainedAssetID != nil {
		out["contained_asset_id"] = *params.ContainedAssetID
	}
	return out
}

func GetUserAssignedBusinessAssetGroupsBatchCall(id string, params GetUserAssignedBusinessAssetGroupsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "assigned_business_asset_groups"), params.ToParams(), options...)
}

func NewGetUserAssignedBusinessAssetGroupsBatchRequest(id string, params GetUserAssignedBusinessAssetGroupsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BusinessAssetGroup]] {
	return core.NewBatchRequest[core.Cursor[objects.BusinessAssetGroup]](GetUserAssignedBusinessAssetGroupsBatchCall(id, params, options...))
}

func DecodeGetUserAssignedBusinessAssetGroupsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BusinessAssetGroup], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.BusinessAssetGroup]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserAssignedBusinessAssetGroupsWithResponse(ctx context.Context, client *core.Client, id string, params GetUserAssignedBusinessAssetGroupsParams) (*core.Cursor[objects.BusinessAssetGroup], *core.Response, error) {
	var out core.Cursor[objects.BusinessAssetGroup]
	call := GetUserAssignedBusinessAssetGroupsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserAssignedBusinessAssetGroups(ctx context.Context, client *core.Client, id string, params GetUserAssignedBusinessAssetGroupsParams) (*core.Cursor[objects.BusinessAssetGroup], error) {
	out, _, err := GetUserAssignedBusinessAssetGroupsWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserAssignedPagesParams struct {
	Pages *[]uint64   `facebook:"pages"`
	Extra core.Params `facebook:"-"`
}

func (params GetUserAssignedPagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Pages != nil {
		out["pages"] = *params.Pages
	}
	return out
}

func GetUserAssignedPagesBatchCall(id string, params GetUserAssignedPagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "assigned_pages"), params.ToParams(), options...)
}

func NewGetUserAssignedPagesBatchRequest(id string, params GetUserAssignedPagesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Page]] {
	return core.NewBatchRequest[core.Cursor[objects.Page]](GetUserAssignedPagesBatchCall(id, params, options...))
}

func DecodeGetUserAssignedPagesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Page], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Page]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserAssignedPagesWithResponse(ctx context.Context, client *core.Client, id string, params GetUserAssignedPagesParams) (*core.Cursor[objects.Page], *core.Response, error) {
	var out core.Cursor[objects.Page]
	call := GetUserAssignedPagesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserAssignedPages(ctx context.Context, client *core.Client, id string, params GetUserAssignedPagesParams) (*core.Cursor[objects.Page], error) {
	out, _, err := GetUserAssignedPagesWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserAssignedProductCatalogsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserAssignedProductCatalogsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserAssignedProductCatalogsBatchCall(id string, params GetUserAssignedProductCatalogsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "assigned_product_catalogs"), params.ToParams(), options...)
}

func NewGetUserAssignedProductCatalogsBatchRequest(id string, params GetUserAssignedProductCatalogsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProductCatalog]] {
	return core.NewBatchRequest[core.Cursor[objects.ProductCatalog]](GetUserAssignedProductCatalogsBatchCall(id, params, options...))
}

func DecodeGetUserAssignedProductCatalogsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProductCatalog], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ProductCatalog]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserAssignedProductCatalogsWithResponse(ctx context.Context, client *core.Client, id string, params GetUserAssignedProductCatalogsParams) (*core.Cursor[objects.ProductCatalog], *core.Response, error) {
	var out core.Cursor[objects.ProductCatalog]
	call := GetUserAssignedProductCatalogsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserAssignedProductCatalogs(ctx context.Context, client *core.Client, id string, params GetUserAssignedProductCatalogsParams) (*core.Cursor[objects.ProductCatalog], error) {
	out, _, err := GetUserAssignedProductCatalogsWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserAssignedWhatsappBusinessAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserAssignedWhatsappBusinessAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserAssignedWhatsappBusinessAccountsBatchCall(id string, params GetUserAssignedWhatsappBusinessAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "assigned_whatsapp_business_accounts"), params.ToParams(), options...)
}

func NewGetUserAssignedWhatsappBusinessAccountsBatchRequest(id string, params GetUserAssignedWhatsappBusinessAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WhatsAppBusinessAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.WhatsAppBusinessAccount]](GetUserAssignedWhatsappBusinessAccountsBatchCall(id, params, options...))
}

func DecodeGetUserAssignedWhatsappBusinessAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WhatsAppBusinessAccount], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WhatsAppBusinessAccount]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserAssignedWhatsappBusinessAccountsWithResponse(ctx context.Context, client *core.Client, id string, params GetUserAssignedWhatsappBusinessAccountsParams) (*core.Cursor[objects.WhatsAppBusinessAccount], *core.Response, error) {
	var out core.Cursor[objects.WhatsAppBusinessAccount]
	call := GetUserAssignedWhatsappBusinessAccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserAssignedWhatsappBusinessAccounts(ctx context.Context, client *core.Client, id string, params GetUserAssignedWhatsappBusinessAccountsParams) (*core.Cursor[objects.WhatsAppBusinessAccount], error) {
	out, _, err := GetUserAssignedWhatsappBusinessAccountsWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserBusinessUsersParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserBusinessUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserBusinessUsersBatchCall(id string, params GetUserBusinessUsersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "business_users"), params.ToParams(), options...)
}

func NewGetUserBusinessUsersBatchRequest(id string, params GetUserBusinessUsersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BusinessUser]] {
	return core.NewBatchRequest[core.Cursor[objects.BusinessUser]](GetUserBusinessUsersBatchCall(id, params, options...))
}

func DecodeGetUserBusinessUsersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BusinessUser], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.BusinessUser]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserBusinessUsersWithResponse(ctx context.Context, client *core.Client, id string, params GetUserBusinessUsersParams) (*core.Cursor[objects.BusinessUser], *core.Response, error) {
	var out core.Cursor[objects.BusinessUser]
	call := GetUserBusinessUsersBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserBusinessUsers(ctx context.Context, client *core.Client, id string, params GetUserBusinessUsersParams) (*core.Cursor[objects.BusinessUser], error) {
	out, _, err := GetUserBusinessUsersWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteUserBusinessesParams struct {
	Business *string     `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params DeleteUserBusinessesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Business != nil {
		out["business"] = *params.Business
	}
	return out
}

func DeleteUserBusinessesBatchCall(id string, params DeleteUserBusinessesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "businesses"), params.ToParams(), options...)
}

func NewDeleteUserBusinessesBatchRequest(id string, params DeleteUserBusinessesParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteUserBusinessesBatchCall(id, params, options...))
}

func DecodeDeleteUserBusinessesBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteUserBusinessesWithResponse(ctx context.Context, client *core.Client, id string, params DeleteUserBusinessesParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteUserBusinessesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteUserBusinesses(ctx context.Context, client *core.Client, id string, params DeleteUserBusinessesParams) (*map[string]interface{}, error) {
	out, _, err := DeleteUserBusinessesWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserBusinessesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserBusinessesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserBusinessesBatchCall(id string, params GetUserBusinessesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "businesses"), params.ToParams(), options...)
}

func NewGetUserBusinessesBatchRequest(id string, params GetUserBusinessesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Business]] {
	return core.NewBatchRequest[core.Cursor[objects.Business]](GetUserBusinessesBatchCall(id, params, options...))
}

func DecodeGetUserBusinessesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Business], error) {
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

func GetUserBusinessesWithResponse(ctx context.Context, client *core.Client, id string, params GetUserBusinessesParams) (*core.Cursor[objects.Business], *core.Response, error) {
	var out core.Cursor[objects.Business]
	call := GetUserBusinessesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserBusinesses(ctx context.Context, client *core.Client, id string, params GetUserBusinessesParams) (*core.Cursor[objects.Business], error) {
	out, _, err := GetUserBusinessesWithResponse(ctx, client, id, params)
	return out, err
}

type CreateUserBusinessesParams struct {
	ChildBusinessExternalID *core.ID                                         `facebook:"child_business_external_id"`
	Email                   *string                                          `facebook:"email"`
	Name                    string                                           `facebook:"name"`
	PrimaryPage             *string                                          `facebook:"primary_page"`
	SalesRepEmail           *string                                          `facebook:"sales_rep_email"`
	SurveyBusinessType      *enums.UserbusinessesSurveyBusinessTypeEnumParam `facebook:"survey_business_type"`
	SurveyNumAssets         *uint64                                          `facebook:"survey_num_assets"`
	SurveyNumPeople         *uint64                                          `facebook:"survey_num_people"`
	TimezoneID              *enums.UserbusinessesTimezoneIDEnumParam         `facebook:"timezone_id"`
	Vertical                enums.UserbusinessesVerticalEnumParam            `facebook:"vertical"`
	Extra                   core.Params                                      `facebook:"-"`
}

func (params CreateUserBusinessesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ChildBusinessExternalID != nil {
		out["child_business_external_id"] = *params.ChildBusinessExternalID
	}
	if params.Email != nil {
		out["email"] = *params.Email
	}
	out["name"] = params.Name
	if params.PrimaryPage != nil {
		out["primary_page"] = *params.PrimaryPage
	}
	if params.SalesRepEmail != nil {
		out["sales_rep_email"] = *params.SalesRepEmail
	}
	if params.SurveyBusinessType != nil {
		out["survey_business_type"] = *params.SurveyBusinessType
	}
	if params.SurveyNumAssets != nil {
		out["survey_num_assets"] = *params.SurveyNumAssets
	}
	if params.SurveyNumPeople != nil {
		out["survey_num_people"] = *params.SurveyNumPeople
	}
	if params.TimezoneID != nil {
		out["timezone_id"] = *params.TimezoneID
	}
	out["vertical"] = params.Vertical
	return out
}

func CreateUserBusinessesBatchCall(id string, params CreateUserBusinessesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "businesses"), params.ToParams(), options...)
}

func NewCreateUserBusinessesBatchRequest(id string, params CreateUserBusinessesParams, options ...core.BatchOption) *core.BatchRequest[objects.Business] {
	return core.NewBatchRequest[objects.Business](CreateUserBusinessesBatchCall(id, params, options...))
}

func DecodeCreateUserBusinessesBatchResponse(response *core.BatchResponse) (*objects.Business, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Business
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateUserBusinessesWithResponse(ctx context.Context, client *core.Client, id string, params CreateUserBusinessesParams) (*objects.Business, *core.Response, error) {
	var out objects.Business
	call := CreateUserBusinessesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateUserBusinesses(ctx context.Context, client *core.Client, id string, params CreateUserBusinessesParams) (*objects.Business, error) {
	out, _, err := CreateUserBusinessesWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserConversationsParams struct {
	Folder   *string                                   `facebook:"folder"`
	Platform *enums.UserconversationsPlatformEnumParam `facebook:"platform"`
	Tags     *[]string                                 `facebook:"tags"`
	UserID   *core.ID                                  `facebook:"user_id"`
	Extra    core.Params                               `facebook:"-"`
}

func (params GetUserConversationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Folder != nil {
		out["folder"] = *params.Folder
	}
	if params.Platform != nil {
		out["platform"] = *params.Platform
	}
	if params.Tags != nil {
		out["tags"] = *params.Tags
	}
	if params.UserID != nil {
		out["user_id"] = *params.UserID
	}
	return out
}

func GetUserConversationsBatchCall(id string, params GetUserConversationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "conversations"), params.ToParams(), options...)
}

func NewGetUserConversationsBatchRequest(id string, params GetUserConversationsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.UnifiedThread]] {
	return core.NewBatchRequest[core.Cursor[objects.UnifiedThread]](GetUserConversationsBatchCall(id, params, options...))
}

func DecodeGetUserConversationsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.UnifiedThread], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.UnifiedThread]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserConversationsWithResponse(ctx context.Context, client *core.Client, id string, params GetUserConversationsParams) (*core.Cursor[objects.UnifiedThread], *core.Response, error) {
	var out core.Cursor[objects.UnifiedThread]
	call := GetUserConversationsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserConversations(ctx context.Context, client *core.Client, id string, params GetUserConversationsParams) (*core.Cursor[objects.UnifiedThread], error) {
	out, _, err := GetUserConversationsWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserCustomLabelsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserCustomLabelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserCustomLabelsBatchCall(id string, params GetUserCustomLabelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "custom_labels"), params.ToParams(), options...)
}

func NewGetUserCustomLabelsBatchRequest(id string, params GetUserCustomLabelsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.PageUserMessageThreadLabel]] {
	return core.NewBatchRequest[core.Cursor[objects.PageUserMessageThreadLabel]](GetUserCustomLabelsBatchCall(id, params, options...))
}

func DecodeGetUserCustomLabelsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.PageUserMessageThreadLabel], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.PageUserMessageThreadLabel]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserCustomLabelsWithResponse(ctx context.Context, client *core.Client, id string, params GetUserCustomLabelsParams) (*core.Cursor[objects.PageUserMessageThreadLabel], *core.Response, error) {
	var out core.Cursor[objects.PageUserMessageThreadLabel]
	call := GetUserCustomLabelsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserCustomLabels(ctx context.Context, client *core.Client, id string, params GetUserCustomLabelsParams) (*core.Cursor[objects.PageUserMessageThreadLabel], error) {
	out, _, err := GetUserCustomLabelsWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserEventsParams struct {
	IncludeCanceled *bool                          `facebook:"include_canceled"`
	Type            *enums.UsereventsTypeEnumParam `facebook:"type"`
	Extra           core.Params                    `facebook:"-"`
}

func (params GetUserEventsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.IncludeCanceled != nil {
		out["include_canceled"] = *params.IncludeCanceled
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	return out
}

func GetUserEventsBatchCall(id string, params GetUserEventsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "events"), params.ToParams(), options...)
}

func NewGetUserEventsBatchRequest(id string, params GetUserEventsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Event]] {
	return core.NewBatchRequest[core.Cursor[objects.Event]](GetUserEventsBatchCall(id, params, options...))
}

func DecodeGetUserEventsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Event], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Event]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserEventsWithResponse(ctx context.Context, client *core.Client, id string, params GetUserEventsParams) (*core.Cursor[objects.Event], *core.Response, error) {
	var out core.Cursor[objects.Event]
	call := GetUserEventsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserEvents(ctx context.Context, client *core.Client, id string, params GetUserEventsParams) (*core.Cursor[objects.Event], error) {
	out, _, err := GetUserEventsWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserFeedParams struct {
	IncludeHidden *bool       `facebook:"include_hidden"`
	Q             *string     `facebook:"q"`
	ShowExpired   *bool       `facebook:"show_expired"`
	Since         *core.Time  `facebook:"since"`
	Until         *core.Time  `facebook:"until"`
	With          *string     `facebook:"with"`
	Extra         core.Params `facebook:"-"`
}

func (params GetUserFeedParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.IncludeHidden != nil {
		out["include_hidden"] = *params.IncludeHidden
	}
	if params.Q != nil {
		out["q"] = *params.Q
	}
	if params.ShowExpired != nil {
		out["show_expired"] = *params.ShowExpired
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	if params.With != nil {
		out["with"] = *params.With
	}
	return out
}

func GetUserFeedBatchCall(id string, params GetUserFeedParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "feed"), params.ToParams(), options...)
}

func NewGetUserFeedBatchRequest(id string, params GetUserFeedParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Post]] {
	return core.NewBatchRequest[core.Cursor[objects.Post]](GetUserFeedBatchCall(id, params, options...))
}

func DecodeGetUserFeedBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Post], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Post]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserFeedWithResponse(ctx context.Context, client *core.Client, id string, params GetUserFeedParams) (*core.Cursor[objects.Post], *core.Response, error) {
	var out core.Cursor[objects.Post]
	call := GetUserFeedBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserFeed(ctx context.Context, client *core.Client, id string, params GetUserFeedParams) (*core.Cursor[objects.Post], error) {
	out, _, err := GetUserFeedWithResponse(ctx, client, id, params)
	return out, err
}

type CreateUserFeedParams struct {
	Actions                   *map[string]interface{}                          `facebook:"actions"`
	AlbumID                   *core.ID                                         `facebook:"album_id"`
	AndroidKeyHash            *string                                          `facebook:"android_key_hash"`
	ApplicationID             *core.ID                                         `facebook:"application_id"`
	AskedFunFactPromptID      *core.ID                                         `facebook:"asked_fun_fact_prompt_id"`
	Asset3dID                 *core.ID                                         `facebook:"asset3d_id"`
	AssociatedID              *core.ID                                         `facebook:"associated_id"`
	AttachPlaceSuggestion     *bool                                            `facebook:"attach_place_suggestion"`
	AttachedMedia             *[]map[string]interface{}                        `facebook:"attached_media"`
	AudienceExp               *bool                                            `facebook:"audience_exp"`
	BackdatedTime             *core.Time                                       `facebook:"backdated_time"`
	BackdatedTimeGranularity  *enums.UserfeedBackdatedTimeGranularityEnumParam `facebook:"backdated_time_granularity"`
	BreakingNews              *bool                                            `facebook:"breaking_news"`
	BreakingNewsExpiration    *uint64                                          `facebook:"breaking_news_expiration"`
	CallToAction              *map[string]interface{}                          `facebook:"call_to_action"`
	Caption                   *string                                          `facebook:"caption"`
	ChildAttachments          *[]map[string]interface{}                        `facebook:"child_attachments"`
	ClientMutationID          *core.ID                                         `facebook:"client_mutation_id"`
	ComposerEntryPicker       *string                                          `facebook:"composer_entry_picker"`
	ComposerEntryPoint        *string                                          `facebook:"composer_entry_point"`
	ComposerEntryTime         *uint64                                          `facebook:"composer_entry_time"`
	ComposerSessionEventsLog  *string                                          `facebook:"composer_session_events_log"`
	ComposerSessionID         *core.ID                                         `facebook:"composer_session_id"`
	ComposerSourceSurface     *string                                          `facebook:"composer_source_surface"`
	ComposerType              *string                                          `facebook:"composer_type"`
	ConnectionClass           *string                                          `facebook:"connection_class"`
	ContentAttachment         *string                                          `facebook:"content_attachment"`
	Coordinates               *map[string]interface{}                          `facebook:"coordinates"`
	CtaLink                   *string                                          `facebook:"cta_link"`
	CtaType                   *string                                          `facebook:"cta_type"`
	Description               *string                                          `facebook:"description"`
	DirectShareStatus         *uint64                                          `facebook:"direct_share_status"`
	ExpandedHeight            *uint64                                          `facebook:"expanded_height"`
	ExpandedWidth             *uint64                                          `facebook:"expanded_width"`
	FeedTargeting             *map[string]interface{}                          `facebook:"feed_targeting"`
	Formatting                *enums.UserfeedFormattingEnumParam               `facebook:"formatting"`
	FunFactPromptID           *core.ID                                         `facebook:"fun_fact_prompt_id"`
	FunFactToasteeID          *core.ID                                         `facebook:"fun_fact_toastee_id"`
	Height                    *uint64                                          `facebook:"height"`
	HomeCheckinCityID         *map[string]interface{}                          `facebook:"home_checkin_city_id"`
	ImageCrops                *map[string]interface{}                          `facebook:"image_crops"`
	ImplicitWithTags          *[]int                                           `facebook:"implicit_with_tags"`
	InstantGameEntryPointData *string                                          `facebook:"instant_game_entry_point_data"`
	IosBundleID               *core.ID                                         `facebook:"ios_bundle_id"`
	IsBackoutDraft            *bool                                            `facebook:"is_backout_draft"`
	IsBoostIntended           *bool                                            `facebook:"is_boost_intended"`
	IsExplicitLocation        *bool                                            `facebook:"is_explicit_location"`
	IsExplicitShare           *bool                                            `facebook:"is_explicit_share"`
	IsGroupLinkingPost        *bool                                            `facebook:"is_group_linking_post"`
	IsPhotoContainer          *bool                                            `facebook:"is_photo_container"`
	Link                      *string                                          `facebook:"link"`
	LocationSourceID          *core.ID                                         `facebook:"location_source_id"`
	ManualPrivacy             *bool                                            `facebook:"manual_privacy"`
	Message                   *string                                          `facebook:"message"`
	MultiShareEndCard         *bool                                            `facebook:"multi_share_end_card"`
	MultiShareOptimized       *bool                                            `facebook:"multi_share_optimized"`
	Name                      *string                                          `facebook:"name"`
	NectarModule              *string                                          `facebook:"nectar_module"`
	ObjectAttachment          *string                                          `facebook:"object_attachment"`
	OgActionTypeID            *core.ID                                         `facebook:"og_action_type_id"`
	OgHideObjectAttachment    *bool                                            `facebook:"og_hide_object_attachment"`
	OgIconID                  *core.ID                                         `facebook:"og_icon_id"`
	OgObjectID                *core.ID                                         `facebook:"og_object_id"`
	OgPhrase                  *string                                          `facebook:"og_phrase"`
	OgSetProfileBadge         *bool                                            `facebook:"og_set_profile_badge"`
	OgSuggestionMechanism     *string                                          `facebook:"og_suggestion_mechanism"`
	PageRecommendation        *string                                          `facebook:"page_recommendation"`
	Picture                   *string                                          `facebook:"picture"`
	Place                     *map[string]interface{}                          `facebook:"place"`
	PlaceAttachmentSetting    *enums.UserfeedPlaceAttachmentSettingEnumParam   `facebook:"place_attachment_setting"`
	PlaceList                 *string                                          `facebook:"place_list"`
	PlaceListData             *[]interface{}                                   `facebook:"place_list_data"`
	PostSurfacesBlacklist     *[]enums.UserfeedPostSurfacesBlacklistEnumParam  `facebook:"post_surfaces_blacklist"`
	PostingToRedspace         *enums.UserfeedPostingToRedspaceEnumParam        `facebook:"posting_to_redspace"`
	Privacy                   *string                                          `facebook:"privacy"`
	PromptID                  *core.ID                                         `facebook:"prompt_id"`
	PromptTrackingString      *string                                          `facebook:"prompt_tracking_string"`
	Properties                *map[string]interface{}                          `facebook:"properties"`
	ProxiedAppID              *core.ID                                         `facebook:"proxied_app_id"`
	PublishEventID            *core.ID                                         `facebook:"publish_event_id"`
	Published                 *bool                                            `facebook:"published"`
	Quote                     *string                                          `facebook:"quote"`
	Ref                       *[]string                                        `facebook:"ref"`
	ReferenceableImageIds     *[]core.ID                                       `facebook:"referenceable_image_ids"`
	ReferralID                *core.ID                                         `facebook:"referral_id"`
	ScheduledPublishTime      *core.Time                                       `facebook:"scheduled_publish_time"`
	Source                    *string                                          `facebook:"source"`
	SponsorID                 *core.ID                                         `facebook:"sponsor_id"`
	SponsorRelationship       *uint64                                          `facebook:"sponsor_relationship"`
	SuggestedPlaceID          *map[string]interface{}                          `facebook:"suggested_place_id"`
	Tags                      *[]int                                           `facebook:"tags"`
	TargetSurface             *enums.UserfeedTargetSurfaceEnumParam            `facebook:"target_surface"`
	Targeting                 *map[string]interface{}                          `facebook:"targeting"`
	TextFormatMetadata        *string                                          `facebook:"text_format_metadata"`
	TextFormatPresetID        *core.ID                                         `facebook:"text_format_preset_id"`
	TextOnlyPlace             *string                                          `facebook:"text_only_place"`
	Thumbnail                 *core.FileParam                                  `facebook:"thumbnail"`
	TimeSinceOriginalPost     *uint64                                          `facebook:"time_since_original_post"`
	Title                     *string                                          `facebook:"title"`
	TrackingInfo              *string                                          `facebook:"tracking_info"`
	UnpublishedContentType    *enums.UserfeedUnpublishedContentTypeEnumParam   `facebook:"unpublished_content_type"`
	UserSelectedTags          *bool                                            `facebook:"user_selected_tags"`
	VideoStartTimeMs          *uint64                                          `facebook:"video_start_time_ms"`
	ViewerCoordinates         *map[string]interface{}                          `facebook:"viewer_coordinates"`
	Width                     *uint64                                          `facebook:"width"`
	Extra                     core.Params                                      `facebook:"-"`
}

func (params CreateUserFeedParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Actions != nil {
		out["actions"] = *params.Actions
	}
	if params.AlbumID != nil {
		out["album_id"] = *params.AlbumID
	}
	if params.AndroidKeyHash != nil {
		out["android_key_hash"] = *params.AndroidKeyHash
	}
	if params.ApplicationID != nil {
		out["application_id"] = *params.ApplicationID
	}
	if params.AskedFunFactPromptID != nil {
		out["asked_fun_fact_prompt_id"] = *params.AskedFunFactPromptID
	}
	if params.Asset3dID != nil {
		out["asset3d_id"] = *params.Asset3dID
	}
	if params.AssociatedID != nil {
		out["associated_id"] = *params.AssociatedID
	}
	if params.AttachPlaceSuggestion != nil {
		out["attach_place_suggestion"] = *params.AttachPlaceSuggestion
	}
	if params.AttachedMedia != nil {
		out["attached_media"] = *params.AttachedMedia
	}
	if params.AudienceExp != nil {
		out["audience_exp"] = *params.AudienceExp
	}
	if params.BackdatedTime != nil {
		out["backdated_time"] = *params.BackdatedTime
	}
	if params.BackdatedTimeGranularity != nil {
		out["backdated_time_granularity"] = *params.BackdatedTimeGranularity
	}
	if params.BreakingNews != nil {
		out["breaking_news"] = *params.BreakingNews
	}
	if params.BreakingNewsExpiration != nil {
		out["breaking_news_expiration"] = *params.BreakingNewsExpiration
	}
	if params.CallToAction != nil {
		out["call_to_action"] = *params.CallToAction
	}
	if params.Caption != nil {
		out["caption"] = *params.Caption
	}
	if params.ChildAttachments != nil {
		out["child_attachments"] = *params.ChildAttachments
	}
	if params.ClientMutationID != nil {
		out["client_mutation_id"] = *params.ClientMutationID
	}
	if params.ComposerEntryPicker != nil {
		out["composer_entry_picker"] = *params.ComposerEntryPicker
	}
	if params.ComposerEntryPoint != nil {
		out["composer_entry_point"] = *params.ComposerEntryPoint
	}
	if params.ComposerEntryTime != nil {
		out["composer_entry_time"] = *params.ComposerEntryTime
	}
	if params.ComposerSessionEventsLog != nil {
		out["composer_session_events_log"] = *params.ComposerSessionEventsLog
	}
	if params.ComposerSessionID != nil {
		out["composer_session_id"] = *params.ComposerSessionID
	}
	if params.ComposerSourceSurface != nil {
		out["composer_source_surface"] = *params.ComposerSourceSurface
	}
	if params.ComposerType != nil {
		out["composer_type"] = *params.ComposerType
	}
	if params.ConnectionClass != nil {
		out["connection_class"] = *params.ConnectionClass
	}
	if params.ContentAttachment != nil {
		out["content_attachment"] = *params.ContentAttachment
	}
	if params.Coordinates != nil {
		out["coordinates"] = *params.Coordinates
	}
	if params.CtaLink != nil {
		out["cta_link"] = *params.CtaLink
	}
	if params.CtaType != nil {
		out["cta_type"] = *params.CtaType
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.DirectShareStatus != nil {
		out["direct_share_status"] = *params.DirectShareStatus
	}
	if params.ExpandedHeight != nil {
		out["expanded_height"] = *params.ExpandedHeight
	}
	if params.ExpandedWidth != nil {
		out["expanded_width"] = *params.ExpandedWidth
	}
	if params.FeedTargeting != nil {
		out["feed_targeting"] = *params.FeedTargeting
	}
	if params.Formatting != nil {
		out["formatting"] = *params.Formatting
	}
	if params.FunFactPromptID != nil {
		out["fun_fact_prompt_id"] = *params.FunFactPromptID
	}
	if params.FunFactToasteeID != nil {
		out["fun_fact_toastee_id"] = *params.FunFactToasteeID
	}
	if params.Height != nil {
		out["height"] = *params.Height
	}
	if params.HomeCheckinCityID != nil {
		out["home_checkin_city_id"] = *params.HomeCheckinCityID
	}
	if params.ImageCrops != nil {
		out["image_crops"] = *params.ImageCrops
	}
	if params.ImplicitWithTags != nil {
		out["implicit_with_tags"] = *params.ImplicitWithTags
	}
	if params.InstantGameEntryPointData != nil {
		out["instant_game_entry_point_data"] = *params.InstantGameEntryPointData
	}
	if params.IosBundleID != nil {
		out["ios_bundle_id"] = *params.IosBundleID
	}
	if params.IsBackoutDraft != nil {
		out["is_backout_draft"] = *params.IsBackoutDraft
	}
	if params.IsBoostIntended != nil {
		out["is_boost_intended"] = *params.IsBoostIntended
	}
	if params.IsExplicitLocation != nil {
		out["is_explicit_location"] = *params.IsExplicitLocation
	}
	if params.IsExplicitShare != nil {
		out["is_explicit_share"] = *params.IsExplicitShare
	}
	if params.IsGroupLinkingPost != nil {
		out["is_group_linking_post"] = *params.IsGroupLinkingPost
	}
	if params.IsPhotoContainer != nil {
		out["is_photo_container"] = *params.IsPhotoContainer
	}
	if params.Link != nil {
		out["link"] = *params.Link
	}
	if params.LocationSourceID != nil {
		out["location_source_id"] = *params.LocationSourceID
	}
	if params.ManualPrivacy != nil {
		out["manual_privacy"] = *params.ManualPrivacy
	}
	if params.Message != nil {
		out["message"] = *params.Message
	}
	if params.MultiShareEndCard != nil {
		out["multi_share_end_card"] = *params.MultiShareEndCard
	}
	if params.MultiShareOptimized != nil {
		out["multi_share_optimized"] = *params.MultiShareOptimized
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.NectarModule != nil {
		out["nectar_module"] = *params.NectarModule
	}
	if params.ObjectAttachment != nil {
		out["object_attachment"] = *params.ObjectAttachment
	}
	if params.OgActionTypeID != nil {
		out["og_action_type_id"] = *params.OgActionTypeID
	}
	if params.OgHideObjectAttachment != nil {
		out["og_hide_object_attachment"] = *params.OgHideObjectAttachment
	}
	if params.OgIconID != nil {
		out["og_icon_id"] = *params.OgIconID
	}
	if params.OgObjectID != nil {
		out["og_object_id"] = *params.OgObjectID
	}
	if params.OgPhrase != nil {
		out["og_phrase"] = *params.OgPhrase
	}
	if params.OgSetProfileBadge != nil {
		out["og_set_profile_badge"] = *params.OgSetProfileBadge
	}
	if params.OgSuggestionMechanism != nil {
		out["og_suggestion_mechanism"] = *params.OgSuggestionMechanism
	}
	if params.PageRecommendation != nil {
		out["page_recommendation"] = *params.PageRecommendation
	}
	if params.Picture != nil {
		out["picture"] = *params.Picture
	}
	if params.Place != nil {
		out["place"] = *params.Place
	}
	if params.PlaceAttachmentSetting != nil {
		out["place_attachment_setting"] = *params.PlaceAttachmentSetting
	}
	if params.PlaceList != nil {
		out["place_list"] = *params.PlaceList
	}
	if params.PlaceListData != nil {
		out["place_list_data"] = *params.PlaceListData
	}
	if params.PostSurfacesBlacklist != nil {
		out["post_surfaces_blacklist"] = *params.PostSurfacesBlacklist
	}
	if params.PostingToRedspace != nil {
		out["posting_to_redspace"] = *params.PostingToRedspace
	}
	if params.Privacy != nil {
		out["privacy"] = *params.Privacy
	}
	if params.PromptID != nil {
		out["prompt_id"] = *params.PromptID
	}
	if params.PromptTrackingString != nil {
		out["prompt_tracking_string"] = *params.PromptTrackingString
	}
	if params.Properties != nil {
		out["properties"] = *params.Properties
	}
	if params.ProxiedAppID != nil {
		out["proxied_app_id"] = *params.ProxiedAppID
	}
	if params.PublishEventID != nil {
		out["publish_event_id"] = *params.PublishEventID
	}
	if params.Published != nil {
		out["published"] = *params.Published
	}
	if params.Quote != nil {
		out["quote"] = *params.Quote
	}
	if params.Ref != nil {
		out["ref"] = *params.Ref
	}
	if params.ReferenceableImageIds != nil {
		out["referenceable_image_ids"] = *params.ReferenceableImageIds
	}
	if params.ReferralID != nil {
		out["referral_id"] = *params.ReferralID
	}
	if params.ScheduledPublishTime != nil {
		out["scheduled_publish_time"] = *params.ScheduledPublishTime
	}
	if params.Source != nil {
		out["source"] = *params.Source
	}
	if params.SponsorID != nil {
		out["sponsor_id"] = *params.SponsorID
	}
	if params.SponsorRelationship != nil {
		out["sponsor_relationship"] = *params.SponsorRelationship
	}
	if params.SuggestedPlaceID != nil {
		out["suggested_place_id"] = *params.SuggestedPlaceID
	}
	if params.Tags != nil {
		out["tags"] = *params.Tags
	}
	if params.TargetSurface != nil {
		out["target_surface"] = *params.TargetSurface
	}
	if params.Targeting != nil {
		out["targeting"] = *params.Targeting
	}
	if params.TextFormatMetadata != nil {
		out["text_format_metadata"] = *params.TextFormatMetadata
	}
	if params.TextFormatPresetID != nil {
		out["text_format_preset_id"] = *params.TextFormatPresetID
	}
	if params.TextOnlyPlace != nil {
		out["text_only_place"] = *params.TextOnlyPlace
	}
	if params.Thumbnail != nil {
		out["thumbnail"] = *params.Thumbnail
	}
	if params.TimeSinceOriginalPost != nil {
		out["time_since_original_post"] = *params.TimeSinceOriginalPost
	}
	if params.Title != nil {
		out["title"] = *params.Title
	}
	if params.TrackingInfo != nil {
		out["tracking_info"] = *params.TrackingInfo
	}
	if params.UnpublishedContentType != nil {
		out["unpublished_content_type"] = *params.UnpublishedContentType
	}
	if params.UserSelectedTags != nil {
		out["user_selected_tags"] = *params.UserSelectedTags
	}
	if params.VideoStartTimeMs != nil {
		out["video_start_time_ms"] = *params.VideoStartTimeMs
	}
	if params.ViewerCoordinates != nil {
		out["viewer_coordinates"] = *params.ViewerCoordinates
	}
	if params.Width != nil {
		out["width"] = *params.Width
	}
	return out
}

func CreateUserFeedBatchCall(id string, params CreateUserFeedParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "feed"), params.ToParams(), options...)
}

func NewCreateUserFeedBatchRequest(id string, params CreateUserFeedParams, options ...core.BatchOption) *core.BatchRequest[objects.Post] {
	return core.NewBatchRequest[objects.Post](CreateUserFeedBatchCall(id, params, options...))
}

func DecodeCreateUserFeedBatchResponse(response *core.BatchResponse) (*objects.Post, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Post
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateUserFeedWithResponse(ctx context.Context, client *core.Client, id string, params CreateUserFeedParams) (*objects.Post, *core.Response, error) {
	var out objects.Post
	call := CreateUserFeedBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateUserFeed(ctx context.Context, client *core.Client, id string, params CreateUserFeedParams) (*objects.Post, error) {
	out, _, err := CreateUserFeedWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserFriendsParams struct {
	UID   *core.ID    `facebook:"uid"`
	Extra core.Params `facebook:"-"`
}

func (params GetUserFriendsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.UID != nil {
		out["uid"] = *params.UID
	}
	return out
}

func GetUserFriendsBatchCall(id string, params GetUserFriendsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "friends"), params.ToParams(), options...)
}

func NewGetUserFriendsBatchRequest(id string, params GetUserFriendsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.User]] {
	return core.NewBatchRequest[core.Cursor[objects.User]](GetUserFriendsBatchCall(id, params, options...))
}

func DecodeGetUserFriendsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.User], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.User]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserFriendsWithResponse(ctx context.Context, client *core.Client, id string, params GetUserFriendsParams) (*core.Cursor[objects.User], *core.Response, error) {
	var out core.Cursor[objects.User]
	call := GetUserFriendsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserFriends(ctx context.Context, client *core.Client, id string, params GetUserFriendsParams) (*core.Cursor[objects.User], error) {
	out, _, err := GetUserFriendsWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserFundraisersParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserFundraisersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserFundraisersBatchCall(id string, params GetUserFundraisersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "fundraisers"), params.ToParams(), options...)
}

func NewGetUserFundraisersBatchRequest(id string, params GetUserFundraisersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.FundraiserPersonToCharity]] {
	return core.NewBatchRequest[core.Cursor[objects.FundraiserPersonToCharity]](GetUserFundraisersBatchCall(id, params, options...))
}

func DecodeGetUserFundraisersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.FundraiserPersonToCharity], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.FundraiserPersonToCharity]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserFundraisersWithResponse(ctx context.Context, client *core.Client, id string, params GetUserFundraisersParams) (*core.Cursor[objects.FundraiserPersonToCharity], *core.Response, error) {
	var out core.Cursor[objects.FundraiserPersonToCharity]
	call := GetUserFundraisersBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserFundraisers(ctx context.Context, client *core.Client, id string, params GetUserFundraisersParams) (*core.Cursor[objects.FundraiserPersonToCharity], error) {
	out, _, err := GetUserFundraisersWithResponse(ctx, client, id, params)
	return out, err
}

type CreateUserFundraisersParams struct {
	CharityID              *core.ID                                     `facebook:"charity_id"`
	CoverPhoto             *core.FileParam                              `facebook:"cover_photo"`
	Currency               string                                       `facebook:"currency"`
	Description            string                                       `facebook:"description"`
	EndTime                core.Time                                    `facebook:"end_time"`
	ExternalEventName      *string                                      `facebook:"external_event_name"`
	ExternalEventStartTime *core.Time                                   `facebook:"external_event_start_time"`
	ExternalEventURI       *string                                      `facebook:"external_event_uri"`
	ExternalFundraiserURI  *string                                      `facebook:"external_fundraiser_uri"`
	ExternalID             core.ID                                      `facebook:"external_id"`
	FundraiserType         enums.UserfundraisersFundraiserTypeEnumParam `facebook:"fundraiser_type"`
	GoalAmount             uint64                                       `facebook:"goal_amount"`
	Name                   string                                       `facebook:"name"`
	PageID                 *core.ID                                     `facebook:"page_id"`
	Extra                  core.Params                                  `facebook:"-"`
}

func (params CreateUserFundraisersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CharityID != nil {
		out["charity_id"] = *params.CharityID
	}
	if params.CoverPhoto != nil {
		out["cover_photo"] = *params.CoverPhoto
	}
	out["currency"] = params.Currency
	out["description"] = params.Description
	out["end_time"] = params.EndTime
	if params.ExternalEventName != nil {
		out["external_event_name"] = *params.ExternalEventName
	}
	if params.ExternalEventStartTime != nil {
		out["external_event_start_time"] = *params.ExternalEventStartTime
	}
	if params.ExternalEventURI != nil {
		out["external_event_uri"] = *params.ExternalEventURI
	}
	if params.ExternalFundraiserURI != nil {
		out["external_fundraiser_uri"] = *params.ExternalFundraiserURI
	}
	out["external_id"] = params.ExternalID
	out["fundraiser_type"] = params.FundraiserType
	out["goal_amount"] = params.GoalAmount
	out["name"] = params.Name
	if params.PageID != nil {
		out["page_id"] = *params.PageID
	}
	return out
}

func CreateUserFundraisersBatchCall(id string, params CreateUserFundraisersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "fundraisers"), params.ToParams(), options...)
}

func NewCreateUserFundraisersBatchRequest(id string, params CreateUserFundraisersParams, options ...core.BatchOption) *core.BatchRequest[objects.FundraiserPersonToCharity] {
	return core.NewBatchRequest[objects.FundraiserPersonToCharity](CreateUserFundraisersBatchCall(id, params, options...))
}

func DecodeCreateUserFundraisersBatchResponse(response *core.BatchResponse) (*objects.FundraiserPersonToCharity, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.FundraiserPersonToCharity
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateUserFundraisersWithResponse(ctx context.Context, client *core.Client, id string, params CreateUserFundraisersParams) (*objects.FundraiserPersonToCharity, *core.Response, error) {
	var out objects.FundraiserPersonToCharity
	call := CreateUserFundraisersBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateUserFundraisers(ctx context.Context, client *core.Client, id string, params CreateUserFundraisersParams) (*objects.FundraiserPersonToCharity, error) {
	out, _, err := CreateUserFundraisersWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserGroupsParams struct {
	AdminOnly *bool       `facebook:"admin_only"`
	Parent    *string     `facebook:"parent"`
	Extra     core.Params `facebook:"-"`
}

func (params GetUserGroupsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdminOnly != nil {
		out["admin_only"] = *params.AdminOnly
	}
	if params.Parent != nil {
		out["parent"] = *params.Parent
	}
	return out
}

func GetUserGroupsBatchCall(id string, params GetUserGroupsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "groups"), params.ToParams(), options...)
}

func NewGetUserGroupsBatchRequest(id string, params GetUserGroupsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Group]] {
	return core.NewBatchRequest[core.Cursor[objects.Group]](GetUserGroupsBatchCall(id, params, options...))
}

func DecodeGetUserGroupsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Group], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Group]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserGroupsWithResponse(ctx context.Context, client *core.Client, id string, params GetUserGroupsParams) (*core.Cursor[objects.Group], *core.Response, error) {
	var out core.Cursor[objects.Group]
	call := GetUserGroupsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserGroups(ctx context.Context, client *core.Client, id string, params GetUserGroupsParams) (*core.Cursor[objects.Group], error) {
	out, _, err := GetUserGroupsWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserIdsForAppsParams struct {
	App   *int        `facebook:"app"`
	Extra core.Params `facebook:"-"`
}

func (params GetUserIdsForAppsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.App != nil {
		out["app"] = *params.App
	}
	return out
}

func GetUserIdsForAppsBatchCall(id string, params GetUserIdsForAppsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ids_for_apps"), params.ToParams(), options...)
}

func NewGetUserIdsForAppsBatchRequest(id string, params GetUserIdsForAppsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.UserIDForApp]] {
	return core.NewBatchRequest[core.Cursor[objects.UserIDForApp]](GetUserIdsForAppsBatchCall(id, params, options...))
}

func DecodeGetUserIdsForAppsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.UserIDForApp], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.UserIDForApp]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserIdsForAppsWithResponse(ctx context.Context, client *core.Client, id string, params GetUserIdsForAppsParams) (*core.Cursor[objects.UserIDForApp], *core.Response, error) {
	var out core.Cursor[objects.UserIDForApp]
	call := GetUserIdsForAppsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserIdsForApps(ctx context.Context, client *core.Client, id string, params GetUserIdsForAppsParams) (*core.Cursor[objects.UserIDForApp], error) {
	out, _, err := GetUserIdsForAppsWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserIdsForBusinessParams struct {
	App   *int        `facebook:"app"`
	Extra core.Params `facebook:"-"`
}

func (params GetUserIdsForBusinessParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.App != nil {
		out["app"] = *params.App
	}
	return out
}

func GetUserIdsForBusinessBatchCall(id string, params GetUserIdsForBusinessParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ids_for_business"), params.ToParams(), options...)
}

func NewGetUserIdsForBusinessBatchRequest(id string, params GetUserIdsForBusinessParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.UserIDForApp]] {
	return core.NewBatchRequest[core.Cursor[objects.UserIDForApp]](GetUserIdsForBusinessBatchCall(id, params, options...))
}

func DecodeGetUserIdsForBusinessBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.UserIDForApp], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.UserIDForApp]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserIdsForBusinessWithResponse(ctx context.Context, client *core.Client, id string, params GetUserIdsForBusinessParams) (*core.Cursor[objects.UserIDForApp], *core.Response, error) {
	var out core.Cursor[objects.UserIDForApp]
	call := GetUserIdsForBusinessBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserIdsForBusiness(ctx context.Context, client *core.Client, id string, params GetUserIdsForBusinessParams) (*core.Cursor[objects.UserIDForApp], error) {
	out, _, err := GetUserIdsForBusinessWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserIdsForPagesParams struct {
	Page  *int        `facebook:"page"`
	Extra core.Params `facebook:"-"`
}

func (params GetUserIdsForPagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Page != nil {
		out["page"] = *params.Page
	}
	return out
}

func GetUserIdsForPagesBatchCall(id string, params GetUserIdsForPagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ids_for_pages"), params.ToParams(), options...)
}

func NewGetUserIdsForPagesBatchRequest(id string, params GetUserIdsForPagesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.UserIDForPage]] {
	return core.NewBatchRequest[core.Cursor[objects.UserIDForPage]](GetUserIdsForPagesBatchCall(id, params, options...))
}

func DecodeGetUserIdsForPagesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.UserIDForPage], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.UserIDForPage]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserIdsForPagesWithResponse(ctx context.Context, client *core.Client, id string, params GetUserIdsForPagesParams) (*core.Cursor[objects.UserIDForPage], *core.Response, error) {
	var out core.Cursor[objects.UserIDForPage]
	call := GetUserIdsForPagesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserIdsForPages(ctx context.Context, client *core.Client, id string, params GetUserIdsForPagesParams) (*core.Cursor[objects.UserIDForPage], error) {
	out, _, err := GetUserIdsForPagesWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserLikesParams struct {
	TargetID *core.ID    `facebook:"target_id"`
	Extra    core.Params `facebook:"-"`
}

func (params GetUserLikesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.TargetID != nil {
		out["target_id"] = *params.TargetID
	}
	return out
}

func GetUserLikesBatchCall(id string, params GetUserLikesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "likes"), params.ToParams(), options...)
}

func NewGetUserLikesBatchRequest(id string, params GetUserLikesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Page]] {
	return core.NewBatchRequest[core.Cursor[objects.Page]](GetUserLikesBatchCall(id, params, options...))
}

func DecodeGetUserLikesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Page], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Page]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserLikesWithResponse(ctx context.Context, client *core.Client, id string, params GetUserLikesParams) (*core.Cursor[objects.Page], *core.Response, error) {
	var out core.Cursor[objects.Page]
	call := GetUserLikesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserLikes(ctx context.Context, client *core.Client, id string, params GetUserLikesParams) (*core.Cursor[objects.Page], error) {
	out, _, err := GetUserLikesWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserLiveVideosParams struct {
	BroadcastStatus *[]enums.UserliveVideosBroadcastStatusEnumParam `facebook:"broadcast_status"`
	Source          *enums.UserliveVideosSourceEnumParam            `facebook:"source"`
	Extra           core.Params                                     `facebook:"-"`
}

func (params GetUserLiveVideosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BroadcastStatus != nil {
		out["broadcast_status"] = *params.BroadcastStatus
	}
	if params.Source != nil {
		out["source"] = *params.Source
	}
	return out
}

func GetUserLiveVideosBatchCall(id string, params GetUserLiveVideosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "live_videos"), params.ToParams(), options...)
}

func NewGetUserLiveVideosBatchRequest(id string, params GetUserLiveVideosParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.LiveVideo]] {
	return core.NewBatchRequest[core.Cursor[objects.LiveVideo]](GetUserLiveVideosBatchCall(id, params, options...))
}

func DecodeGetUserLiveVideosBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.LiveVideo], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.LiveVideo]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserLiveVideosWithResponse(ctx context.Context, client *core.Client, id string, params GetUserLiveVideosParams) (*core.Cursor[objects.LiveVideo], *core.Response, error) {
	var out core.Cursor[objects.LiveVideo]
	call := GetUserLiveVideosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserLiveVideos(ctx context.Context, client *core.Client, id string, params GetUserLiveVideosParams) (*core.Cursor[objects.LiveVideo], error) {
	out, _, err := GetUserLiveVideosWithResponse(ctx, client, id, params)
	return out, err
}

type CreateUserLiveVideosParams struct {
	ContentTags                *[]string                                        `facebook:"content_tags"`
	Description                *string                                          `facebook:"description"`
	EnableBackupIngest         *bool                                            `facebook:"enable_backup_ingest"`
	EncodingSettings           *string                                          `facebook:"encoding_settings"`
	EventParams                *map[string]interface{}                          `facebook:"event_params"`
	FisheyeVideoCropped        *bool                                            `facebook:"fisheye_video_cropped"`
	FrontZRotation             *float64                                         `facebook:"front_z_rotation"`
	IsAudioOnly                *bool                                            `facebook:"is_audio_only"`
	IsSpherical                *bool                                            `facebook:"is_spherical"`
	OriginalFov                *uint64                                          `facebook:"original_fov"`
	Privacy                    *string                                          `facebook:"privacy"`
	Projection                 *enums.UserliveVideosProjectionEnumParam         `facebook:"projection"`
	Published                  *bool                                            `facebook:"published"`
	ScheduleCustomProfileImage *core.FileParam                                  `facebook:"schedule_custom_profile_image"`
	SpatialAudioFormat         *enums.UserliveVideosSpatialAudioFormatEnumParam `facebook:"spatial_audio_format"`
	Status                     *enums.UserliveVideosStatusEnumParam             `facebook:"status"`
	StereoscopicMode           *enums.UserliveVideosStereoscopicModeEnumParam   `facebook:"stereoscopic_mode"`
	StopOnDeleteStream         *bool                                            `facebook:"stop_on_delete_stream"`
	StreamType                 *enums.UserliveVideosStreamTypeEnumParam         `facebook:"stream_type"`
	Title                      *string                                          `facebook:"title"`
	Extra                      core.Params                                      `facebook:"-"`
}

func (params CreateUserLiveVideosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ContentTags != nil {
		out["content_tags"] = *params.ContentTags
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.EnableBackupIngest != nil {
		out["enable_backup_ingest"] = *params.EnableBackupIngest
	}
	if params.EncodingSettings != nil {
		out["encoding_settings"] = *params.EncodingSettings
	}
	if params.EventParams != nil {
		out["event_params"] = *params.EventParams
	}
	if params.FisheyeVideoCropped != nil {
		out["fisheye_video_cropped"] = *params.FisheyeVideoCropped
	}
	if params.FrontZRotation != nil {
		out["front_z_rotation"] = *params.FrontZRotation
	}
	if params.IsAudioOnly != nil {
		out["is_audio_only"] = *params.IsAudioOnly
	}
	if params.IsSpherical != nil {
		out["is_spherical"] = *params.IsSpherical
	}
	if params.OriginalFov != nil {
		out["original_fov"] = *params.OriginalFov
	}
	if params.Privacy != nil {
		out["privacy"] = *params.Privacy
	}
	if params.Projection != nil {
		out["projection"] = *params.Projection
	}
	if params.Published != nil {
		out["published"] = *params.Published
	}
	if params.ScheduleCustomProfileImage != nil {
		out["schedule_custom_profile_image"] = *params.ScheduleCustomProfileImage
	}
	if params.SpatialAudioFormat != nil {
		out["spatial_audio_format"] = *params.SpatialAudioFormat
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	if params.StereoscopicMode != nil {
		out["stereoscopic_mode"] = *params.StereoscopicMode
	}
	if params.StopOnDeleteStream != nil {
		out["stop_on_delete_stream"] = *params.StopOnDeleteStream
	}
	if params.StreamType != nil {
		out["stream_type"] = *params.StreamType
	}
	if params.Title != nil {
		out["title"] = *params.Title
	}
	return out
}

func CreateUserLiveVideosBatchCall(id string, params CreateUserLiveVideosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "live_videos"), params.ToParams(), options...)
}

func NewCreateUserLiveVideosBatchRequest(id string, params CreateUserLiveVideosParams, options ...core.BatchOption) *core.BatchRequest[objects.LiveVideo] {
	return core.NewBatchRequest[objects.LiveVideo](CreateUserLiveVideosBatchCall(id, params, options...))
}

func DecodeCreateUserLiveVideosBatchResponse(response *core.BatchResponse) (*objects.LiveVideo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.LiveVideo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateUserLiveVideosWithResponse(ctx context.Context, client *core.Client, id string, params CreateUserLiveVideosParams) (*objects.LiveVideo, *core.Response, error) {
	var out objects.LiveVideo
	call := CreateUserLiveVideosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateUserLiveVideos(ctx context.Context, client *core.Client, id string, params CreateUserLiveVideosParams) (*objects.LiveVideo, error) {
	out, _, err := CreateUserLiveVideosWithResponse(ctx, client, id, params)
	return out, err
}

type CreateUserMessengerDesktopPerformanceTracesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params CreateUserMessengerDesktopPerformanceTracesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func CreateUserMessengerDesktopPerformanceTracesBatchCall(id string, params CreateUserMessengerDesktopPerformanceTracesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "messenger_desktop_performance_traces"), params.ToParams(), options...)
}

func NewCreateUserMessengerDesktopPerformanceTracesBatchRequest(id string, params CreateUserMessengerDesktopPerformanceTracesParams, options ...core.BatchOption) *core.BatchRequest[objects.User] {
	return core.NewBatchRequest[objects.User](CreateUserMessengerDesktopPerformanceTracesBatchCall(id, params, options...))
}

func DecodeCreateUserMessengerDesktopPerformanceTracesBatchResponse(response *core.BatchResponse) (*objects.User, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.User
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateUserMessengerDesktopPerformanceTracesWithResponse(ctx context.Context, client *core.Client, id string, params CreateUserMessengerDesktopPerformanceTracesParams) (*objects.User, *core.Response, error) {
	var out objects.User
	call := CreateUserMessengerDesktopPerformanceTracesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateUserMessengerDesktopPerformanceTraces(ctx context.Context, client *core.Client, id string, params CreateUserMessengerDesktopPerformanceTracesParams) (*objects.User, error) {
	out, _, err := CreateUserMessengerDesktopPerformanceTracesWithResponse(ctx, client, id, params)
	return out, err
}

type CreateUserMessengerKidsAccountsUnreadBadgeParams struct {
	ProxiedAppID core.ID     `facebook:"proxied_app_id"`
	Extra        core.Params `facebook:"-"`
}

func (params CreateUserMessengerKidsAccountsUnreadBadgeParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["proxied_app_id"] = params.ProxiedAppID
	return out
}

func CreateUserMessengerKidsAccountsUnreadBadgeBatchCall(id string, params CreateUserMessengerKidsAccountsUnreadBadgeParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "messenger_kids_accounts_unread_badge"), params.ToParams(), options...)
}

func NewCreateUserMessengerKidsAccountsUnreadBadgeBatchRequest(id string, params CreateUserMessengerKidsAccountsUnreadBadgeParams, options ...core.BatchOption) *core.BatchRequest[objects.User] {
	return core.NewBatchRequest[objects.User](CreateUserMessengerKidsAccountsUnreadBadgeBatchCall(id, params, options...))
}

func DecodeCreateUserMessengerKidsAccountsUnreadBadgeBatchResponse(response *core.BatchResponse) (*objects.User, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.User
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateUserMessengerKidsAccountsUnreadBadgeWithResponse(ctx context.Context, client *core.Client, id string, params CreateUserMessengerKidsAccountsUnreadBadgeParams) (*objects.User, *core.Response, error) {
	var out objects.User
	call := CreateUserMessengerKidsAccountsUnreadBadgeBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateUserMessengerKidsAccountsUnreadBadge(ctx context.Context, client *core.Client, id string, params CreateUserMessengerKidsAccountsUnreadBadgeParams) (*objects.User, error) {
	out, _, err := CreateUserMessengerKidsAccountsUnreadBadgeWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserMusicParams struct {
	TargetID *core.ID    `facebook:"target_id"`
	Extra    core.Params `facebook:"-"`
}

func (params GetUserMusicParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.TargetID != nil {
		out["target_id"] = *params.TargetID
	}
	return out
}

func GetUserMusicBatchCall(id string, params GetUserMusicParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "music"), params.ToParams(), options...)
}

func NewGetUserMusicBatchRequest(id string, params GetUserMusicParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Page]] {
	return core.NewBatchRequest[core.Cursor[objects.Page]](GetUserMusicBatchCall(id, params, options...))
}

func DecodeGetUserMusicBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Page], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Page]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserMusicWithResponse(ctx context.Context, client *core.Client, id string, params GetUserMusicParams) (*core.Cursor[objects.Page], *core.Response, error) {
	var out core.Cursor[objects.Page]
	call := GetUserMusicBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserMusic(ctx context.Context, client *core.Client, id string, params GetUserMusicParams) (*core.Cursor[objects.Page], error) {
	out, _, err := GetUserMusicWithResponse(ctx, client, id, params)
	return out, err
}

type CreateUserNotificationsParams struct {
	BotMessagePayloadElements *string                                      `facebook:"bot_message_payload_elements"`
	Filtering                 *[]enums.UsernotificationsFilteringEnumParam `facebook:"filtering"`
	GameMetadata              *string                                      `facebook:"game_metadata"`
	Href                      *map[string]interface{}                      `facebook:"href"`
	Label                     *string                                      `facebook:"label"`
	Message                   *map[string]interface{}                      `facebook:"message"`
	NotifIds                  *[]core.ID                                   `facebook:"notif_ids"`
	NotificationContentID     *core.ID                                     `facebook:"notification_content_id"`
	Payload                   *string                                      `facebook:"payload"`
	Read                      *bool                                        `facebook:"read"`
	Ref                       *string                                      `facebook:"ref"`
	ScheduleInterval          *uint64                                      `facebook:"schedule_interval"`
	Seen                      *bool                                        `facebook:"seen"`
	Template                  *map[string]interface{}                      `facebook:"template"`
	Type                      *enums.UsernotificationsTypeEnumParam        `facebook:"type"`
	Extra                     core.Params                                  `facebook:"-"`
}

func (params CreateUserNotificationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BotMessagePayloadElements != nil {
		out["bot_message_payload_elements"] = *params.BotMessagePayloadElements
	}
	if params.Filtering != nil {
		out["filtering"] = *params.Filtering
	}
	if params.GameMetadata != nil {
		out["game_metadata"] = *params.GameMetadata
	}
	if params.Href != nil {
		out["href"] = *params.Href
	}
	if params.Label != nil {
		out["label"] = *params.Label
	}
	if params.Message != nil {
		out["message"] = *params.Message
	}
	if params.NotifIds != nil {
		out["notif_ids"] = *params.NotifIds
	}
	if params.NotificationContentID != nil {
		out["notification_content_id"] = *params.NotificationContentID
	}
	if params.Payload != nil {
		out["payload"] = *params.Payload
	}
	if params.Read != nil {
		out["read"] = *params.Read
	}
	if params.Ref != nil {
		out["ref"] = *params.Ref
	}
	if params.ScheduleInterval != nil {
		out["schedule_interval"] = *params.ScheduleInterval
	}
	if params.Seen != nil {
		out["seen"] = *params.Seen
	}
	if params.Template != nil {
		out["template"] = *params.Template
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	return out
}

func CreateUserNotificationsBatchCall(id string, params CreateUserNotificationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "notifications"), params.ToParams(), options...)
}

func NewCreateUserNotificationsBatchRequest(id string, params CreateUserNotificationsParams, options ...core.BatchOption) *core.BatchRequest[objects.User] {
	return core.NewBatchRequest[objects.User](CreateUserNotificationsBatchCall(id, params, options...))
}

func DecodeCreateUserNotificationsBatchResponse(response *core.BatchResponse) (*objects.User, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.User
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateUserNotificationsWithResponse(ctx context.Context, client *core.Client, id string, params CreateUserNotificationsParams) (*objects.User, *core.Response, error) {
	var out objects.User
	call := CreateUserNotificationsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateUserNotifications(ctx context.Context, client *core.Client, id string, params CreateUserNotificationsParams) (*objects.User, error) {
	out, _, err := CreateUserNotificationsWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserPaymentTransactionsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserPaymentTransactionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserPaymentTransactionsBatchCall(id string, params GetUserPaymentTransactionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "payment_transactions"), params.ToParams(), options...)
}

func NewGetUserPaymentTransactionsBatchRequest(id string, params GetUserPaymentTransactionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.PaymentEnginePayment]] {
	return core.NewBatchRequest[core.Cursor[objects.PaymentEnginePayment]](GetUserPaymentTransactionsBatchCall(id, params, options...))
}

func DecodeGetUserPaymentTransactionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.PaymentEnginePayment], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.PaymentEnginePayment]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserPaymentTransactionsWithResponse(ctx context.Context, client *core.Client, id string, params GetUserPaymentTransactionsParams) (*core.Cursor[objects.PaymentEnginePayment], *core.Response, error) {
	var out core.Cursor[objects.PaymentEnginePayment]
	call := GetUserPaymentTransactionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserPaymentTransactions(ctx context.Context, client *core.Client, id string, params GetUserPaymentTransactionsParams) (*core.Cursor[objects.PaymentEnginePayment], error) {
	out, _, err := GetUserPaymentTransactionsWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteUserPermissionsParams struct {
	Permission *string     `facebook:"permission"`
	Extra      core.Params `facebook:"-"`
}

func (params DeleteUserPermissionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Permission != nil {
		out["permission"] = *params.Permission
	}
	return out
}

func DeleteUserPermissionsBatchCall(id string, params DeleteUserPermissionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "permissions"), params.ToParams(), options...)
}

func NewDeleteUserPermissionsBatchRequest(id string, params DeleteUserPermissionsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteUserPermissionsBatchCall(id, params, options...))
}

func DecodeDeleteUserPermissionsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteUserPermissionsWithResponse(ctx context.Context, client *core.Client, id string, params DeleteUserPermissionsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteUserPermissionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteUserPermissions(ctx context.Context, client *core.Client, id string, params DeleteUserPermissionsParams) (*map[string]interface{}, error) {
	out, _, err := DeleteUserPermissionsWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserPermissionsParams struct {
	Permission *string                               `facebook:"permission"`
	Status     *enums.UserpermissionsStatusEnumParam `facebook:"status"`
	Extra      core.Params                           `facebook:"-"`
}

func (params GetUserPermissionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Permission != nil {
		out["permission"] = *params.Permission
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	return out
}

func GetUserPermissionsBatchCall(id string, params GetUserPermissionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "permissions"), params.ToParams(), options...)
}

func NewGetUserPermissionsBatchRequest(id string, params GetUserPermissionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Permission]] {
	return core.NewBatchRequest[core.Cursor[objects.Permission]](GetUserPermissionsBatchCall(id, params, options...))
}

func DecodeGetUserPermissionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Permission], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Permission]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserPermissionsWithResponse(ctx context.Context, client *core.Client, id string, params GetUserPermissionsParams) (*core.Cursor[objects.Permission], *core.Response, error) {
	var out core.Cursor[objects.Permission]
	call := GetUserPermissionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserPermissions(ctx context.Context, client *core.Client, id string, params GetUserPermissionsParams) (*core.Cursor[objects.Permission], error) {
	out, _, err := GetUserPermissionsWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserPersonalAdAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserPersonalAdAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserPersonalAdAccountsBatchCall(id string, params GetUserPersonalAdAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "personal_ad_accounts"), params.ToParams(), options...)
}

func NewGetUserPersonalAdAccountsBatchRequest(id string, params GetUserPersonalAdAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccount]](GetUserPersonalAdAccountsBatchCall(id, params, options...))
}

func DecodeGetUserPersonalAdAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccount], error) {
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

func GetUserPersonalAdAccountsWithResponse(ctx context.Context, client *core.Client, id string, params GetUserPersonalAdAccountsParams) (*core.Cursor[objects.AdAccount], *core.Response, error) {
	var out core.Cursor[objects.AdAccount]
	call := GetUserPersonalAdAccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserPersonalAdAccounts(ctx context.Context, client *core.Client, id string, params GetUserPersonalAdAccountsParams) (*core.Cursor[objects.AdAccount], error) {
	out, _, err := GetUserPersonalAdAccountsWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserPhotosParams struct {
	Type  *enums.UserphotosTypeEnumParam `facebook:"type"`
	Extra core.Params                    `facebook:"-"`
}

func (params GetUserPhotosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	return out
}

func GetUserPhotosBatchCall(id string, params GetUserPhotosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "photos"), params.ToParams(), options...)
}

func NewGetUserPhotosBatchRequest(id string, params GetUserPhotosParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Photo]] {
	return core.NewBatchRequest[core.Cursor[objects.Photo]](GetUserPhotosBatchCall(id, params, options...))
}

func DecodeGetUserPhotosBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Photo], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Photo]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserPhotosWithResponse(ctx context.Context, client *core.Client, id string, params GetUserPhotosParams) (*core.Cursor[objects.Photo], *core.Response, error) {
	var out core.Cursor[objects.Photo]
	call := GetUserPhotosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserPhotos(ctx context.Context, client *core.Client, id string, params GetUserPhotosParams) (*core.Cursor[objects.Photo], error) {
	out, _, err := GetUserPhotosWithResponse(ctx, client, id, params)
	return out, err
}

type CreateUserPhotosParams struct {
	Aid                                   *string                                            `facebook:"aid"`
	AllowSphericalPhoto                   *bool                                              `facebook:"allow_spherical_photo"`
	AltTextCustom                         *string                                            `facebook:"alt_text_custom"`
	AndroidKeyHash                        *string                                            `facebook:"android_key_hash"`
	ApplicationID                         *core.ID                                           `facebook:"application_id"`
	Attempt                               *uint64                                            `facebook:"attempt"`
	AudienceExp                           *bool                                              `facebook:"audience_exp"`
	BackdatedTime                         *core.Time                                         `facebook:"backdated_time"`
	BackdatedTimeGranularity              *enums.UserphotosBackdatedTimeGranularityEnumParam `facebook:"backdated_time_granularity"`
	Caption                               *string                                            `facebook:"caption"`
	ComposerSessionID                     *core.ID                                           `facebook:"composer_session_id"`
	DirectShareStatus                     *uint64                                            `facebook:"direct_share_status"`
	FeedTargeting                         *map[string]interface{}                            `facebook:"feed_targeting"`
	FilterType                            *uint64                                            `facebook:"filter_type"`
	FullResIsComingLater                  *bool                                              `facebook:"full_res_is_coming_later"`
	InitialViewHeadingOverrideDegrees     *uint64                                            `facebook:"initial_view_heading_override_degrees"`
	InitialViewPitchOverrideDegrees       *uint64                                            `facebook:"initial_view_pitch_override_degrees"`
	InitialViewVerticalFovOverrideDegrees *uint64                                            `facebook:"initial_view_vertical_fov_override_degrees"`
	IosBundleID                           *core.ID                                           `facebook:"ios_bundle_id"`
	IsExplicitLocation                    *bool                                              `facebook:"is_explicit_location"`
	IsExplicitPlace                       *bool                                              `facebook:"is_explicit_place"`
	ManualPrivacy                         *bool                                              `facebook:"manual_privacy"`
	Message                               *string                                            `facebook:"message"`
	Name                                  *string                                            `facebook:"name"`
	NoStory                               *bool                                              `facebook:"no_story"`
	OfflineID                             *core.ID                                           `facebook:"offline_id"`
	OgActionTypeID                        *core.ID                                           `facebook:"og_action_type_id"`
	OgIconID                              *core.ID                                           `facebook:"og_icon_id"`
	OgObjectID                            *core.ID                                           `facebook:"og_object_id"`
	OgPhrase                              *string                                            `facebook:"og_phrase"`
	OgSetProfileBadge                     *bool                                              `facebook:"og_set_profile_badge"`
	OgSuggestionMechanism                 *string                                            `facebook:"og_suggestion_mechanism"`
	Place                                 *map[string]interface{}                            `facebook:"place"`
	Privacy                               *string                                            `facebook:"privacy"`
	ProfileID                             *core.ID                                           `facebook:"profile_id"`
	ProvenanceInfo                        *map[string]interface{}                            `facebook:"provenance_info"`
	ProxiedAppID                          *core.ID                                           `facebook:"proxied_app_id"`
	Published                             *bool                                              `facebook:"published"`
	Qn                                    *string                                            `facebook:"qn"`
	ScheduledPublishTime                  *uint64                                            `facebook:"scheduled_publish_time"`
	SphericalMetadata                     *map[string]interface{}                            `facebook:"spherical_metadata"`
	SponsorID                             *core.ID                                           `facebook:"sponsor_id"`
	SponsorRelationship                   *uint64                                            `facebook:"sponsor_relationship"`
	Tags                                  *[]map[string]interface{}                          `facebook:"tags"`
	TargetID                              *core.ID                                           `facebook:"target_id"`
	Targeting                             *map[string]interface{}                            `facebook:"targeting"`
	TimeSinceOriginalPost                 *uint64                                            `facebook:"time_since_original_post"`
	UID                                   *core.ID                                           `facebook:"uid"`
	UnpublishedContentType                *enums.UserphotosUnpublishedContentTypeEnumParam   `facebook:"unpublished_content_type"`
	URL                                   *string                                            `facebook:"url"`
	UserSelectedTags                      *bool                                              `facebook:"user_selected_tags"`
	VaultImageID                          *core.ID                                           `facebook:"vault_image_id"`
	Extra                                 core.Params                                        `facebook:"-"`
}

func (params CreateUserPhotosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Aid != nil {
		out["aid"] = *params.Aid
	}
	if params.AllowSphericalPhoto != nil {
		out["allow_spherical_photo"] = *params.AllowSphericalPhoto
	}
	if params.AltTextCustom != nil {
		out["alt_text_custom"] = *params.AltTextCustom
	}
	if params.AndroidKeyHash != nil {
		out["android_key_hash"] = *params.AndroidKeyHash
	}
	if params.ApplicationID != nil {
		out["application_id"] = *params.ApplicationID
	}
	if params.Attempt != nil {
		out["attempt"] = *params.Attempt
	}
	if params.AudienceExp != nil {
		out["audience_exp"] = *params.AudienceExp
	}
	if params.BackdatedTime != nil {
		out["backdated_time"] = *params.BackdatedTime
	}
	if params.BackdatedTimeGranularity != nil {
		out["backdated_time_granularity"] = *params.BackdatedTimeGranularity
	}
	if params.Caption != nil {
		out["caption"] = *params.Caption
	}
	if params.ComposerSessionID != nil {
		out["composer_session_id"] = *params.ComposerSessionID
	}
	if params.DirectShareStatus != nil {
		out["direct_share_status"] = *params.DirectShareStatus
	}
	if params.FeedTargeting != nil {
		out["feed_targeting"] = *params.FeedTargeting
	}
	if params.FilterType != nil {
		out["filter_type"] = *params.FilterType
	}
	if params.FullResIsComingLater != nil {
		out["full_res_is_coming_later"] = *params.FullResIsComingLater
	}
	if params.InitialViewHeadingOverrideDegrees != nil {
		out["initial_view_heading_override_degrees"] = *params.InitialViewHeadingOverrideDegrees
	}
	if params.InitialViewPitchOverrideDegrees != nil {
		out["initial_view_pitch_override_degrees"] = *params.InitialViewPitchOverrideDegrees
	}
	if params.InitialViewVerticalFovOverrideDegrees != nil {
		out["initial_view_vertical_fov_override_degrees"] = *params.InitialViewVerticalFovOverrideDegrees
	}
	if params.IosBundleID != nil {
		out["ios_bundle_id"] = *params.IosBundleID
	}
	if params.IsExplicitLocation != nil {
		out["is_explicit_location"] = *params.IsExplicitLocation
	}
	if params.IsExplicitPlace != nil {
		out["is_explicit_place"] = *params.IsExplicitPlace
	}
	if params.ManualPrivacy != nil {
		out["manual_privacy"] = *params.ManualPrivacy
	}
	if params.Message != nil {
		out["message"] = *params.Message
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.NoStory != nil {
		out["no_story"] = *params.NoStory
	}
	if params.OfflineID != nil {
		out["offline_id"] = *params.OfflineID
	}
	if params.OgActionTypeID != nil {
		out["og_action_type_id"] = *params.OgActionTypeID
	}
	if params.OgIconID != nil {
		out["og_icon_id"] = *params.OgIconID
	}
	if params.OgObjectID != nil {
		out["og_object_id"] = *params.OgObjectID
	}
	if params.OgPhrase != nil {
		out["og_phrase"] = *params.OgPhrase
	}
	if params.OgSetProfileBadge != nil {
		out["og_set_profile_badge"] = *params.OgSetProfileBadge
	}
	if params.OgSuggestionMechanism != nil {
		out["og_suggestion_mechanism"] = *params.OgSuggestionMechanism
	}
	if params.Place != nil {
		out["place"] = *params.Place
	}
	if params.Privacy != nil {
		out["privacy"] = *params.Privacy
	}
	if params.ProfileID != nil {
		out["profile_id"] = *params.ProfileID
	}
	if params.ProvenanceInfo != nil {
		out["provenance_info"] = *params.ProvenanceInfo
	}
	if params.ProxiedAppID != nil {
		out["proxied_app_id"] = *params.ProxiedAppID
	}
	if params.Published != nil {
		out["published"] = *params.Published
	}
	if params.Qn != nil {
		out["qn"] = *params.Qn
	}
	if params.ScheduledPublishTime != nil {
		out["scheduled_publish_time"] = *params.ScheduledPublishTime
	}
	if params.SphericalMetadata != nil {
		out["spherical_metadata"] = *params.SphericalMetadata
	}
	if params.SponsorID != nil {
		out["sponsor_id"] = *params.SponsorID
	}
	if params.SponsorRelationship != nil {
		out["sponsor_relationship"] = *params.SponsorRelationship
	}
	if params.Tags != nil {
		out["tags"] = *params.Tags
	}
	if params.TargetID != nil {
		out["target_id"] = *params.TargetID
	}
	if params.Targeting != nil {
		out["targeting"] = *params.Targeting
	}
	if params.TimeSinceOriginalPost != nil {
		out["time_since_original_post"] = *params.TimeSinceOriginalPost
	}
	if params.UID != nil {
		out["uid"] = *params.UID
	}
	if params.UnpublishedContentType != nil {
		out["unpublished_content_type"] = *params.UnpublishedContentType
	}
	if params.URL != nil {
		out["url"] = *params.URL
	}
	if params.UserSelectedTags != nil {
		out["user_selected_tags"] = *params.UserSelectedTags
	}
	if params.VaultImageID != nil {
		out["vault_image_id"] = *params.VaultImageID
	}
	return out
}

func CreateUserPhotosBatchCall(id string, params CreateUserPhotosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "photos"), params.ToParams(), options...)
}

func NewCreateUserPhotosBatchRequest(id string, params CreateUserPhotosParams, options ...core.BatchOption) *core.BatchRequest[objects.Photo] {
	return core.NewBatchRequest[objects.Photo](CreateUserPhotosBatchCall(id, params, options...))
}

func DecodeCreateUserPhotosBatchResponse(response *core.BatchResponse) (*objects.Photo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Photo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateUserPhotosWithResponse(ctx context.Context, client *core.Client, id string, params CreateUserPhotosParams) (*objects.Photo, *core.Response, error) {
	var out objects.Photo
	call := CreateUserPhotosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateUserPhotos(ctx context.Context, client *core.Client, id string, params CreateUserPhotosParams) (*objects.Photo, error) {
	out, _, err := CreateUserPhotosWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserPictureParams struct {
	Height   *int                            `facebook:"height"`
	Redirect *bool                           `facebook:"redirect"`
	Type     *enums.UserpictureTypeEnumParam `facebook:"type"`
	Width    *int                            `facebook:"width"`
	Extra    core.Params                     `facebook:"-"`
}

func (params GetUserPictureParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Height != nil {
		out["height"] = *params.Height
	}
	if params.Redirect != nil {
		out["redirect"] = *params.Redirect
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	if params.Width != nil {
		out["width"] = *params.Width
	}
	return out
}

func GetUserPictureBatchCall(id string, params GetUserPictureParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "picture"), params.ToParams(), options...)
}

func NewGetUserPictureBatchRequest(id string, params GetUserPictureParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProfilePictureSource]] {
	return core.NewBatchRequest[core.Cursor[objects.ProfilePictureSource]](GetUserPictureBatchCall(id, params, options...))
}

func DecodeGetUserPictureBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProfilePictureSource], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ProfilePictureSource]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserPictureWithResponse(ctx context.Context, client *core.Client, id string, params GetUserPictureParams) (*core.Cursor[objects.ProfilePictureSource], *core.Response, error) {
	var out core.Cursor[objects.ProfilePictureSource]
	call := GetUserPictureBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserPicture(ctx context.Context, client *core.Client, id string, params GetUserPictureParams) (*core.Cursor[objects.ProfilePictureSource], error) {
	out, _, err := GetUserPictureWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserPostsParams struct {
	IncludeHidden *bool       `facebook:"include_hidden"`
	Q             *string     `facebook:"q"`
	ShowExpired   *bool       `facebook:"show_expired"`
	Since         *core.Time  `facebook:"since"`
	Until         *core.Time  `facebook:"until"`
	With          *string     `facebook:"with"`
	Extra         core.Params `facebook:"-"`
}

func (params GetUserPostsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.IncludeHidden != nil {
		out["include_hidden"] = *params.IncludeHidden
	}
	if params.Q != nil {
		out["q"] = *params.Q
	}
	if params.ShowExpired != nil {
		out["show_expired"] = *params.ShowExpired
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	if params.With != nil {
		out["with"] = *params.With
	}
	return out
}

func GetUserPostsBatchCall(id string, params GetUserPostsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "posts"), params.ToParams(), options...)
}

func NewGetUserPostsBatchRequest(id string, params GetUserPostsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Post]] {
	return core.NewBatchRequest[core.Cursor[objects.Post]](GetUserPostsBatchCall(id, params, options...))
}

func DecodeGetUserPostsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Post], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Post]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserPostsWithResponse(ctx context.Context, client *core.Client, id string, params GetUserPostsParams) (*core.Cursor[objects.Post], *core.Response, error) {
	var out core.Cursor[objects.Post]
	call := GetUserPostsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserPosts(ctx context.Context, client *core.Client, id string, params GetUserPostsParams) (*core.Cursor[objects.Post], error) {
	out, _, err := GetUserPostsWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserRichMediaDocumentsParams struct {
	Query *string     `facebook:"query"`
	Extra core.Params `facebook:"-"`
}

func (params GetUserRichMediaDocumentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Query != nil {
		out["query"] = *params.Query
	}
	return out
}

func GetUserRichMediaDocumentsBatchCall(id string, params GetUserRichMediaDocumentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "rich_media_documents"), params.ToParams(), options...)
}

func NewGetUserRichMediaDocumentsBatchRequest(id string, params GetUserRichMediaDocumentsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Canvas]] {
	return core.NewBatchRequest[core.Cursor[objects.Canvas]](GetUserRichMediaDocumentsBatchCall(id, params, options...))
}

func DecodeGetUserRichMediaDocumentsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Canvas], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Canvas]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserRichMediaDocumentsWithResponse(ctx context.Context, client *core.Client, id string, params GetUserRichMediaDocumentsParams) (*core.Cursor[objects.Canvas], *core.Response, error) {
	var out core.Cursor[objects.Canvas]
	call := GetUserRichMediaDocumentsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserRichMediaDocuments(ctx context.Context, client *core.Client, id string, params GetUserRichMediaDocumentsParams) (*core.Cursor[objects.Canvas], error) {
	out, _, err := GetUserRichMediaDocumentsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateUserStagingResourcesParams struct {
	File  *core.FileParam `facebook:"file"`
	Extra core.Params     `facebook:"-"`
}

func (params CreateUserStagingResourcesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.File != nil {
		out["file"] = *params.File
	}
	return out
}

func CreateUserStagingResourcesBatchCall(id string, params CreateUserStagingResourcesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "staging_resources"), params.ToParams(), options...)
}

func NewCreateUserStagingResourcesBatchRequest(id string, params CreateUserStagingResourcesParams, options ...core.BatchOption) *core.BatchRequest[objects.User] {
	return core.NewBatchRequest[objects.User](CreateUserStagingResourcesBatchCall(id, params, options...))
}

func DecodeCreateUserStagingResourcesBatchResponse(response *core.BatchResponse) (*objects.User, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.User
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateUserStagingResourcesWithResponse(ctx context.Context, client *core.Client, id string, params CreateUserStagingResourcesParams) (*objects.User, *core.Response, error) {
	var out objects.User
	call := CreateUserStagingResourcesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateUserStagingResources(ctx context.Context, client *core.Client, id string, params CreateUserStagingResourcesParams) (*objects.User, error) {
	out, _, err := CreateUserStagingResourcesWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserVideosParams struct {
	Type  *enums.UservideosTypeEnumParam `facebook:"type"`
	Extra core.Params                    `facebook:"-"`
}

func (params GetUserVideosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	return out
}

func GetUserVideosBatchCall(id string, params GetUserVideosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "videos"), params.ToParams(), options...)
}

func NewGetUserVideosBatchRequest(id string, params GetUserVideosParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdVideo]] {
	return core.NewBatchRequest[core.Cursor[objects.AdVideo]](GetUserVideosBatchCall(id, params, options...))
}

func DecodeGetUserVideosBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdVideo], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdVideo]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserVideosWithResponse(ctx context.Context, client *core.Client, id string, params GetUserVideosParams) (*core.Cursor[objects.AdVideo], *core.Response, error) {
	var out core.Cursor[objects.AdVideo]
	call := GetUserVideosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserVideos(ctx context.Context, client *core.Client, id string, params GetUserVideosParams) (*core.Cursor[objects.AdVideo], error) {
	out, _, err := GetUserVideosWithResponse(ctx, client, id, params)
	return out, err
}

type CreateUserVideosParams struct {
	ApplicationID                 *core.ID                                         `facebook:"application_id"`
	AskedFunFactPromptID          *core.ID                                         `facebook:"asked_fun_fact_prompt_id"`
	AudioStoryWaveAnimationHandle *string                                          `facebook:"audio_story_wave_animation_handle"`
	ComposerEntryPicker           *string                                          `facebook:"composer_entry_picker"`
	ComposerEntryPoint            *string                                          `facebook:"composer_entry_point"`
	ComposerEntryTime             *uint64                                          `facebook:"composer_entry_time"`
	ComposerSessionEventsLog      *string                                          `facebook:"composer_session_events_log"`
	ComposerSessionID             *core.ID                                         `facebook:"composer_session_id"`
	ComposerSourceSurface         *string                                          `facebook:"composer_source_surface"`
	ComposerType                  *string                                          `facebook:"composer_type"`
	ContainerType                 *enums.UservideosContainerTypeEnumParam          `facebook:"container_type"`
	ContentCategory               *enums.UservideosContentCategoryEnumParam        `facebook:"content_category"`
	CreativeTools                 *string                                          `facebook:"creative_tools"`
	Description                   *string                                          `facebook:"description"`
	DirectShareStatus             *uint64                                          `facebook:"direct_share_status"`
	EditDescriptionSpec           *map[string]interface{}                          `facebook:"edit_description_spec"`
	Embeddable                    *bool                                            `facebook:"embeddable"`
	EndOffset                     *uint64                                          `facebook:"end_offset"`
	FbuploaderVideoFileChunk      *string                                          `facebook:"fbuploader_video_file_chunk"`
	FileSize                      *uint64                                          `facebook:"file_size"`
	FileURL                       *string                                          `facebook:"file_url"`
	FisheyeVideoCropped           *bool                                            `facebook:"fisheye_video_cropped"`
	Formatting                    *enums.UservideosFormattingEnumParam             `facebook:"formatting"`
	Fov                           *uint64                                          `facebook:"fov"`
	FrontZRotation                *float64                                         `facebook:"front_z_rotation"`
	FunFactPromptID               *core.ID                                         `facebook:"fun_fact_prompt_id"`
	FunFactToasteeID              *core.ID                                         `facebook:"fun_fact_toastee_id"`
	Guide                         *[][]uint64                                      `facebook:"guide"`
	GuideEnabled                  *bool                                            `facebook:"guide_enabled"`
	InitialHeading                *uint64                                          `facebook:"initial_heading"`
	InitialPitch                  *uint64                                          `facebook:"initial_pitch"`
	InstantGameEntryPointData     *string                                          `facebook:"instant_game_entry_point_data"`
	IsBoostIntended               *bool                                            `facebook:"is_boost_intended"`
	IsExplicitShare               *bool                                            `facebook:"is_explicit_share"`
	IsGroupLinkingPost            *bool                                            `facebook:"is_group_linking_post"`
	IsPartnershipAd               *bool                                            `facebook:"is_partnership_ad"`
	IsVoiceClip                   *bool                                            `facebook:"is_voice_clip"`
	LocationSourceID              *core.ID                                         `facebook:"location_source_id"`
	ManualPrivacy                 *bool                                            `facebook:"manual_privacy"`
	NoStory                       *bool                                            `facebook:"no_story"`
	OgActionTypeID                *core.ID                                         `facebook:"og_action_type_id"`
	OgIconID                      *core.ID                                         `facebook:"og_icon_id"`
	OgObjectID                    *core.ID                                         `facebook:"og_object_id"`
	OgPhrase                      *string                                          `facebook:"og_phrase"`
	OgSuggestionMechanism         *string                                          `facebook:"og_suggestion_mechanism"`
	OriginalFov                   *uint64                                          `facebook:"original_fov"`
	OriginalProjectionType        *enums.UservideosOriginalProjectionTypeEnumParam `facebook:"original_projection_type"`
	PartnershipAdAdCode           *string                                          `facebook:"partnership_ad_ad_code"`
	Privacy                       *string                                          `facebook:"privacy"`
	PublishEventID                *core.ID                                         `facebook:"publish_event_id"`
	ReferencedStickerID           *core.ID                                         `facebook:"referenced_sticker_id"`
	ReplaceVideoID                *core.ID                                         `facebook:"replace_video_id"`
	SelectedAudioSpec             *map[string]interface{}                          `facebook:"selected_audio_spec"`
	SlideshowSpec                 *map[string]interface{}                          `facebook:"slideshow_spec"`
	Source                        *string                                          `facebook:"source"`
	SourceInstagramMediaID        *core.ID                                         `facebook:"source_instagram_media_id"`
	Spherical                     *bool                                            `facebook:"spherical"`
	SponsorID                     *core.ID                                         `facebook:"sponsor_id"`
	StartOffset                   *uint64                                          `facebook:"start_offset"`
	SwapMode                      *enums.UservideosSwapModeEnumParam               `facebook:"swap_mode"`
	TextFormatMetadata            *string                                          `facebook:"text_format_metadata"`
	Thumb                         *core.FileParam                                  `facebook:"thumb"`
	TimeSinceOriginalPost         *uint64                                          `facebook:"time_since_original_post"`
	Title                         *string                                          `facebook:"title"`
	TranscodeSettingProperties    *string                                          `facebook:"transcode_setting_properties"`
	UnpublishedContentType        *enums.UservideosUnpublishedContentTypeEnumParam `facebook:"unpublished_content_type"`
	UploadPhase                   *enums.UservideosUploadPhaseEnumParam            `facebook:"upload_phase"`
	UploadSessionID               *core.ID                                         `facebook:"upload_session_id"`
	UploadSettingProperties       *string                                          `facebook:"upload_setting_properties"`
	VideoFileChunk                *string                                          `facebook:"video_file_chunk"`
	VideoIDOriginal               *string                                          `facebook:"video_id_original"`
	VideoStartTimeMs              *uint64                                          `facebook:"video_start_time_ms"`
	WaterfallID                   *core.ID                                         `facebook:"waterfall_id"`
	Extra                         core.Params                                      `facebook:"-"`
}

func (params CreateUserVideosParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ApplicationID != nil {
		out["application_id"] = *params.ApplicationID
	}
	if params.AskedFunFactPromptID != nil {
		out["asked_fun_fact_prompt_id"] = *params.AskedFunFactPromptID
	}
	if params.AudioStoryWaveAnimationHandle != nil {
		out["audio_story_wave_animation_handle"] = *params.AudioStoryWaveAnimationHandle
	}
	if params.ComposerEntryPicker != nil {
		out["composer_entry_picker"] = *params.ComposerEntryPicker
	}
	if params.ComposerEntryPoint != nil {
		out["composer_entry_point"] = *params.ComposerEntryPoint
	}
	if params.ComposerEntryTime != nil {
		out["composer_entry_time"] = *params.ComposerEntryTime
	}
	if params.ComposerSessionEventsLog != nil {
		out["composer_session_events_log"] = *params.ComposerSessionEventsLog
	}
	if params.ComposerSessionID != nil {
		out["composer_session_id"] = *params.ComposerSessionID
	}
	if params.ComposerSourceSurface != nil {
		out["composer_source_surface"] = *params.ComposerSourceSurface
	}
	if params.ComposerType != nil {
		out["composer_type"] = *params.ComposerType
	}
	if params.ContainerType != nil {
		out["container_type"] = *params.ContainerType
	}
	if params.ContentCategory != nil {
		out["content_category"] = *params.ContentCategory
	}
	if params.CreativeTools != nil {
		out["creative_tools"] = *params.CreativeTools
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.DirectShareStatus != nil {
		out["direct_share_status"] = *params.DirectShareStatus
	}
	if params.EditDescriptionSpec != nil {
		out["edit_description_spec"] = *params.EditDescriptionSpec
	}
	if params.Embeddable != nil {
		out["embeddable"] = *params.Embeddable
	}
	if params.EndOffset != nil {
		out["end_offset"] = *params.EndOffset
	}
	if params.FbuploaderVideoFileChunk != nil {
		out["fbuploader_video_file_chunk"] = *params.FbuploaderVideoFileChunk
	}
	if params.FileSize != nil {
		out["file_size"] = *params.FileSize
	}
	if params.FileURL != nil {
		out["file_url"] = *params.FileURL
	}
	if params.FisheyeVideoCropped != nil {
		out["fisheye_video_cropped"] = *params.FisheyeVideoCropped
	}
	if params.Formatting != nil {
		out["formatting"] = *params.Formatting
	}
	if params.Fov != nil {
		out["fov"] = *params.Fov
	}
	if params.FrontZRotation != nil {
		out["front_z_rotation"] = *params.FrontZRotation
	}
	if params.FunFactPromptID != nil {
		out["fun_fact_prompt_id"] = *params.FunFactPromptID
	}
	if params.FunFactToasteeID != nil {
		out["fun_fact_toastee_id"] = *params.FunFactToasteeID
	}
	if params.Guide != nil {
		out["guide"] = *params.Guide
	}
	if params.GuideEnabled != nil {
		out["guide_enabled"] = *params.GuideEnabled
	}
	if params.InitialHeading != nil {
		out["initial_heading"] = *params.InitialHeading
	}
	if params.InitialPitch != nil {
		out["initial_pitch"] = *params.InitialPitch
	}
	if params.InstantGameEntryPointData != nil {
		out["instant_game_entry_point_data"] = *params.InstantGameEntryPointData
	}
	if params.IsBoostIntended != nil {
		out["is_boost_intended"] = *params.IsBoostIntended
	}
	if params.IsExplicitShare != nil {
		out["is_explicit_share"] = *params.IsExplicitShare
	}
	if params.IsGroupLinkingPost != nil {
		out["is_group_linking_post"] = *params.IsGroupLinkingPost
	}
	if params.IsPartnershipAd != nil {
		out["is_partnership_ad"] = *params.IsPartnershipAd
	}
	if params.IsVoiceClip != nil {
		out["is_voice_clip"] = *params.IsVoiceClip
	}
	if params.LocationSourceID != nil {
		out["location_source_id"] = *params.LocationSourceID
	}
	if params.ManualPrivacy != nil {
		out["manual_privacy"] = *params.ManualPrivacy
	}
	if params.NoStory != nil {
		out["no_story"] = *params.NoStory
	}
	if params.OgActionTypeID != nil {
		out["og_action_type_id"] = *params.OgActionTypeID
	}
	if params.OgIconID != nil {
		out["og_icon_id"] = *params.OgIconID
	}
	if params.OgObjectID != nil {
		out["og_object_id"] = *params.OgObjectID
	}
	if params.OgPhrase != nil {
		out["og_phrase"] = *params.OgPhrase
	}
	if params.OgSuggestionMechanism != nil {
		out["og_suggestion_mechanism"] = *params.OgSuggestionMechanism
	}
	if params.OriginalFov != nil {
		out["original_fov"] = *params.OriginalFov
	}
	if params.OriginalProjectionType != nil {
		out["original_projection_type"] = *params.OriginalProjectionType
	}
	if params.PartnershipAdAdCode != nil {
		out["partnership_ad_ad_code"] = *params.PartnershipAdAdCode
	}
	if params.Privacy != nil {
		out["privacy"] = *params.Privacy
	}
	if params.PublishEventID != nil {
		out["publish_event_id"] = *params.PublishEventID
	}
	if params.ReferencedStickerID != nil {
		out["referenced_sticker_id"] = *params.ReferencedStickerID
	}
	if params.ReplaceVideoID != nil {
		out["replace_video_id"] = *params.ReplaceVideoID
	}
	if params.SelectedAudioSpec != nil {
		out["selected_audio_spec"] = *params.SelectedAudioSpec
	}
	if params.SlideshowSpec != nil {
		out["slideshow_spec"] = *params.SlideshowSpec
	}
	if params.Source != nil {
		out["source"] = *params.Source
	}
	if params.SourceInstagramMediaID != nil {
		out["source_instagram_media_id"] = *params.SourceInstagramMediaID
	}
	if params.Spherical != nil {
		out["spherical"] = *params.Spherical
	}
	if params.SponsorID != nil {
		out["sponsor_id"] = *params.SponsorID
	}
	if params.StartOffset != nil {
		out["start_offset"] = *params.StartOffset
	}
	if params.SwapMode != nil {
		out["swap_mode"] = *params.SwapMode
	}
	if params.TextFormatMetadata != nil {
		out["text_format_metadata"] = *params.TextFormatMetadata
	}
	if params.Thumb != nil {
		out["thumb"] = *params.Thumb
	}
	if params.TimeSinceOriginalPost != nil {
		out["time_since_original_post"] = *params.TimeSinceOriginalPost
	}
	if params.Title != nil {
		out["title"] = *params.Title
	}
	if params.TranscodeSettingProperties != nil {
		out["transcode_setting_properties"] = *params.TranscodeSettingProperties
	}
	if params.UnpublishedContentType != nil {
		out["unpublished_content_type"] = *params.UnpublishedContentType
	}
	if params.UploadPhase != nil {
		out["upload_phase"] = *params.UploadPhase
	}
	if params.UploadSessionID != nil {
		out["upload_session_id"] = *params.UploadSessionID
	}
	if params.UploadSettingProperties != nil {
		out["upload_setting_properties"] = *params.UploadSettingProperties
	}
	if params.VideoFileChunk != nil {
		out["video_file_chunk"] = *params.VideoFileChunk
	}
	if params.VideoIDOriginal != nil {
		out["video_id_original"] = *params.VideoIDOriginal
	}
	if params.VideoStartTimeMs != nil {
		out["video_start_time_ms"] = *params.VideoStartTimeMs
	}
	if params.WaterfallID != nil {
		out["waterfall_id"] = *params.WaterfallID
	}
	return out
}

func CreateUserVideosBatchCall(id string, params CreateUserVideosParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "videos"), params.ToParams(), options...)
}

func NewCreateUserVideosBatchRequest(id string, params CreateUserVideosParams, options ...core.BatchOption) *core.BatchRequest[objects.AdVideo] {
	return core.NewBatchRequest[objects.AdVideo](CreateUserVideosBatchCall(id, params, options...))
}

func DecodeCreateUserVideosBatchResponse(response *core.BatchResponse) (*objects.AdVideo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdVideo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateUserVideosWithResponse(ctx context.Context, client *core.Client, id string, params CreateUserVideosParams) (*objects.AdVideo, *core.Response, error) {
	var out objects.AdVideo
	call := CreateUserVideosBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateUserVideos(ctx context.Context, client *core.Client, id string, params CreateUserVideosParams) (*objects.AdVideo, error) {
	out, _, err := CreateUserVideosWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteUserParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteUserParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteUserBatchCall(id string, params DeleteUserParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteUserBatchRequest(id string, params DeleteUserParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteUserBatchCall(id, params, options...))
}

func DecodeDeleteUserBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteUserWithResponse(ctx context.Context, client *core.Client, id string, params DeleteUserParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteUserBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteUser(ctx context.Context, client *core.Client, id string, params DeleteUserParams) (*map[string]interface{}, error) {
	out, _, err := DeleteUserWithResponse(ctx, client, id, params)
	return out, err
}

type GetUserParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserBatchCall(id string, params GetUserParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetUserBatchRequest(id string, params GetUserParams, options ...core.BatchOption) *core.BatchRequest[objects.User] {
	return core.NewBatchRequest[objects.User](GetUserBatchCall(id, params, options...))
}

func DecodeGetUserBatchResponse(response *core.BatchResponse) (*objects.User, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.User
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserWithResponse(ctx context.Context, client *core.Client, id string, params GetUserParams) (*objects.User, *core.Response, error) {
	var out objects.User
	call := GetUserBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUser(ctx context.Context, client *core.Client, id string, params GetUserParams) (*objects.User, error) {
	out, _, err := GetUserWithResponse(ctx, client, id, params)
	return out, err
}

type UpdateUserParams struct {
	EmojiColorPref                  *uint64                                    `facebook:"emoji_color_pref"`
	Firstname                       *string                                    `facebook:"firstname"`
	Lastname                        *string                                    `facebook:"lastname"`
	LocalNewsMegaphoneDismissStatus *enums.UserLocalNewsMegaphoneDismissStatus `facebook:"local_news_megaphone_dismiss_status"`
	LocalNewsSubscriptionStatus     *enums.UserLocalNewsSubscriptionStatus     `facebook:"local_news_subscription_status"`
	Name                            *string                                    `facebook:"name"`
	Password                        *string                                    `facebook:"password"`
	Extra                           core.Params                                `facebook:"-"`
}

func (params UpdateUserParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.EmojiColorPref != nil {
		out["emoji_color_pref"] = *params.EmojiColorPref
	}
	if params.Firstname != nil {
		out["firstname"] = *params.Firstname
	}
	if params.Lastname != nil {
		out["lastname"] = *params.Lastname
	}
	if params.LocalNewsMegaphoneDismissStatus != nil {
		out["local_news_megaphone_dismiss_status"] = *params.LocalNewsMegaphoneDismissStatus
	}
	if params.LocalNewsSubscriptionStatus != nil {
		out["local_news_subscription_status"] = *params.LocalNewsSubscriptionStatus
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.Password != nil {
		out["password"] = *params.Password
	}
	return out
}

func UpdateUserBatchCall(id string, params UpdateUserParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateUserBatchRequest(id string, params UpdateUserParams, options ...core.BatchOption) *core.BatchRequest[objects.User] {
	return core.NewBatchRequest[objects.User](UpdateUserBatchCall(id, params, options...))
}

func DecodeUpdateUserBatchResponse(response *core.BatchResponse) (*objects.User, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.User
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateUserWithResponse(ctx context.Context, client *core.Client, id string, params UpdateUserParams) (*objects.User, *core.Response, error) {
	var out objects.User
	call := UpdateUserBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateUser(ctx context.Context, client *core.Client, id string, params UpdateUserParams) (*objects.User, error) {
	out, _, err := UpdateUserWithResponse(ctx, client, id, params)
	return out, err
}
