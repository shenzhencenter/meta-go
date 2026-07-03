package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetUserPageOneTimeOptInTokenSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserPageOneTimeOptInTokenSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserPageOneTimeOptInTokenSettingsBatchCall(id string, params GetUserPageOneTimeOptInTokenSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetUserPageOneTimeOptInTokenSettingsBatchRequest(id string, params GetUserPageOneTimeOptInTokenSettingsParams, options ...core.BatchOption) *core.BatchRequest[objects.UserPageOneTimeOptInTokenSettings] {
	return core.NewBatchRequest[objects.UserPageOneTimeOptInTokenSettings](GetUserPageOneTimeOptInTokenSettingsBatchCall(id, params, options...))
}

func DecodeGetUserPageOneTimeOptInTokenSettingsBatchResponse(response *core.BatchResponse) (*objects.UserPageOneTimeOptInTokenSettings, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.UserPageOneTimeOptInTokenSettings
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetUserPageOneTimeOptInTokenSettingsWithResponse(ctx context.Context, client *core.Client, id string, params GetUserPageOneTimeOptInTokenSettingsParams) (*objects.UserPageOneTimeOptInTokenSettings, *core.Response, error) {
	var out objects.UserPageOneTimeOptInTokenSettings
	call := GetUserPageOneTimeOptInTokenSettingsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetUserPageOneTimeOptInTokenSettings(ctx context.Context, client *core.Client, id string, params GetUserPageOneTimeOptInTokenSettingsParams) (*objects.UserPageOneTimeOptInTokenSettings, error) {
	out, _, err := GetUserPageOneTimeOptInTokenSettingsWithResponse(ctx, client, id, params)
	return out, err
}
