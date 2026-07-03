package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func GetIGMediaBoostAdsList(ctx context.Context, client *core.Client, id string, params GetIGMediaBoostAdsListParams) (*core.Cursor[objects.IGBoostMediaAd], error) {
	var out core.Cursor[objects.IGBoostMediaAd]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "boost_ads_list"), params.ToParams(), &out)
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

func GetIGMediaBrandedContentPartnerPromote(ctx context.Context, client *core.Client, id string, params GetIGMediaBrandedContentPartnerPromoteParams) (*core.Cursor[objects.BrandedContentShadowIGUserID], error) {
	var out core.Cursor[objects.BrandedContentShadowIGUserID]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "branded_content_partner_promote"), params.ToParams(), &out)
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

func CreateIGMediaBrandedContentPartnerPromote(ctx context.Context, client *core.Client, id string, params CreateIGMediaBrandedContentPartnerPromoteParams) (*objects.BrandedContentShadowIGUserID, error) {
	var out objects.BrandedContentShadowIGUserID
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "branded_content_partner_promote"), params.ToParams(), &out)
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

func GetIGMediaChildren(ctx context.Context, client *core.Client, id string, params GetIGMediaChildrenParams) (*core.Cursor[objects.IGMedia], error) {
	var out core.Cursor[objects.IGMedia]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "children"), params.ToParams(), &out)
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

func GetIGMediaCollaborators(ctx context.Context, client *core.Client, id string, params GetIGMediaCollaboratorsParams) (*core.Cursor[objects.ShadowIGMediaCollaborators], error) {
	var out core.Cursor[objects.ShadowIGMediaCollaborators]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "collaborators"), params.ToParams(), &out)
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

func GetIGMediaComments(ctx context.Context, client *core.Client, id string, params GetIGMediaCommentsParams) (*core.Cursor[objects.IGComment], error) {
	var out core.Cursor[objects.IGComment]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "comments"), params.ToParams(), &out)
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

func CreateIGMediaComments(ctx context.Context, client *core.Client, id string, params CreateIGMediaCommentsParams) (*objects.IGComment, error) {
	var out objects.IGComment
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "comments"), params.ToParams(), &out)
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

func GetIGMediaInsights(ctx context.Context, client *core.Client, id string, params GetIGMediaInsightsParams) (*core.Cursor[objects.InstagramInsightsResult], error) {
	var out core.Cursor[objects.InstagramInsightsResult]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "insights"), params.ToParams(), &out)
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

func DeleteIGMediaPartnershipAdCode(ctx context.Context, client *core.Client, id string, params DeleteIGMediaPartnershipAdCodeParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "partnership_ad_code"), params.ToParams(), &out)
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

func CreateIGMediaPartnershipAdCode(ctx context.Context, client *core.Client, id string, params CreateIGMediaPartnershipAdCodeParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "partnership_ad_code"), params.ToParams(), &out)
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

func GetIGMediaProductTags(ctx context.Context, client *core.Client, id string, params GetIGMediaProductTagsParams) (*core.Cursor[objects.ShadowIGMediaProductTags], error) {
	var out core.Cursor[objects.ShadowIGMediaProductTags]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "product_tags"), params.ToParams(), &out)
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

func CreateIGMediaProductTags(ctx context.Context, client *core.Client, id string, params CreateIGMediaProductTagsParams) (*objects.ShadowIGMediaProductTags, error) {
	var out objects.ShadowIGMediaProductTags
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "product_tags"), params.ToParams(), &out)
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

func DeleteIGMedia(ctx context.Context, client *core.Client, id string, params DeleteIGMediaParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
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

func GetIGMedia(ctx context.Context, client *core.Client, id string, params GetIGMediaParams) (*objects.IGMedia, error) {
	var out objects.IGMedia
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
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

func UpdateIGMedia(ctx context.Context, client *core.Client, id string, params UpdateIGMediaParams) (*objects.IGMedia, error) {
	var out objects.IGMedia
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
