package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
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

func GetIGUserForIGOnlyAPIBusinessMessagingFeatureStatusBatchCall(id string, params GetIGUserForIGOnlyAPIBusinessMessagingFeatureStatusParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "business_messaging_feature_status"), params.ToParams(), options...)
}

func NewGetIGUserForIGOnlyAPIBusinessMessagingFeatureStatusBatchRequest(id string, params GetIGUserForIGOnlyAPIBusinessMessagingFeatureStatusParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BusinessMessagingFeatureStatus]] {
	return core.NewBatchRequest[core.Cursor[objects.BusinessMessagingFeatureStatus]](GetIGUserForIGOnlyAPIBusinessMessagingFeatureStatusBatchCall(id, params, options...))
}

func DecodeGetIGUserForIGOnlyAPIBusinessMessagingFeatureStatusBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BusinessMessagingFeatureStatus], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.BusinessMessagingFeatureStatus]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserForIGOnlyAPIBusinessMessagingFeatureStatusWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIBusinessMessagingFeatureStatusParams) (*core.Cursor[objects.BusinessMessagingFeatureStatus], *core.Response, error) {
	var out core.Cursor[objects.BusinessMessagingFeatureStatus]
	call := GetIGUserForIGOnlyAPIBusinessMessagingFeatureStatusBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserForIGOnlyAPIBusinessMessagingFeatureStatus(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIBusinessMessagingFeatureStatusParams) (*core.Cursor[objects.BusinessMessagingFeatureStatus], error) {
	out, _, err := GetIGUserForIGOnlyAPIBusinessMessagingFeatureStatusWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserForIGOnlyAPIContentPublishingLimitParams struct {
	Since *core.Time  `facebook:"since"`
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

func GetIGUserForIGOnlyAPIContentPublishingLimitBatchCall(id string, params GetIGUserForIGOnlyAPIContentPublishingLimitParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "content_publishing_limit"), params.ToParams(), options...)
}

func NewGetIGUserForIGOnlyAPIContentPublishingLimitBatchRequest(id string, params GetIGUserForIGOnlyAPIContentPublishingLimitParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ContentPublishingLimitResponse]] {
	return core.NewBatchRequest[core.Cursor[objects.ContentPublishingLimitResponse]](GetIGUserForIGOnlyAPIContentPublishingLimitBatchCall(id, params, options...))
}

func DecodeGetIGUserForIGOnlyAPIContentPublishingLimitBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ContentPublishingLimitResponse], error) {
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

func GetIGUserForIGOnlyAPIContentPublishingLimitWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIContentPublishingLimitParams) (*core.Cursor[objects.ContentPublishingLimitResponse], *core.Response, error) {
	var out core.Cursor[objects.ContentPublishingLimitResponse]
	call := GetIGUserForIGOnlyAPIContentPublishingLimitBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserForIGOnlyAPIContentPublishingLimit(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIContentPublishingLimitParams) (*core.Cursor[objects.ContentPublishingLimitResponse], error) {
	out, _, err := GetIGUserForIGOnlyAPIContentPublishingLimitWithResponse(ctx, client, id, params)
	return out, err
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

func GetIGUserForIGOnlyAPIConversationsBatchCall(id string, params GetIGUserForIGOnlyAPIConversationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "conversations"), params.ToParams(), options...)
}

func NewGetIGUserForIGOnlyAPIConversationsBatchRequest(id string, params GetIGUserForIGOnlyAPIConversationsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.UnifiedThread]] {
	return core.NewBatchRequest[core.Cursor[objects.UnifiedThread]](GetIGUserForIGOnlyAPIConversationsBatchCall(id, params, options...))
}

