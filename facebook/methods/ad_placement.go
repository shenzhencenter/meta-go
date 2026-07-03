package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdPlacementParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdPlacementParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdPlacementBatchCall(id string, params GetAdPlacementParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdPlacementBatchRequest(id string, params GetAdPlacementParams, options ...core.BatchOption) *core.BatchRequest[objects.AdPlacement] {
	return core.NewBatchRequest[objects.AdPlacement](GetAdPlacementBatchCall(id, params, options...))
}

func DecodeGetAdPlacementBatchResponse(response *core.BatchResponse) (*objects.AdPlacement, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdPlacement
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdPlacementWithResponse(ctx context.Context, client *core.Client, id string, params GetAdPlacementParams) (*objects.AdPlacement, *core.Response, error) {
	var out objects.AdPlacement
	call := GetAdPlacementBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdPlacement(ctx context.Context, client *core.Client, id string, params GetAdPlacementParams) (*objects.AdPlacement, error) {
	out, _, err := GetAdPlacementWithResponse(ctx, client, id, params)
	return out, err
}
