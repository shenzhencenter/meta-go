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

func GetAnalyticsSegment(ctx context.Context, client *core.Client, id string, params GetAnalyticsSegmentParams) (*objects.AnalyticsSegment, error) {
	var out objects.AnalyticsSegment
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
