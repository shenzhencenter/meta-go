package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type DeleteStoreCatalogSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteStoreCatalogSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteStoreCatalogSettings(ctx context.Context, client *core.Client, id string, params DeleteStoreCatalogSettingsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetStoreCatalogSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetStoreCatalogSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetStoreCatalogSettings(ctx context.Context, client *core.Client, id string, params GetStoreCatalogSettingsParams) (*objects.StoreCatalogSettings, error) {
	var out objects.StoreCatalogSettings
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
