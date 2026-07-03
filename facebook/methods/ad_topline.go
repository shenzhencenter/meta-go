package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetAdToplineParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdToplineParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdTopline(ctx context.Context, client *core.Client, id string, params GetAdToplineParams) (*objects.AdTopline, error) {
	var out objects.AdTopline
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
