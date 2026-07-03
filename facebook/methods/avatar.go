package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAvatarParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAvatarParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAvatar(ctx context.Context, client *core.Client, id string, params GetAvatarParams) (*objects.Avatar, error) {
	var out objects.Avatar
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
