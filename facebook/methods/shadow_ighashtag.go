package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetShadowIGHashtagRecentMediaParams struct {
	UserID core.ID     `facebook:"user_id"`
	Extra  core.Params `facebook:"-"`
}

func (params GetShadowIGHashtagRecentMediaParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["user_id"] = params.UserID
	return out
}

func GetShadowIGHashtagRecentMediaBatchCall(id string, params GetShadowIGHashtagRecentMediaParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "recent_media"), params.ToParams(), options...)
}

func NewGetShadowIGHashtagRecentMediaBatchRequest(id string, params GetShadowIGHashtagRecentMediaParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGMedia]] {
	return core.NewBatchRequest[core.Cursor[objects.IGMedia]](GetShadowIGHashtagRecentMediaBatchCall(id, params, options...))
}

func DecodeGetShadowIGHashtagRecentMediaBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGMedia], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGMedia]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetShadowIGHashtagRecentMedia(ctx context.Context, client *core.Client, id string, params GetShadowIGHashtagRecentMediaParams) (*core.Cursor[objects.IGMedia], error) {
	var out core.Cursor[objects.IGMedia]
	call := GetShadowIGHashtagRecentMediaBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetShadowIGHashtagTopMediaParams struct {
	UserID core.ID     `facebook:"user_id"`
	Extra  core.Params `facebook:"-"`
}

func (params GetShadowIGHashtagTopMediaParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["user_id"] = params.UserID
	return out
}

func GetShadowIGHashtagTopMediaBatchCall(id string, params GetShadowIGHashtagTopMediaParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "top_media"), params.ToParams(), options...)
}

func NewGetShadowIGHashtagTopMediaBatchRequest(id string, params GetShadowIGHashtagTopMediaParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGMedia]] {
	return core.NewBatchRequest[core.Cursor[objects.IGMedia]](GetShadowIGHashtagTopMediaBatchCall(id, params, options...))
}

func DecodeGetShadowIGHashtagTopMediaBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGMedia], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGMedia]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetShadowIGHashtagTopMedia(ctx context.Context, client *core.Client, id string, params GetShadowIGHashtagTopMediaParams) (*core.Cursor[objects.IGMedia], error) {
	var out core.Cursor[objects.IGMedia]
	call := GetShadowIGHashtagTopMediaBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetShadowIGHashtagParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetShadowIGHashtagParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetShadowIGHashtagBatchCall(id string, params GetShadowIGHashtagParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetShadowIGHashtagBatchRequest(id string, params GetShadowIGHashtagParams, options ...core.BatchOption) *core.BatchRequest[objects.ShadowIGHashtag] {
	return core.NewBatchRequest[objects.ShadowIGHashtag](GetShadowIGHashtagBatchCall(id, params, options...))
}

func DecodeGetShadowIGHashtagBatchResponse(response *core.BatchResponse) (*objects.ShadowIGHashtag, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ShadowIGHashtag
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetShadowIGHashtag(ctx context.Context, client *core.Client, id string, params GetShadowIGHashtagParams) (*objects.ShadowIGHashtag, error) {
	var out objects.ShadowIGHashtag
	call := GetShadowIGHashtagBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
