package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAnalyticsUserConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAnalyticsUserConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAnalyticsUserConfigBatchCall(id string, params GetAnalyticsUserConfigParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAnalyticsUserConfigBatchRequest(id string, params GetAnalyticsUserConfigParams, options ...core.BatchOption) *core.BatchRequest[objects.AnalyticsUserConfig] {
	return core.NewBatchRequest[objects.AnalyticsUserConfig](GetAnalyticsUserConfigBatchCall(id, params, options...))
}

func DecodeGetAnalyticsUserConfigBatchResponse(response *core.BatchResponse) (*objects.AnalyticsUserConfig, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AnalyticsUserConfig
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAnalyticsUserConfigWithResponse(ctx context.Context, client *core.Client, id string, params GetAnalyticsUserConfigParams) (*objects.AnalyticsUserConfig, *core.Response, error) {
	var out objects.AnalyticsUserConfig
	call := GetAnalyticsUserConfigBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAnalyticsUserConfig(ctx context.Context, client *core.Client, id string, params GetAnalyticsUserConfigParams) (*objects.AnalyticsUserConfig, error) {
	out, _, err := GetAnalyticsUserConfigWithResponse(ctx, client, id, params)
	return out, err
}
