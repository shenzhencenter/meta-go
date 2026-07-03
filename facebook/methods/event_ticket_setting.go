package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetEventTicketSettingParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetEventTicketSettingParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetEventTicketSettingBatchCall(id string, params GetEventTicketSettingParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetEventTicketSettingBatchRequest(id string, params GetEventTicketSettingParams, options ...core.BatchOption) *core.BatchRequest[objects.EventTicketSetting] {
	return core.NewBatchRequest[objects.EventTicketSetting](GetEventTicketSettingBatchCall(id, params, options...))
}

func DecodeGetEventTicketSettingBatchResponse(response *core.BatchResponse) (*objects.EventTicketSetting, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.EventTicketSetting
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetEventTicketSettingWithResponse(ctx context.Context, client *core.Client, id string, params GetEventTicketSettingParams) (*objects.EventTicketSetting, *core.Response, error) {
	var out objects.EventTicketSetting
	call := GetEventTicketSettingBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetEventTicketSetting(ctx context.Context, client *core.Client, id string, params GetEventTicketSettingParams) (*objects.EventTicketSetting, error) {
	out, _, err := GetEventTicketSettingWithResponse(ctx, client, id, params)
	return out, err
}
