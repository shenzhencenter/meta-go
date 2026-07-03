package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetShopParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetShopParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetShopBatchCall(id string, params GetShopParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetShopBatchRequest(id string, params GetShopParams, options ...core.BatchOption) *core.BatchRequest[objects.Shop] {
	return core.NewBatchRequest[objects.Shop](GetShopBatchCall(id, params, options...))
}

func DecodeGetShopBatchResponse(response *core.BatchResponse) (*objects.Shop, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Shop
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetShop(ctx context.Context, client *core.Client, id string, params GetShopParams) (*objects.Shop, error) {
	var out objects.Shop
	call := GetShopBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
