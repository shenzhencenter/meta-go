package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
	"time"
)

type GetAdVideoBoostAdsListParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdVideoBoostAdsListParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdVideoBoostAdsList(ctx context.Context, client *core.Client, id string, params GetAdVideoBoostAdsListParams) (*core.Cursor[objects.VideoBoostMediaAd], error) {
	var out core.Cursor[objects.VideoBoostMediaAd]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "boost_ads_list"), params.ToParams(), &out)
	return &out, err
}

type GetAdVideoCaptionsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdVideoCaptionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdVideoCaptions(ctx context.Context, client *core.Client, id string, params GetAdVideoCaptionsParams) (*core.Cursor[objects.VideoCaption], error) {
	var out core.Cursor[objects.VideoCaption]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "captions"), params.ToParams(), &out)
	return &out, err
}

type CreateAdVideoCaptionsParams struct {
	CaptionsFile    *core.FileParam `facebook:"captions_file"`
	DefaultLocale   *string         `facebook:"default_locale"`
	LocalesToDelete *[]string       `facebook:"locales_to_delete"`
	Extra           core.Params     `facebook:"-"`
}

func (params CreateAdVideoCaptionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CaptionsFile != nil {
		out["captions_file"] = *params.CaptionsFile
	}
	if params.DefaultLocale != nil {
		out["default_locale"] = *params.DefaultLocale
	}
	if params.LocalesToDelete != nil {
		out["locales_to_delete"] = *params.LocalesToDelete
	}
	return out
}

func CreateAdVideoCaptions(ctx context.Context, client *core.Client, id string, params CreateAdVideoCaptionsParams) (*objects.AdVideo, error) {
	var out objects.AdVideo
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "captions"), params.ToParams(), &out)
	return &out, err
}

type GetAdVideoCollaboratorsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdVideoCollaboratorsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdVideoCollaborators(ctx context.Context, client *core.Client, id string, params GetAdVideoCollaboratorsParams) (*core.Cursor[objects.VideoCollaborators], error) {
	var out core.Cursor[objects.VideoCollaborators]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "collaborators"), params.ToParams(), &out)
	return &out, err
}

type CreateAdVideoCollaboratorsParams struct {
	TargetID core.ID     `facebook:"target_id"`
	Extra    core.Params `facebook:"-"`
}

func (params CreateAdVideoCollaboratorsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["target_id"] = params.TargetID
	return out
}

func CreateAdVideoCollaborators(ctx context.Context, client *core.Client, id string, params CreateAdVideoCollaboratorsParams) (*objects.AdVideo, error) {
	var out objects.AdVideo
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "collaborators"), params.ToParams(), &out)
	return &out, err
}

type GetAdVideoCommentsParams struct {
	Filter     *enums.VideocommentsFilterEnumParam     `facebook:"filter"`
	LiveFilter *enums.VideocommentsLiveFilterEnumParam `facebook:"live_filter"`
	Order      *enums.VideocommentsOrderEnumParam      `facebook:"order"`
	Since      *time.Time                              `facebook:"since"`
	Extra      core.Params                             `facebook:"-"`
}

func (params GetAdVideoCommentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Filter != nil {
		out["filter"] = *params.Filter
	}
	if params.LiveFilter != nil {
		out["live_filter"] = *params.LiveFilter
	}
	if params.Order != nil {
		out["order"] = *params.Order
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	return out
}

func GetAdVideoComments(ctx context.Context, client *core.Client, id string, params GetAdVideoCommentsParams) (*core.Cursor[objects.Comment], error) {
	var out core.Cursor[objects.Comment]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "comments"), params.ToParams(), &out)
	return &out, err
}

type CreateAdVideoCommentsParams struct {
	AttachmentID         *core.ID                                         `facebook:"attachment_id"`
	AttachmentShareURL   *string                                          `facebook:"attachment_share_url"`
	AttachmentURL        *string                                          `facebook:"attachment_url"`
	CommentPrivacyValue  *enums.VideocommentsCommentPrivacyValueEnumParam `facebook:"comment_privacy_value"`
	FacepileMentionedIds *[]core.ID                                       `facebook:"facepile_mentioned_ids"`
	FeedbackSource       *string                                          `facebook:"feedback_source"`
	IsOffline            *bool                                            `facebook:"is_offline"`
	Message              *string                                          `facebook:"message"`
	NectarModule         *string                                          `facebook:"nectar_module"`
	ObjectID             *core.ID                                         `facebook:"object_id"`
	ParentCommentID      *map[string]interface{}                          `facebook:"parent_comment_id"`
	Text                 *string                                          `facebook:"text"`
	Tracking             *string                                          `facebook:"tracking"`
	Extra                core.Params                                      `facebook:"-"`
}

