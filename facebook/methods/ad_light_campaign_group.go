package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdLightCampaignGroupParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdLightCampaignGroupParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdLightCampaignGroupBatchCall(id string, params GetAdLightCampaignGroupParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdLightCampaignGroupBatchRequest(id string, params GetAdLightCampaignGroupParams, options ...core.BatchOption) *core.BatchRequest[objects.AdLightCampaignGroup] {
	return core.NewBatchRequest[objects.AdLightCampaignGroup](GetAdLightCampaignGroupBatchCall(id, params, options...))
}

func DecodeGetAdLightCampaignGroupBatchResponse(response *core.BatchResponse) (*objects.AdLightCampaignGroup, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdLightCampaignGroup
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdLightCampaignGroupWithResponse(ctx context.Context, client *core.Client, id string, params GetAdLightCampaignGroupParams) (*objects.AdLightCampaignGroup, *core.Response, error) {
	var out objects.AdLightCampaignGroup
	call := GetAdLightCampaignGroupBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdLightCampaignGroup(ctx context.Context, client *core.Client, id string, params GetAdLightCampaignGroupParams) (*objects.AdLightCampaignGroup, error) {
	out, _, err := GetAdLightCampaignGroupWithResponse(ctx, client, id, params)
	return out, err
}
