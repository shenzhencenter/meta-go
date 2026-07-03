package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetExternalMerchantSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetExternalMerchantSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetExternalMerchantSettingsBatchCall(id string, params GetExternalMerchantSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetExternalMerchantSettingsBatchRequest(id string, params GetExternalMerchantSettingsParams, options ...core.BatchOption) *core.BatchRequest[objects.ExternalMerchantSettings] {
	return core.NewBatchRequest[objects.ExternalMerchantSettings](GetExternalMerchantSettingsBatchCall(id, params, options...))
}

func DecodeGetExternalMerchantSettingsBatchResponse(response *core.BatchResponse) (*objects.ExternalMerchantSettings, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ExternalMerchantSettings
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetExternalMerchantSettings(ctx context.Context, client *core.Client, id string, params GetExternalMerchantSettingsParams) (*objects.ExternalMerchantSettings, error) {
	var out objects.ExternalMerchantSettings
	call := GetExternalMerchantSettingsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
