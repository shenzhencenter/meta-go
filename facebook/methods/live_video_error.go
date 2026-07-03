package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetLiveVideoErrorParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLiveVideoErrorParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLiveVideoError(ctx context.Context, client *core.Client, id string, params GetLiveVideoErrorParams) (*objects.LiveVideoError, error) {
	var out objects.LiveVideoError
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
