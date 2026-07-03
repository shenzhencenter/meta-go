package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type CreatePaymentEnginePaymentDisputeParams struct {
	Reason enums.PaymentenginepaymentdisputeReasonEnumParam `facebook:"reason"`
	Extra  core.Params                                      `facebook:"-"`
}

func (params CreatePaymentEnginePaymentDisputeParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["reason"] = params.Reason
	return out
}

func CreatePaymentEnginePaymentDisputeBatchCall(id string, params CreatePaymentEnginePaymentDisputeParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "dispute"), params.ToParams(), options...)
}

func NewCreatePaymentEnginePaymentDisputeBatchRequest(id string, params CreatePaymentEnginePaymentDisputeParams, options ...core.BatchOption) *core.BatchRequest[objects.PaymentEnginePayment] {
	return core.NewBatchRequest[objects.PaymentEnginePayment](CreatePaymentEnginePaymentDisputeBatchCall(id, params, options...))
}

func DecodeCreatePaymentEnginePaymentDisputeBatchResponse(response *core.BatchResponse) (*objects.PaymentEnginePayment, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PaymentEnginePayment
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePaymentEnginePaymentDisputeWithResponse(ctx context.Context, client *core.Client, id string, params CreatePaymentEnginePaymentDisputeParams) (*objects.PaymentEnginePayment, *core.Response, error) {
	var out objects.PaymentEnginePayment
	call := CreatePaymentEnginePaymentDisputeBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePaymentEnginePaymentDispute(ctx context.Context, client *core.Client, id string, params CreatePaymentEnginePaymentDisputeParams) (*objects.PaymentEnginePayment, error) {
	out, _, err := CreatePaymentEnginePaymentDisputeWithResponse(ctx, client, id, params)
	return out, err
}

type CreatePaymentEnginePaymentRefundsParams struct {
	Amount   float64                                           `facebook:"amount"`
	Currency string                                            `facebook:"currency"`
	Reason   *enums.PaymentenginepaymentrefundsReasonEnumParam `facebook:"reason"`
	Extra    core.Params                                       `facebook:"-"`
}

func (params CreatePaymentEnginePaymentRefundsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["amount"] = params.Amount
	out["currency"] = params.Currency
	if params.Reason != nil {
		out["reason"] = *params.Reason
	}
	return out
}

func CreatePaymentEnginePaymentRefundsBatchCall(id string, params CreatePaymentEnginePaymentRefundsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "refunds"), params.ToParams(), options...)
}

func NewCreatePaymentEnginePaymentRefundsBatchRequest(id string, params CreatePaymentEnginePaymentRefundsParams, options ...core.BatchOption) *core.BatchRequest[objects.PaymentEnginePayment] {
	return core.NewBatchRequest[objects.PaymentEnginePayment](CreatePaymentEnginePaymentRefundsBatchCall(id, params, options...))
}

func DecodeCreatePaymentEnginePaymentRefundsBatchResponse(response *core.BatchResponse) (*objects.PaymentEnginePayment, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PaymentEnginePayment
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreatePaymentEnginePaymentRefundsWithResponse(ctx context.Context, client *core.Client, id string, params CreatePaymentEnginePaymentRefundsParams) (*objects.PaymentEnginePayment, *core.Response, error) {
	var out objects.PaymentEnginePayment
	call := CreatePaymentEnginePaymentRefundsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreatePaymentEnginePaymentRefunds(ctx context.Context, client *core.Client, id string, params CreatePaymentEnginePaymentRefundsParams) (*objects.PaymentEnginePayment, error) {
	out, _, err := CreatePaymentEnginePaymentRefundsWithResponse(ctx, client, id, params)
	return out, err
}

type GetPaymentEnginePaymentParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPaymentEnginePaymentParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPaymentEnginePaymentBatchCall(id string, params GetPaymentEnginePaymentParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetPaymentEnginePaymentBatchRequest(id string, params GetPaymentEnginePaymentParams, options ...core.BatchOption) *core.BatchRequest[objects.PaymentEnginePayment] {
	return core.NewBatchRequest[objects.PaymentEnginePayment](GetPaymentEnginePaymentBatchCall(id, params, options...))
}

func DecodeGetPaymentEnginePaymentBatchResponse(response *core.BatchResponse) (*objects.PaymentEnginePayment, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PaymentEnginePayment
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPaymentEnginePaymentWithResponse(ctx context.Context, client *core.Client, id string, params GetPaymentEnginePaymentParams) (*objects.PaymentEnginePayment, *core.Response, error) {
	var out objects.PaymentEnginePayment
	call := GetPaymentEnginePaymentBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPaymentEnginePayment(ctx context.Context, client *core.Client, id string, params GetPaymentEnginePaymentParams) (*objects.PaymentEnginePayment, error) {
	out, _, err := GetPaymentEnginePaymentWithResponse(ctx, client, id, params)
	return out, err
}