func (params CreateAdVideoCommentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AttachmentID != nil {
		out["attachment_id"] = *params.AttachmentID
	}
	if params.AttachmentShareURL != nil {
		out["attachment_share_url"] = *params.AttachmentShareURL
	}
	if params.AttachmentURL != nil {
		out["attachment_url"] = *params.AttachmentURL
	}
	if params.CommentPrivacyValue != nil {
		out["comment_privacy_value"] = *params.CommentPrivacyValue
	}
	if params.FacepileMentionedIds != nil {
		out["facepile_mentioned_ids"] = *params.FacepileMentionedIds
	}
	if params.FeedbackSource != nil {
		out["feedback_source"] = *params.FeedbackSource
	}
	if params.IsOffline != nil {
		out["is_offline"] = *params.IsOffline
	}
	if params.Message != nil {
		out["message"] = *params.Message
	}
	if params.NectarModule != nil {
		out["nectar_module"] = *params.NectarModule
	}
	if params.ObjectID != nil {
		out["object_id"] = *params.ObjectID
	}
	if params.ParentCommentID != nil {
		out["parent_comment_id"] = *params.ParentCommentID
	}
	if params.Text != nil {
		out["text"] = *params.Text
	}
	if params.Tracking != nil {
		out["tracking"] = *params.Tracking
	}
	return out
}

func CreateAdVideoComments(ctx context.Context, client *core.Client, id string, params CreateAdVideoCommentsParams) (*objects.Comment, error) {
	var out objects.Comment
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "comments"), params.ToParams(), &out)
	return &out, err
}

type GetAdVideoCrosspostSharedPagesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdVideoCrosspostSharedPagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdVideoCrosspostSharedPages(ctx context.Context, client *core.Client, id string, params GetAdVideoCrosspostSharedPagesParams) (*core.Cursor[objects.Page], error) {
	var out core.Cursor[objects.Page]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "crosspost_shared_pages"), params.ToParams(), &out)
	return &out, err
}

type GetAdVideoLikesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdVideoLikesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdVideoLikes(ctx context.Context, client *core.Client, id string, params GetAdVideoLikesParams) (*core.Cursor[objects.Profile], error) {
	var out core.Cursor[objects.Profile]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "likes"), params.ToParams(), &out)
	return &out, err
}

type CreateAdVideoLikesParams struct {
	AttributionIDV2 *string     `facebook:"attribution_id_v2"`
	FeedbackSource  *string     `facebook:"feedback_source"`
	NectarModule    *string     `facebook:"nectar_module"`
	Notify          *bool       `facebook:"notify"`
	Tracking        *string     `facebook:"tracking"`
	Extra           core.Params `facebook:"-"`
}

func (params CreateAdVideoLikesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AttributionIDV2 != nil {
		out["attribution_id_v2"] = *params.AttributionIDV2
	}
	if params.FeedbackSource != nil {
		out["feedback_source"] = *params.FeedbackSource
	}
	if params.NectarModule != nil {
		out["nectar_module"] = *params.NectarModule
	}
	if params.Notify != nil {
		out["notify"] = *params.Notify
	}
	if params.Tracking != nil {
		out["tracking"] = *params.Tracking
	}
	return out
}

func CreateAdVideoLikes(ctx context.Context, client *core.Client, id string, params CreateAdVideoLikesParams) (*objects.AdVideo, error) {
	var out objects.AdVideo
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "likes"), params.ToParams(), &out)
	return &out, err
}

type GetAdVideoPollSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdVideoPollSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdVideoPollSettings(ctx context.Context, client *core.Client, id string, params GetAdVideoPollSettingsParams) (*core.Cursor[objects.PollSettings], error) {
	var out core.Cursor[objects.PollSettings]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "poll_settings"), params.ToParams(), &out)
	return &out, err
}

type GetAdVideoPollsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdVideoPollsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdVideoPolls(ctx context.Context, client *core.Client, id string, params GetAdVideoPollsParams) (*core.Cursor[objects.VideoPoll], error) {
	var out core.Cursor[objects.VideoPoll]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "polls"), params.ToParams(), &out)
	return &out, err
}

