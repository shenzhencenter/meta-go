package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetAdRuleEvaluationSpecParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdRuleEvaluationSpecParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdRuleEvaluationSpec(ctx context.Context, client *core.Client, id string, params GetAdRuleEvaluationSpecParams) (*objects.AdRuleEvaluationSpec, error) {
	var out objects.AdRuleEvaluationSpec
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
