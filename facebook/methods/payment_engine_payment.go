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

func CreatePaymentEnginePaymentDispute(ctx context.Context, client *core.Client, id string, params CreatePaymentEnginePaymentDisputeParams) (*objects.PaymentEnginePayment, error) {
	var out objects.PaymentEnginePayment
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "dispute"), params.ToParams(), &out)
	return &out, err
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

func CreatePaymentEnginePaymentRefunds(ctx context.Context, client *core.Client, id string, params CreatePaymentEnginePaymentRefundsParams) (*objects.PaymentEnginePayment, error) {
	var out objects.PaymentEnginePayment
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "refunds"), params.ToParams(), &out)
	return &out, err
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

func GetPaymentEnginePayment(ctx context.Context, client *core.Client, id string, params GetPaymentEnginePaymentParams) (*objects.PaymentEnginePayment, error) {
	var out objects.PaymentEnginePayment
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
