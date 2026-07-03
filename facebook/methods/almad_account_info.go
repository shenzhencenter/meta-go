package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetALMAdAccountInfoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetALMAdAccountInfoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetALMAdAccountInfoBatchCall(id string, params GetALMAdAccountInfoParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetALMAdAccountInfoBatchRequest(id string, params GetALMAdAccountInfoParams, options ...core.BatchOption) *core.BatchRequest[objects.ALMAdAccountInfo] {
	return core.NewBatchRequest[objects.ALMAdAccountInfo](GetALMAdAccountInfoBatchCall(id, params, options...))
}

func DecodeGetALMAdAccountInfoBatchResponse(response *core.BatchResponse) (*objects.ALMAdAccountInfo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ALMAdAccountInfo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetALMAdAccountInfoWithResponse(ctx context.Context, client *core.Client, id string, params GetALMAdAccountInfoParams) (*objects.ALMAdAccountInfo, *core.Response, error) {
	var out objects.ALMAdAccountInfo
	call := GetALMAdAccountInfoBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetALMAdAccountInfo(ctx context.Context, client *core.Client, id string, params GetALMAdAccountInfoParams) (*objects.ALMAdAccountInfo, error) {
	out, _, err := GetALMAdAccountInfoWithResponse(ctx, client, id, params)
	return out, err
}
