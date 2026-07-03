package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAppEventConfigParams struct {
	EventName *string     `facebook:"event_name"`
	Extra     core.Params `facebook:"-"`
}

func (params GetAppEventConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.EventName != nil {
		out["event_name"] = *params.EventName
	}
	return out
}

func GetAppEventConfigBatchCall(id string, params GetAppEventConfigParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAppEventConfigBatchRequest(id string, params GetAppEventConfigParams, options ...core.BatchOption) *core.BatchRequest[objects.AppEventConfig] {
	return core.NewBatchRequest[objects.AppEventConfig](GetAppEventConfigBatchCall(id, params, options...))
}

func DecodeGetAppEventConfigBatchResponse(response *core.BatchResponse) (*objects.AppEventConfig, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AppEventConfig
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAppEventConfig(ctx context.Context, client *core.Client, id string, params GetAppEventConfigParams) (*objects.AppEventConfig, error) {
	var out objects.AppEventConfig
	call := GetAppEventConfigBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
