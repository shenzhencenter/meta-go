package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
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

func GetCatalogWebsiteOnboardingSettingsBatchCall(id string, params GetCatalogWebsiteOnboardingSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCatalogWebsiteOnboardingSettingsBatchRequest(id string, params GetCatalogWebsiteOnboardingSettingsParams, options ...core.BatchOption) *core.BatchRequest[objects.CatalogWebsiteOnboardingSettings] {
	return core.NewBatchRequest[objects.CatalogWebsiteOnboardingSettings](GetCatalogWebsiteOnboardingSettingsBatchCall(id, params, options...))
}

func DecodeGetCatalogWebsiteOnboardingSettingsBatchResponse(response *core.BatchResponse) (*objects.CatalogWebsiteOnboardingSettings, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CatalogWebsiteOnboardingSettings
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCatalogWebsiteOnboardingSettings(ctx context.Context, client *core.Client, id string, params GetCatalogWebsiteOnboardingSettingsParams) (*objects.CatalogWebsiteOnboardingSettings, error) {
	var out objects.CatalogWebsiteOnboardingSettings
	call := GetCatalogWebsiteOnboardingSettingsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
