package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdCustomDerivedMetricsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdCustomDerivedMetricsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdCustomDerivedMetrics(ctx context.Context, client *core.Client, id string, params GetAdCustomDerivedMetricsParams) (*objects.AdCustomDerivedMetrics, error) {
	var out objects.AdCustomDerivedMetrics
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
