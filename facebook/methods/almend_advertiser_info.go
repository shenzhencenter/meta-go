package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetALMEndAdvertiserInfoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetALMEndAdvertiserInfoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetALMEndAdvertiserInfo(ctx context.Context, client *core.Client, id string, params GetALMEndAdvertiserInfoParams) (*objects.ALMEndAdvertiserInfo, error) {
	var out objects.ALMEndAdvertiserInfo
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
