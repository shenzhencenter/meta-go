package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetCloudGameParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCloudGameParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCloudGame(ctx context.Context, client *core.Client, id string, params GetCloudGameParams) (*objects.CloudGame, error) {
	var out objects.CloudGame
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
