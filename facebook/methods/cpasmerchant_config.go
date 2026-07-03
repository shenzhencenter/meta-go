package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCPASMerchantConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCPASMerchantConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCPASMerchantConfigBatchCall(id string, params GetCPASMerchantConfigParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCPASMerchantConfigBatchRequest(id string, params GetCPASMerchantConfigParams, options ...core.BatchOption) *core.BatchRequest[objects.CPASMerchantConfig] {
	return core.NewBatchRequest[objects.CPASMerchantConfig](GetCPASMerchantConfigBatchCall(id, params, options...))
}

func DecodeGetCPASMerchantConfigBatchResponse(response *core.BatchResponse) (*objects.CPASMerchantConfig, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CPASMerchantConfig
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCPASMerchantConfigWithResponse(ctx context.Context, client *core.Client, id string, params GetCPASMerchantConfigParams) (*objects.CPASMerchantConfig, *core.Response, error) {
	var out objects.CPASMerchantConfig
	call := GetCPASMerchantConfigBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCPASMerchantConfig(ctx context.Context, client *core.Client, id string, params GetCPASMerchantConfigParams) (*objects.CPASMerchantConfig, error) {
	out, _, err := GetCPASMerchantConfigWithResponse(ctx, client, id, params)
	return out, err
}
