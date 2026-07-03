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

func GetHoursWithResponse(ctx context.Context, client *core.Client, id string, params GetHoursParams) (*objects.Hours, *core.Response, error) {
	var out objects.Hours
	call := GetHoursBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetHours(ctx context.Context, client *core.Client, id string, params GetHoursParams) (*objects.Hours, error) {
	out, _, err := GetHoursWithResponse(ctx, client, id, params)
	return out, err
}
