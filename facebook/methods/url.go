package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
	"time"
)

type GetURLParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetURLParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetURLBatchCall(id string, params GetURLParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetURLBatchRequest(id string, params GetURLParams, options ...core.BatchOption) *core.BatchRequest[objects.URL] {
	return core.NewBatchRequest[objects.URL](GetURLBatchCall(id, params, options...))
}

func DecodeGetURLBatchResponse(response *core.BatchResponse) (*objects.URL, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.URL
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetURL(ctx context.Context, client *core.Client, id string, params GetURLParams) (*objects.URL, error) {
	var out objects.URL
	call := GetURLBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type UpdateURLParams struct {
	Blacklist *bool              `facebook:"blacklist"`
	Denylist  *bool              `facebook:"denylist"`
	Hmac      *string            `facebook:"hmac"`
	Locale    *[]string          `facebook:"locale"`
	Scopes    *[]enums.URLScopes `facebook:"scopes"`
	Ts        *time.Time         `facebook:"ts"`
	Extra     core.Params        `facebook:"-"`
}

func (params UpdateURLParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Blacklist != nil {
		out["blacklist"] = *params.Blacklist
	}
	if params.Denylist != nil {
		out["denylist"] = *params.Denylist
	}
	if params.Hmac != nil {
		out["hmac"] = *params.Hmac
	}
	if params.Locale != nil {
		out["locale"] = *params.Locale
	}
	if params.Scopes != nil {
		out["scopes"] = *params.Scopes
	}
	if params.Ts != nil {
		out["ts"] = *params.Ts
	}
	return out
}

func UpdateURLBatchCall(id string, params UpdateURLParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateURLBatchRequest(id string, params UpdateURLParams, options ...core.BatchOption) *core.BatchRequest[objects.URL] {
	return core.NewBatchRequest[objects.URL](UpdateURLBatchCall(id, params, options...))
}

func DecodeUpdateURLBatchResponse(response *core.BatchResponse) (*objects.URL, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.URL
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateURL(ctx context.Context, client *core.Client, id string, params UpdateURLParams) (*objects.URL, error) {
	var out objects.URL
	call := UpdateURLBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
