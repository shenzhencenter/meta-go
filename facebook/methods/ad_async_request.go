package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type DeleteAdAsyncRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteAdAsyncRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteAdAsyncRequestBatchCall(id string, params DeleteAdAsyncRequestParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteAdAsyncRequestBatchRequest(id string, params DeleteAdAsyncRequestParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteAdAsyncRequestBatchCall(id, params, options...))
}

func DecodeDeleteAdAsyncRequestBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func DeleteAdAsyncRequestWithResponse(ctx context.Context, client *core.Client, id string, params DeleteAdAsyncRequestParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteAdAsyncRequestBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteAdAsyncRequest(ctx context.Context, client *core.Client, id string, params DeleteAdAsyncRequestParams) (*map[string]interface{}, error) {
	out, _, err := DeleteAdAsyncRequestWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdAsyncRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAsyncRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAsyncRequestBatchCall(id string, params GetAdAsyncRequestParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdAsyncRequestBatchRequest(id string, params GetAdAsyncRequestParams, options ...core.BatchOption) *core.BatchRequest[objects.AdAsyncRequest] {
	return core.NewBatchRequest[objects.AdAsyncRequest](GetAdAsyncRequestBatchCall(id, params, options...))
}

func DecodeGetAdAsyncRequestBatchResponse(response *core.BatchResponse) (*objects.AdAsyncRequest, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdAsyncRequest
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAsyncRequestWithResponse(ctx context.Context, client *core.Client, id string, params GetAdAsyncRequestParams) (*objects.AdAsyncRequest, *core.Response, error) {
	var out objects.AdAsyncRequest
	call := GetAdAsyncRequestBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdAsyncRequest(ctx context.Context, client *core.Client, id string, params GetAdAsyncRequestParams) (*objects.AdAsyncRequest, error) {
	out, _, err := GetAdAsyncRequestWithResponse(ctx, client, id, params)
	return out, err
}
