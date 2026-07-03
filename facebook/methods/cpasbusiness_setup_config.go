package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCPASBusinessSetupConfigAdAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCPASBusinessSetupConfigAdAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCPASBusinessSetupConfigAdAccountsBatchCall(id string, params GetCPASBusinessSetupConfigAdAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ad_accounts"), params.ToParams(), options...)
}

func NewGetCPASBusinessSetupConfigAdAccountsBatchRequest(id string, params GetCPASBusinessSetupConfigAdAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccount]](GetCPASBusinessSetupConfigAdAccountsBatchCall(id, params, options...))
}

func DecodeGetCPASBusinessSetupConfigAdAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccount], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccount]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCPASBusinessSetupConfigAdAccountsWithResponse(ctx context.Context, client *core.Client, id string, params GetCPASBusinessSetupConfigAdAccountsParams) (*core.Cursor[objects.AdAccount], *core.Response, error) {
	var out core.Cursor[objects.AdAccount]
	call := GetCPASBusinessSetupConfigAdAccountsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCPASBusinessSetupConfigAdAccounts(ctx context.Context, client *core.Client, id string, params GetCPASBusinessSetupConfigAdAccountsParams) (*core.Cursor[objects.AdAccount], error) {
	out, _, err := GetCPASBusinessSetupConfigAdAccountsWithResponse(ctx, client, id, params)
	return out, err
}

type GetCPASBusinessSetupConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCPASBusinessSetupConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCPASBusinessSetupConfigBatchCall(id string, params GetCPASBusinessSetupConfigParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCPASBusinessSetupConfigBatchRequest(id string, params GetCPASBusinessSetupConfigParams, options ...core.BatchOption) *core.BatchRequest[objects.CPASBusinessSetupConfig] {
	return core.NewBatchRequest[objects.CPASBusinessSetupConfig](GetCPASBusinessSetupConfigBatchCall(id, params, options...))
}

func DecodeGetCPASBusinessSetupConfigBatchResponse(response *core.BatchResponse) (*objects.CPASBusinessSetupConfig, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CPASBusinessSetupConfig
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCPASBusinessSetupConfigWithResponse(ctx context.Context, client *core.Client, id string, params GetCPASBusinessSetupConfigParams) (*objects.CPASBusinessSetupConfig, *core.Response, error) {
	var out objects.CPASBusinessSetupConfig
	call := GetCPASBusinessSetupConfigBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCPASBusinessSetupConfig(ctx context.Context, client *core.Client, id string, params GetCPASBusinessSetupConfigParams) (*objects.CPASBusinessSetupConfig, error) {
	out, _, err := GetCPASBusinessSetupConfigWithResponse(ctx, client, id, params)
	return out, err
}
