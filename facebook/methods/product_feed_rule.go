package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type DeleteProductFeedRuleParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteProductFeedRuleParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteProductFeedRule(ctx context.Context, client *core.Client, id string, params DeleteProductFeedRuleParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetProductFeedRuleParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductFeedRuleParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductFeedRule(ctx context.Context, client *core.Client, id string, params GetProductFeedRuleParams) (*objects.ProductFeedRule, error) {
	var out objects.ProductFeedRule
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateProductFeedRuleParams struct {
	Params map[string]interface{} `facebook:"params"`
	Extra  core.Params            `facebook:"-"`
}

func (params UpdateProductFeedRuleParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["params"] = params.Params
	return out
}

func UpdateProductFeedRule(ctx context.Context, client *core.Client, id string, params UpdateProductFeedRuleParams) (*objects.ProductFeedRule, error) {
	var out objects.ProductFeedRule
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
