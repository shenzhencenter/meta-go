package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
	"time"
)

type GetIGUserForIGOnlyAPIBusinessMessagingFeatureStatusParams struct {
	Feature string      `facebook:"feature"`
	Extra   core.Params `facebook:"-"`
}

func (params GetIGUserForIGOnlyAPIBusinessMessagingFeatureStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["feature"] = params.Feature
	return out
}

func GetIGUserForIGOnlyAPIBusinessMessagingFeatureStatus(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIBusinessMessagingFeatureStatusParams) (*core.Cursor[objects.BusinessMessagingFeatureStatus], error) {
	var out core.Cursor[objects.BusinessMessagingFeatureStatus]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "business_messaging_feature_status"), params.ToParams(), &out)
	return &out, err
}

type GetIGUserForIGOnlyAPIContentPublishingLimitParams struct {
	Since *time.Time  `facebook:"since"`
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserForIGOnlyAPIContentPublishingLimitParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	return out
}

func GetIGUserForIGOnlyAPIContentPublishingLimit(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIContentPublishingLimitParams) (*core.Cursor[objects.ContentPublishingLimitResponse], error) {
	var out core.Cursor[objects.ContentPublishingLimitResponse]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "content_publishing_limit"), params.ToParams(), &out)
	return &out, err
}

type GetIGUserForIGOnlyAPIConversationsParams struct {
	Folder   *string                                   `facebook:"folder"`
	Platform *enums.UserconversationsPlatformEnumParam `facebook:"platform"`
	Tags     *[]string                                 `facebook:"tags"`
	UserID   *core.ID                                  `facebook:"user_id"`
	Extra    core.Params                               `facebook:"-"`
}

func (params GetIGUserForIGOnlyAPIConversationsParams) ToParams() core.Params {
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

func GetIGUserForIGOnlyAPIConversations(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIConversationsParams) (*core.Cursor[objects.UnifiedThread], error) {
	var out core.Cursor[objects.UnifiedThread]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "conversations"), params.ToParams(), &out)
	return &out, err
}

type GetIGUserForIGOnlyAPIInsightsParams struct {
	Breakdown  *[]enums.UserinsightsBreakdownEnumParam `facebook:"breakdown"`
	Metric     []enums.UserinsightsMetricEnumParam     `facebook:"metric"`
	MetricType *enums.UserinsightsMetricTypeEnumParam  `facebook:"metric_type"`
	Period     []enums.UserinsightsPeriodEnumParam     `facebook:"period"`
	Since      *time.Time                              `facebook:"since"`
	Timeframe  *enums.UserinsightsTimeframeEnumParam   `facebook:"timeframe"`
	Until      *time.Time                              `facebook:"until"`
	Extra      core.Params                             `facebook:"-"`
}

func (params GetIGUserForIGOnlyAPIInsightsParams) ToParams() core.Params {
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

func GetIGUserForIGOnlyAPIInsights(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIInsightsParams) (*core.Cursor[objects.InsightsResult], error) {
	var out core.Cursor[objects.InsightsResult]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "insights"), params.ToParams(), &out)
	return &out, err
}

type GetIGUserForIGOnlyAPILiveMediaParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserForIGOnlyAPILiveMediaParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGUserForIGOnlyAPILiveMedia(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPILiveMediaParams) (*core.Cursor[objects.Media], error) {
	var out core.Cursor[objects.Media]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "live_media"), params.ToParams(), &out)
	return &out, err
}

type GetIGUserForIGOnlyAPIMediaParams struct {
	Since *time.Time  `facebook:"since"`
	Until *time.Time  `facebook:"until"`
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserForIGOnlyAPIMediaParams) ToParams() core.Params {
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

func GetIGUserForIGOnlyAPIMedia(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIMediaParams) (*core.Cursor[objects.Media], error) {
	var out core.Cursor[objects.Media]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "media"), params.ToParams(), &out)
	return &out, err
}

type CreateIGUserForIGOnlyAPIMediaParams struct {
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

func (params CreateIGUserForIGOnlyAPIMediaParams) ToParams() core.Params {
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

func CreateIGUserForIGOnlyAPIMedia(ctx context.Context, client *core.Client, id string, params CreateIGUserForIGOnlyAPIMediaParams) (*objects.IGGraphMedia, error) {
	var out objects.IGGraphMedia
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "media"), params.ToParams(), &out)
	return &out, err
}

type CreateIGUserForIGOnlyAPIMediapublishParams struct {
	CreationID core.ID     `facebook:"creation_id"`
	Extra      core.Params `facebook:"-"`
}

func (params CreateIGUserForIGOnlyAPIMediapublishParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["creation_id"] = params.CreationID
	return out
}

func CreateIGUserForIGOnlyAPIMediapublish(ctx context.Context, client *core.Client, id string, params CreateIGUserForIGOnlyAPIMediapublishParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "mediapublish"), params.ToParams(), &out)
	return &out, err
}

