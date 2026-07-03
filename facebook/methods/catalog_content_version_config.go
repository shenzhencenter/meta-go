package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCatalogContentVersionConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCatalogContentVersionConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCatalogContentVersionConfigBatchCall(id string, params GetCatalogContentVersionConfigParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCatalogContentVersionConfigBatchRequest(id string, params GetCatalogContentVersionConfigParams, options ...core.BatchOption) *core.BatchRequest[objects.CatalogContentVersionConfig] {
	return core.NewBatchRequest[objects.CatalogContentVersionConfig](GetCatalogContentVersionConfigBatchCall(id, params, options...))
}

func DecodeGetCatalogContentVersionConfigBatchResponse(response *core.BatchResponse) (*objects.CatalogContentVersionConfig, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CatalogContentVersionConfig
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCatalogContentVersionConfigWithResponse(ctx context.Context, client *core.Client, id string, params GetCatalogContentVersionConfigParams) (*objects.CatalogContentVersionConfig, *core.Response, error) {
	var out objects.CatalogContentVersionConfig
	call := GetCatalogContentVersionConfigBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCatalogContentVersionConfig(ctx context.Context, client *core.Client, id string, params GetCatalogContentVersionConfigParams) (*objects.CatalogContentVersionConfig, error) {
	out, _, err := GetCatalogContentVersionConfigWithResponse(ctx, client, id, params)
	return out, err
}
