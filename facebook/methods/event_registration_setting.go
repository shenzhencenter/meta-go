package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetEventRegistrationSettingParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetEventRegistrationSettingParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetEventRegistrationSettingBatchCall(id string, params GetEventRegistrationSettingParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetEventRegistrationSettingBatchRequest(id string, params GetEventRegistrationSettingParams, options ...core.BatchOption) *core.BatchRequest[objects.EventRegistrationSetting] {
	return core.NewBatchRequest[objects.EventRegistrationSetting](GetEventRegistrationSettingBatchCall(id, params, options...))
}

func DecodeGetEventRegistrationSettingBatchResponse(response *core.BatchResponse) (*objects.EventRegistrationSetting, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.EventRegistrationSetting
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetEventRegistrationSettingWithResponse(ctx context.Context, client *core.Client, id string, params GetEventRegistrationSettingParams) (*objects.EventRegistrationSetting, *core.Response, error) {
	var out objects.EventRegistrationSetting
	call := GetEventRegistrationSettingBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetEventRegistrationSetting(ctx context.Context, client *core.Client, id string, params GetEventRegistrationSettingParams) (*objects.EventRegistrationSetting, error) {
	out, _, err := GetEventRegistrationSettingWithResponse(ctx, client, id, params)
	return out, err
}
