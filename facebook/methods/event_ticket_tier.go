package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetEventTicketTierParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetEventTicketTierParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetEventTicketTierBatchCall(id string, params GetEventTicketTierParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetEventTicketTierBatchRequest(id string, params GetEventTicketTierParams, options ...core.BatchOption) *core.BatchRequest[objects.EventTicketTier] {
	return core.NewBatchRequest[objects.EventTicketTier](GetEventTicketTierBatchCall(id, params, options...))
}

func DecodeGetEventTicketTierBatchResponse(response *core.BatchResponse) (*objects.EventTicketTier, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.EventTicketTier
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetEventTicketTierWithResponse(ctx context.Context, client *core.Client, id string, params GetEventTicketTierParams) (*objects.EventTicketTier, *core.Response, error) {
	var out objects.EventTicketTier
	call := GetEventTicketTierBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetEventTicketTier(ctx context.Context, client *core.Client, id string, params GetEventTicketTierParams) (*objects.EventTicketTier, error) {
	out, _, err := GetEventTicketTierWithResponse(ctx, client, id, params)
	return out, err
}
