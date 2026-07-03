package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetBrandRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBrandRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBrandRequest(ctx context.Context, client *core.Client, id string, params GetBrandRequestParams) (*objects.BrandRequest, error) {
	var out objects.BrandRequest
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
