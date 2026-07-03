package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetAdColumnSizesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdColumnSizesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdColumnSizes(ctx context.Context, client *core.Client, id string, params GetAdColumnSizesParams) (*objects.AdColumnSizes, error) {
	var out objects.AdColumnSizes
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
