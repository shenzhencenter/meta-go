package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetProductItemOfferParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductItemOfferParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductItemOfferBatchCall(id string, params GetProductItemOfferParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetProductItemOfferBatchRequest(id string, params GetProductItemOfferParams, options ...core.BatchOption) *core.BatchRequest[objects.ProductItemOffer] {
	return core.NewBatchRequest[objects.ProductItemOffer](GetProductItemOfferBatchCall(id, params, options...))
}

func DecodeGetProductItemOfferBatchResponse(response *core.BatchResponse) (*objects.ProductItemOffer, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ProductItemOffer
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductItemOfferWithResponse(ctx context.Context, client *core.Client, id string, params GetProductItemOfferParams) (*objects.ProductItemOffer, *core.Response, error) {
	var out objects.ProductItemOffer
	call := GetProductItemOfferBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetProductItemOffer(ctx context.Context, client *core.Client, id string, params GetProductItemOfferParams) (*objects.ProductItemOffer, error) {
	out, _, err := GetProductItemOfferWithResponse(ctx, client, id, params)
	return out, err
}
