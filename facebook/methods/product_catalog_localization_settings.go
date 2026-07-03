package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetProductCatalogLocalizationSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductCatalogLocalizationSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductCatalogLocalizationSettings(ctx context.Context, client *core.Client, id string, params GetProductCatalogLocalizationSettingsParams) (*objects.ProductCatalogLocalizationSettings, error) {
	var out objects.ProductCatalogLocalizationSettings
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
