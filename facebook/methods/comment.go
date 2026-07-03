package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCommentCommentsParams struct {
	Filter     *enums.CommentcommentsFilterEnumParam     `facebook:"filter"`
	LiveFilter *enums.CommentcommentsLiveFilterEnumParam `facebook:"live_filter"`
	Order      *enums.CommentcommentsOrderEnumParam      `facebook:"order"`
	Since      *core.Time                                `facebook:"since"`
	Extra      core.Params                               `facebook:"-"`
}

func (params GetCommentCommentsParams) ToParams() core.Params {
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

func GetCommentCommentsBatchCall(id string, params GetCommentCommentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "comments"), params.ToParams(), options...)
}

func NewGetCommentCommentsBatchRequest(id string, params GetCommentCommentsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Comment]] {
	return core.NewBatchRequest[core.Cursor[objects.Comment]](GetCommentCommentsBatchCall(id, params, options...))
}

func DecodeGetCommentCommentsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Comment], error) {
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

func GetCommentCommentsWithResponse(ctx context.Context, client *core.Client, id string, params GetCommentCommentsParams) (*core.Cursor[objects.Comment], *core.Response, error) {
	var out core.Cursor[objects.Comment]
	call := GetCommentCommentsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCommentComments(ctx context.Context, client *core.Client, id string, params GetCommentCommentsParams) (*core.Cursor[objects.Comment], error) {
	out, _, err := GetCommentCommentsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateCommentCommentsParams struct {
	AttachmentID         *core.ID                                           `facebook:"attachment_id"`
	AttachmentShareURL   *string                                            `facebook:"attachment_share_url"`
	AttachmentURL        *string                                            `facebook:"attachment_url"`
	CommentPrivacyValue  *enums.CommentcommentsCommentPrivacyValueEnumParam `facebook:"comment_privacy_value"`
	FacepileMentionedIds *[]core.ID                                         `facebook:"facepile_mentioned_ids"`
	FeedbackSource       *string                                            `facebook:"feedback_source"`
	IsOffline            *bool                                              `facebook:"is_offline"`
	Message              *string                                            `facebook:"message"`
	NectarModule         *string                                            `facebook:"nectar_module"`
	ObjectID             *core.ID                                           `facebook:"object_id"`
	ParentCommentID      *map[string]interface{}                            `facebook:"parent_comment_id"`
	Text                 *string                                            `facebook:"text"`
	Tracking             *string                                            `facebook:"tracking"`
	Extra                core.Params                                        `facebook:"-"`
}

func (params CreateCommentCommentsParams) ToParams() core.Params {
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

func CreateCommentCommentsBatchCall(id string, params CreateCommentCommentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "comments"), params.ToParams(), options...)
}

func NewCreateCommentCommentsBatchRequest(id string, params CreateCommentCommentsParams, options ...core.BatchOption) *core.BatchRequest[objects.Comment] {
	return core.NewBatchRequest[objects.Comment](CreateCommentCommentsBatchCall(id, params, options...))
}

func DecodeCreateCommentCommentsBatchResponse(response *core.BatchResponse) (*objects.Comment, error) {
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

func CreateCommentCommentsWithResponse(ctx context.Context, client *core.Client, id string, params CreateCommentCommentsParams) (*objects.Comment, *core.Response, error) {
	var out objects.Comment
	call := CreateCommentCommentsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateCommentComments(ctx context.Context, client *core.Client, id string, params CreateCommentCommentsParams) (*objects.Comment, error) {
	out, _, err := CreateCommentCommentsWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteCommentLikesParams struct {
	AttributionIDV2 *string     `facebook:"attribution_id_v2"`
	FeedbackSource  *string     `facebook:"feedback_source"`
	NectarModule    *string     `facebook:"nectar_module"`
	Tracking        *string     `facebook:"tracking"`
	Extra           core.Params `facebook:"-"`
}

func (params DeleteCommentLikesParams) ToParams() core.Params {
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

func DeleteCommentLikesBatchCall(id string, params DeleteCommentLikesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id, "likes"), params.ToParams(), options...)
}

func NewDeleteCommentLikesBatchRequest(id string, params DeleteCommentLikesParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteCommentLikesBatchCall(id, params, options...))
}

func DecodeDeleteCommentLikesBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteCommentLikesWithResponse(ctx context.Context, client *core.Client, id string, params DeleteCommentLikesParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteCommentLikesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteCommentLikes(ctx context.Context, client *core.Client, id string, params DeleteCommentLikesParams) (*map[string]interface{}, error) {
	out, _, err := DeleteCommentLikesWithResponse(ctx, client, id, params)
	return out, err
}

type GetCommentLikesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCommentLikesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCommentLikesBatchCall(id string, params GetCommentLikesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "likes"), params.ToParams(), options...)
}

func NewGetCommentLikesBatchRequest(id string, params GetCommentLikesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Profile]] {
	return core.NewBatchRequest[core.Cursor[objects.Profile]](GetCommentLikesBatchCall(id, params, options...))
}

func DecodeGetCommentLikesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Profile], error) {
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

func GetCommentLikesWithResponse(ctx context.Context, client *core.Client, id string, params GetCommentLikesParams) (*core.Cursor[objects.Profile], *core.Response, error) {
	var out core.Cursor[objects.Profile]
	call := GetCommentLikesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCommentLikes(ctx context.Context, client *core.Client, id string, params GetCommentLikesParams) (*core.Cursor[objects.Profile], error) {
	out, _, err := GetCommentLikesWithResponse(ctx, client, id, params)
	return out, err
}

type CreateCommentLikesParams struct {
	AttributionIDV2 *string     `facebook:"attribution_id_v2"`
	FeedbackSource  *string     `facebook:"feedback_source"`
	NectarModule    *string     `facebook:"nectar_module"`
	Tracking        *string     `facebook:"tracking"`
	Extra           core.Params `facebook:"-"`
}

func (params CreateCommentLikesParams) ToParams() core.Params {
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

func CreateCommentLikesBatchCall(id string, params CreateCommentLikesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "likes"), params.ToParams(), options...)
}

func NewCreateCommentLikesBatchRequest(id string, params CreateCommentLikesParams, options ...core.BatchOption) *core.BatchRequest[objects.Comment] {
	return core.NewBatchRequest[objects.Comment](CreateCommentLikesBatchCall(id, params, options...))
}

func DecodeCreateCommentLikesBatchResponse(response *core.BatchResponse) (*objects.Comment, error) {
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

func CreateCommentLikesWithResponse(ctx context.Context, client *core.Client, id string, params CreateCommentLikesParams) (*objects.Comment, *core.Response, error) {
	var out objects.Comment
	call := CreateCommentLikesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateCommentLikes(ctx context.Context, client *core.Client, id string, params CreateCommentLikesParams) (*objects.Comment, error) {
	out, _, err := CreateCommentLikesWithResponse(ctx, client, id, params)
	return out, err
}

type GetCommentReactionsParams struct {
	Type  *enums.CommentreactionsTypeEnumParam `facebook:"type"`
	Extra core.Params                          `facebook:"-"`
}

func (params GetCommentReactionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	return out
}

func GetCommentReactionsBatchCall(id string, params GetCommentReactionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "reactions"), params.ToParams(), options...)
}

func NewGetCommentReactionsBatchRequest(id string, params GetCommentReactionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Profile]] {
	return core.NewBatchRequest[core.Cursor[objects.Profile]](GetCommentReactionsBatchCall(id, params, options...))
}

func DecodeGetCommentReactionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Profile], error) {
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

func GetCommentReactionsWithResponse(ctx context.Context, client *core.Client, id string, params GetCommentReactionsParams) (*core.Cursor[objects.Profile], *core.Response, error) {
	var out core.Cursor[objects.Profile]
	call := GetCommentReactionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCommentReactions(ctx context.Context, client *core.Client, id string, params GetCommentReactionsParams) (*core.Cursor[objects.Profile], error) {
	out, _, err := GetCommentReactionsWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteCommentParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteCommentParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteCommentBatchCall(id string, params DeleteCommentParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteCommentBatchRequest(id string, params DeleteCommentParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteCommentBatchCall(id, params, options...))
}

func DecodeDeleteCommentBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteCommentWithResponse(ctx context.Context, client *core.Client, id string, params DeleteCommentParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteCommentBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteComment(ctx context.Context, client *core.Client, id string, params DeleteCommentParams) (*map[string]interface{}, error) {
	out, _, err := DeleteCommentWithResponse(ctx, client, id, params)
	return out, err
}

type GetCommentParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCommentParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCommentBatchCall(id string, params GetCommentParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCommentBatchRequest(id string, params GetCommentParams, options ...core.BatchOption) *core.BatchRequest[objects.Comment] {
	return core.NewBatchRequest[objects.Comment](GetCommentBatchCall(id, params, options...))
}

func DecodeGetCommentBatchResponse(response *core.BatchResponse) (*objects.Comment, error) {
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

func GetCommentWithResponse(ctx context.Context, client *core.Client, id string, params GetCommentParams) (*objects.Comment, *core.Response, error) {
	var out objects.Comment
	call := GetCommentBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetComment(ctx context.Context, client *core.Client, id string, params GetCommentParams) (*objects.Comment, error) {
	out, _, err := GetCommentWithResponse(ctx, client, id, params)
	return out, err
}

type UpdateCommentParams struct {
	AttachmentID       *core.ID    `facebook:"attachment_id"`
	AttachmentShareURL *string     `facebook:"attachment_share_url"`
	AttachmentURL      *string     `facebook:"attachment_url"`
	IsHidden           *bool       `facebook:"is_hidden"`
	Message            *string     `facebook:"message"`
	Extra              core.Params `facebook:"-"`
}

func (params UpdateCommentParams) ToParams() core.Params {
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
	if params.IsHidden != nil {
		out["is_hidden"] = *params.IsHidden
	}
	if params.Message != nil {
		out["message"] = *params.Message
	}
	return out
}

func UpdateCommentBatchCall(id string, params UpdateCommentParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateCommentBatchRequest(id string, params UpdateCommentParams, options ...core.BatchOption) *core.BatchRequest[objects.Comment] {
	return core.NewBatchRequest[objects.Comment](UpdateCommentBatchCall(id, params, options...))
}

func DecodeUpdateCommentBatchResponse(response *core.BatchResponse) (*objects.Comment, error) {
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

func UpdateCommentWithResponse(ctx context.Context, client *core.Client, id string, params UpdateCommentParams) (*objects.Comment, *core.Response, error) {
	var out objects.Comment
	call := UpdateCommentBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateComment(ctx context.Context, client *core.Client, id string, params UpdateCommentParams) (*objects.Comment, error) {
	out, _, err := UpdateCommentWithResponse(ctx, client, id, params)
	return out, err
}
