package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdAccountCreationRequestAdaccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountCreationRequestAdaccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountCreationRequestAdaccountsBatchCall(id string, params GetAdAccountCreationRequestAdaccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adaccounts"), params.ToParams(), options...)
}

func NewGetAdAccountCreationRequestAdaccountsBatchRequest(id string, params GetAdAccountCreationRequestAdaccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccount]](GetAdAccountCreationRequestAdaccountsBatchCall(id, params, options...))
}

func DecodeGetAdAccountCreationRequestAdaccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccount], error) {
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

func GetAdAccountCreationRequestAdaccounts(ctx context.Context, client *core.Client, id string, params GetAdAccountCreationRequestAdaccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	call := GetAdAccountCreationRequestAdaccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetAdAccountCreationRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountCreationRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountCreationRequestBatchCall(id string, params GetAdAccountCreationRequestParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdAccountCreationRequestBatchRequest(id string, params GetAdAccountCreationRequestParams, options ...core.BatchOption) *core.BatchRequest[objects.AdAccountCreationRequest] {
	return core.NewBatchRequest[objects.AdAccountCreationRequest](GetAdAccountCreationRequestBatchCall(id, params, options...))
}

func DecodeGetAdAccountCreationRequestBatchResponse(response *core.BatchResponse) (*objects.AdAccountCreationRequest, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdAccountCreationRequest
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountCreationRequest(ctx context.Context, client *core.Client, id string, params GetAdAccountCreationRequestParams) (*objects.AdAccountCreationRequest, error) {
	var out objects.AdAccountCreationRequest
	call := GetAdAccountCreationRequestBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
