package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdAccountUserSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountUserSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountUserSettingsBatchCall(id string, params GetAdAccountUserSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdAccountUserSettingsBatchRequest(id string, params GetAdAccountUserSettingsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdAccountUserSettings] {
	return core.NewBatchRequest[objects.AdAccountUserSettings](GetAdAccountUserSettingsBatchCall(id, params, options...))
}

func DecodeGetAdAccountUserSettingsBatchResponse(response *core.BatchResponse) (*objects.AdAccountUserSettings, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdAccountUserSettings
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAccountUserSettings(ctx context.Context, client *core.Client, id string, params GetAdAccountUserSettingsParams) (*objects.AdAccountUserSettings, error) {
	var out objects.AdAccountUserSettings
	call := GetAdAccountUserSettingsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
