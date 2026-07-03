package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCatalogItemOverrideParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCatalogItemOverrideParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCatalogItemOverrideBatchCall(id string, params GetCatalogItemOverrideParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCatalogItemOverrideBatchRequest(id string, params GetCatalogItemOverrideParams, options ...core.BatchOption) *core.BatchRequest[objects.CatalogItemOverride] {
	return core.NewBatchRequest[objects.CatalogItemOverride](GetCatalogItemOverrideBatchCall(id, params, options...))
}

func DecodeGetCatalogItemOverrideBatchResponse(response *core.BatchResponse) (*objects.CatalogItemOverride, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CatalogItemOverride
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCatalogItemOverrideWithResponse(ctx context.Context, client *core.Client, id string, params GetCatalogItemOverrideParams) (*objects.CatalogItemOverride, *core.Response, error) {
	var out objects.CatalogItemOverride
	call := GetCatalogItemOverrideBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCatalogItemOverride(ctx context.Context, client *core.Client, id string, params GetCatalogItemOverrideParams) (*objects.CatalogItemOverride, error) {
	out, _, err := GetCatalogItemOverrideWithResponse(ctx, client, id, params)
	return out, err
}
