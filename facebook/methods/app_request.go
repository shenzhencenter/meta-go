package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type DeleteAppRequestParams struct {
	Ids   []core.ID   `facebook:"ids"`
	Extra core.Params `facebook:"-"`
}

func (params DeleteAppRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["ids"] = params.Ids
	return out
}

func DeleteAppRequestBatchCall(id string, params DeleteAppRequestParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteAppRequestBatchRequest(id string, params DeleteAppRequestParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteAppRequestBatchCall(id, params, options...))
}

func DecodeDeleteAppRequestBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteAppRequestWithResponse(ctx context.Context, client *core.Client, id string, params DeleteAppRequestParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteAppRequestBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteAppRequest(ctx context.Context, client *core.Client, id string, params DeleteAppRequestParams) (*map[string]interface{}, error) {
	out, _, err := DeleteAppRequestWithResponse(ctx, client, id, params)
	return out, err
}

type GetAppRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAppRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAppRequestBatchCall(id string, params GetAppRequestParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAppRequestBatchRequest(id string, params GetAppRequestParams, options ...core.BatchOption) *core.BatchRequest[objects.AppRequest] {
	return core.NewBatchRequest[objects.AppRequest](GetAppRequestBatchCall(id, params, options...))
}

func DecodeGetAppRequestBatchResponse(response *core.BatchResponse) (*objects.AppRequest, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AppRequest
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAppRequestWithResponse(ctx context.Context, client *core.Client, id string, params GetAppRequestParams) (*objects.AppRequest, *core.Response, error) {
	var out objects.AppRequest
	call := GetAppRequestBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAppRequest(ctx context.Context, client *core.Client, id string, params GetAppRequestParams) (*objects.AppRequest, error) {
	out, _, err := GetAppRequestWithResponse(ctx, client, id, params)
	return out, err
}
