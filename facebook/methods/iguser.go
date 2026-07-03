package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
	"time"
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

func GetIGUserAgencies(ctx context.Context, client *core.Client, id string, params GetIGUserAgenciesParams) (*core.Cursor[objects.Business], error) {
	var out core.Cursor[objects.Business]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "agencies"), params.ToParams(), &out)
	return &out, err
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

func GetIGUserAuthorizedAdaccounts(ctx context.Context, client *core.Client, id string, params GetIGUserAuthorizedAdaccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "authorized_adaccounts"), params.ToParams(), &out)
	return &out, err
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

func CreateIGUserAuthorizedAdaccounts(ctx context.Context, client *core.Client, id string, params CreateIGUserAuthorizedAdaccountsParams) (*objects.IGUser, error) {
	var out objects.IGUser
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "authorized_adaccounts"), params.ToParams(), &out)
	return &out, err
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

func GetIGUserAvailableCatalogs(ctx context.Context, client *core.Client, id string, params GetIGUserAvailableCatalogsParams) (*core.Cursor[objects.UserAvailableCatalogs], error) {
	var out core.Cursor[objects.UserAvailableCatalogs]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "available_catalogs"), params.ToParams(), &out)
	return &out, err
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

func GetIGUserBrandedContentAdPermissions(ctx context.Context, client *core.Client, id string, params GetIGUserBrandedContentAdPermissionsParams) (*core.Cursor[objects.IGBCAdsPermission], error) {
	var out core.Cursor[objects.IGBCAdsPermission]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "branded_content_ad_permissions"), params.ToParams(), &out)
	return &out, err
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

func CreateIGUserBrandedContentAdPermissions(ctx context.Context, client *core.Client, id string, params CreateIGUserBrandedContentAdPermissionsParams) (*objects.IGBCAdsPermission, error) {
	var out objects.IGBCAdsPermission
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "branded_content_ad_permissions"), params.ToParams(), &out)
	return &out, err
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

func GetIGUserBrandedContentAdvertisableMedias(ctx context.Context, client *core.Client, id string, params GetIGUserBrandedContentAdvertisableMediasParams) (*core.Cursor[objects.BrandedContentShadowIGMediaID], error) {
	var out core.Cursor[objects.BrandedContentShadowIGMediaID]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "branded_content_advertisable_medias"), params.ToParams(), &out)
	return &out, err
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

func DeleteIGUserBrandedContentTagApproval(ctx context.Context, client *core.Client, id string, params DeleteIGUserBrandedContentTagApprovalParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "branded_content_tag_approval"), params.ToParams(), &out)
	return &out, err
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

func GetIGUserBrandedContentTagApproval(ctx context.Context, client *core.Client, id string, params GetIGUserBrandedContentTagApprovalParams) (*core.Cursor[objects.BrandedContentShadowIGUserID], error) {
	var out core.Cursor[objects.BrandedContentShadowIGUserID]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "branded_content_tag_approval"), params.ToParams(), &out)
	return &out, err
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

func CreateIGUserBrandedContentTagApproval(ctx context.Context, client *core.Client, id string, params CreateIGUserBrandedContentTagApprovalParams) (*objects.BrandedContentShadowIGUserID, error) {
	var out objects.BrandedContentShadowIGUserID
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "branded_content_tag_approval"), params.ToParams(), &out)
	return &out, err
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

func CreateIGUserBusinessMessagingFeatureStatus(ctx context.Context, client *core.Client, id string, params CreateIGUserBusinessMessagingFeatureStatusParams) (*objects.IGUser, error) {
	var out objects.IGUser
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "business_messaging_feature_status"), params.ToParams(), &out)
	return &out, err
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

func GetIGUserCatalogProductSearch(ctx context.Context, client *core.Client, id string, params GetIGUserCatalogProductSearchParams) (*core.Cursor[objects.ShadowIGUserCatalogProductSearch], error) {
	var out core.Cursor[objects.ShadowIGUserCatalogProductSearch]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "catalog_product_search"), params.ToParams(), &out)
	return &out, err
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

func GetIGUserCollaborationInvites(ctx context.Context, client *core.Client, id string, params GetIGUserCollaborationInvitesParams) (*core.Cursor[objects.ShadowIGUserCollaborationInvites], error) {
	var out core.Cursor[objects.ShadowIGUserCollaborationInvites]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "collaboration_invites"), params.ToParams(), &out)
	return &out, err
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

func CreateIGUserCollaborationInvites(ctx context.Context, client *core.Client, id string, params CreateIGUserCollaborationInvitesParams) (*objects.ShadowIGUserCollaborationInvites, error) {
	var out objects.ShadowIGUserCollaborationInvites
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "collaboration_invites"), params.ToParams(), &out)
	return &out, err
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

