package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetLifeEventLikesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLifeEventLikesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLifeEventLikesBatchCall(id string, params GetLifeEventLikesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "likes"), params.ToParams(), options...)
}

func NewGetLifeEventLikesBatchRequest(id string, params GetLifeEventLikesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Profile]] {
	return core.NewBatchRequest[core.Cursor[objects.Profile]](GetLifeEventLikesBatchCall(id, params, options...))
}

func DecodeGetLifeEventLikesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Profile], error) {
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

func GetLifeEventLikes(ctx context.Context, client *core.Client, id string, params GetLifeEventLikesParams) (*core.Cursor[objects.Profile], error) {
	var out core.Cursor[objects.Profile]
	call := GetLifeEventLikesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetLifeEventParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLifeEventParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLifeEventBatchCall(id string, params GetLifeEventParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetLifeEventBatchRequest(id string, params GetLifeEventParams, options ...core.BatchOption) *core.BatchRequest[objects.LifeEvent] {
	return core.NewBatchRequest[objects.LifeEvent](GetLifeEventBatchCall(id, params, options...))
}

func DecodeGetLifeEventBatchResponse(response *core.BatchResponse) (*objects.LifeEvent, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.LifeEvent
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetLifeEvent(ctx context.Context, client *core.Client, id string, params GetLifeEventParams) (*objects.LifeEvent, error) {
	var out objects.LifeEvent
	call := GetLifeEventBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