func DecodeGetIGUserForIGOnlyAPIConversationsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.UnifiedThread], error) {
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

func GetIGUserForIGOnlyAPIConversationsWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIConversationsParams) (*core.Cursor[objects.UnifiedThread], *core.Response, error) {
	var out core.Cursor[objects.UnifiedThread]
	call := GetIGUserForIGOnlyAPIConversationsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserForIGOnlyAPIConversations(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIConversationsParams) (*core.Cursor[objects.UnifiedThread], error) {
	out, _, err := GetIGUserForIGOnlyAPIConversationsWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserForIGOnlyAPIInsightsParams struct {
	Breakdown  *[]enums.UserinsightsBreakdownEnumParam `facebook:"breakdown"`
	Metric     []enums.UserinsightsMetricEnumParam     `facebook:"metric"`
	MetricType *enums.UserinsightsMetricTypeEnumParam  `facebook:"metric_type"`
	Period     []enums.UserinsightsPeriodEnumParam     `facebook:"period"`
	Since      *core.Time                              `facebook:"since"`
	Timeframe  *enums.UserinsightsTimeframeEnumParam   `facebook:"timeframe"`
	Until      *core.Time                              `facebook:"until"`
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

func GetIGUserForIGOnlyAPIInsightsBatchCall(id string, params GetIGUserForIGOnlyAPIInsightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "insights"), params.ToParams(), options...)
}

func NewGetIGUserForIGOnlyAPIInsightsBatchRequest(id string, params GetIGUserForIGOnlyAPIInsightsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.InsightsResult]] {
	return core.NewBatchRequest[core.Cursor[objects.InsightsResult]](GetIGUserForIGOnlyAPIInsightsBatchCall(id, params, options...))
}

func DecodeGetIGUserForIGOnlyAPIInsightsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.InsightsResult], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.InsightsResult]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserForIGOnlyAPIInsightsWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIInsightsParams) (*core.Cursor[objects.InsightsResult], *core.Response, error) {
	var out core.Cursor[objects.InsightsResult]
	call := GetIGUserForIGOnlyAPIInsightsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserForIGOnlyAPIInsights(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIInsightsParams) (*core.Cursor[objects.InsightsResult], error) {
	out, _, err := GetIGUserForIGOnlyAPIInsightsWithResponse(ctx, client, id, params)
	return out, err
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

func GetIGUserForIGOnlyAPILiveMediaBatchCall(id string, params GetIGUserForIGOnlyAPILiveMediaParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "live_media"), params.ToParams(), options...)
}

func NewGetIGUserForIGOnlyAPILiveMediaBatchRequest(id string, params GetIGUserForIGOnlyAPILiveMediaParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Media]] {
	return core.NewBatchRequest[core.Cursor[objects.Media]](GetIGUserForIGOnlyAPILiveMediaBatchCall(id, params, options...))
}

func DecodeGetIGUserForIGOnlyAPILiveMediaBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Media], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Media]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserForIGOnlyAPILiveMediaWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPILiveMediaParams) (*core.Cursor[objects.Media], *core.Response, error) {
	var out core.Cursor[objects.Media]
	call := GetIGUserForIGOnlyAPILiveMediaBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserForIGOnlyAPILiveMedia(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPILiveMediaParams) (*core.Cursor[objects.Media], error) {
	out, _, err := GetIGUserForIGOnlyAPILiveMediaWithResponse(ctx, client, id, params)
	return out, err
}

type GetIGUserForIGOnlyAPIMediaParams struct {
	Since *core.Time  `facebook:"since"`
	Until *core.Time  `facebook:"until"`
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

func GetIGUserForIGOnlyAPIMediaBatchCall(id string, params GetIGUserForIGOnlyAPIMediaParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "media"), params.ToParams(), options...)
}

func NewGetIGUserForIGOnlyAPIMediaBatchRequest(id string, params GetIGUserForIGOnlyAPIMediaParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Media]] {
	return core.NewBatchRequest[core.Cursor[objects.Media]](GetIGUserForIGOnlyAPIMediaBatchCall(id, params, options...))
}

func DecodeGetIGUserForIGOnlyAPIMediaBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Media], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Media]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserForIGOnlyAPIMediaWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIMediaParams) (*core.Cursor[objects.Media], *core.Response, error) {
	var out core.Cursor[objects.Media]
	call := GetIGUserForIGOnlyAPIMediaBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserForIGOnlyAPIMedia(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIMediaParams) (*core.Cursor[objects.Media], error) {
	out, _, err := GetIGUserForIGOnlyAPIMediaWithResponse(ctx, client, id, params)
	return out, err
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

func CreateIGUserForIGOnlyAPIMediaBatchCall(id string, params CreateIGUserForIGOnlyAPIMediaParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "media"), params.ToParams(), options...)
}

