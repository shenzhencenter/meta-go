package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetHoursParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetHoursParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetHoursBatchCall(id string, params GetHoursParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetHoursBatchRequest(id string, params GetHoursParams, options ...core.BatchOption) *core.BatchRequest[objects.Hours] {
	return core.NewBatchRequest[objects.Hours](GetHoursBatchCall(id, params, options...))
}

func DecodeGetHoursBatchResponse(response *core.BatchResponse) (*objects.Hours, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Hours
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetHours(ctx context.Context, client *core.Client, id string, params GetHoursParams) (*objects.Hours, error) {
	var out objects.Hours
	call := GetHoursBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
