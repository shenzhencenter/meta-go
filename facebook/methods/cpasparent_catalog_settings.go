package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCPASParentCatalogSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCPASParentCatalogSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCPASParentCatalogSettingsBatchCall(id string, params GetCPASParentCatalogSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCPASParentCatalogSettingsBatchRequest(id string, params GetCPASParentCatalogSettingsParams, options ...core.BatchOption) *core.BatchRequest[objects.CPASParentCatalogSettings] {
	return core.NewBatchRequest[objects.CPASParentCatalogSettings](GetCPASParentCatalogSettingsBatchCall(id, params, options...))
}

func DecodeGetCPASParentCatalogSettingsBatchResponse(response *core.BatchResponse) (*objects.CPASParentCatalogSettings, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CPASParentCatalogSettings
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCPASParentCatalogSettingsWithResponse(ctx context.Context, client *core.Client, id string, params GetCPASParentCatalogSettingsParams) (*objects.CPASParentCatalogSettings, *core.Response, error) {
	var out objects.CPASParentCatalogSettings
	call := GetCPASParentCatalogSettingsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCPASParentCatalogSettings(ctx context.Context, client *core.Client, id string, params GetCPASParentCatalogSettingsParams) (*objects.CPASParentCatalogSettings, error) {
	out, _, err := GetCPASParentCatalogSettingsWithResponse(ctx, client, id, params)
	return out, err
}
