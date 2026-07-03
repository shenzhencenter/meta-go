package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetLiveVideoAdCampaignConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLiveVideoAdCampaignConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLiveVideoAdCampaignConfigBatchCall(id string, params GetLiveVideoAdCampaignConfigParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetLiveVideoAdCampaignConfigBatchRequest(id string, params GetLiveVideoAdCampaignConfigParams, options ...core.BatchOption) *core.BatchRequest[objects.LiveVideoAdCampaignConfig] {
	return core.NewBatchRequest[objects.LiveVideoAdCampaignConfig](GetLiveVideoAdCampaignConfigBatchCall(id, params, options...))
}

func DecodeGetLiveVideoAdCampaignConfigBatchResponse(response *core.BatchResponse) (*objects.LiveVideoAdCampaignConfig, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.LiveVideoAdCampaignConfig
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetLiveVideoAdCampaignConfigWithResponse(ctx context.Context, client *core.Client, id string, params GetLiveVideoAdCampaignConfigParams) (*objects.LiveVideoAdCampaignConfig, *core.Response, error) {
	var out objects.LiveVideoAdCampaignConfig
	call := GetLiveVideoAdCampaignConfigBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetLiveVideoAdCampaignConfig(ctx context.Context, client *core.Client, id string, params GetLiveVideoAdCampaignConfigParams) (*objects.LiveVideoAdCampaignConfig, error) {
	out, _, err := GetLiveVideoAdCampaignConfigWithResponse(ctx, client, id, params)
	return out, err
}
