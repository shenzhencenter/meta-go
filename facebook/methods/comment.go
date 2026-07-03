package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
	"time"
)

type GetCommentCommentsParams struct {
	Filter     *enums.CommentcommentsFilterEnumParam     `facebook:"filter"`
	LiveFilter *enums.CommentcommentsLiveFilterEnumParam `facebook:"live_filter"`
	Order      *enums.CommentcommentsOrderEnumParam      `facebook:"order"`
	Since      *time.Time                                `facebook:"since"`
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

func GetCommentComments(ctx context.Context, client *core.Client, id string, params GetCommentCommentsParams) (*core.Cursor[objects.Comment], error) {
	var out core.Cursor[objects.Comment]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "comments"), params.ToParams(), &out)
	return &out, err
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

func CreateCommentComments(ctx context.Context, client *core.Client, id string, params CreateCommentCommentsParams) (*objects.Comment, error) {
	var out objects.Comment
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "comments"), params.ToParams(), &out)
	return &out, err
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

func DeleteCommentLikes(ctx context.Context, client *core.Client, id string, params DeleteCommentLikesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "likes"), params.ToParams(), &out)
	return &out, err
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

func GetCommentLikes(ctx context.Context, client *core.Client, id string, params GetCommentLikesParams) (*core.Cursor[objects.Profile], error) {
	var out core.Cursor[objects.Profile]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "likes"), params.ToParams(), &out)
	return &out, err
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

func CreateCommentLikes(ctx context.Context, client *core.Client, id string, params CreateCommentLikesParams) (*objects.Comment, error) {
	var out objects.Comment
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "likes"), params.ToParams(), &out)
	return &out, err
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

func GetCommentReactions(ctx context.Context, client *core.Client, id string, params GetCommentReactionsParams) (*core.Cursor[objects.Profile], error) {
	var out core.Cursor[objects.Profile]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "reactions"), params.ToParams(), &out)
	return &out, err
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

func DeleteComment(ctx context.Context, client *core.Client, id string, params DeleteCommentParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
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

func GetComment(ctx context.Context, client *core.Client, id string, params GetCommentParams) (*objects.Comment, error) {
	var out objects.Comment
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
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

func UpdateComment(ctx context.Context, client *core.Client, id string, params UpdateCommentParams) (*objects.Comment, error) {
	var out objects.Comment
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
