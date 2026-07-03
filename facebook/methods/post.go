package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
	"time"
)

type GetPostAttachmentsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPostAttachmentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPostAttachmentsBatchCall(id string, params GetPostAttachmentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "attachments"), params.ToParams(), options...)
}

func NewGetPostAttachmentsBatchRequest(id string, params GetPostAttachmentsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.StoryAttachment]] {
	return core.NewBatchRequest[core.Cursor[objects.StoryAttachment]](GetPostAttachmentsBatchCall(id, params, options...))
}

func DecodeGetPostAttachmentsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.StoryAttachment], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.StoryAttachment]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPostAttachments(ctx context.Context, client *core.Client, id string, params GetPostAttachmentsParams) (*core.Cursor[objects.StoryAttachment], error) {
	var out core.Cursor[objects.StoryAttachment]
	call := GetPostAttachmentsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetPostCommentsParams struct {
	Filter     *enums.PostcommentsFilterEnumParam     `facebook:"filter"`
	LiveFilter *enums.PostcommentsLiveFilterEnumParam `facebook:"live_filter"`
	Order      *enums.PostcommentsOrderEnumParam      `facebook:"order"`
	Since      *time.Time                             `facebook:"since"`
	Extra      core.Params                            `facebook:"-"`
}

func (params GetPostCommentsParams) ToParams() core.Params {
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

func GetPostCommentsBatchCall(id string, params GetPostCommentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "comments"), params.ToParams(), options...)
}

func NewGetPostCommentsBatchRequest(id string, params GetPostCommentsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Comment]] {
	return core.NewBatchRequest[core.Cursor[objects.Comment]](GetPostCommentsBatchCall(id, params, options...))
}

func DecodeGetPostCommentsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Comment], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Comment]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPostComments(ctx context.Context, client *core.Client, id string, params GetPostCommentsParams) (*core.Cursor[objects.Comment], error) {
	var out core.Cursor[objects.Comment]
	call := GetPostCommentsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreatePostCommentsParams struct {
	AttachmentID        *core.ID                                        `facebook:"attachment_id"`
	AttachmentShareURL  *string                                         `facebook:"attachment_share_url"`
	AttachmentURL       *string                                         `facebook:"attachment_url"`
	Comment             *string                                         `facebook:"comment"`
	CommentPrivacyValue *enums.PostcommentsCommentPrivacyValueEnumParam `facebook:"comment_privacy_value"`
	FeedbackSource      *string                                         `facebook:"feedback_source"`
	Message             *string                                         `facebook:"message"`
	NectarModule        *string                                         `facebook:"nectar_module"`
	ParentCommentID     *map[string]interface{}                         `facebook:"parent_comment_id"`
	PostID              *core.ID                                        `facebook:"post_id"`
	Tracking            *string                                         `facebook:"tracking"`
	Extra               core.Params                                     `facebook:"-"`
}

func (params CreatePostCommentsParams) ToParams() core.Params {
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

func CreatePostCommentsBatchCall(id string, params CreatePostCommentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "comments"), params.ToParams(), options...)
}

func NewCreatePostCommentsBatchRequest(id string, params CreatePostCommentsParams, options ...core.BatchOption) *core.BatchRequest[objects.Comment] {
	return core.NewBatchRequest[objects.Comment](CreatePostCommentsBatchCall(id, params, options...))
}

func DecodeCreatePostCommentsBatchResponse(response *core.BatchResponse) (*objects.Comment, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Comment
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePostComments(ctx context.Context, client *core.Client, id string, params CreatePostCommentsParams) (*objects.Comment, error) {
	var out objects.Comment
	call := CreatePostCommentsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetPostDynamicPostsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPostDynamicPostsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPostDynamicPostsBatchCall(id string, params GetPostDynamicPostsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "dynamic_posts"), params.ToParams(), options...)
}

func NewGetPostDynamicPostsBatchRequest(id string, params GetPostDynamicPostsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.RTBDynamicPost]] {
	return core.NewBatchRequest[core.Cursor[objects.RTBDynamicPost]](GetPostDynamicPostsBatchCall(id, params, options...))
}

func DecodeGetPostDynamicPostsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.RTBDynamicPost], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.RTBDynamicPost]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPostDynamicPosts(ctx context.Context, client *core.Client, id string, params GetPostDynamicPostsParams) (*core.Cursor[objects.RTBDynamicPost], error) {
	var out core.Cursor[objects.RTBDynamicPost]
	call := GetPostDynamicPostsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetPostInsightsParams struct {
	DatePreset *enums.PostinsightsDatePresetEnumParam `facebook:"date_preset"`
	Metric     *[]map[string]interface{}              `facebook:"metric"`
	Period     *enums.PostinsightsPeriodEnumParam     `facebook:"period"`
	Since      *time.Time                             `facebook:"since"`
	Until      *time.Time                             `facebook:"until"`
	Extra      core.Params                            `facebook:"-"`
}

func (params GetPostInsightsParams) ToParams() core.Params {
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

func GetPostInsightsBatchCall(id string, params GetPostInsightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "insights"), params.ToParams(), options...)
}

