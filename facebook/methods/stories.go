package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetStoriesInsightsParams struct {
	Metric *[]enums.StoriesinsightsMetricEnumParam `facebook:"metric"`
	Extra  core.Params                             `facebook:"-"`
}

func (params GetStoriesInsightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Metric != nil {
		out["metric"] = *params.Metric
	}
	return out
}

func GetStoriesInsights(ctx context.Context, client *core.Client, id string, params GetStoriesInsightsParams) (*core.Cursor[objects.InsightsResult], error) {
	var out core.Cursor[objects.InsightsResult]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "insights"), params.ToParams(), &out)
	return &out, err
}

type GetStoriesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetStoriesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetStories(ctx context.Context, client *core.Client, id string, params GetStoriesParams) (*objects.Stories, error) {
	var out objects.Stories
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
