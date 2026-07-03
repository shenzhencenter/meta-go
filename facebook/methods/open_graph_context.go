package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetOpenGraphContextParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOpenGraphContextParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOpenGraphContext(ctx context.Context, client *core.Client, id string, params GetOpenGraphContextParams) (*objects.OpenGraphContext, error) {
	var out objects.OpenGraphContext
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
