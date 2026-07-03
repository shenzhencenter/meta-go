package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCatalogSmartPixelSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCatalogSmartPixelSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCatalogSmartPixelSettings(ctx context.Context, client *core.Client, id string, params GetCatalogSmartPixelSettingsParams) (*objects.CatalogSmartPixelSettings, error) {
	var out objects.CatalogSmartPixelSettings
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
