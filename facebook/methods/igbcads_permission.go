package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetIGBCAdsPermissionParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGBCAdsPermissionParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGBCAdsPermission(ctx context.Context, client *core.Client, id string, params GetIGBCAdsPermissionParams) (*objects.IGBCAdsPermission, error) {
	var out objects.IGBCAdsPermission
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
