package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetProductItemOfferParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductItemOfferParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductItemOffer(ctx context.Context, client *core.Client, id string, params GetProductItemOfferParams) (*objects.ProductItemOffer, error) {
	var out objects.ProductItemOffer
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