func NewGetPostInsightsBatchRequest(id string, params GetPostInsightsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.InsightsResult]] {
	return core.NewBatchRequest[core.Cursor[objects.InsightsResult]](GetPostInsightsBatchCall(id, params, options...))
}

func DecodeGetPostInsightsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.InsightsResult], error) {
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

func GetPostInsights(ctx context.Context, client *core.Client, id string, params GetPostInsightsParams) (*core.Cursor[objects.InsightsResult], error) {
	var out core.Cursor[objects.InsightsResult]
	call := GetPostInsightsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeletePostLikesParams struct {
	AttributionIDV2 *string     `facebook:"attribution_id_v2"`
	NectarModule    *string     `facebook:"nectar_module"`
	Tracking        *string     `facebook:"tracking"`
	Extra           core.Params `facebook:"-"`
}

func (params DeletePostLikesParams) ToParams() core.Params {
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

func DeletePostLikesBatchCall(id string, params DeletePostLikesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "likes"), params.ToParams(), options...)
}

func NewDeletePostLikesBatchRequest(id string, params DeletePostLikesParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeletePostLikesBatchCall(id, params, options...))
}

func DecodeDeletePostLikesBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeletePostLikes(ctx context.Context, client *core.Client, id string, params DeletePostLikesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeletePostLikesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type CreatePostLikesParams struct {
	AttributionIDV2 *string     `facebook:"attribution_id_v2"`
	FeedbackSource  *string     `facebook:"feedback_source"`
	NectarModule    *string     `facebook:"nectar_module"`
	Tracking        *string     `facebook:"tracking"`
	Extra           core.Params `facebook:"-"`
}

func (params CreatePostLikesParams) ToParams() core.Params {
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

func CreatePostLikesBatchCall(id string, params CreatePostLikesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "likes"), params.ToParams(), options...)
}

func NewCreatePostLikesBatchRequest(id string, params CreatePostLikesParams, options ...core.BatchOption) *core.BatchRequest[objects.Post] {
	return core.NewBatchRequest[objects.Post](CreatePostLikesBatchCall(id, params, options...))
}

func DecodeCreatePostLikesBatchResponse(response *core.BatchResponse) (*objects.Post, error) {
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

func CreatePostLikes(ctx context.Context, client *core.Client, id string, params CreatePostLikesParams) (*objects.Post, error) {
	var out objects.Post
	call := CreatePostLikesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetPostReactionsParams struct {
	Type  *enums.PostreactionsTypeEnumParam `facebook:"type"`
	Extra core.Params                       `facebook:"-"`
}

func (params GetPostReactionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	return out
}

func GetPostReactionsBatchCall(id string, params GetPostReactionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "reactions"), params.ToParams(), options...)
}

func NewGetPostReactionsBatchRequest(id string, params GetPostReactionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Profile]] {
	return core.NewBatchRequest[core.Cursor[objects.Profile]](GetPostReactionsBatchCall(id, params, options...))
}

func DecodeGetPostReactionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Profile], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Profile]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPostReactions(ctx context.Context, client *core.Client, id string, params GetPostReactionsParams) (*core.Cursor[objects.Profile], error) {
	var out core.Cursor[objects.Profile]
	call := GetPostReactionsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetPostSharedpostsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPostSharedpostsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPostSharedpostsBatchCall(id string, params GetPostSharedpostsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "sharedposts"), params.ToParams(), options...)
}

func NewGetPostSharedpostsBatchRequest(id string, params GetPostSharedpostsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Post]] {
	return core.NewBatchRequest[core.Cursor[objects.Post]](GetPostSharedpostsBatchCall(id, params, options...))
}

func DecodeGetPostSharedpostsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Post], error) {
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

func GetPostSharedposts(ctx context.Context, client *core.Client, id string, params GetPostSharedpostsParams) (*core.Cursor[objects.Post], error) {
	var out core.Cursor[objects.Post]
	call := GetPostSharedpostsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetPostSponsorTagsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPostSponsorTagsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPostSponsorTagsBatchCall(id string, params GetPostSponsorTagsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "sponsor_tags"), params.ToParams(), options...)
}

func NewGetPostSponsorTagsBatchRequest(id string, params GetPostSponsorTagsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Page]] {
	return core.NewBatchRequest[core.Cursor[objects.Page]](GetPostSponsorTagsBatchCall(id, params, options...))
}

func DecodeGetPostSponsorTagsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Page], error) {
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

func GetPostSponsorTags(ctx context.Context, client *core.Client, id string, params GetPostSponsorTagsParams) (*core.Cursor[objects.Page], error) {
	var out core.Cursor[objects.Page]
	call := GetPostSponsorTagsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetPostToParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPostToParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPostToBatchCall(id string, params GetPostToParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "to"), params.ToParams(), options...)
}

