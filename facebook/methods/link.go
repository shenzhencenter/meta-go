package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type CreateLinkCommentsParams struct {
	AttachmentID         *core.ID                                        `facebook:"attachment_id"`
	AttachmentShareURL   *string                                         `facebook:"attachment_share_url"`
	AttachmentURL        *string                                         `facebook:"attachment_url"`
	CommentPrivacyValue  *enums.LinkcommentsCommentPrivacyValueEnumParam `facebook:"comment_privacy_value"`
	FacepileMentionedIds *[]core.ID                                      `facebook:"facepile_mentioned_ids"`
	FeedbackSource       *string                                         `facebook:"feedback_source"`
	IsOffline            *bool                                           `facebook:"is_offline"`
	Message              *string                                         `facebook:"message"`
	NectarModule         *string                                         `facebook:"nectar_module"`
	ObjectID             *core.ID                                        `facebook:"object_id"`
	ParentCommentID      *map[string]interface{}                         `facebook:"parent_comment_id"`
	Text                 *string                                         `facebook:"text"`
	Tracking             *string                                         `facebook:"tracking"`
	Extra                core.Params                                     `facebook:"-"`
}

func (params CreateLinkCommentsParams) ToParams() core.Params {
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

func CreateLinkCommentsBatchCall(id string, params CreateLinkCommentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "comments"), params.ToParams(), options...)
}

func NewCreateLinkCommentsBatchRequest(id string, params CreateLinkCommentsParams, options ...core.BatchOption) *core.BatchRequest[objects.Comment] {
	return core.NewBatchRequest[objects.Comment](CreateLinkCommentsBatchCall(id, params, options...))
}

func DecodeCreateLinkCommentsBatchResponse(response *core.BatchResponse) (*objects.Comment, error) {
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

func CreateLinkCommentsWithResponse(ctx context.Context, client *core.Client, id string, params CreateLinkCommentsParams) (*objects.Comment, *core.Response, error) {
	var out objects.Comment
	call := CreateLinkCommentsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateLinkComments(ctx context.Context, client *core.Client, id string, params CreateLinkCommentsParams) (*objects.Comment, error) {
	out, _, err := CreateLinkCommentsWithResponse(ctx, client, id, params)
	return out, err
}

type GetLinkLikesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLinkLikesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLinkLikesBatchCall(id string, params GetLinkLikesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "likes"), params.ToParams(), options...)
}

func NewGetLinkLikesBatchRequest(id string, params GetLinkLikesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Profile]] {
	return core.NewBatchRequest[core.Cursor[objects.Profile]](GetLinkLikesBatchCall(id, params, options...))
}

func DecodeGetLinkLikesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Profile], error) {
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

func GetLinkLikesWithResponse(ctx context.Context, client *core.Client, id string, params GetLinkLikesParams) (*core.Cursor[objects.Profile], *core.Response, error) {
	var out core.Cursor[objects.Profile]
	call := GetLinkLikesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetLinkLikes(ctx context.Context, client *core.Client, id string, params GetLinkLikesParams) (*core.Cursor[objects.Profile], error) {
	out, _, err := GetLinkLikesWithResponse(ctx, client, id, params)
	return out, err
}

type GetLinkParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLinkParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLinkBatchCall(id string, params GetLinkParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetLinkBatchRequest(id string, params GetLinkParams, options ...core.BatchOption) *core.BatchRequest[objects.Link] {
	return core.NewBatchRequest[objects.Link](GetLinkBatchCall(id, params, options...))
}

func DecodeGetLinkBatchResponse(response *core.BatchResponse) (*objects.Link, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Link
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetLinkWithResponse(ctx context.Context, client *core.Client, id string, params GetLinkParams) (*objects.Link, *core.Response, error) {
	var out objects.Link
	call := GetLinkBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetLink(ctx context.Context, client *core.Client, id string, params GetLinkParams) (*objects.Link, error) {
	out, _, err := GetLinkWithResponse(ctx, client, id, params)
	return out, err
}
