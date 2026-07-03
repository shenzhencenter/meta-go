package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetBusinessRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessRequest(ctx context.Context, client *core.Client, id string, params GetBusinessRequestParams) (*objects.BusinessRequest, error) {
	var out objects.BusinessRequest
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
