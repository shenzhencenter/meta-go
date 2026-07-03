package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCTXPartnerAppWelcomeMessageFlowParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCTXPartnerAppWelcomeMessageFlowParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCTXPartnerAppWelcomeMessageFlowBatchCall(id string, params GetCTXPartnerAppWelcomeMessageFlowParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCTXPartnerAppWelcomeMessageFlowBatchRequest(id string, params GetCTXPartnerAppWelcomeMessageFlowParams, options ...core.BatchOption) *core.BatchRequest[objects.CTXPartnerAppWelcomeMessageFlow] {
	return core.NewBatchRequest[objects.CTXPartnerAppWelcomeMessageFlow](GetCTXPartnerAppWelcomeMessageFlowBatchCall(id, params, options...))
}

func DecodeGetCTXPartnerAppWelcomeMessageFlowBatchResponse(response *core.BatchResponse) (*objects.CTXPartnerAppWelcomeMessageFlow, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CTXPartnerAppWelcomeMessageFlow
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCTXPartnerAppWelcomeMessageFlow(ctx context.Context, client *core.Client, id string, params GetCTXPartnerAppWelcomeMessageFlowParams) (*objects.CTXPartnerAppWelcomeMessageFlow, error) {
	var out objects.CTXPartnerAppWelcomeMessageFlow
	call := GetCTXPartnerAppWelcomeMessageFlowBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