func NewCreateIGUserForIGOnlyAPIMediaBatchRequest(id string, params CreateIGUserForIGOnlyAPIMediaParams, options ...core.BatchOption) *core.BatchRequest[objects.IGGraphMedia] {
	return core.NewBatchRequest[objects.IGGraphMedia](CreateIGUserForIGOnlyAPIMediaBatchCall(id, params, options...))
}

func DecodeCreateIGUserForIGOnlyAPIMediaBatchResponse(response *core.BatchResponse) (*objects.IGGraphMedia, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.IGGraphMedia
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateIGUserForIGOnlyAPIMediaWithResponse(ctx context.Context, client *core.Client, id string, params CreateIGUserForIGOnlyAPIMediaParams) (*objects.IGGraphMedia, *core.Response, error) {
	var out objects.IGGraphMedia
	call := CreateIGUserForIGOnlyAPIMediaBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateIGUserForIGOnlyAPIMedia(ctx context.Context, client *core.Client, id string, params CreateIGUserForIGOnlyAPIMediaParams) (*objects.IGGraphMedia, error) {
	out, _, err := CreateIGUserForIGOnlyAPIMediaWithResponse(ctx, client, id, params)
	return out, err
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

func CreateIGUserForIGOnlyAPIMediapublishBatchCall(id string, params CreateIGUserForIGOnlyAPIMediapublishParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "mediapublish"), params.ToParams(), options...)
}

func NewCreateIGUserForIGOnlyAPIMediapublishBatchRequest(id string, params CreateIGUserForIGOnlyAPIMediapublishParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateIGUserForIGOnlyAPIMediapublishBatchCall(id, params, options...))
}

func DecodeCreateIGUserForIGOnlyAPIMediapublishBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreateIGUserForIGOnlyAPIMediapublishWithResponse(ctx context.Context, client *core.Client, id string, params CreateIGUserForIGOnlyAPIMediapublishParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreateIGUserForIGOnlyAPIMediapublishBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateIGUserForIGOnlyAPIMediapublish(ctx context.Context, client *core.Client, id string, params CreateIGUserForIGOnlyAPIMediapublishParams) (*map[string]interface{}, error) {
	out, _, err := CreateIGUserForIGOnlyAPIMediapublishWithResponse(ctx, client, id, params)
	return out, err
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

func CreateIGUserForIGOnlyAPIMentionsBatchCall(id string, params CreateIGUserForIGOnlyAPIMentionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "mentions"), params.ToParams(), options...)
}

func NewCreateIGUserForIGOnlyAPIMentionsBatchRequest(id string, params CreateIGUserForIGOnlyAPIMentionsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateIGUserForIGOnlyAPIMentionsBatchCall(id, params, options...))
}

func DecodeCreateIGUserForIGOnlyAPIMentionsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreateIGUserForIGOnlyAPIMentionsWithResponse(ctx context.Context, client *core.Client, id string, params CreateIGUserForIGOnlyAPIMentionsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreateIGUserForIGOnlyAPIMentionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateIGUserForIGOnlyAPIMentions(ctx context.Context, client *core.Client, id string, params CreateIGUserForIGOnlyAPIMentionsParams) (*map[string]interface{}, error) {
	out, _, err := CreateIGUserForIGOnlyAPIMentionsWithResponse(ctx, client, id, params)
	return out, err
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

func CreateIGUserForIGOnlyAPIMessageattachmentsBatchCall(id string, params CreateIGUserForIGOnlyAPIMessageattachmentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "messageattachments"), params.ToParams(), options...)
}

func NewCreateIGUserForIGOnlyAPIMessageattachmentsBatchRequest(id string, params CreateIGUserForIGOnlyAPIMessageattachmentsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateIGUserForIGOnlyAPIMessageattachmentsBatchCall(id, params, options...))
}

func DecodeCreateIGUserForIGOnlyAPIMessageattachmentsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreateIGUserForIGOnlyAPIMessageattachmentsWithResponse(ctx context.Context, client *core.Client, id string, params CreateIGUserForIGOnlyAPIMessageattachmentsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreateIGUserForIGOnlyAPIMessageattachmentsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateIGUserForIGOnlyAPIMessageattachments(ctx context.Context, client *core.Client, id string, params CreateIGUserForIGOnlyAPIMessageattachmentsParams) (*map[string]interface{}, error) {
	out, _, err := CreateIGUserForIGOnlyAPIMessageattachmentsWithResponse(ctx, client, id, params)
	return out, err
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

func CreateIGUserForIGOnlyAPIMessagesBatchCall(id string, params CreateIGUserForIGOnlyAPIMessagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "messages"), params.ToParams(), options...)
}