func NewGetPostToBatchRequest(id string, params GetPostToParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Profile]] {
	return core.NewBatchRequest[core.Cursor[objects.Profile]](GetPostToBatchCall(id, params, options...))
}

func DecodeGetPostToBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Profile], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Profile]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPostTo(ctx context.Context, client *core.Client, id string, params GetPostToParams) (*core.Cursor[objects.Profile], error) {
	var out core.Cursor[objects.Profile]
	call := GetPostToBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeletePostParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeletePostParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeletePostBatchCall(id string, params DeletePostParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeletePostBatchRequest(id string, params DeletePostParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeletePostBatchCall(id, params, options...))
}

func DecodeDeletePostBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeletePost(ctx context.Context, client *core.Client, id string, params DeletePostParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeletePostBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetPostParams struct {
	PrimaryFbPageID *core.ID    `facebook:"primary_fb_page_id"`
	PrimaryIgUserID *core.ID    `facebook:"primary_ig_user_id"`
	Extra           core.Params `facebook:"-"`
}

func (params GetPostParams) ToParams() core.Params {
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

func GetPostBatchCall(id string, params GetPostParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetPostBatchRequest(id string, params GetPostParams, options ...core.BatchOption) *core.BatchRequest[objects.Post] {
	return core.NewBatchRequest[objects.Post](GetPostBatchCall(id, params, options...))
}

func DecodeGetPostBatchResponse(response *core.BatchResponse) (*objects.Post, error) {
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

func GetPost(ctx context.Context, client *core.Client, id string, params GetPostParams) (*objects.Post, error) {
	var out objects.Post
	call := GetPostBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type UpdatePostParams struct {
	AttachedMedia               *[]map[string]interface{}           `facebook:"attached_media"`
	BackdatedTime               *time.Time                          `facebook:"backdated_time"`
	BackdatedTimeGranularity    *enums.PostBackdatedTimeGranularity `facebook:"backdated_time_granularity"`
	ComposerSessionID           *core.ID                            `facebook:"composer_session_id"`
	DirectShareStatus           *uint64                             `facebook:"direct_share_status"`
	ExplicitlyAddedMentioneeIds *[]core.ID                          `facebook:"explicitly_added_mentionee_ids"`
	FeedStoryVisibility         *enums.PostFeedStoryVisibility      `facebook:"feed_story_visibility"`
	IsExplicitLocation          *bool                               `facebook:"is_explicit_location"`
	IsHidden                    *bool                               `facebook:"is_hidden"`
	IsPinned                    *bool                               `facebook:"is_pinned"`
	IsPublished                 *bool                               `facebook:"is_published"`
	Message                     *string                             `facebook:"message"`
	OgActionTypeID              *core.ID                            `facebook:"og_action_type_id"`
	OgHideObjectAttachment      *bool                               `facebook:"og_hide_object_attachment"`
	OgIconID                    *core.ID                            `facebook:"og_icon_id"`
	OgObjectID                  *core.ID                            `facebook:"og_object_id"`
	OgPhrase                    *string                             `facebook:"og_phrase"`
	OgSetProfileBadge           *bool                               `facebook:"og_set_profile_badge"`
	OgSuggestionMechanism       *string                             `facebook:"og_suggestion_mechanism"`
	Place                       *map[string]interface{}             `facebook:"place"`
	Privacy                     *string                             `facebook:"privacy"`
	ProductItem                 *map[string]interface{}             `facebook:"product_item"`
	ScheduledPublishTime        *uint64                             `facebook:"scheduled_publish_time"`
	ShouldSyncProductEdit       *bool                               `facebook:"should_sync_product_edit"`
	SourceType                  *string                             `facebook:"source_type"`
	SponsorID                   *core.ID                            `facebook:"sponsor_id"`
	SponsorRelationship         *uint64                             `facebook:"sponsor_relationship"`
	Tags                        *[]int                              `facebook:"tags"`
	TextFormatPresetID          *core.ID                            `facebook:"text_format_preset_id"`
	TimelineVisibility          *enums.PostTimelineVisibility       `facebook:"timeline_visibility"`
	Tracking                    *string                             `facebook:"tracking"`
	Extra                       core.Params                         `facebook:"-"`
}

func (params UpdatePostParams) ToParams() core.Params {
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

func UpdatePostBatchCall(id string, params UpdatePostParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdatePostBatchRequest(id string, params UpdatePostParams, options ...core.BatchOption) *core.BatchRequest[objects.Post] {
	return core.NewBatchRequest[objects.Post](UpdatePostBatchCall(id, params, options...))
}

func DecodeUpdatePostBatchResponse(response *core.BatchResponse) (*objects.Post, error) {
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

func UpdatePost(ctx context.Context, client *core.Client, id string, params UpdatePostParams) (*objects.Post, error) {
	var out objects.Post
	call := UpdatePostBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