type CreateAdVideoPollsParams struct {
	CloseAfterVoting *bool       `facebook:"close_after_voting"`
	CorrectOption    *uint64     `facebook:"correct_option"`
	DefaultOpen      *bool       `facebook:"default_open"`
	Options          []string    `facebook:"options"`
	Question         string      `facebook:"question"`
	ShowGradient     *bool       `facebook:"show_gradient"`
	ShowResults      *bool       `facebook:"show_results"`
	Extra            core.Params `facebook:"-"`
}

func (params CreateAdVideoPollsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CloseAfterVoting != nil {
		out["close_after_voting"] = *params.CloseAfterVoting
	}
	if params.CorrectOption != nil {
		out["correct_option"] = *params.CorrectOption
	}
	if params.DefaultOpen != nil {
		out["default_open"] = *params.DefaultOpen
	}
	out["options"] = params.Options
	out["question"] = params.Question
	if params.ShowGradient != nil {
		out["show_gradient"] = *params.ShowGradient
	}
	if params.ShowResults != nil {
		out["show_results"] = *params.ShowResults
	}
	return out
}

func CreateAdVideoPolls(ctx context.Context, client *core.Client, id string, params CreateAdVideoPollsParams) (*objects.VideoPoll, error) {
	var out objects.VideoPoll
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "polls"), params.ToParams(), &out)
	return &out, err
}

type GetAdVideoSponsorTagsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdVideoSponsorTagsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdVideoSponsorTags(ctx context.Context, client *core.Client, id string, params GetAdVideoSponsorTagsParams) (*core.Cursor[objects.Page], error) {
	var out core.Cursor[objects.Page]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "sponsor_tags"), params.ToParams(), &out)
	return &out, err
}

type GetAdVideoTagsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdVideoTagsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdVideoTags(ctx context.Context, client *core.Client, id string, params GetAdVideoTagsParams) (*core.Cursor[objects.TaggableSubject], error) {
	var out core.Cursor[objects.TaggableSubject]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "tags"), params.ToParams(), &out)
	return &out, err
}

type GetAdVideoThumbnailsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdVideoThumbnailsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdVideoThumbnails(ctx context.Context, client *core.Client, id string, params GetAdVideoThumbnailsParams) (*core.Cursor[objects.VideoThumbnail], error) {
	var out core.Cursor[objects.VideoThumbnail]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "thumbnails"), params.ToParams(), &out)
	return &out, err
}

type CreateAdVideoThumbnailsParams struct {
	IsPreferred *bool          `facebook:"is_preferred"`
	Source      core.FileParam `facebook:"source"`
	Extra       core.Params    `facebook:"-"`
}

func (params CreateAdVideoThumbnailsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.IsPreferred != nil {
		out["is_preferred"] = *params.IsPreferred
	}
	out["source"] = params.Source
	return out
}

func CreateAdVideoThumbnails(ctx context.Context, client *core.Client, id string, params CreateAdVideoThumbnailsParams) (*objects.AdVideo, error) {
	var out objects.AdVideo
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "thumbnails"), params.ToParams(), &out)
	return &out, err
}

type GetAdVideoVideoInsightsParams struct {
	Metric *[]map[string]interface{}                `facebook:"metric"`
	Period *enums.VideovideoInsightsPeriodEnumParam `facebook:"period"`
	Since  *time.Time                               `facebook:"since"`
	Until  *time.Time                               `facebook:"until"`
	Extra  core.Params                              `facebook:"-"`
}

func (params GetAdVideoVideoInsightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Metric != nil {
		out["metric"] = *params.Metric
	}
	if params.Period != nil {
		out["period"] = *params.Period
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	return out
}

func GetAdVideoVideoInsights(ctx context.Context, client *core.Client, id string, params GetAdVideoVideoInsightsParams) (*core.Cursor[objects.InsightsResult], error) {
	var out core.Cursor[objects.InsightsResult]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "video_insights"), params.ToParams(), &out)
	return &out, err
}

type DeleteAdVideoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteAdVideoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteAdVideo(ctx context.Context, client *core.Client, id string, params DeleteAdVideoParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetAdVideoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdVideoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdVideo(ctx context.Context, client *core.Client, id string, params GetAdVideoParams) (*objects.AdVideo, error) {
	var out objects.AdVideo
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateAdVideoParams struct {
	AdBreaks                  *[]interface{}                       `facebook:"ad_breaks"`
	AllowBmCrossposting       *bool                                `facebook:"allow_bm_crossposting"`
	AllowCrosspostingForPages *[]map[string]interface{}            `facebook:"allow_crossposting_for_pages"`
	BackdatedTime             *time.Time                           `facebook:"backdated_time"`
	BackdatedTimeGranularity  *enums.VideoBackdatedTimeGranularity `facebook:"backdated_time_granularity"`
	CallToAction              *map[string]interface{}              `facebook:"call_to_action"`
	ContentCategory           *enums.VideoContentCategory          `facebook:"content_category"`
	ContentTags               *[]string                            `facebook:"content_tags"`
	CustomLabels              *[]string                            `facebook:"custom_labels"`
	Description               *string                              `facebook:"description"`
	DirectShareStatus         *uint64                              `facebook:"direct_share_status"`
	Embeddable                *bool                                `facebook:"embeddable"`
	Expiration                *map[string]interface{}              `facebook:"expiration"`
	ExpireNow                 *bool                                `facebook:"expire_now"`
	IncrementPlayCount        *bool                                `facebook:"increment_play_count"`
	Name                      *string                              `facebook:"name"`
	PreferredThumbnailID      *core.ID                             `facebook:"preferred_thumbnail_id"`
	Privacy                   *string                              `facebook:"privacy"`
	PublishToNewsFeed         *bool                                `facebook:"publish_to_news_feed"`
	PublishToVideosTab        *bool                                `facebook:"publish_to_videos_tab"`
	Published                 *bool                                `facebook:"published"`
	ScheduledPublishTime      *uint64                              `facebook:"scheduled_publish_time"`
	SocialActions             *bool                                `facebook:"social_actions"`
	SponsorID                 *core.ID                             `facebook:"sponsor_id"`
	SponsorRelationship       *uint64                              `facebook:"sponsor_relationship"`
	Tags                      *[]string                            `facebook:"tags"`
	Target                    *string                              `facebook:"target"`
	UniversalVideoID          *core.ID                             `facebook:"universal_video_id"`
	Extra                     core.Params                          `facebook:"-"`
}

func (params UpdateAdVideoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdBreaks != nil {
		out["ad_breaks"] = *params.AdBreaks
	}
	if params.AllowBmCrossposting != nil {
		out["allow_bm_crossposting"] = *params.AllowBmCrossposting
	}
	if params.AllowCrosspostingForPages != nil {
		out["allow_crossposting_for_pages"] = *params.AllowCrosspostingForPages
	}
	if params.BackdatedTime != nil {
		out["backdated_time"] = *params.BackdatedTime
	}
	if params.BackdatedTimeGranularity != nil {
		out["backdated_time_granularity"] = *params.BackdatedTimeGranularity
	}
	if params.CallToAction != nil {
		out["call_to_action"] = *params.CallToAction
	}
	if params.ContentCategory != nil {
		out["content_category"] = *params.ContentCategory
	}
	if params.ContentTags != nil {
		out["content_tags"] = *params.ContentTags
	}
	if params.CustomLabels != nil {
		out["custom_labels"] = *params.CustomLabels
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.DirectShareStatus != nil {
		out["direct_share_status"] = *params.DirectShareStatus
	}
	if params.Embeddable != nil {
		out["embeddable"] = *params.Embeddable
	}
	if params.Expiration != nil {
		out["expiration"] = *params.Expiration
	}
	if params.ExpireNow != nil {
		out["expire_now"] = *params.ExpireNow
	}
	if params.IncrementPlayCount != nil {
		out["increment_play_count"] = *params.IncrementPlayCount
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.PreferredThumbnailID != nil {
		out["preferred_thumbnail_id"] = *params.PreferredThumbnailID
	}
	if params.Privacy != nil {
		out["privacy"] = *params.Privacy
	}
	if params.PublishToNewsFeed != nil {
		out["publish_to_news_feed"] = *params.PublishToNewsFeed
	}
	if params.PublishToVideosTab != nil {
		out["publish_to_videos_tab"] = *params.PublishToVideosTab
	}
	if params.Published != nil {
		out["published"] = *params.Published
	}
	if params.ScheduledPublishTime != nil {
		out["scheduled_publish_time"] = *params.ScheduledPublishTime
	}
	if params.SocialActions != nil {
		out["social_actions"] = *params.SocialActions
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
	if params.Target != nil {
		out["target"] = *params.Target
	}
	if params.UniversalVideoID != nil {
		out["universal_video_id"] = *params.UniversalVideoID
	}
	return out
}

func UpdateAdVideo(ctx context.Context, client *core.Client, id string, params UpdateAdVideoParams) (*objects.AdVideo, error) {
	var out objects.AdVideo
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
