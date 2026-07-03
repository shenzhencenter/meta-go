package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetIGMediaBoostAdsListParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGMediaBoostAdsListParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGMediaBoostAdsListBatchCall(id string, params GetIGMediaBoostAdsListParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "boost_ads_list"), params.ToParams(), options...)
}

func NewGetIGMediaBoostAdsListBatchRequest(id string, params GetIGMediaBoostAdsListParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGBoostMediaAd]] {
	return core.NewBatchRequest[core.Cursor[objects.IGBoostMediaAd]](GetIGMediaBoostAdsListBatchCall(id, params, options...))
}

func DecodeGetIGMediaBoostAdsListBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGBoostMediaAd], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGBoostMediaAd]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGMediaBoostAdsList(ctx context.Context, client *core.Client, id string, params GetIGMediaBoostAdsListParams) (*core.Cursor[objects.IGBoostMediaAd], error) {
	var out core.Cursor[objects.IGBoostMediaAd]
	call := GetIGMediaBoostAdsListBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetIGMediaBrandedContentPartnerPromoteParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGMediaBrandedContentPartnerPromoteParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGMediaBrandedContentPartnerPromoteBatchCall(id string, params GetIGMediaBrandedContentPartnerPromoteParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "branded_content_partner_promote"), params.ToParams(), options...)
}

func NewGetIGMediaBrandedContentPartnerPromoteBatchRequest(id string, params GetIGMediaBrandedContentPartnerPromoteParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BrandedContentShadowIGUserID]] {
	return core.NewBatchRequest[core.Cursor[objects.BrandedContentShadowIGUserID]](GetIGMediaBrandedContentPartnerPromoteBatchCall(id, params, options...))
}

func DecodeGetIGMediaBrandedContentPartnerPromoteBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BrandedContentShadowIGUserID], error) {
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

func GetIGMediaBrandedContentPartnerPromote(ctx context.Context, client *core.Client, id string, params GetIGMediaBrandedContentPartnerPromoteParams) (*core.Cursor[objects.BrandedContentShadowIGUserID], error) {
	var out core.Cursor[objects.BrandedContentShadowIGUserID]
	call := GetIGMediaBrandedContentPartnerPromoteBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateIGMediaBrandedContentPartnerPromoteParams struct {
	Permission bool        `facebook:"permission"`
	SponsorID  core.ID     `facebook:"sponsor_id"`
	Extra      core.Params `facebook:"-"`
}

func (params CreateIGMediaBrandedContentPartnerPromoteParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["permission"] = params.Permission
	out["sponsor_id"] = params.SponsorID
	return out
}

func CreateIGMediaBrandedContentPartnerPromoteBatchCall(id string, params CreateIGMediaBrandedContentPartnerPromoteParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "branded_content_partner_promote"), params.ToParams(), options...)
}

func NewCreateIGMediaBrandedContentPartnerPromoteBatchRequest(id string, params CreateIGMediaBrandedContentPartnerPromoteParams, options ...core.BatchOption) *core.BatchRequest[objects.BrandedContentShadowIGUserID] {
	return core.NewBatchRequest[objects.BrandedContentShadowIGUserID](CreateIGMediaBrandedContentPartnerPromoteBatchCall(id, params, options...))
}

func DecodeCreateIGMediaBrandedContentPartnerPromoteBatchResponse(response *core.BatchResponse) (*objects.BrandedContentShadowIGUserID, error) {
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

func CreateIGMediaBrandedContentPartnerPromote(ctx context.Context, client *core.Client, id string, params CreateIGMediaBrandedContentPartnerPromoteParams) (*objects.BrandedContentShadowIGUserID, error) {
	var out objects.BrandedContentShadowIGUserID
	call := CreateIGMediaBrandedContentPartnerPromoteBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetIGMediaChildrenParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGMediaChildrenParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGMediaChildrenBatchCall(id string, params GetIGMediaChildrenParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "children"), params.ToParams(), options...)
}

