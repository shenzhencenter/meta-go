package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetCatalogWebsiteOnboardingSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCatalogWebsiteOnboardingSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCatalogWebsiteOnboardingSettings(ctx context.Context, client *core.Client, id string, params GetCatalogWebsiteOnboardingSettingsParams) (*objects.CatalogWebsiteOnboardingSettings, error) {
	var out objects.CatalogWebsiteOnboardingSettings
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
