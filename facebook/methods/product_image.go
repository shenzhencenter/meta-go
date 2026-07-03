package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetProductImageParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductImageParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductImage(ctx context.Context, client *core.Client, id string, params GetProductImageParams) (*objects.ProductImage, error) {
	var out objects.ProductImage
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
