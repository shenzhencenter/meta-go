package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
	"time"
)

type GetPagePostAttachmentsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPagePostAttachmentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPagePostAttachments(ctx context.Context, client *core.Client, id string, params GetPagePostAttachmentsParams) (*core.Cursor[objects.StoryAttachment], error) {
	var out core.Cursor[objects.StoryAttachment]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "attachments"), params.ToParams(), &out)
	return &out, err
}

type GetPagePostCommentsParams struct {
	Filter     *enums.PagepostcommentsFilterEnumParam     `facebook:"filter"`
	LiveFilter *enums.PagepostcommentsLiveFilterEnumParam `facebook:"live_filter"`
	Order      *enums.PagepostcommentsOrderEnumParam      `facebook:"order"`
	Since      *time.Time                                 `facebook:"since"`
	Extra      core.Params                                `facebook:"-"`
}

func (params GetPagePostCommentsParams) ToParams() core.Params {
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

func GetPagePostComments(ctx context.Context, client *core.Client, id string, params GetPagePostCommentsParams) (*core.Cursor[objects.Comment], error) {
	var out core.Cursor[objects.Comment]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "comments"), params.ToParams(), &out)
	return &out, err
}

type CreatePagePostCommentsParams struct {
	AttachmentID        *core.ID                                            `facebook:"attachment_id"`
	AttachmentShareURL  *string                                             `facebook:"attachment_share_url"`
	AttachmentURL       *string                                             `facebook:"attachment_url"`
	Comment             *string                                             `facebook:"comment"`
	CommentPrivacyValue *enums.PagepostcommentsCommentPrivacyValueEnumParam `facebook:"comment_privacy_value"`
	FeedbackSource      *string                                             `facebook:"feedback_source"`
	Message             *string                                             `facebook:"message"`
	NectarModule        *string                                             `facebook:"nectar_module"`
	ParentCommentID     *map[string]interface{}                             `facebook:"parent_comment_id"`
	PostID              *core.ID                                            `facebook:"post_id"`
	Tracking            *string                                             `facebook:"tracking"`
	Extra               core.Params                                         `facebook:"-"`
}

func (params CreatePagePostCommentsParams) ToParams() core.Params {
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
	if params.Comment != nil {
		out["comment"] = *params.Comment
	}
	if params.CommentPrivacyValue != nil {
		out["comment_privacy_value"] = *params.CommentPrivacyValue
	}
	if params.FeedbackSource != nil {
		out["feedback_source"] = *params.FeedbackSource
	}
	if params.Message != nil {
		out["message"] = *params.Message
	}
	if params.NectarModule != nil {
		out["nectar_module"] = *params.NectarModule
	}
	if params.ParentCommentID != nil {
		out["parent_comment_id"] = *params.ParentCommentID
	}
	if params.PostID != nil {
		out["post_id"] = *params.PostID
	}
	if params.Tracking != nil {
		out["tracking"] = *params.Tracking
	}
	return out
}

func CreatePagePostComments(ctx context.Context, client *core.Client, id string, params CreatePagePostCommentsParams) (*objects.Comment, error) {
	var out objects.Comment
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "comments"), params.ToParams(), &out)
	return &out, err
}

type GetPagePostDynamicPostsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPagePostDynamicPostsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPagePostDynamicPosts(ctx context.Context, client *core.Client, id string, params GetPagePostDynamicPostsParams) (*core.Cursor[objects.RTBDynamicPost], error) {
	var out core.Cursor[objects.RTBDynamicPost]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "dynamic_posts"), params.ToParams(), &out)
	return &out, err
}

type GetPagePostInsightsParams struct {
	DatePreset *enums.PagepostinsightsDatePresetEnumParam `facebook:"date_preset"`
	Metric     *[]map[string]interface{}                  `facebook:"metric"`
	Period     *enums.PagepostinsightsPeriodEnumParam     `facebook:"period"`
	Since      *time.Time                                 `facebook:"since"`
	Until      *time.Time                                 `facebook:"until"`
	Extra      core.Params                                `facebook:"-"`
}

func (params GetPagePostInsightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.DatePreset != nil {
		out["date_preset"] = *params.DatePreset
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

func GetPagePostInsights(ctx context.Context, client *core.Client, id string, params GetPagePostInsightsParams) (*core.Cursor[objects.InsightsResult], error) {
	var out core.Cursor[objects.InsightsResult]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "insights"), params.ToParams(), &out)
	return &out, err
}

