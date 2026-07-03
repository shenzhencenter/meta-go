package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetALMEndAdvertiserInfoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetALMEndAdvertiserInfoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetALMEndAdvertiserInfoBatchCall(id string, params GetALMEndAdvertiserInfoParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetALMEndAdvertiserInfoBatchRequest(id string, params GetALMEndAdvertiserInfoParams, options ...core.BatchOption) *core.BatchRequest[objects.ALMEndAdvertiserInfo] {
	return core.NewBatchRequest[objects.ALMEndAdvertiserInfo](GetALMEndAdvertiserInfoBatchCall(id, params, options...))
}

func DecodeGetALMEndAdvertiserInfoBatchResponse(response *core.BatchResponse) (*objects.ALMEndAdvertiserInfo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ALMEndAdvertiserInfo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetALMEndAdvertiserInfo(ctx context.Context, client *core.Client, id string, params GetALMEndAdvertiserInfoParams) (*objects.ALMEndAdvertiserInfo, error) {
	var out objects.ALMEndAdvertiserInfo
	call := GetALMEndAdvertiserInfoBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
