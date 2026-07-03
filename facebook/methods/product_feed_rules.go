package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetProductFeedRulesRulesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductFeedRulesRulesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductFeedRulesRules(ctx context.Context, client *core.Client, id string, params GetProductFeedRulesRulesParams) (*core.Cursor[objects.ProductFeedRulesGet], error) {
	var out core.Cursor[objects.ProductFeedRulesGet]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "rules"), params.ToParams(), &out)
	return &out, err
}
