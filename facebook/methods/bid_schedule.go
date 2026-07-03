package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetBidScheduleParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBidScheduleParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBidScheduleBatchCall(id string, params GetBidScheduleParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetBidScheduleBatchRequest(id string, params GetBidScheduleParams, options ...core.BatchOption) *core.BatchRequest[objects.BidSchedule] {
	return core.NewBatchRequest[objects.BidSchedule](GetBidScheduleBatchCall(id, params, options...))
}

func DecodeGetBidScheduleBatchResponse(response *core.BatchResponse) (*objects.BidSchedule, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BidSchedule
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBidScheduleWithResponse(ctx context.Context, client *core.Client, id string, params GetBidScheduleParams) (*objects.BidSchedule, *core.Response, error) {
	var out objects.BidSchedule
	call := GetBidScheduleBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetBidSchedule(ctx context.Context, client *core.Client, id string, params GetBidScheduleParams) (*objects.BidSchedule, error) {
	out, _, err := GetBidScheduleWithResponse(ctx, client, id, params)
	return out, err
}
