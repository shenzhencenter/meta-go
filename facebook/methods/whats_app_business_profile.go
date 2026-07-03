package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetWhatsAppBusinessProfileParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWhatsAppBusinessProfileParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWhatsAppBusinessProfileBatchCall(id string, params GetWhatsAppBusinessProfileParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetWhatsAppBusinessProfileBatchRequest(id string, params GetWhatsAppBusinessProfileParams, options ...core.BatchOption) *core.BatchRequest[objects.WhatsAppBusinessProfile] {
	return core.NewBatchRequest[objects.WhatsAppBusinessProfile](GetWhatsAppBusinessProfileBatchCall(id, params, options...))
}

func DecodeGetWhatsAppBusinessProfileBatchResponse(response *core.BatchResponse) (*objects.WhatsAppBusinessProfile, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WhatsAppBusinessProfile
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWhatsAppBusinessProfile(ctx context.Context, client *core.Client, id string, params GetWhatsAppBusinessProfileParams) (*objects.WhatsAppBusinessProfile, error) {
	var out objects.WhatsAppBusinessProfile
	call := GetWhatsAppBusinessProfileBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type UpdateWhatsAppBusinessProfileParams struct {
	Extra core.Params `facebook:"-"`
}

func (params UpdateWhatsAppBusinessProfileParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func UpdateWhatsAppBusinessProfileBatchCall(id string, params UpdateWhatsAppBusinessProfileParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateWhatsAppBusinessProfileBatchRequest(id string, params UpdateWhatsAppBusinessProfileParams, options ...core.BatchOption) *core.BatchRequest[objects.WhatsAppBusinessProfile] {
	return core.NewBatchRequest[objects.WhatsAppBusinessProfile](UpdateWhatsAppBusinessProfileBatchCall(id, params, options...))
}

func DecodeUpdateWhatsAppBusinessProfileBatchResponse(response *core.BatchResponse) (*objects.WhatsAppBusinessProfile, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WhatsAppBusinessProfile
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateWhatsAppBusinessProfile(ctx context.Context, client *core.Client, id string, params UpdateWhatsAppBusinessProfileParams) (*objects.WhatsAppBusinessProfile, error) {
	var out objects.WhatsAppBusinessProfile
	call := UpdateWhatsAppBusinessProfileBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
