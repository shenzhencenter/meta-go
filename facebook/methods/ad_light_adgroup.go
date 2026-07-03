package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetAdLightAdgroupParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdLightAdgroupParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdLightAdgroup(ctx context.Context, client *core.Client, id string, params GetAdLightAdgroupParams) (*objects.AdLightAdgroup, error) {
	var out objects.AdLightAdgroup
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
