package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func GetOmegaCustomerTrxCampaigns(ctx context.Context, client *core.Client, id string, params GetOmegaCustomerTrxCampaignsParams) (*core.Cursor[objects.InvoiceCampaign], error) {
	var out core.Cursor[objects.InvoiceCampaign]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "campaigns"), params.ToParams(), &out)
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

func GetOmegaCustomerTrx(ctx context.Context, client *core.Client, id string, params GetOmegaCustomerTrxParams) (*objects.OmegaCustomerTrx, error) {
	var out objects.OmegaCustomerTrx
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
