package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCatalogItemOverrideParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCatalogItemOverrideParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCatalogItemOverride(ctx context.Context, client *core.Client, id string, params GetCatalogItemOverrideParams) (*objects.CatalogItemOverride, error) {
	var out objects.CatalogItemOverride
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
