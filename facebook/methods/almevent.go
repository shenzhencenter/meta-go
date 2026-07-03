package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetALMEventParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetALMEventParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetALMEventBatchCall(id string, params GetALMEventParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetALMEventBatchRequest(id string, params GetALMEventParams, options ...core.BatchOption) *core.BatchRequest[objects.ALMEvent] {
	return core.NewBatchRequest[objects.ALMEvent](GetALMEventBatchCall(id, params, options...))
}

func DecodeGetALMEventBatchResponse(response *core.BatchResponse) (*objects.ALMEvent, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ALMEvent
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetALMEventWithResponse(ctx context.Context, client *core.Client, id string, params GetALMEventParams) (*objects.ALMEvent, *core.Response, error) {
	var out objects.ALMEvent
	call := GetALMEventBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetALMEvent(ctx context.Context, client *core.Client, id string, params GetALMEventParams) (*objects.ALMEvent, error) {
	out, _, err := GetALMEventWithResponse(ctx, client, id, params)
	return out, err
}
