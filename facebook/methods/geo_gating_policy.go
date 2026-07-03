package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetGeoGatingPolicyParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetGeoGatingPolicyParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetGeoGatingPolicy(ctx context.Context, client *core.Client, id string, params GetGeoGatingPolicyParams) (*objects.GeoGatingPolicy, error) {
	var out objects.GeoGatingPolicy
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
