package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetAdvAInstanceParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdvAInstanceParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdvAInstance(ctx context.Context, client *core.Client, id string, params GetAdvAInstanceParams) (*objects.AdvAInstance, error) {
	var out objects.AdvAInstance
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