type CreateIGUserForIGOnlyAPIMentionsParams struct {
	CommentID *core.ID    `facebook:"comment_id"`
	MediaID   core.ID     `facebook:"media_id"`
	Message   string      `facebook:"message"`
	Extra     core.Params `facebook:"-"`
}

func (params CreateIGUserForIGOnlyAPIMentionsParams) ToParams() core.Params {
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

func CreateIGUserForIGOnlyAPIMentions(ctx context.Context, client *core.Client, id string, params CreateIGUserForIGOnlyAPIMentionsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "mentions"), params.ToParams(), &out)
	return &out, err
}

type CreateIGUserForIGOnlyAPIMessageattachmentsParams struct {
	Message map[string]interface{} `facebook:"message"`
	Extra   core.Params            `facebook:"-"`
}

func (params CreateIGUserForIGOnlyAPIMessageattachmentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["message"] = params.Message
	return out
}

func CreateIGUserForIGOnlyAPIMessageattachments(ctx context.Context, client *core.Client, id string, params CreateIGUserForIGOnlyAPIMessageattachmentsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "messageattachments"), params.ToParams(), &out)
	return &out, err
}

type CreateIGUserForIGOnlyAPIMessagesParams struct {
	Folder        *enums.IggraphusermessagesFolderEnumParam        `facebook:"folder"`
	Message       *map[string]interface{}                          `facebook:"message"`
	MessagingType *enums.IggraphusermessagesMessagingTypeEnumParam `facebook:"messaging_type"`
	Payload       *string                                          `facebook:"payload"`
	Recipient     *map[string]interface{}                          `facebook:"recipient"`
	ReplyTo       *map[string]interface{}                          `facebook:"reply_to"`
	SenderAction  *enums.IggraphusermessagesSenderActionEnumParam  `facebook:"sender_action"`
	Tag           *map[string]interface{}                          `facebook:"tag"`
	ThreadControl *map[string]interface{}                          `facebook:"thread_control"`
	Extra         core.Params                                      `facebook:"-"`
}

func (params CreateIGUserForIGOnlyAPIMessagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Folder != nil {
		out["folder"] = *params.Folder
	}
	if params.Message != nil {
		out["message"] = *params.Message
	}
	if params.MessagingType != nil {
		out["messaging_type"] = *params.MessagingType
	}
	if params.Payload != nil {
		out["payload"] = *params.Payload
	}
	if params.Recipient != nil {
		out["recipient"] = *params.Recipient
	}
	if params.ReplyTo != nil {
		out["reply_to"] = *params.ReplyTo
	}
	if params.SenderAction != nil {
		out["sender_action"] = *params.SenderAction
	}
	if params.Tag != nil {
		out["tag"] = *params.Tag
	}
	if params.ThreadControl != nil {
		out["thread_control"] = *params.ThreadControl
	}
	return out
}

func CreateIGUserForIGOnlyAPIMessages(ctx context.Context, client *core.Client, id string, params CreateIGUserForIGOnlyAPIMessagesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "messages"), params.ToParams(), &out)
	return &out, err
}

type DeleteIGUserForIGOnlyAPIMessengerProfileParams struct {
	Fields []enums.IggraphusermessengerProfileFieldsEnumParam `facebook:"fields"`
	Extra  core.Params                                        `facebook:"-"`
}

func (params DeleteIGUserForIGOnlyAPIMessengerProfileParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["fields"] = params.Fields
	return out
}

func DeleteIGUserForIGOnlyAPIMessengerProfile(ctx context.Context, client *core.Client, id string, params DeleteIGUserForIGOnlyAPIMessengerProfileParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "messenger_profile"), params.ToParams(), &out)
	return &out, err
}

type GetIGUserForIGOnlyAPIMessengerProfileParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserForIGOnlyAPIMessengerProfileParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGUserForIGOnlyAPIMessengerProfile(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIMessengerProfileParams) (*core.Cursor[objects.UserProfile], error) {
	var out core.Cursor[objects.UserProfile]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "messenger_profile"), params.ToParams(), &out)
	return &out, err
}

type CreateIGUserForIGOnlyAPIMessengerProfileParams struct {
	IceBreakers    *[]map[string]interface{} `facebook:"ice_breakers"`
	PersistentMenu *[]map[string]interface{} `facebook:"persistent_menu"`
	Extra          core.Params               `facebook:"-"`
}

func (params CreateIGUserForIGOnlyAPIMessengerProfileParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.IceBreakers != nil {
		out["ice_breakers"] = *params.IceBreakers
	}
	if params.PersistentMenu != nil {
		out["persistent_menu"] = *params.PersistentMenu
	}
	return out
}

func CreateIGUserForIGOnlyAPIMessengerProfile(ctx context.Context, client *core.Client, id string, params CreateIGUserForIGOnlyAPIMessengerProfileParams) (*objects.IGGraphUser, error) {
	var out objects.IGGraphUser
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "messenger_profile"), params.ToParams(), &out)
	return &out, err
}

type GetIGUserForIGOnlyAPIStoriesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserForIGOnlyAPIStoriesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGUserForIGOnlyAPIStories(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIStoriesParams) (*core.Cursor[objects.Media], error) {
	var out core.Cursor[objects.Media]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "stories"), params.ToParams(), &out)
	return &out, err
}

type DeleteIGUserForIGOnlyAPISubscribedAppsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteIGUserForIGOnlyAPISubscribedAppsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteIGUserForIGOnlyAPISubscribedApps(ctx context.Context, client *core.Client, id string, params DeleteIGUserForIGOnlyAPISubscribedAppsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "subscribed_apps"), params.ToParams(), &out)
	return &out, err
}

type GetIGUserForIGOnlyAPISubscribedAppsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserForIGOnlyAPISubscribedAppsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGUserForIGOnlyAPISubscribedApps(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPISubscribedAppsParams) (*core.Cursor[objects.UserSubscribedAppsData], error) {
	var out core.Cursor[objects.UserSubscribedAppsData]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "subscribed_apps"), params.ToParams(), &out)
	return &out, err
}

type CreateIGUserForIGOnlyAPISubscribedAppsParams struct {
	SubscribedFields []enums.IggraphusersubscribedAppsSubscribedFieldsEnumParam `facebook:"subscribed_fields"`
	Extra            core.Params                                                `facebook:"-"`
}

func (params CreateIGUserForIGOnlyAPISubscribedAppsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["subscribed_fields"] = params.SubscribedFields
	return out
}

func CreateIGUserForIGOnlyAPISubscribedApps(ctx context.Context, client *core.Client, id string, params CreateIGUserForIGOnlyAPISubscribedAppsParams) (*objects.IGGraphUser, error) {
	var out objects.IGGraphUser
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "subscribed_apps"), params.ToParams(), &out)
	return &out, err
}

type GetIGUserForIGOnlyAPITagsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserForIGOnlyAPITagsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGUserForIGOnlyAPITags(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPITagsParams) (*core.Cursor[objects.Media], error) {
	var out core.Cursor[objects.Media]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "tags"), params.ToParams(), &out)
	return &out, err
}

type DeleteIGUserForIGOnlyAPIWelcomeMessageFlowsParams struct {
	FlowID *core.ID    `facebook:"flow_id"`
	Extra  core.Params `facebook:"-"`
}

func (params DeleteIGUserForIGOnlyAPIWelcomeMessageFlowsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.FlowID != nil {
		out["flow_id"] = *params.FlowID
	}
	return out
}

func DeleteIGUserForIGOnlyAPIWelcomeMessageFlows(ctx context.Context, client *core.Client, id string, params DeleteIGUserForIGOnlyAPIWelcomeMessageFlowsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "welcome_message_flows"), params.ToParams(), &out)
	return &out, err
}

type GetIGUserForIGOnlyAPIWelcomeMessageFlowsParams struct {
	AppID  *core.ID    `facebook:"app_id"`
	FlowID *core.ID    `facebook:"flow_id"`
	Extra  core.Params `facebook:"-"`
}

func (params GetIGUserForIGOnlyAPIWelcomeMessageFlowsParams) ToParams() core.Params {
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

func GetIGUserForIGOnlyAPIWelcomeMessageFlows(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIWelcomeMessageFlowsParams) (*core.Cursor[objects.CTXPartnerAppWelcomeMessageFlow], error) {
	var out core.Cursor[objects.CTXPartnerAppWelcomeMessageFlow]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "welcome_message_flows"), params.ToParams(), &out)
	return &out, err
}

type CreateIGUserForIGOnlyAPIWelcomeMessageFlowsParams struct {
	EligiblePlatforms  *[]enums.IggraphuserwelcomeMessageFlowsEligiblePlatformsEnumParam `facebook:"eligible_platforms"`
	FlowID             *core.ID                                                          `facebook:"flow_id"`
	Name               *string                                                           `facebook:"name"`
	WelcomeMessageFlow *[]map[string]interface{}                                         `facebook:"welcome_message_flow"`
	Extra              core.Params                                                       `facebook:"-"`
}

func (params CreateIGUserForIGOnlyAPIWelcomeMessageFlowsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.EligiblePlatforms != nil {
		out["eligible_platforms"] = *params.EligiblePlatforms
	}
	if params.FlowID != nil {
		out["flow_id"] = *params.FlowID
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.WelcomeMessageFlow != nil {
		out["welcome_message_flow"] = *params.WelcomeMessageFlow
	}
	return out
}

func CreateIGUserForIGOnlyAPIWelcomeMessageFlows(ctx context.Context, client *core.Client, id string, params CreateIGUserForIGOnlyAPIWelcomeMessageFlowsParams) (*objects.IGGraphCTXPartnerAppWelcomeMessageFlow, error) {
	var out objects.IGGraphCTXPartnerAppWelcomeMessageFlow
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "welcome_message_flows"), params.ToParams(), &out)
	return &out, err
}

type GetIGUserForIGOnlyAPIParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserForIGOnlyAPIParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGUserForIGOnlyAPI(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIParams) (*objects.User, error) {
	var out objects.User
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
