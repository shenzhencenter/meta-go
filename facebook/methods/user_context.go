package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetUserContextParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserContextParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserContext(ctx context.Context, client *core.Client, id string, params GetUserContextParams) (*objects.UserContext, error) {
	var out objects.UserContext
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
