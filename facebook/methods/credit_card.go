package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCreditCardParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCreditCardParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCreditCardBatchCall(id string, params GetCreditCardParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCreditCardBatchRequest(id string, params GetCreditCardParams, options ...core.BatchOption) *core.BatchRequest[objects.CreditCard] {
	return core.NewBatchRequest[objects.CreditCard](GetCreditCardBatchCall(id, params, options...))
}

func DecodeGetCreditCardBatchResponse(response *core.BatchResponse) (*objects.CreditCard, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CreditCard
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCreditCard(ctx context.Context, client *core.Client, id string, params GetCreditCardParams) (*objects.CreditCard, error) {
	var out objects.CreditCard
	call := GetCreditCardBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
