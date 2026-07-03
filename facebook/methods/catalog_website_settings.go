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

func GetCatalogWebsiteSettingsBatchCall(id string, params GetCatalogWebsiteSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCatalogWebsiteSettingsBatchRequest(id string, params GetCatalogWebsiteSettingsParams, options ...core.BatchOption) *core.BatchRequest[objects.CatalogWebsiteSettings] {
	return core.NewBatchRequest[objects.CatalogWebsiteSettings](GetCatalogWebsiteSettingsBatchCall(id, params, options...))
}

func DecodeGetCatalogWebsiteSettingsBatchResponse(response *core.BatchResponse) (*objects.CatalogWebsiteSettings, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CatalogWebsiteSettings
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCatalogWebsiteSettings(ctx context.Context, client *core.Client, id string, params GetCatalogWebsiteSettingsParams) (*objects.CatalogWebsiteSettings, error) {
	var out objects.CatalogWebsiteSettings
	call := GetCatalogWebsiteSettingsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