func NewGetIGMediaChildrenBatchRequest(id string, params GetIGMediaChildrenParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGMedia]] {
	return core.NewBatchRequest[core.Cursor[objects.IGMedia]](GetIGMediaChildrenBatchCall(id, params, options...))
}

func DecodeGetIGMediaChildrenBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGMedia], error) {
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

func GetIGMediaChildren(ctx context.Context, client *core.Client, id string, params GetIGMediaChildrenParams) (*core.Cursor[objects.IGMedia], error) {
	var out core.Cursor[objects.IGMedia]
	call := GetIGMediaChildrenBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetIGMediaCollaboratorsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGMediaCollaboratorsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGMediaCollaboratorsBatchCall(id string, params GetIGMediaCollaboratorsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "collaborators"), params.ToParams(), options...)
}

func NewGetIGMediaCollaboratorsBatchRequest(id string, params GetIGMediaCollaboratorsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ShadowIGMediaCollaborators]] {
	return core.NewBatchRequest[core.Cursor[objects.ShadowIGMediaCollaborators]](GetIGMediaCollaboratorsBatchCall(id, params, options...))
}

func DecodeGetIGMediaCollaboratorsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ShadowIGMediaCollaborators], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ShadowIGMediaCollaborators]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGMediaCollaborators(ctx context.Context, client *core.Client, id string, params GetIGMediaCollaboratorsParams) (*core.Cursor[objects.ShadowIGMediaCollaborators], error) {
	var out core.Cursor[objects.ShadowIGMediaCollaborators]
	call := GetIGMediaCollaboratorsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetIGMediaCommentsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGMediaCommentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGMediaCommentsBatchCall(id string, params GetIGMediaCommentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "comments"), params.ToParams(), options...)
}

func NewGetIGMediaCommentsBatchRequest(id string, params GetIGMediaCommentsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGComment]] {
	return core.NewBatchRequest[core.Cursor[objects.IGComment]](GetIGMediaCommentsBatchCall(id, params, options...))
}

func DecodeGetIGMediaCommentsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGComment], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGComment]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGMediaComments(ctx context.Context, client *core.Client, id string, params GetIGMediaCommentsParams) (*core.Cursor[objects.IGComment], error) {
	var out core.Cursor[objects.IGComment]
	call := GetIGMediaCommentsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateIGMediaCommentsParams struct {
	AdID    *core.ID    `facebook:"ad_id"`
	Message *string     `facebook:"message"`
	Extra   core.Params `facebook:"-"`
}

func (params CreateIGMediaCommentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdID != nil {
		out["ad_id"] = *params.AdID
	}
	if params.Message != nil {
		out["message"] = *params.Message
	}
	return out
}

func CreateIGMediaCommentsBatchCall(id string, params CreateIGMediaCommentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "comments"), params.ToParams(), options...)
}

func NewCreateIGMediaCommentsBatchRequest(id string, params CreateIGMediaCommentsParams, options ...core.BatchOption) *core.BatchRequest[objects.IGComment] {
	return core.NewBatchRequest[objects.IGComment](CreateIGMediaCommentsBatchCall(id, params, options...))
}

func DecodeCreateIGMediaCommentsBatchResponse(response *core.BatchResponse) (*objects.IGComment, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.IGComment
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateIGMediaComments(ctx context.Context, client *core.Client, id string, params CreateIGMediaCommentsParams) (*objects.IGComment, error) {
	var out objects.IGComment
	call := CreateIGMediaCommentsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetIGMediaInsightsParams struct {
	Breakdown *[]enums.ShadowigmediainsightsBreakdownEnumParam `facebook:"breakdown"`
	Metric    []enums.ShadowigmediainsightsMetricEnumParam     `facebook:"metric"`
	Period    *[]enums.ShadowigmediainsightsPeriodEnumParam    `facebook:"period"`
	Extra     core.Params                                      `facebook:"-"`
}

func (params GetIGMediaInsightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Breakdown != nil {
		out["breakdown"] = *params.Breakdown
	}
	out["metric"] = params.Metric
	if params.Period != nil {
		out["period"] = *params.Period
	}
	return out
}

func GetIGMediaInsightsBatchCall(id string, params GetIGMediaInsightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "insights"), params.ToParams(), options...)
}

