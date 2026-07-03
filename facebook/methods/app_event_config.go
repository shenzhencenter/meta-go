package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAppEventConfigParams struct {
	EventName *string     `facebook:"event_name"`
	Extra     core.Params `facebook:"-"`
}

func (params GetAppEventConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.EventName != nil {
		out["event_name"] = *params.EventName
	}
	return out
}

func GetAppEventConfig(ctx context.Context, client *core.Client, id string, params GetAppEventConfigParams) (*objects.AppEventConfig, error) {
	var out objects.AppEventConfig
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
