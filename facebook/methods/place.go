package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetPlaceParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPlaceParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPlaceBatchCall(id string, params GetPlaceParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetPlaceBatchRequest(id string, params GetPlaceParams, options ...core.BatchOption) *core.BatchRequest[objects.Place] {
	return core.NewBatchRequest[objects.Place](GetPlaceBatchCall(id, params, options...))
}

func DecodeGetPlaceBatchResponse(response *core.BatchResponse) (*objects.Place, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Place
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPlace(ctx context.Context, client *core.Client, id string, params GetPlaceParams) (*objects.Place, error) {
	var out objects.Place
	call := GetPlaceBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