func NewCreateIGUserForIGOnlyAPIMessagesBatchRequest(id string, params CreateIGUserForIGOnlyAPIMessagesParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateIGUserForIGOnlyAPIMessagesBatchCall(id, params, options...))
}

func DecodeCreateIGUserForIGOnlyAPIMessagesBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreateIGUserForIGOnlyAPIMessagesWithResponse(ctx context.Context, client *core.Client, id string, params CreateIGUserForIGOnlyAPIMessagesParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := CreateIGUserForIGOnlyAPIMessagesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateIGUserForIGOnlyAPIMessages(ctx context.Context, client *core.Client, id string, params CreateIGUserForIGOnlyAPIMessagesParams) (*map[string]interface{}, error) {
	out, _, err := CreateIGUserForIGOnlyAPIMessagesWithResponse(ctx, client, id, params)
	return out, err
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

func DeleteIGUserForIGOnlyAPIMessengerProfileBatchCall(id string, params DeleteIGUserForIGOnlyAPIMessengerProfileParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "messenger_profile"), params.ToParams(), options...)
}

func NewDeleteIGUserForIGOnlyAPIMessengerProfileBatchRequest(id string, params DeleteIGUserForIGOnlyAPIMessengerProfileParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteIGUserForIGOnlyAPIMessengerProfileBatchCall(id, params, options...))
}

func DecodeDeleteIGUserForIGOnlyAPIMessengerProfileBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteIGUserForIGOnlyAPIMessengerProfileWithResponse(ctx context.Context, client *core.Client, id string, params DeleteIGUserForIGOnlyAPIMessengerProfileParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteIGUserForIGOnlyAPIMessengerProfileBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteIGUserForIGOnlyAPIMessengerProfile(ctx context.Context, client *core.Client, id string, params DeleteIGUserForIGOnlyAPIMessengerProfileParams) (*map[string]interface{}, error) {
	out, _, err := DeleteIGUserForIGOnlyAPIMessengerProfileWithResponse(ctx, client, id, params)
	return out, err
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

func GetIGUserForIGOnlyAPIMessengerProfileBatchCall(id string, params GetIGUserForIGOnlyAPIMessengerProfileParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "messenger_profile"), params.ToParams(), options...)
}

func NewGetIGUserForIGOnlyAPIMessengerProfileBatchRequest(id string, params GetIGUserForIGOnlyAPIMessengerProfileParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.UserProfile]] {
	return core.NewBatchRequest[core.Cursor[objects.UserProfile]](GetIGUserForIGOnlyAPIMessengerProfileBatchCall(id, params, options...))
}

func DecodeGetIGUserForIGOnlyAPIMessengerProfileBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.UserProfile], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.UserProfile]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserForIGOnlyAPIMessengerProfileWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIMessengerProfileParams) (*core.Cursor[objects.UserProfile], *core.Response, error) {
	var out core.Cursor[objects.UserProfile]
	call := GetIGUserForIGOnlyAPIMessengerProfileBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserForIGOnlyAPIMessengerProfile(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIMessengerProfileParams) (*core.Cursor[objects.UserProfile], error) {
	out, _, err := GetIGUserForIGOnlyAPIMessengerProfileWithResponse(ctx, client, id, params)
	return out, err
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

func CreateIGUserForIGOnlyAPIMessengerProfileBatchCall(id string, params CreateIGUserForIGOnlyAPIMessengerProfileParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "messenger_profile"), params.ToParams(), options...)
}

func NewCreateIGUserForIGOnlyAPIMessengerProfileBatchRequest(id string, params CreateIGUserForIGOnlyAPIMessengerProfileParams, options ...core.BatchOption) *core.BatchRequest[objects.IGGraphUser] {
	return core.NewBatchRequest[objects.IGGraphUser](CreateIGUserForIGOnlyAPIMessengerProfileBatchCall(id, params, options...))
}

