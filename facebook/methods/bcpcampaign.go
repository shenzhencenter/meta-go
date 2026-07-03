package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetBCPCampaignParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBCPCampaignParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBCPCampaignBatchCall(id string, params GetBCPCampaignParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetBCPCampaignBatchRequest(id string, params GetBCPCampaignParams, options ...core.BatchOption) *core.BatchRequest[objects.BCPCampaign] {
	return core.NewBatchRequest[objects.BCPCampaign](GetBCPCampaignBatchCall(id, params, options...))
}

func DecodeGetBCPCampaignBatchResponse(response *core.BatchResponse) (*objects.BCPCampaign, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BCPCampaign
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBCPCampaignWithResponse(ctx context.Context, client *core.Client, id string, params GetBCPCampaignParams) (*objects.BCPCampaign, *core.Response, error) {
	var out objects.BCPCampaign
	call := GetBCPCampaignBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetBCPCampaign(ctx context.Context, client *core.Client, id string, params GetBCPCampaignParams) (*objects.BCPCampaign, error) {
	out, _, err := GetBCPCampaignWithResponse(ctx, client, id, params)
	return out, err
}