func GetIGUserCollaborativeMedia(ctx context.Context, client *core.Client, id string, params GetIGUserCollaborativeMediaParams) (*core.Cursor[objects.ShadowIGUserCollaborativeMedia], error) {
	var out core.Cursor[objects.ShadowIGUserCollaborativeMedia]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "collaborative_media"), params.ToParams(), &out)
	return &out, err
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

func GetIGUserConnectedThreadsUser(ctx context.Context, client *core.Client, id string, params GetIGUserConnectedThreadsUserParams) (*core.Cursor[objects.ThreadsUser], error) {
	var out core.Cursor[objects.ThreadsUser]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "connected_threads_user"), params.ToParams(), &out)
	return &out, err
}

type GetIGUserContentPublishingLimitParams struct {
	Since *time.Time  `facebook:"since"`
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

func GetIGUserContentPublishingLimit(ctx context.Context, client *core.Client, id string, params GetIGUserContentPublishingLimitParams) (*core.Cursor[objects.ContentPublishingLimitResponse], error) {
	var out core.Cursor[objects.ContentPublishingLimitResponse]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "content_publishing_limit"), params.ToParams(), &out)
	return &out, err
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

func GetIGUserCreatorMarketplaceCreators(ctx context.Context, client *core.Client, id string, params GetIGUserCreatorMarketplaceCreatorsParams) (*core.Cursor[objects.IGUserExportForCAM], error) {
	var out core.Cursor[objects.IGUserExportForCAM]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "creator_marketplace_creators"), params.ToParams(), &out)
	return &out, err
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

func GetIGUserDataset(ctx context.Context, client *core.Client, id string, params GetIGUserDatasetParams) (*core.Cursor[objects.Dataset], error) {
	var out core.Cursor[objects.Dataset]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "dataset"), params.ToParams(), &out)
	return &out, err
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

func CreateIGUserDataset(ctx context.Context, client *core.Client, id string, params CreateIGUserDatasetParams) (*objects.Dataset, error) {
	var out objects.Dataset
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "dataset"), params.ToParams(), &out)
	return &out, err
}

type GetIGUserInsightsParams struct {
	Breakdown  *[]enums.ShadowiguserinsightsBreakdownEnumParam `facebook:"breakdown"`
	Metric     []enums.ShadowiguserinsightsMetricEnumParam     `facebook:"metric"`
	MetricType *enums.ShadowiguserinsightsMetricTypeEnumParam  `facebook:"metric_type"`
	Period     []enums.ShadowiguserinsightsPeriodEnumParam     `facebook:"period"`
	Since      *time.Time                                      `facebook:"since"`
	Timeframe  *enums.ShadowiguserinsightsTimeframeEnumParam   `facebook:"timeframe"`
	Until      *time.Time                                      `facebook:"until"`
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

func GetIGUserInsights(ctx context.Context, client *core.Client, id string, params GetIGUserInsightsParams) (*core.Cursor[objects.InstagramInsightsResult], error) {
	var out core.Cursor[objects.InstagramInsightsResult]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "insights"), params.ToParams(), &out)
	return &out, err
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

func GetIGUserInstagramBackedThreadsUser(ctx context.Context, client *core.Client, id string, params GetIGUserInstagramBackedThreadsUserParams) (*core.Cursor[objects.ThreadsUser], error) {
	var out core.Cursor[objects.ThreadsUser]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "instagram_backed_threads_user"), params.ToParams(), &out)
	return &out, err
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

func CreateIGUserInstagramBackedThreadsUser(ctx context.Context, client *core.Client, id string, params CreateIGUserInstagramBackedThreadsUserParams) (*objects.ThreadsUser, error) {
	var out objects.ThreadsUser
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "instagram_backed_threads_user"), params.ToParams(), &out)
	return &out, err
}

type GetIGUserLiveMediaParams struct {
	Since *time.Time  `facebook:"since"`
	Until *time.Time  `facebook:"until"`
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

func GetIGUserLiveMedia(ctx context.Context, client *core.Client, id string, params GetIGUserLiveMediaParams) (*core.Cursor[objects.IGMedia], error) {
	var out core.Cursor[objects.IGMedia]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "live_media"), params.ToParams(), &out)
	return &out, err
}

type GetIGUserMediaParams struct {
	Since *time.Time  `facebook:"since"`
	Until *time.Time  `facebook:"until"`
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

func GetIGUserMedia(ctx context.Context, client *core.Client, id string, params GetIGUserMediaParams) (*core.Cursor[objects.IGMedia], error) {
	var out core.Cursor[objects.IGMedia]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "media"), params.ToParams(), &out)
	return &out, err
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

