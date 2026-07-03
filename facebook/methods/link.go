package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func CreateLinkComments(ctx context.Context, client *core.Client, id string, params CreateLinkCommentsParams) (*objects.Comment, error) {
	var out objects.Comment
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "comments"), params.ToParams(), &out)
	return &out, err
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

func GetLinkLikes(ctx context.Context, client *core.Client, id string, params GetLinkLikesParams) (*core.Cursor[objects.Profile], error) {
	var out core.Cursor[objects.Profile]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "likes"), params.ToParams(), &out)
	return &out, err
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

func GetLink(ctx context.Context, client *core.Client, id string, params GetLinkParams) (*objects.Link, error) {
	var out objects.Link
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
