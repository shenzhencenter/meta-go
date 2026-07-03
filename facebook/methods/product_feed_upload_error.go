package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetProductFeedUploadErrorSamplesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductFeedUploadErrorSamplesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductFeedUploadErrorSamples(ctx context.Context, client *core.Client, id string, params GetProductFeedUploadErrorSamplesParams) (*core.Cursor[objects.ProductFeedUploadErrorSample], error) {
	var out core.Cursor[objects.ProductFeedUploadErrorSample]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "samples"), params.ToParams(), &out)
	return &out, err
}

type GetProductFeedUploadErrorSuggestedRulesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductFeedUploadErrorSuggestedRulesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductFeedUploadErrorSuggestedRules(ctx context.Context, client *core.Client, id string, params GetProductFeedUploadErrorSuggestedRulesParams) (*core.Cursor[objects.ProductFeedRuleSuggestion], error) {
	var out core.Cursor[objects.ProductFeedRuleSuggestion]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "suggested_rules"), params.ToParams(), &out)
	return &out, err
}

type GetProductFeedUploadErrorParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductFeedUploadErrorParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductFeedUploadError(ctx context.Context, client *core.Client, id string, params GetProductFeedUploadErrorParams) (*objects.ProductFeedUploadError, error) {
	var out objects.ProductFeedUploadError
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