func CreateIGUserMedia(ctx context.Context, client *core.Client, id string, params CreateIGUserMediaParams) (*objects.IGMedia, error) {
	var out objects.IGMedia
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "media"), params.ToParams(), &out)
	return &out, err
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

func CreateIGUserMediaPublish(ctx context.Context, client *core.Client, id string, params CreateIGUserMediaPublishParams) (*objects.IGMedia, error) {
	var out objects.IGMedia
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "media_publish"), params.ToParams(), &out)
	return &out, err
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

func CreateIGUserMentions(ctx context.Context, client *core.Client, id string, params CreateIGUserMentionsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "mentions"), params.ToParams(), &out)
	return &out, err
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

func CreateIGUserModerateConversations(ctx context.Context, client *core.Client, id string, params CreateIGUserModerateConversationsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "moderate_conversations"), params.ToParams(), &out)
	return &out, err
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

func GetIGUserNotificationMessageTokens(ctx context.Context, client *core.Client, id string, params GetIGUserNotificationMessageTokensParams) (*core.Cursor[objects.UserPageOneTimeOptInTokenSettings], error) {
	var out core.Cursor[objects.UserPageOneTimeOptInTokenSettings]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "notification_message_tokens"), params.ToParams(), &out)
	return &out, err
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

func GetIGUserProductAppeal(ctx context.Context, client *core.Client, id string, params GetIGUserProductAppealParams) (*core.Cursor[objects.IGShoppingProductAppeal], error) {
	var out core.Cursor[objects.IGShoppingProductAppeal]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "product_appeal"), params.ToParams(), &out)
	return &out, err
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

func CreateIGUserProductAppeal(ctx context.Context, client *core.Client, id string, params CreateIGUserProductAppealParams) (*objects.IGShoppingProductAppeal, error) {
	var out objects.IGShoppingProductAppeal
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "product_appeal"), params.ToParams(), &out)
	return &out, err
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

func GetIGUserRecentlySearchedHashtags(ctx context.Context, client *core.Client, id string, params GetIGUserRecentlySearchedHashtagsParams) (*core.Cursor[objects.ShadowIGHashtag], error) {
	var out core.Cursor[objects.ShadowIGHashtag]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "recently_searched_hashtags"), params.ToParams(), &out)
	return &out, err
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

func GetIGUserScheduledMedia(ctx context.Context, client *core.Client, id string, params GetIGUserScheduledMediaParams) (*core.Cursor[objects.ShadowIGScheduledMedia], error) {
	var out core.Cursor[objects.ShadowIGScheduledMedia]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "scheduled_media"), params.ToParams(), &out)
	return &out, err
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

func GetIGUserStories(ctx context.Context, client *core.Client, id string, params GetIGUserStoriesParams) (*core.Cursor[objects.IGMedia], error) {
	var out core.Cursor[objects.IGMedia]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "stories"), params.ToParams(), &out)
	return &out, err
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

func GetIGUserTags(ctx context.Context, client *core.Client, id string, params GetIGUserTagsParams) (*core.Cursor[objects.IGMedia], error) {
	var out core.Cursor[objects.IGMedia]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "tags"), params.ToParams(), &out)
	return &out, err
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

func GetIGUserUpcomingEvents(ctx context.Context, client *core.Client, id string, params GetIGUserUpcomingEventsParams) (*core.Cursor[objects.IGUpcomingEvent], error) {
	var out core.Cursor[objects.IGUpcomingEvent]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "upcoming_events"), params.ToParams(), &out)
	return &out, err
}

type CreateIGUserUpcomingEventsParams struct {
	EndTime              *time.Time                                                       `facebook:"end_time"`
	NotificationSubtypes *[]enums.ShadowiguserupcomingEventsNotificationSubtypesEnumParam `facebook:"notification_subtypes"`
	StartTime            time.Time                                                        `facebook:"start_time"`
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

func CreateIGUserUpcomingEvents(ctx context.Context, client *core.Client, id string, params CreateIGUserUpcomingEventsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "upcoming_events"), params.ToParams(), &out)
	return &out, err
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

func GetIGUserWelcomeMessageFlows(ctx context.Context, client *core.Client, id string, params GetIGUserWelcomeMessageFlowsParams) (*core.Cursor[objects.ShadowIGUserCTXPartnerAppWelcomeMessageFlow], error) {
	var out core.Cursor[objects.ShadowIGUserCTXPartnerAppWelcomeMessageFlow]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "welcome_message_flows"), params.ToParams(), &out)
	return &out, err
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

func GetIGUser(ctx context.Context, client *core.Client, id string, params GetIGUserParams) (*objects.IGUser, error) {
	var out objects.IGUser
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
