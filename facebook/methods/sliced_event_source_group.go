package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetSlicedEventSourceGroupParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetSlicedEventSourceGroupParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetSlicedEventSourceGroupBatchCall(id string, params GetSlicedEventSourceGroupParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetSlicedEventSourceGroupBatchRequest(id string, params GetSlicedEventSourceGroupParams, options ...core.BatchOption) *core.BatchRequest[objects.SlicedEventSourceGroup] {
	return core.NewBatchRequest[objects.SlicedEventSourceGroup](GetSlicedEventSourceGroupBatchCall(id, params, options...))
}

func DecodeGetSlicedEventSourceGroupBatchResponse(response *core.BatchResponse) (*objects.SlicedEventSourceGroup, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.SlicedEventSourceGroup
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetSlicedEventSourceGroup(ctx context.Context, client *core.Client, id string, params GetSlicedEventSourceGroupParams) (*objects.SlicedEventSourceGroup, error) {
	var out objects.SlicedEventSourceGroup
	call := GetSlicedEventSourceGroupBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
