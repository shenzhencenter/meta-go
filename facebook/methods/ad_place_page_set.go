package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdPlacePageSetParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdPlacePageSetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdPlacePageSetBatchCall(id string, params GetAdPlacePageSetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdPlacePageSetBatchRequest(id string, params GetAdPlacePageSetParams, options ...core.BatchOption) *core.BatchRequest[objects.AdPlacePageSet] {
	return core.NewBatchRequest[objects.AdPlacePageSet](GetAdPlacePageSetBatchCall(id, params, options...))
}

func DecodeGetAdPlacePageSetBatchResponse(response *core.BatchResponse) (*objects.AdPlacePageSet, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdPlacePageSet
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdPlacePageSetWithResponse(ctx context.Context, client *core.Client, id string, params GetAdPlacePageSetParams) (*objects.AdPlacePageSet, *core.Response, error) {
	var out objects.AdPlacePageSet
	call := GetAdPlacePageSetBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdPlacePageSet(ctx context.Context, client *core.Client, id string, params GetAdPlacePageSetParams) (*objects.AdPlacePageSet, error) {
	out, _, err := GetAdPlacePageSetWithResponse(ctx, client, id, params)
	return out, err
}
