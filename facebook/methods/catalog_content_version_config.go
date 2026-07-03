package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCatalogContentVersionConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCatalogContentVersionConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCatalogContentVersionConfig(ctx context.Context, client *core.Client, id string, params GetCatalogContentVersionConfigParams) (*objects.CatalogContentVersionConfig, error) {
	var out objects.CatalogContentVersionConfig
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
