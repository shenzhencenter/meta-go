package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetShadowIGMediaBuilderParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetShadowIGMediaBuilderParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetShadowIGMediaBuilder(ctx context.Context, client *core.Client, id string, params GetShadowIGMediaBuilderParams) (*objects.ShadowIGMediaBuilder, error) {
	var out objects.ShadowIGMediaBuilder
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