func NewGetIGMediaInsightsBatchRequest(id string, params GetIGMediaInsightsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.InstagramInsightsResult]] {
	return core.NewBatchRequest[core.Cursor[objects.InstagramInsightsResult]](GetIGMediaInsightsBatchCall(id, params, options...))
}

func DecodeGetIGMediaInsightsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.InstagramInsightsResult], error) {
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

func GetIGMediaInsights(ctx context.Context, client *core.Client, id string, params GetIGMediaInsightsParams) (*core.Cursor[objects.InstagramInsightsResult], error) {
	var out core.Cursor[objects.InstagramInsightsResult]
	call := GetIGMediaInsightsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeleteIGMediaPartnershipAdCodeParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteIGMediaPartnershipAdCodeParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteIGMediaPartnershipAdCodeBatchCall(id string, params DeleteIGMediaPartnershipAdCodeParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "partnership_ad_code"), params.ToParams(), options...)
}

func NewDeleteIGMediaPartnershipAdCodeBatchRequest(id string, params DeleteIGMediaPartnershipAdCodeParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteIGMediaPartnershipAdCodeBatchCall(id, params, options...))
}

func DecodeDeleteIGMediaPartnershipAdCodeBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteIGMediaPartnershipAdCode(ctx context.Context, client *core.Client, id string, params DeleteIGMediaPartnershipAdCodeParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteIGMediaPartnershipAdCodeBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateIGMediaPartnershipAdCodeParams struct {
	Extra core.Params `facebook:"-"`
}

func (params CreateIGMediaPartnershipAdCodeParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func CreateIGMediaPartnershipAdCodeBatchCall(id string, params CreateIGMediaPartnershipAdCodeParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "partnership_ad_code"), params.ToParams(), options...)
}

func NewCreateIGMediaPartnershipAdCodeBatchRequest(id string, params CreateIGMediaPartnershipAdCodeParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateIGMediaPartnershipAdCodeBatchCall(id, params, options...))
}

func DecodeCreateIGMediaPartnershipAdCodeBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreateIGMediaPartnershipAdCode(ctx context.Context, client *core.Client, id string, params CreateIGMediaPartnershipAdCodeParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := CreateIGMediaPartnershipAdCodeBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetIGMediaProductTagsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGMediaProductTagsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGMediaProductTagsBatchCall(id string, params GetIGMediaProductTagsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "product_tags"), params.ToParams(), options...)
}

func NewGetIGMediaProductTagsBatchRequest(id string, params GetIGMediaProductTagsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ShadowIGMediaProductTags]] {
	return core.NewBatchRequest[core.Cursor[objects.ShadowIGMediaProductTags]](GetIGMediaProductTagsBatchCall(id, params, options...))
}

func DecodeGetIGMediaProductTagsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ShadowIGMediaProductTags], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ShadowIGMediaProductTags]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGMediaProductTags(ctx context.Context, client *core.Client, id string, params GetIGMediaProductTagsParams) (*core.Cursor[objects.ShadowIGMediaProductTags], error) {
	var out core.Cursor[objects.ShadowIGMediaProductTags]
	call := GetIGMediaProductTagsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreateIGMediaProductTagsParams struct {
	ChildIndex  *uint64                  `facebook:"child_index"`
	UpdatedTags []map[string]interface{} `facebook:"updated_tags"`
	Extra       core.Params              `facebook:"-"`
}

func (params CreateIGMediaProductTagsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ChildIndex != nil {
		out["child_index"] = *params.ChildIndex
	}
	out["updated_tags"] = params.UpdatedTags
	return out
}

func CreateIGMediaProductTagsBatchCall(id string, params CreateIGMediaProductTagsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "product_tags"), params.ToParams(), options...)
}

func NewCreateIGMediaProductTagsBatchRequest(id string, params CreateIGMediaProductTagsParams, options ...core.BatchOption) *core.BatchRequest[objects.ShadowIGMediaProductTags] {
	return core.NewBatchRequest[objects.ShadowIGMediaProductTags](CreateIGMediaProductTagsBatchCall(id, params, options...))
}

