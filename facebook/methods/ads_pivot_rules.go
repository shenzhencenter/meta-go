package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetAdsPivotRulesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsPivotRulesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsPivotRules(ctx context.Context, client *core.Client, id string, params GetAdsPivotRulesParams) (*objects.AdsPivotRules, error) {
	var out objects.AdsPivotRules
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
