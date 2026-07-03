package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetShopParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetShopParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetShop(ctx context.Context, client *core.Client, id string, params GetShopParams) (*objects.Shop, error) {
	var out objects.Shop
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
