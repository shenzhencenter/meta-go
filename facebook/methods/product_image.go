package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetProductImageParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductImageParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductImageBatchCall(id string, params GetProductImageParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetProductImageBatchRequest(id string, params GetProductImageParams, options ...core.BatchOption) *core.BatchRequest[objects.ProductImage] {
	return core.NewBatchRequest[objects.ProductImage](GetProductImageBatchCall(id, params, options...))
}

func DecodeGetProductImageBatchResponse(response *core.BatchResponse) (*objects.ProductImage, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ProductImage
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductImage(ctx context.Context, client *core.Client, id string, params GetProductImageParams) (*objects.ProductImage, error) {
	var out objects.ProductImage
	call := GetProductImageBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
