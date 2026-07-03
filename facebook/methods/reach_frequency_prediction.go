package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func GetReachFrequencyPrediction(ctx context.Context, client *core.Client, id string, params GetReachFrequencyPredictionParams) (*objects.ReachFrequencyPrediction, error) {
	var out objects.ReachFrequencyPrediction
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
