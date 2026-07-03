package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetBusinessProjectParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessProjectParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessProject(ctx context.Context, client *core.Client, id string, params GetBusinessProjectParams) (*objects.BusinessProject, error) {
	var out objects.BusinessProject
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