func DecodeCreateIGUserForIGOnlyAPIMessengerProfileBatchResponse(response *core.BatchResponse) (*objects.IGGraphUser, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.IGGraphUser
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateIGUserForIGOnlyAPIMessengerProfileWithResponse(ctx context.Context, client *core.Client, id string, params CreateIGUserForIGOnlyAPIMessengerProfileParams) (*objects.IGGraphUser, *core.Response, error) {
	var out objects.IGGraphUser
	call := CreateIGUserForIGOnlyAPIMessengerProfileBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateIGUserForIGOnlyAPIMessengerProfile(ctx context.Context, client *core.Client, id string, params CreateIGUserForIGOnlyAPIMessengerProfileParams) (*objects.IGGraphUser, error) {
	out, _, err := CreateIGUserForIGOnlyAPIMessengerProfileWithResponse(ctx, client, id, params)
	return out, err
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

func GetIGUserForIGOnlyAPIStoriesBatchCall(id string, params GetIGUserForIGOnlyAPIStoriesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "stories"), params.ToParams(), options...)
}

func NewGetIGUserForIGOnlyAPIStoriesBatchRequest(id string, params GetIGUserForIGOnlyAPIStoriesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Media]] {
	return core.NewBatchRequest[core.Cursor[objects.Media]](GetIGUserForIGOnlyAPIStoriesBatchCall(id, params, options...))
}

func DecodeGetIGUserForIGOnlyAPIStoriesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Media], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Media]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserForIGOnlyAPIStoriesWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIStoriesParams) (*core.Cursor[objects.Media], *core.Response, error) {
	var out core.Cursor[objects.Media]
	call := GetIGUserForIGOnlyAPIStoriesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserForIGOnlyAPIStories(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIStoriesParams) (*core.Cursor[objects.Media], error) {
	out, _, err := GetIGUserForIGOnlyAPIStoriesWithResponse(ctx, client, id, params)
	return out, err
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

func DeleteIGUserForIGOnlyAPISubscribedAppsBatchCall(id string, params DeleteIGUserForIGOnlyAPISubscribedAppsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "subscribed_apps"), params.ToParams(), options...)
}

func NewDeleteIGUserForIGOnlyAPISubscribedAppsBatchRequest(id string, params DeleteIGUserForIGOnlyAPISubscribedAppsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteIGUserForIGOnlyAPISubscribedAppsBatchCall(id, params, options...))
}

func DecodeDeleteIGUserForIGOnlyAPISubscribedAppsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteIGUserForIGOnlyAPISubscribedAppsWithResponse(ctx context.Context, client *core.Client, id string, params DeleteIGUserForIGOnlyAPISubscribedAppsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteIGUserForIGOnlyAPISubscribedAppsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteIGUserForIGOnlyAPISubscribedApps(ctx context.Context, client *core.Client, id string, params DeleteIGUserForIGOnlyAPISubscribedAppsParams) (*map[string]interface{}, error) {
	out, _, err := DeleteIGUserForIGOnlyAPISubscribedAppsWithResponse(ctx, client, id, params)
	return out, err
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

func GetIGUserForIGOnlyAPISubscribedAppsBatchCall(id string, params GetIGUserForIGOnlyAPISubscribedAppsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "subscribed_apps"), params.ToParams(), options...)
}

func NewGetIGUserForIGOnlyAPISubscribedAppsBatchRequest(id string, params GetIGUserForIGOnlyAPISubscribedAppsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.UserSubscribedAppsData]] {
	return core.NewBatchRequest[core.Cursor[objects.UserSubscribedAppsData]](GetIGUserForIGOnlyAPISubscribedAppsBatchCall(id, params, options...))
}

func DecodeGetIGUserForIGOnlyAPISubscribedAppsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.UserSubscribedAppsData], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.UserSubscribedAppsData]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserForIGOnlyAPISubscribedAppsWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPISubscribedAppsParams) (*core.Cursor[objects.UserSubscribedAppsData], *core.Response, error) {
	var out core.Cursor[objects.UserSubscribedAppsData]
	call := GetIGUserForIGOnlyAPISubscribedAppsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserForIGOnlyAPISubscribedApps(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPISubscribedAppsParams) (*core.Cursor[objects.UserSubscribedAppsData], error) {
	out, _, err := GetIGUserForIGOnlyAPISubscribedAppsWithResponse(ctx, client, id, params)
	return out, err
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

func CreateIGUserForIGOnlyAPISubscribedAppsBatchCall(id string, params CreateIGUserForIGOnlyAPISubscribedAppsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "subscribed_apps"), params.ToParams(), options...)
}

