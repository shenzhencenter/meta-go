package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetALMEventParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetALMEventParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetALMEvent(ctx context.Context, client *core.Client, id string, params GetALMEventParams) (*objects.ALMEvent, error) {
	var out objects.ALMEvent
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
