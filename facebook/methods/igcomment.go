package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetIGCommentRepliesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGCommentRepliesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGCommentReplies(ctx context.Context, client *core.Client, id string, params GetIGCommentRepliesParams) (*core.Cursor[objects.IGComment], error) {
	var out core.Cursor[objects.IGComment]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "replies"), params.ToParams(), &out)
	return &out, err
}

type CreateIGCommentRepliesParams struct {
	Message *string     `facebook:"message"`
	Extra   core.Params `facebook:"-"`
}

func (params CreateIGCommentRepliesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Message != nil {
		out["message"] = *params.Message
	}
	return out
}

func CreateIGCommentReplies(ctx context.Context, client *core.Client, id string, params CreateIGCommentRepliesParams) (*objects.IGComment, error) {
	var out objects.IGComment
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "replies"), params.ToParams(), &out)
	return &out, err
}

type DeleteIGCommentParams struct {
	AdID  *core.ID    `facebook:"ad_id"`
	Extra core.Params `facebook:"-"`
}

func (params DeleteIGCommentParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdID != nil {
		out["ad_id"] = *params.AdID
	}
	return out
}

func DeleteIGComment(ctx context.Context, client *core.Client, id string, params DeleteIGCommentParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetIGCommentParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGCommentParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGComment(ctx context.Context, client *core.Client, id string, params GetIGCommentParams) (*objects.IGComment, error) {
	var out objects.IGComment
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateIGCommentParams struct {
	AdID  *core.ID    `facebook:"ad_id"`
	Hide  bool        `facebook:"hide"`
	Extra core.Params `facebook:"-"`
}

func (params UpdateIGCommentParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdID != nil {
		out["ad_id"] = *params.AdID
	}
	out["hide"] = params.Hide
	return out
}

func UpdateIGComment(ctx context.Context, client *core.Client, id string, params UpdateIGCommentParams) (*objects.IGComment, error) {
	var out objects.IGComment
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