func NewCreateIGUserForIGOnlyAPISubscribedAppsBatchRequest(id string, params CreateIGUserForIGOnlyAPISubscribedAppsParams, options ...core.BatchOption) *core.BatchRequest[objects.IGGraphUser] {
	return core.NewBatchRequest[objects.IGGraphUser](CreateIGUserForIGOnlyAPISubscribedAppsBatchCall(id, params, options...))
}

func DecodeCreateIGUserForIGOnlyAPISubscribedAppsBatchResponse(response *core.BatchResponse) (*objects.IGGraphUser, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.IGGraphUser
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateIGUserForIGOnlyAPISubscribedAppsWithResponse(ctx context.Context, client *core.Client, id string, params CreateIGUserForIGOnlyAPISubscribedAppsParams) (*objects.IGGraphUser, *core.Response, error) {
	var out objects.IGGraphUser
	call := CreateIGUserForIGOnlyAPISubscribedAppsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateIGUserForIGOnlyAPISubscribedApps(ctx context.Context, client *core.Client, id string, params CreateIGUserForIGOnlyAPISubscribedAppsParams) (*objects.IGGraphUser, error) {
	out, _, err := CreateIGUserForIGOnlyAPISubscribedAppsWithResponse(ctx, client, id, params)
	return out, err
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

func GetIGUserForIGOnlyAPITagsBatchCall(id string, params GetIGUserForIGOnlyAPITagsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "tags"), params.ToParams(), options...)
}

func NewGetIGUserForIGOnlyAPITagsBatchRequest(id string, params GetIGUserForIGOnlyAPITagsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Media]] {
	return core.NewBatchRequest[core.Cursor[objects.Media]](GetIGUserForIGOnlyAPITagsBatchCall(id, params, options...))
}

func DecodeGetIGUserForIGOnlyAPITagsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Media], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Media]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserForIGOnlyAPITagsWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPITagsParams) (*core.Cursor[objects.Media], *core.Response, error) {
	var out core.Cursor[objects.Media]
	call := GetIGUserForIGOnlyAPITagsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserForIGOnlyAPITags(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPITagsParams) (*core.Cursor[objects.Media], error) {
	out, _, err := GetIGUserForIGOnlyAPITagsWithResponse(ctx, client, id, params)
	return out, err
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

func DeleteIGUserForIGOnlyAPIWelcomeMessageFlowsBatchCall(id string, params DeleteIGUserForIGOnlyAPIWelcomeMessageFlowsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "welcome_message_flows"), params.ToParams(), options...)
}

func NewDeleteIGUserForIGOnlyAPIWelcomeMessageFlowsBatchRequest(id string, params DeleteIGUserForIGOnlyAPIWelcomeMessageFlowsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteIGUserForIGOnlyAPIWelcomeMessageFlowsBatchCall(id, params, options...))
}

func DecodeDeleteIGUserForIGOnlyAPIWelcomeMessageFlowsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteIGUserForIGOnlyAPIWelcomeMessageFlowsWithResponse(ctx context.Context, client *core.Client, id string, params DeleteIGUserForIGOnlyAPIWelcomeMessageFlowsParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteIGUserForIGOnlyAPIWelcomeMessageFlowsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteIGUserForIGOnlyAPIWelcomeMessageFlows(ctx context.Context, client *core.Client, id string, params DeleteIGUserForIGOnlyAPIWelcomeMessageFlowsParams) (*map[string]interface{}, error) {
	out, _, err := DeleteIGUserForIGOnlyAPIWelcomeMessageFlowsWithResponse(ctx, client, id, params)
	return out, err
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

func GetIGUserForIGOnlyAPIWelcomeMessageFlowsBatchCall(id string, params GetIGUserForIGOnlyAPIWelcomeMessageFlowsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "welcome_message_flows"), params.ToParams(), options...)
}

