package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetProductFeedScheduleParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductFeedScheduleParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductFeedScheduleBatchCall(id string, params GetProductFeedScheduleParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetProductFeedScheduleBatchRequest(id string, params GetProductFeedScheduleParams, options ...core.BatchOption) *core.BatchRequest[objects.ProductFeedSchedule] {
	return core.NewBatchRequest[objects.ProductFeedSchedule](GetProductFeedScheduleBatchCall(id, params, options...))
}

func DecodeGetProductFeedScheduleBatchResponse(response *core.BatchResponse) (*objects.ProductFeedSchedule, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ProductFeedSchedule
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductFeedSchedule(ctx context.Context, client *core.Client, id string, params GetProductFeedScheduleParams) (*objects.ProductFeedSchedule, error) {
	var out objects.ProductFeedSchedule
	call := GetProductFeedScheduleBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
