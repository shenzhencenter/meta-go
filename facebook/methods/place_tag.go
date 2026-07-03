package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetPlaceTagParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPlaceTagParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPlaceTagBatchCall(id string, params GetPlaceTagParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetPlaceTagBatchRequest(id string, params GetPlaceTagParams, options ...core.BatchOption) *core.BatchRequest[objects.PlaceTag] {
	return core.NewBatchRequest[objects.PlaceTag](GetPlaceTagBatchCall(id, params, options...))
}

func DecodeGetPlaceTagBatchResponse(response *core.BatchResponse) (*objects.PlaceTag, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PlaceTag
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPlaceTag(ctx context.Context, client *core.Client, id string, params GetPlaceTagParams) (*objects.PlaceTag, error) {
	var out objects.PlaceTag
	call := GetPlaceTagBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
