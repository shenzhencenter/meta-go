package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetPagePostExperimentVideoInsightsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPagePostExperimentVideoInsightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPagePostExperimentVideoInsights(ctx context.Context, client *core.Client, id string, params GetPagePostExperimentVideoInsightsParams) (*core.Cursor[objects.PagePostExperimentVideoInsightsQueryResult], error) {
	var out core.Cursor[objects.PagePostExperimentVideoInsightsQueryResult]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "video_insights"), params.ToParams(), &out)
	return &out, err
}

type DeletePagePostExperimentParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeletePagePostExperimentParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeletePagePostExperiment(ctx context.Context, client *core.Client, id string, params DeletePagePostExperimentParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetPagePostExperimentParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPagePostExperimentParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPagePostExperiment(ctx context.Context, client *core.Client, id string, params GetPagePostExperimentParams) (*objects.PagePostExperiment, error) {
	var out objects.PagePostExperiment
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
