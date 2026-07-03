package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetDynamicItemDisplayBundleParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetDynamicItemDisplayBundleParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetDynamicItemDisplayBundleBatchCall(id string, params GetDynamicItemDisplayBundleParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetDynamicItemDisplayBundleBatchRequest(id string, params GetDynamicItemDisplayBundleParams, options ...core.BatchOption) *core.BatchRequest[objects.DynamicItemDisplayBundle] {
	return core.NewBatchRequest[objects.DynamicItemDisplayBundle](GetDynamicItemDisplayBundleBatchCall(id, params, options...))
}

func DecodeGetDynamicItemDisplayBundleBatchResponse(response *core.BatchResponse) (*objects.DynamicItemDisplayBundle, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.DynamicItemDisplayBundle
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetDynamicItemDisplayBundleWithResponse(ctx context.Context, client *core.Client, id string, params GetDynamicItemDisplayBundleParams) (*objects.DynamicItemDisplayBundle, *core.Response, error) {
	var out objects.DynamicItemDisplayBundle
	call := GetDynamicItemDisplayBundleBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetDynamicItemDisplayBundle(ctx context.Context, client *core.Client, id string, params GetDynamicItemDisplayBundleParams) (*objects.DynamicItemDisplayBundle, error) {
	out, _, err := GetDynamicItemDisplayBundleWithResponse(ctx, client, id, params)
	return out, err
}
