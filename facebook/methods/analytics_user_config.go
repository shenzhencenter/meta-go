package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAnalyticsUserConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAnalyticsUserConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAnalyticsUserConfig(ctx context.Context, client *core.Client, id string, params GetAnalyticsUserConfigParams) (*objects.AnalyticsUserConfig, error) {
	var out objects.AnalyticsUserConfig
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
