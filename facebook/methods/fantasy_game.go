package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetFantasyGameParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetFantasyGameParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetFantasyGame(ctx context.Context, client *core.Client, id string, params GetFantasyGameParams) (*objects.FantasyGame, error) {
	var out objects.FantasyGame
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
