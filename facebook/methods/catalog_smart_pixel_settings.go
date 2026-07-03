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

func GetCatalogSmartPixelSettingsBatchCall(id string, params GetCatalogSmartPixelSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCatalogSmartPixelSettingsBatchRequest(id string, params GetCatalogSmartPixelSettingsParams, options ...core.BatchOption) *core.BatchRequest[objects.CatalogSmartPixelSettings] {
	return core.NewBatchRequest[objects.CatalogSmartPixelSettings](GetCatalogSmartPixelSettingsBatchCall(id, params, options...))
}

func DecodeGetCatalogSmartPixelSettingsBatchResponse(response *core.BatchResponse) (*objects.CatalogSmartPixelSettings, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CatalogSmartPixelSettings
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCatalogSmartPixelSettingsWithResponse(ctx context.Context, client *core.Client, id string, params GetCatalogSmartPixelSettingsParams) (*objects.CatalogSmartPixelSettings, *core.Response, error) {
	var out objects.CatalogSmartPixelSettings
	call := GetCatalogSmartPixelSettingsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCatalogSmartPixelSettings(ctx context.Context, client *core.Client, id string, params GetCatalogSmartPixelSettingsParams) (*objects.CatalogSmartPixelSettings, error) {
	out, _, err := GetCatalogSmartPixelSettingsWithResponse(ctx, client, id, params)
	return out, err
}
