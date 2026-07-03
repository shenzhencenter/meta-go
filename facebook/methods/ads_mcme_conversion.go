package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdsMcmeConversionParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsMcmeConversionParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsMcmeConversion(ctx context.Context, client *core.Client, id string, params GetAdsMcmeConversionParams) (*objects.AdsMcmeConversion, error) {
	var out objects.AdsMcmeConversion
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
