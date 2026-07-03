package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAnalyticsSegmentParams struct {
	AsyncTaskID *core.ID    `facebook:"async_task_id"`
	EndDate     *int        `facebook:"end_date"`
	StartDate   *int        `facebook:"start_date"`
	Extra       core.Params `facebook:"-"`
}

func (params GetAnalyticsSegmentParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AsyncTaskID != nil {
		out["async_task_id"] = *params.AsyncTaskID
	}
	if params.EndDate != nil {
		out["end_date"] = *params.EndDate
	}
	if params.StartDate != nil {
		out["start_date"] = *params.StartDate
	}
	return out
}

func GetAnalyticsSegmentBatchCall(id string, params GetAnalyticsSegmentParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAnalyticsSegmentBatchRequest(id string, params GetAnalyticsSegmentParams, options ...core.BatchOption) *core.BatchRequest[objects.AnalyticsSegment] {
	return core.NewBatchRequest[objects.AnalyticsSegment](GetAnalyticsSegmentBatchCall(id, params, options...))
}

func DecodeGetAnalyticsSegmentBatchResponse(response *core.BatchResponse) (*objects.AnalyticsSegment, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AnalyticsSegment
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAnalyticsSegmentWithResponse(ctx context.Context, client *core.Client, id string, params GetAnalyticsSegmentParams) (*objects.AnalyticsSegment, *core.Response, error) {
	var out objects.AnalyticsSegment
	call := GetAnalyticsSegmentBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAnalyticsSegment(ctx context.Context, client *core.Client, id string, params GetAnalyticsSegmentParams) (*objects.AnalyticsSegment, error) {
	out, _, err := GetAnalyticsSegmentWithResponse(ctx, client, id, params)
	return out, err
}
