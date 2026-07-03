package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdRuleExecutionSpecParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdRuleExecutionSpecParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdRuleExecutionSpec(ctx context.Context, client *core.Client, id string, params GetAdRuleExecutionSpecParams) (*objects.AdRuleExecutionSpec, error) {
	var out objects.AdRuleExecutionSpec
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
