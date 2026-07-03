package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdLightCampaignParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdLightCampaignParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdLightCampaignBatchCall(id string, params GetAdLightCampaignParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdLightCampaignBatchRequest(id string, params GetAdLightCampaignParams, options ...core.BatchOption) *core.BatchRequest[objects.AdLightCampaign] {
	return core.NewBatchRequest[objects.AdLightCampaign](GetAdLightCampaignBatchCall(id, params, options...))
}

func DecodeGetAdLightCampaignBatchResponse(response *core.BatchResponse) (*objects.AdLightCampaign, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdLightCampaign
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdLightCampaign(ctx context.Context, client *core.Client, id string, params GetAdLightCampaignParams) (*objects.AdLightCampaign, error) {
	var out objects.AdLightCampaign
	call := GetAdLightCampaignBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
