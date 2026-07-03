package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetPaymentSubscriptionParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPaymentSubscriptionParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPaymentSubscriptionBatchCall(id string, params GetPaymentSubscriptionParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetPaymentSubscriptionBatchRequest(id string, params GetPaymentSubscriptionParams, options ...core.BatchOption) *core.BatchRequest[objects.PaymentSubscription] {
	return core.NewBatchRequest[objects.PaymentSubscription](GetPaymentSubscriptionBatchCall(id, params, options...))
}

func DecodeGetPaymentSubscriptionBatchResponse(response *core.BatchResponse) (*objects.PaymentSubscription, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PaymentSubscription
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPaymentSubscriptionWithResponse(ctx context.Context, client *core.Client, id string, params GetPaymentSubscriptionParams) (*objects.PaymentSubscription, *core.Response, error) {
	var out objects.PaymentSubscription
	call := GetPaymentSubscriptionBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPaymentSubscription(ctx context.Context, client *core.Client, id string, params GetPaymentSubscriptionParams) (*objects.PaymentSubscription, error) {
	out, _, err := GetPaymentSubscriptionWithResponse(ctx, client, id, params)
	return out, err
}
