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

func DeleteStoreCatalogSettingsBatchCall(id string, params DeleteStoreCatalogSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteStoreCatalogSettingsBatchRequest(id string, params DeleteStoreCatalogSettingsParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteStoreCatalogSettingsBatchCall(id, params, options...))
}

func DecodeDeleteStoreCatalogSettingsBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func DeleteStoreCatalogSettings(ctx context.Context, client *core.Client, id string, params DeleteStoreCatalogSettingsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteStoreCatalogSettingsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetStoreCatalogSettingsBatchCall(id string, params GetStoreCatalogSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetStoreCatalogSettingsBatchRequest(id string, params GetStoreCatalogSettingsParams, options ...core.BatchOption) *core.BatchRequest[objects.StoreCatalogSettings] {
	return core.NewBatchRequest[objects.StoreCatalogSettings](GetStoreCatalogSettingsBatchCall(id, params, options...))
}

func DecodeGetStoreCatalogSettingsBatchResponse(response *core.BatchResponse) (*objects.StoreCatalogSettings, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.StoreCatalogSettings
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetStoreCatalogSettings(ctx context.Context, client *core.Client, id string, params GetStoreCatalogSettingsParams) (*objects.StoreCatalogSettings, error) {
	var out objects.StoreCatalogSettings
	call := GetStoreCatalogSettingsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