func DecodeCreateIGMediaProductTagsBatchResponse(response *core.BatchResponse) (*objects.ShadowIGMediaProductTags, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ShadowIGMediaProductTags
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateIGMediaProductTags(ctx context.Context, client *core.Client, id string, params CreateIGMediaProductTagsParams) (*objects.ShadowIGMediaProductTags, error) {
	var out objects.ShadowIGMediaProductTags
	call := CreateIGMediaProductTagsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeleteIGMediaParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteIGMediaParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteIGMediaBatchCall(id string, params DeleteIGMediaParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteIGMediaBatchRequest(id string, params DeleteIGMediaParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteIGMediaBatchCall(id, params, options...))
}

func DecodeDeleteIGMediaBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteIGMedia(ctx context.Context, client *core.Client, id string, params DeleteIGMediaParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteIGMediaBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetIGMediaParams struct {
	AdAccountID            *core.ID                                   `facebook:"ad_account_id"`
	BoostableMediaCallsite *enums.ShadowigmediaBoostableMediaCallsite `facebook:"boostable_media_callsite"`
	BusinessID             *core.ID                                   `facebook:"business_id"`
	PrimaryFbPageID        *core.ID                                   `facebook:"primary_fb_page_id"`
	PrimaryIgUserID        *core.ID                                   `facebook:"primary_ig_user_id"`
	SecondaryFbPageID      *core.ID                                   `facebook:"secondary_fb_page_id"`
	SecondaryIgUserID      *core.ID                                   `facebook:"secondary_ig_user_id"`
	Extra                  core.Params                                `facebook:"-"`
}

func (params GetIGMediaParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdAccountID != nil {
		out["ad_account_id"] = *params.AdAccountID
	}
	if params.BoostableMediaCallsite != nil {
		out["boostable_media_callsite"] = *params.BoostableMediaCallsite
	}
	if params.BusinessID != nil {
		out["business_id"] = *params.BusinessID
	}
	if params.PrimaryFbPageID != nil {
		out["primary_fb_page_id"] = *params.PrimaryFbPageID
	}
	if params.PrimaryIgUserID != nil {
		out["primary_ig_user_id"] = *params.PrimaryIgUserID
	}
	if params.SecondaryFbPageID != nil {
		out["secondary_fb_page_id"] = *params.SecondaryFbPageID
	}
	if params.SecondaryIgUserID != nil {
		out["secondary_ig_user_id"] = *params.SecondaryIgUserID
	}
	return out
}

func GetIGMediaBatchCall(id string, params GetIGMediaParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetIGMediaBatchRequest(id string, params GetIGMediaParams, options ...core.BatchOption) *core.BatchRequest[objects.IGMedia] {
	return core.NewBatchRequest[objects.IGMedia](GetIGMediaBatchCall(id, params, options...))
}

func DecodeGetIGMediaBatchResponse(response *core.BatchResponse) (*objects.IGMedia, error) {
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

func GetIGMedia(ctx context.Context, client *core.Client, id string, params GetIGMediaParams) (*objects.IGMedia, error) {
	var out objects.IGMedia
	call := GetIGMediaBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type UpdateIGMediaParams struct {
	CommentEnabled bool        `facebook:"comment_enabled"`
	Extra          core.Params `facebook:"-"`
}

func (params UpdateIGMediaParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["comment_enabled"] = params.CommentEnabled
	return out
}

func UpdateIGMediaBatchCall(id string, params UpdateIGMediaParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateIGMediaBatchRequest(id string, params UpdateIGMediaParams, options ...core.BatchOption) *core.BatchRequest[objects.IGMedia] {
	return core.NewBatchRequest[objects.IGMedia](UpdateIGMediaBatchCall(id, params, options...))
}

func DecodeUpdateIGMediaBatchResponse(response *core.BatchResponse) (*objects.IGMedia, error) {
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

func UpdateIGMedia(ctx context.Context, client *core.Client, id string, params UpdateIGMediaParams) (*objects.IGMedia, error) {
	var out objects.IGMedia
	call := UpdateIGMediaBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
