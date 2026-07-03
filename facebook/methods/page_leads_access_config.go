package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetPageLeadsAccessConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageLeadsAccessConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageLeadsAccessConfig(ctx context.Context, client *core.Client, id string, params GetPageLeadsAccessConfigParams) (*objects.PageLeadsAccessConfig, error) {
	var out objects.PageLeadsAccessConfig
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
