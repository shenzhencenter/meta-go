package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdsUserSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsUserSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsUserSettingsBatchCall(id string, params GetAdsUserSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdsUserSettingsBatchRequest(id string, params GetAdsUserSettingsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdsUserSettings] {
	return core.NewBatchRequest[objects.AdsUserSettings](GetAdsUserSettingsBatchCall(id, params, options...))
}

func DecodeGetAdsUserSettingsBatchResponse(response *core.BatchResponse) (*objects.AdsUserSettings, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdsUserSettings
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdsUserSettings(ctx context.Context, client *core.Client, id string, params GetAdsUserSettingsParams) (*objects.AdsUserSettings, error) {
	var out objects.AdsUserSettings
	call := GetAdsUserSettingsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
