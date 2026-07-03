package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
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

func GetPagePostAttachmentsBatchCall(id string, params GetPagePostAttachmentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "attachments"), params.ToParams(), options...)
}

func NewGetPagePostAttachmentsBatchRequest(id string, params GetPagePostAttachmentsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.StoryAttachment]] {
	return core.NewBatchRequest[core.Cursor[objects.StoryAttachment]](GetPagePostAttachmentsBatchCall(id, params, options...))
}

func DecodeGetPagePostAttachmentsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.StoryAttachment], error) {
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

func GetPagePostAttachmentsWithResponse(ctx context.Context, client *core.Client, id string, params GetPagePostAttachmentsParams) (*core.Cursor[objects.StoryAttachment], *core.Response, error) {
	var out core.Cursor[objects.StoryAttachment]
	call := GetPagePostAttachmentsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPagePostAttachments(ctx context.Context, client *core.Client, id string, params GetPagePostAttachmentsParams) (*core.Cursor[objects.StoryAttachment], error) {
	out, _, err := GetPagePostAttachmentsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPagePostCommentsParams struct {
	Filter     *enums.PagepostcommentsFilterEnumParam     `facebook:"filter"`
	LiveFilter *enums.PagepostcommentsLiveFilterEnumParam `facebook:"live_filter"`
	Order      *enums.PagepostcommentsOrderEnumParam      `facebook:"order"`
	Since      *core.Time                                 `facebook:"since"`
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

func GetPagePostCommentsBatchCall(id string, params GetPagePostCommentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "comments"), params.ToParams(), options...)
}

func NewGetPagePostCommentsBatchRequest(id string, params GetPagePostCommentsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Comment]] {
	return core.NewBatchRequest[core.Cursor[objects.Comment]](GetPagePostCommentsBatchCall(id, params, options...))
}

func DecodeGetPagePostCommentsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Comment], error) {
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

func GetPagePostCommentsWithResponse(ctx context.Context, client *core.Client, id string, params GetPagePostCommentsParams) (*core.Cursor[objects.Comment], *core.Response, error) {
	var out core.Cursor[objects.Comment]
	call := GetPagePostCommentsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPagePostComments(ctx context.Context, client *core.Client, id string, params GetPagePostCommentsParams) (*core.Cursor[objects.Comment], error) {
	out, _, err := GetPagePostCommentsWithResponse(ctx, client, id, params)
	return out, err
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

func CreatePagePostCommentsBatchCall(id string, params CreatePagePostCommentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "comments"), params.ToParams(), options...)
}

func NewCreatePagePostCommentsBatchRequest(id string, params CreatePagePostCommentsParams, options ...core.BatchOption) *core.BatchRequest[objects.Comment] {
	return core.NewBatchRequest[objects.Comment](CreatePagePostCommentsBatchCall(id, params, options...))
}

func DecodeCreatePagePostCommentsBatchResponse(response *core.BatchResponse) (*objects.Comment, error) {
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

func CreatePagePostCommentsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePagePostCommentsParams) (*objects.Comment, *core.Response, error) {
	var out objects.Comment
	call := CreatePagePostCommentsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePagePostComments(ctx context.Context, client *core.Client, id string, params CreatePagePostCommentsParams) (*objects.Comment, error) {
	out, _, err := CreatePagePostCommentsWithResponse(ctx, client, id, params)
	return out, err
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

func GetPagePostDynamicPostsBatchCall(id string, params GetPagePostDynamicPostsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "dynamic_posts"), params.ToParams(), options...)
}

func NewGetPagePostDynamicPostsBatchRequest(id string, params GetPagePostDynamicPostsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.RTBDynamicPost]] {
	return core.NewBatchRequest[core.Cursor[objects.RTBDynamicPost]](GetPagePostDynamicPostsBatchCall(id, params, options...))
}

func DecodeGetPagePostDynamicPostsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.RTBDynamicPost], error) {
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

func GetPagePostDynamicPostsWithResponse(ctx context.Context, client *core.Client, id string, params GetPagePostDynamicPostsParams) (*core.Cursor[objects.RTBDynamicPost], *core.Response, error) {
	var out core.Cursor[objects.RTBDynamicPost]
	call := GetPagePostDynamicPostsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPagePostDynamicPosts(ctx context.Context, client *core.Client, id string, params GetPagePostDynamicPostsParams) (*core.Cursor[objects.RTBDynamicPost], error) {
	out, _, err := GetPagePostDynamicPostsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPagePostInsightsParams struct {
	DatePreset *enums.PagepostinsightsDatePresetEnumParam `facebook:"date_preset"`
	Metric     *[]map[string]interface{}                  `facebook:"metric"`
	Period     *enums.PagepostinsightsPeriodEnumParam     `facebook:"period"`
	Since      *core.Time                                 `facebook:"since"`
	Until      *core.Time                                 `facebook:"until"`
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

func GetPagePostInsightsBatchCall(id string, params GetPagePostInsightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "insights"), params.ToParams(), options...)
}

func NewGetPagePostInsightsBatchRequest(id string, params GetPagePostInsightsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.InsightsResult]] {
	return core.NewBatchRequest[core.Cursor[objects.InsightsResult]](GetPagePostInsightsBatchCall(id, params, options...))
}

func DecodeGetPagePostInsightsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.InsightsResult], error) {
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

func GetPagePostInsightsWithResponse(ctx context.Context, client *core.Client, id string, params GetPagePostInsightsParams) (*core.Cursor[objects.InsightsResult], *core.Response, error) {
	var out core.Cursor[objects.InsightsResult]
	call := GetPagePostInsightsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPagePostInsights(ctx context.Context, client *core.Client, id string, params GetPagePostInsightsParams) (*core.Cursor[objects.InsightsResult], error) {
	out, _, err := GetPagePostInsightsWithResponse(ctx, client, id, params)
	return out, err
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

func DeletePagePostLikesBatchCall(id string, params DeletePagePostLikesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "likes"), params.ToParams(), options...)
}

func NewDeletePagePostLikesBatchRequest(id string, params DeletePagePostLikesParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeletePagePostLikesBatchCall(id, params, options...))
}

func DecodeDeletePagePostLikesBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeletePagePostLikesWithResponse(ctx context.Context, client *core.Client, id string, params DeletePagePostLikesParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeletePagePostLikesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeletePagePostLikes(ctx context.Context, client *core.Client, id string, params DeletePagePostLikesParams) (*map[string]interface{}, error) {
	out, _, err := DeletePagePostLikesWithResponse(ctx, client, id, params)
	return out, err
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

func GetPagePostLikesBatchCall(id string, params GetPagePostLikesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "likes"), params.ToParams(), options...)
}

func NewGetPagePostLikesBatchRequest(id string, params GetPagePostLikesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Profile]] {
	return core.NewBatchRequest[core.Cursor[objects.Profile]](GetPagePostLikesBatchCall(id, params, options...))
}

func DecodeGetPagePostLikesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Profile], error) {
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

func GetPagePostLikesWithResponse(ctx context.Context, client *core.Client, id string, params GetPagePostLikesParams) (*core.Cursor[objects.Profile], *core.Response, error) {
	var out core.Cursor[objects.Profile]
	call := GetPagePostLikesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPagePostLikes(ctx context.Context, client *core.Client, id string, params GetPagePostLikesParams) (*core.Cursor[objects.Profile], error) {
	out, _, err := GetPagePostLikesWithResponse(ctx, client, id, params)
	return out, err
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

func CreatePagePostLikesBatchCall(id string, params CreatePagePostLikesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "likes"), params.ToParams(), options...)
}

func NewCreatePagePostLikesBatchRequest(id string, params CreatePagePostLikesParams, options ...core.BatchOption) *core.BatchRequest[objects.PagePost] {
	return core.NewBatchRequest[objects.PagePost](CreatePagePostLikesBatchCall(id, params, options...))
}

func DecodeCreatePagePostLikesBatchResponse(response *core.BatchResponse) (*objects.PagePost, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PagePost
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePagePostLikesWithResponse(ctx context.Context, client *core.Client, id string, params CreatePagePostLikesParams) (*objects.PagePost, *core.Response, error) {
	var out objects.PagePost
	call := CreatePagePostLikesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePagePostLikes(ctx context.Context, client *core.Client, id string, params CreatePagePostLikesParams) (*objects.PagePost, error) {
	out, _, err := CreatePagePostLikesWithResponse(ctx, client, id, params)
	return out, err
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

func GetPagePostReactionsBatchCall(id string, params GetPagePostReactionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "reactions"), params.ToParams(), options...)
}

func NewGetPagePostReactionsBatchRequest(id string, params GetPagePostReactionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Profile]] {
	return core.NewBatchRequest[core.Cursor[objects.Profile]](GetPagePostReactionsBatchCall(id, params, options...))
}

func DecodeGetPagePostReactionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Profile], error) {
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

func GetPagePostReactionsWithResponse(ctx context.Context, client *core.Client, id string, params GetPagePostReactionsParams) (*core.Cursor[objects.Profile], *core.Response, error) {
	var out core.Cursor[objects.Profile]
	call := GetPagePostReactionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPagePostReactions(ctx context.Context, client *core.Client, id string, params GetPagePostReactionsParams) (*core.Cursor[objects.Profile], error) {
	out, _, err := GetPagePostReactionsWithResponse(ctx, client, id, params)
	return out, err
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

func GetPagePostSharedpostsBatchCall(id string, params GetPagePostSharedpostsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "sharedposts"), params.ToParams(), options...)
}

func NewGetPagePostSharedpostsBatchRequest(id string, params GetPagePostSharedpostsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Post]] {
	return core.NewBatchRequest[core.Cursor[objects.Post]](GetPagePostSharedpostsBatchCall(id, params, options...))
}

func DecodeGetPagePostSharedpostsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Post], error) {
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

func GetPagePostSharedpostsWithResponse(ctx context.Context, client *core.Client, id string, params GetPagePostSharedpostsParams) (*core.Cursor[objects.Post], *core.Response, error) {
	var out core.Cursor[objects.Post]
	call := GetPagePostSharedpostsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPagePostSharedposts(ctx context.Context, client *core.Client, id string, params GetPagePostSharedpostsParams) (*core.Cursor[objects.Post], error) {
	out, _, err := GetPagePostSharedpostsWithResponse(ctx, client, id, params)
	return out, err
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

func GetPagePostSponsorTagsBatchCall(id string, params GetPagePostSponsorTagsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "sponsor_tags"), params.ToParams(), options...)
}

func NewGetPagePostSponsorTagsBatchRequest(id string, params GetPagePostSponsorTagsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Page]] {
	return core.NewBatchRequest[core.Cursor[objects.Page]](GetPagePostSponsorTagsBatchCall(id, params, options...))
}

func DecodeGetPagePostSponsorTagsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Page], error) {
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

func GetPagePostSponsorTagsWithResponse(ctx context.Context, client *core.Client, id string, params GetPagePostSponsorTagsParams) (*core.Cursor[objects.Page], *core.Response, error) {
	var out core.Cursor[objects.Page]
	call := GetPagePostSponsorTagsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPagePostSponsorTags(ctx context.Context, client *core.Client, id string, params GetPagePostSponsorTagsParams) (*core.Cursor[objects.Page], error) {
	out, _, err := GetPagePostSponsorTagsWithResponse(ctx, client, id, params)
	return out, err
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

func GetPagePostToBatchCall(id string, params GetPagePostToParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "to"), params.ToParams(), options...)
}

func NewGetPagePostToBatchRequest(id string, params GetPagePostToParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Profile]] {
	return core.NewBatchRequest[core.Cursor[objects.Profile]](GetPagePostToBatchCall(id, params, options...))
}

func DecodeGetPagePostToBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Profile], error) {
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

func GetPagePostToWithResponse(ctx context.Context, client *core.Client, id string, params GetPagePostToParams) (*core.Cursor[objects.Profile], *core.Response, error) {
	var out core.Cursor[objects.Profile]
	call := GetPagePostToBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPagePostTo(ctx context.Context, client *core.Client, id string, params GetPagePostToParams) (*core.Cursor[objects.Profile], error) {
	out, _, err := GetPagePostToWithResponse(ctx, client, id, params)
	return out, err
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

func DeletePagePostBatchCall(id string, params DeletePagePostParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeletePagePostBatchRequest(id string, params DeletePagePostParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeletePagePostBatchCall(id, params, options...))
}

func DecodeDeletePagePostBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeletePagePostWithResponse(ctx context.Context, client *core.Client, id string, params DeletePagePostParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeletePagePostBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeletePagePost(ctx context.Context, client *core.Client, id string, params DeletePagePostParams) (*map[string]interface{}, error) {
	out, _, err := DeletePagePostWithResponse(ctx, client, id, params)
	return out, err
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

func GetPagePostBatchCall(id string, params GetPagePostParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetPagePostBatchRequest(id string, params GetPagePostParams, options ...core.BatchOption) *core.BatchRequest[objects.PagePost] {
	return core.NewBatchRequest[objects.PagePost](GetPagePostBatchCall(id, params, options...))
}

func DecodeGetPagePostBatchResponse(response *core.BatchResponse) (*objects.PagePost, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PagePost
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPagePostWithResponse(ctx context.Context, client *core.Client, id string, params GetPagePostParams) (*objects.PagePost, *core.Response, error) {
	var out objects.PagePost
	call := GetPagePostBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPagePost(ctx context.Context, client *core.Client, id string, params GetPagePostParams) (*objects.PagePost, error) {
	out, _, err := GetPagePostWithResponse(ctx, client, id, params)
	return out, err
}

type UpdatePagePostParams struct {
	AttachedMedia               *[]map[string]interface{}               `facebook:"attached_media"`
	BackdatedTime               *core.Time                              `facebook:"backdated_time"`
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

func UpdatePagePostBatchCall(id string, params UpdatePagePostParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdatePagePostBatchRequest(id string, params UpdatePagePostParams, options ...core.BatchOption) *core.BatchRequest[objects.PagePost] {
	return core.NewBatchRequest[objects.PagePost](UpdatePagePostBatchCall(id, params, options...))
}

func DecodeUpdatePagePostBatchResponse(response *core.BatchResponse) (*objects.PagePost, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PagePost
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdatePagePostWithResponse(ctx context.Context, client *core.Client, id string, params UpdatePagePostParams) (*objects.PagePost, *core.Response, error) {
	var out objects.PagePost
	call := UpdatePagePostBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdatePagePost(ctx context.Context, client *core.Client, id string, params UpdatePagePostParams) (*objects.PagePost, error) {
	out, _, err := UpdatePagePostWithResponse(ctx, client, id, params)
	return out, err
}
