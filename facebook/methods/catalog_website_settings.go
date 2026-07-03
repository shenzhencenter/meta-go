package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCatalogWebsiteSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCatalogWebsiteSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCatalogWebsiteSettings(ctx context.Context, client *core.Client, id string, params GetCatalogWebsiteSettingsParams) (*objects.CatalogWebsiteSettings, error) {
	var out objects.CatalogWebsiteSettings
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
