package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetReachFrequencyPredictionParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetReachFrequencyPredictionParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetReachFrequencyPredictionBatchCall(id string, params GetReachFrequencyPredictionParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetReachFrequencyPredictionBatchRequest(id string, params GetReachFrequencyPredictionParams, options ...core.BatchOption) *core.BatchRequest[objects.ReachFrequencyPrediction] {
	return core.NewBatchRequest[objects.ReachFrequencyPrediction](GetReachFrequencyPredictionBatchCall(id, params, options...))
}

func DecodeGetReachFrequencyPredictionBatchResponse(response *core.BatchResponse) (*objects.ReachFrequencyPrediction, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ReachFrequencyPrediction
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetReachFrequencyPrediction(ctx context.Context, client *core.Client, id string, params GetReachFrequencyPredictionParams) (*objects.ReachFrequencyPrediction, error) {
	var out objects.ReachFrequencyPrediction
	call := GetReachFrequencyPredictionBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
