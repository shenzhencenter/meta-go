package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetOmegaCustomerTrxCampaignsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOmegaCustomerTrxCampaignsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOmegaCustomerTrxCampaignsBatchCall(id string, params GetOmegaCustomerTrxCampaignsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "campaigns"), params.ToParams(), options...)
}

func NewGetOmegaCustomerTrxCampaignsBatchRequest(id string, params GetOmegaCustomerTrxCampaignsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.InvoiceCampaign]] {
	return core.NewBatchRequest[core.Cursor[objects.InvoiceCampaign]](GetOmegaCustomerTrxCampaignsBatchCall(id, params, options...))
}

func DecodeGetOmegaCustomerTrxCampaignsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.InvoiceCampaign], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.InvoiceCampaign]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOmegaCustomerTrxCampaigns(ctx context.Context, client *core.Client, id string, params GetOmegaCustomerTrxCampaignsParams) (*core.Cursor[objects.InvoiceCampaign], error) {
	var out core.Cursor[objects.InvoiceCampaign]
	call := GetOmegaCustomerTrxCampaignsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetOmegaCustomerTrxParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOmegaCustomerTrxParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOmegaCustomerTrxBatchCall(id string, params GetOmegaCustomerTrxParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetOmegaCustomerTrxBatchRequest(id string, params GetOmegaCustomerTrxParams, options ...core.BatchOption) *core.BatchRequest[objects.OmegaCustomerTrx] {
	return core.NewBatchRequest[objects.OmegaCustomerTrx](GetOmegaCustomerTrxBatchCall(id, params, options...))
}

func DecodeGetOmegaCustomerTrxBatchResponse(response *core.BatchResponse) (*objects.OmegaCustomerTrx, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.OmegaCustomerTrx
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOmegaCustomerTrx(ctx context.Context, client *core.Client, id string, params GetOmegaCustomerTrxParams) (*objects.OmegaCustomerTrx, error) {
	var out objects.OmegaCustomerTrx
	call := GetOmegaCustomerTrxBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
