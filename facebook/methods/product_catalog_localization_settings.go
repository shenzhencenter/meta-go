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

func GetProductCatalogLocalizationSettingsBatchCall(id string, params GetProductCatalogLocalizationSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetProductCatalogLocalizationSettingsBatchRequest(id string, params GetProductCatalogLocalizationSettingsParams, options ...core.BatchOption) *core.BatchRequest[objects.ProductCatalogLocalizationSettings] {
	return core.NewBatchRequest[objects.ProductCatalogLocalizationSettings](GetProductCatalogLocalizationSettingsBatchCall(id, params, options...))
}

func DecodeGetProductCatalogLocalizationSettingsBatchResponse(response *core.BatchResponse) (*objects.ProductCatalogLocalizationSettings, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ProductCatalogLocalizationSettings
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductCatalogLocalizationSettingsWithResponse(ctx context.Context, client *core.Client, id string, params GetProductCatalogLocalizationSettingsParams) (*objects.ProductCatalogLocalizationSettings, *core.Response, error) {
	var out objects.ProductCatalogLocalizationSettings
	call := GetProductCatalogLocalizationSettingsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetProductCatalogLocalizationSettings(ctx context.Context, client *core.Client, id string, params GetProductCatalogLocalizationSettingsParams) (*objects.ProductCatalogLocalizationSettings, error) {
	out, _, err := GetProductCatalogLocalizationSettingsWithResponse(ctx, client, id, params)
	return out, err
}
