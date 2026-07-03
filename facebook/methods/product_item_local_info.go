package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetProductItemLocalInfoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductItemLocalInfoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductItemLocalInfo(ctx context.Context, client *core.Client, id string, params GetProductItemLocalInfoParams) (*objects.ProductItemLocalInfo, error) {
	var out objects.ProductItemLocalInfo
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
