package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetWoodhengePurchasedPAYGReceiptParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWoodhengePurchasedPAYGReceiptParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWoodhengePurchasedPAYGReceipt(ctx context.Context, client *core.Client, id string, params GetWoodhengePurchasedPAYGReceiptParams) (*objects.WoodhengePurchasedPAYGReceipt, error) {
	var out objects.WoodhengePurchasedPAYGReceipt
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
