package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetJobOpeningParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetJobOpeningParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetJobOpeningBatchCall(id string, params GetJobOpeningParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetJobOpeningBatchRequest(id string, params GetJobOpeningParams, options ...core.BatchOption) *core.BatchRequest[objects.JobOpening] {
	return core.NewBatchRequest[objects.JobOpening](GetJobOpeningBatchCall(id, params, options...))
}

func DecodeGetJobOpeningBatchResponse(response *core.BatchResponse) (*objects.JobOpening, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.JobOpening
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetJobOpeningWithResponse(ctx context.Context, client *core.Client, id string, params GetJobOpeningParams) (*objects.JobOpening, *core.Response, error) {
	var out objects.JobOpening
	call := GetJobOpeningBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetJobOpening(ctx context.Context, client *core.Client, id string, params GetJobOpeningParams) (*objects.JobOpening, error) {
	out, _, err := GetJobOpeningWithResponse(ctx, client, id, params)
	return out, err
}
