package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetRTBDynamicPostCommentsParams struct {
	Filter     *enums.RtbdynamicpostcommentsFilterEnumParam     `facebook:"filter"`
	LiveFilter *enums.RtbdynamicpostcommentsLiveFilterEnumParam `facebook:"live_filter"`
	Order      *enums.RtbdynamicpostcommentsOrderEnumParam      `facebook:"order"`
	Since      *core.Time                                       `facebook:"since"`
	Extra      core.Params                                      `facebook:"-"`
}

func (params GetRTBDynamicPostCommentsParams) ToParams() core.Params {
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

func GetRTBDynamicPostCommentsBatchCall(id string, params GetRTBDynamicPostCommentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "comments"), params.ToParams(), options...)
}

func NewGetRTBDynamicPostCommentsBatchRequest(id string, params GetRTBDynamicPostCommentsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Comment]] {
	return core.NewBatchRequest[core.Cursor[objects.Comment]](GetRTBDynamicPostCommentsBatchCall(id, params, options...))
}

func DecodeGetRTBDynamicPostCommentsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Comment], error) {
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

func GetRTBDynamicPostCommentsWithResponse(ctx context.Context, client *core.Client, id string, params GetRTBDynamicPostCommentsParams) (*core.Cursor[objects.Comment], *core.Response, error) {
	var out core.Cursor[objects.Comment]
	call := GetRTBDynamicPostCommentsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetRTBDynamicPostComments(ctx context.Context, client *core.Client, id string, params GetRTBDynamicPostCommentsParams) (*core.Cursor[objects.Comment], error) {
	out, _, err := GetRTBDynamicPostCommentsWithResponse(ctx, client, id, params)
	return out, err
}

type GetRTBDynamicPostLikesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetRTBDynamicPostLikesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetRTBDynamicPostLikesBatchCall(id string, params GetRTBDynamicPostLikesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "likes"), params.ToParams(), options...)
}

func NewGetRTBDynamicPostLikesBatchRequest(id string, params GetRTBDynamicPostLikesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Profile]] {
	return core.NewBatchRequest[core.Cursor[objects.Profile]](GetRTBDynamicPostLikesBatchCall(id, params, options...))
}

func DecodeGetRTBDynamicPostLikesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Profile], error) {
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

func GetRTBDynamicPostLikesWithResponse(ctx context.Context, client *core.Client, id string, params GetRTBDynamicPostLikesParams) (*core.Cursor[objects.Profile], *core.Response, error) {
	var out core.Cursor[objects.Profile]
	call := GetRTBDynamicPostLikesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetRTBDynamicPostLikes(ctx context.Context, client *core.Client, id string, params GetRTBDynamicPostLikesParams) (*core.Cursor[objects.Profile], error) {
	out, _, err := GetRTBDynamicPostLikesWithResponse(ctx, client, id, params)
	return out, err
}

type GetRTBDynamicPostParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetRTBDynamicPostParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetRTBDynamicPostBatchCall(id string, params GetRTBDynamicPostParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetRTBDynamicPostBatchRequest(id string, params GetRTBDynamicPostParams, options ...core.BatchOption) *core.BatchRequest[objects.RTBDynamicPost] {
	return core.NewBatchRequest[objects.RTBDynamicPost](GetRTBDynamicPostBatchCall(id, params, options...))
}

func DecodeGetRTBDynamicPostBatchResponse(response *core.BatchResponse) (*objects.RTBDynamicPost, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.RTBDynamicPost
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetRTBDynamicPostWithResponse(ctx context.Context, client *core.Client, id string, params GetRTBDynamicPostParams) (*objects.RTBDynamicPost, *core.Response, error) {
	var out objects.RTBDynamicPost
	call := GetRTBDynamicPostBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetRTBDynamicPost(ctx context.Context, client *core.Client, id string, params GetRTBDynamicPostParams) (*objects.RTBDynamicPost, error) {
	out, _, err := GetRTBDynamicPostWithResponse(ctx, client, id, params)
	return out, err
}