type DeletePagePostLikesParams struct {
	AttributionIDV2 *string     `facebook:"attribution_id_v2"`
	NectarModule    *string     `facebook:"nectar_module"`
	Tracking        *string     `facebook:"tracking"`
	Extra           core.Params `facebook:"-"`
}

func (params DeletePagePostLikesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AttributionIDV2 != nil {
		out["attribution_id_v2"] = *params.AttributionIDV2
	}
	if params.NectarModule != nil {
		out["nectar_module"] = *params.NectarModule
	}
	if params.Tracking != nil {
		out["tracking"] = *params.Tracking
	}
	return out
}

func DeletePagePostLikes(ctx context.Context, client *core.Client, id string, params DeletePagePostLikesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "likes"), params.ToParams(), &out)
	return &out, err
}

type GetPagePostLikesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPagePostLikesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPagePostLikes(ctx context.Context, client *core.Client, id string, params GetPagePostLikesParams) (*core.Cursor[objects.Profile], error) {
	var out core.Cursor[objects.Profile]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "likes"), params.ToParams(), &out)
	return &out, err
}

type CreatePagePostLikesParams struct {
	AttributionIDV2 *string     `facebook:"attribution_id_v2"`
	FeedbackSource  *string     `facebook:"feedback_source"`
	NectarModule    *string     `facebook:"nectar_module"`
	Tracking        *string     `facebook:"tracking"`
	Extra           core.Params `facebook:"-"`
}

func (params CreatePagePostLikesParams) ToParams() core.Params {
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
	if params.Tracking != nil {
		out["tracking"] = *params.Tracking
	}
	return out
}

func CreatePagePostLikes(ctx context.Context, client *core.Client, id string, params CreatePagePostLikesParams) (*objects.PagePost, error) {
	var out objects.PagePost
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "likes"), params.ToParams(), &out)
	return &out, err
}

type GetPagePostReactionsParams struct {
	Type  *enums.PagepostreactionsTypeEnumParam `facebook:"type"`
	Extra core.Params                           `facebook:"-"`
}

func (params GetPagePostReactionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	return out
}

func GetPagePostReactions(ctx context.Context, client *core.Client, id string, params GetPagePostReactionsParams) (*core.Cursor[objects.Profile], error) {
	var out core.Cursor[objects.Profile]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "reactions"), params.ToParams(), &out)
	return &out, err
}

type GetPagePostSharedpostsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPagePostSharedpostsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPagePostSharedposts(ctx context.Context, client *core.Client, id string, params GetPagePostSharedpostsParams) (*core.Cursor[objects.Post], error) {
	var out core.Cursor[objects.Post]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "sharedposts"), params.ToParams(), &out)
	return &out, err
}

type GetPagePostSponsorTagsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPagePostSponsorTagsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPagePostSponsorTags(ctx context.Context, client *core.Client, id string, params GetPagePostSponsorTagsParams) (*core.Cursor[objects.Page], error) {
	var out core.Cursor[objects.Page]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "sponsor_tags"), params.ToParams(), &out)
	return &out, err
}

type GetPagePostToParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPagePostToParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPagePostTo(ctx context.Context, client *core.Client, id string, params GetPagePostToParams) (*core.Cursor[objects.Profile], error) {
	var out core.Cursor[objects.Profile]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "to"), params.ToParams(), &out)
	return &out, err
}

type DeletePagePostParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeletePagePostParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeletePagePost(ctx context.Context, client *core.Client, id string, params DeletePagePostParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetPagePostParams struct {
	PrimaryFbPageID *core.ID    `facebook:"primary_fb_page_id"`
	PrimaryIgUserID *core.ID    `facebook:"primary_ig_user_id"`
	Extra           core.Params `facebook:"-"`
}

func (params GetPagePostParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.PrimaryFbPageID != nil {
		out["primary_fb_page_id"] = *params.PrimaryFbPageID
	}
	if params.PrimaryIgUserID != nil {
		out["primary_ig_user_id"] = *params.PrimaryIgUserID
	}
	return out
}

func GetPagePost(ctx context.Context, client *core.Client, id string, params GetPagePostParams) (*objects.PagePost, error) {
	var out objects.PagePost
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdatePagePostParams struct {
	AttachedMedia               *[]map[string]interface{}               `facebook:"attached_media"`
	BackdatedTime               *time.Time                              `facebook:"backdated_time"`
	BackdatedTimeGranularity    *enums.PagepostBackdatedTimeGranularity `facebook:"backdated_time_granularity"`
	ComposerSessionID           *core.ID                                `facebook:"composer_session_id"`
	DirectShareStatus           *uint64                                 `facebook:"direct_share_status"`
	ExplicitlyAddedMentioneeIds *[]core.ID                              `facebook:"explicitly_added_mentionee_ids"`
	FeedStoryVisibility         *enums.PagepostFeedStoryVisibility      `facebook:"feed_story_visibility"`
	IsExplicitLocation          *bool                                   `facebook:"is_explicit_location"`
	IsHidden                    *bool                                   `facebook:"is_hidden"`
	IsPinned                    *bool                                   `facebook:"is_pinned"`
	IsPublished                 *bool                                   `facebook:"is_published"`
	Message                     *string                                 `facebook:"message"`
	OgActionTypeID              *core.ID                                `facebook:"og_action_type_id"`
	OgHideObjectAttachment      *bool                                   `facebook:"og_hide_object_attachment"`
	OgIconID                    *core.ID                                `facebook:"og_icon_id"`
	OgObjectID                  *core.ID                                `facebook:"og_object_id"`
	OgPhrase                    *string                                 `facebook:"og_phrase"`
	OgSetProfileBadge           *bool                                   `facebook:"og_set_profile_badge"`
	OgSuggestionMechanism       *string                                 `facebook:"og_suggestion_mechanism"`
	Place                       *map[string]interface{}                 `facebook:"place"`
	Privacy                     *string                                 `facebook:"privacy"`
	ProductItem                 *map[string]interface{}                 `facebook:"product_item"`
	ScheduledPublishTime        *uint64                                 `facebook:"scheduled_publish_time"`
	ShouldSyncProductEdit       *bool                                   `facebook:"should_sync_product_edit"`
	SourceType                  *string                                 `facebook:"source_type"`
	SponsorID                   *core.ID                                `facebook:"sponsor_id"`
	SponsorRelationship         *uint64                                 `facebook:"sponsor_relationship"`
	Tags                        *[]int                                  `facebook:"tags"`
	TextFormatPresetID          *core.ID                                `facebook:"text_format_preset_id"`
	TimelineVisibility          *enums.PagepostTimelineVisibility       `facebook:"timeline_visibility"`
	Tracking                    *string                                 `facebook:"tracking"`
	Extra                       core.Params                             `facebook:"-"`
}

func (params UpdatePagePostParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AttachedMedia != nil {
		out["attached_media"] = *params.AttachedMedia
	}
	if params.BackdatedTime != nil {
		out["backdated_time"] = *params.BackdatedTime
	}
	if params.BackdatedTimeGranularity != nil {
		out["backdated_time_granularity"] = *params.BackdatedTimeGranularity
	}
	if params.ComposerSessionID != nil {
		out["composer_session_id"] = *params.ComposerSessionID
	}
	if params.DirectShareStatus != nil {
		out["direct_share_status"] = *params.DirectShareStatus
	}
	if params.ExplicitlyAddedMentioneeIds != nil {
		out["explicitly_added_mentionee_ids"] = *params.ExplicitlyAddedMentioneeIds
	}
	if params.FeedStoryVisibility != nil {
		out["feed_story_visibility"] = *params.FeedStoryVisibility
	}
	if params.IsExplicitLocation != nil {
		out["is_explicit_location"] = *params.IsExplicitLocation
	}
	if params.IsHidden != nil {
		out["is_hidden"] = *params.IsHidden
	}
	if params.IsPinned != nil {
		out["is_pinned"] = *params.IsPinned
	}
	if params.IsPublished != nil {
		out["is_published"] = *params.IsPublished
	}
	if params.Message != nil {
		out["message"] = *params.Message
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
	if params.Place != nil {
		out["place"] = *params.Place
	}
	if params.Privacy != nil {
		out["privacy"] = *params.Privacy
	}
	if params.ProductItem != nil {
		out["product_item"] = *params.ProductItem
	}
	if params.ScheduledPublishTime != nil {
		out["scheduled_publish_time"] = *params.ScheduledPublishTime
	}
	if params.ShouldSyncProductEdit != nil {
		out["should_sync_product_edit"] = *params.ShouldSyncProductEdit
	}
	if params.SourceType != nil {
		out["source_type"] = *params.SourceType
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
	if params.TextFormatPresetID != nil {
		out["text_format_preset_id"] = *params.TextFormatPresetID
	}
	if params.TimelineVisibility != nil {
		out["timeline_visibility"] = *params.TimelineVisibility
	}
	if params.Tracking != nil {
		out["tracking"] = *params.Tracking
	}
	return out
}

func UpdatePagePost(ctx context.Context, client *core.Client, id string, params UpdatePagePostParams) (*objects.PagePost, error) {
	var out objects.PagePost
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