func NewGetIGUserForIGOnlyAPIWelcomeMessageFlowsBatchRequest(id string, params GetIGUserForIGOnlyAPIWelcomeMessageFlowsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CTXPartnerAppWelcomeMessageFlow]] {
	return core.NewBatchRequest[core.Cursor[objects.CTXPartnerAppWelcomeMessageFlow]](GetIGUserForIGOnlyAPIWelcomeMessageFlowsBatchCall(id, params, options...))
}

func DecodeGetIGUserForIGOnlyAPIWelcomeMessageFlowsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CTXPartnerAppWelcomeMessageFlow], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CTXPartnerAppWelcomeMessageFlow]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserForIGOnlyAPIWelcomeMessageFlowsWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIWelcomeMessageFlowsParams) (*core.Cursor[objects.CTXPartnerAppWelcomeMessageFlow], *core.Response, error) {
	var out core.Cursor[objects.CTXPartnerAppWelcomeMessageFlow]
	call := GetIGUserForIGOnlyAPIWelcomeMessageFlowsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserForIGOnlyAPIWelcomeMessageFlows(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIWelcomeMessageFlowsParams) (*core.Cursor[objects.CTXPartnerAppWelcomeMessageFlow], error) {
	out, _, err := GetIGUserForIGOnlyAPIWelcomeMessageFlowsWithResponse(ctx, client, id, params)
	return out, err
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

func CreateIGUserForIGOnlyAPIWelcomeMessageFlowsBatchCall(id string, params CreateIGUserForIGOnlyAPIWelcomeMessageFlowsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "welcome_message_flows"), params.ToParams(), options...)
}

func NewCreateIGUserForIGOnlyAPIWelcomeMessageFlowsBatchRequest(id string, params CreateIGUserForIGOnlyAPIWelcomeMessageFlowsParams, options ...core.BatchOption) *core.BatchRequest[objects.IGGraphCTXPartnerAppWelcomeMessageFlow] {
	return core.NewBatchRequest[objects.IGGraphCTXPartnerAppWelcomeMessageFlow](CreateIGUserForIGOnlyAPIWelcomeMessageFlowsBatchCall(id, params, options...))
}

func DecodeCreateIGUserForIGOnlyAPIWelcomeMessageFlowsBatchResponse(response *core.BatchResponse) (*objects.IGGraphCTXPartnerAppWelcomeMessageFlow, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.IGGraphCTXPartnerAppWelcomeMessageFlow
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateIGUserForIGOnlyAPIWelcomeMessageFlowsWithResponse(ctx context.Context, client *core.Client, id string, params CreateIGUserForIGOnlyAPIWelcomeMessageFlowsParams) (*objects.IGGraphCTXPartnerAppWelcomeMessageFlow, *core.Response, error) {
	var out objects.IGGraphCTXPartnerAppWelcomeMessageFlow
	call := CreateIGUserForIGOnlyAPIWelcomeMessageFlowsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateIGUserForIGOnlyAPIWelcomeMessageFlows(ctx context.Context, client *core.Client, id string, params CreateIGUserForIGOnlyAPIWelcomeMessageFlowsParams) (*objects.IGGraphCTXPartnerAppWelcomeMessageFlow, error) {
	out, _, err := CreateIGUserForIGOnlyAPIWelcomeMessageFlowsWithResponse(ctx, client, id, params)
	return out, err
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

func GetIGUserForIGOnlyAPIBatchCall(id string, params GetIGUserForIGOnlyAPIParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetIGUserForIGOnlyAPIBatchRequest(id string, params GetIGUserForIGOnlyAPIParams, options ...core.BatchOption) *core.BatchRequest[objects.User] {
	return core.NewBatchRequest[objects.User](GetIGUserForIGOnlyAPIBatchCall(id, params, options...))
}

func DecodeGetIGUserForIGOnlyAPIBatchResponse(response *core.BatchResponse) (*objects.User, error) {
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

func GetIGUserForIGOnlyAPIWithResponse(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIParams) (*objects.User, *core.Response, error) {
	var out objects.User
	call := GetIGUserForIGOnlyAPIBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetIGUserForIGOnlyAPI(ctx context.Context, client *core.Client, id string, params GetIGUserForIGOnlyAPIParams) (*objects.User, error) {
	out, _, err := GetIGUserForIGOnlyAPIWithResponse(ctx, client, id, params)
	return out, err
}
