package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdCampaignPlacementParams struct {
	AccountID          core.ID     `facebook:"account_id"`
	BillingEvent       string      `facebook:"billing_event"`
	BuyingType         string      `facebook:"buying_type"`
	DestinationType    *string     `facebook:"destination_type"`
	HasDynamicCreative *bool       `facebook:"has_dynamic_creative"`
	Metadata           *bool       `facebook:"metadata"`
	Objective          string      `facebook:"objective"`
	OptimizationGoal   string      `facebook:"optimization_goal"`
	PromotedObject     *string     `facebook:"promoted_object"`
	Extra              core.Params `facebook:"-"`
}

func (params GetAdCampaignPlacementParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["account_id"] = params.AccountID
	out["billing_event"] = params.BillingEvent
	out["buying_type"] = params.BuyingType
	if params.DestinationType != nil {
		out["destination_type"] = *params.DestinationType
	}
	if params.HasDynamicCreative != nil {
		out["has_dynamic_creative"] = *params.HasDynamicCreative
	}
	if params.Metadata != nil {
		out["metadata"] = *params.Metadata
	}
	out["objective"] = params.Objective
	out["optimization_goal"] = params.OptimizationGoal
	if params.PromotedObject != nil {
		out["promoted_object"] = *params.PromotedObject
	}
	return out
}

func GetAdCampaignPlacementBatchCall(params GetAdCampaignPlacementParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath("ad_campaign_placement"), params.ToParams(), options...)
}

func NewGetAdCampaignPlacementBatchRequest(params GetAdCampaignPlacementParams, options ...core.BatchOption) *core.BatchRequest[objects.AdCampaignPlacementGet] {
	return core.NewBatchRequest[objects.AdCampaignPlacementGet](GetAdCampaignPlacementBatchCall(params, options...))
}

func DecodeGetAdCampaignPlacementBatchResponse(response *core.BatchResponse) (*objects.AdCampaignPlacementGet, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdCampaignPlacementGet
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdCampaignPlacementWithResponse(ctx context.Context, client *core.Client, params GetAdCampaignPlacementParams) (*objects.AdCampaignPlacementGet, *core.Response, error) {
	var out objects.AdCampaignPlacementGet
	call := GetAdCampaignPlacementBatchCall(params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdCampaignPlacement(ctx context.Context, client *core.Client, params GetAdCampaignPlacementParams) (*objects.AdCampaignPlacementGet, error) {
	out, _, err := GetAdCampaignPlacementWithResponse(ctx, client, params)
	return out, err
}
