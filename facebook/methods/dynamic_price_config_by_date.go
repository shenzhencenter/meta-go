package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetDynamicPriceConfigByDateParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetDynamicPriceConfigByDateParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetDynamicPriceConfigByDate(ctx context.Context, client *core.Client, id string, params GetDynamicPriceConfigByDateParams) (*objects.DynamicPriceConfigByDate, error) {
	var out objects.DynamicPriceConfigByDate
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
