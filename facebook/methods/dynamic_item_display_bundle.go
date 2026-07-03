package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetDynamicItemDisplayBundleParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetDynamicItemDisplayBundleParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetDynamicItemDisplayBundle(ctx context.Context, client *core.Client, id string, params GetDynamicItemDisplayBundleParams) (*objects.DynamicItemDisplayBundle, error) {
	var out objects.DynamicItemDisplayBundle
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
