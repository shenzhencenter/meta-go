package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetIGMediaForIGOnlyAPIChildrenParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGMediaForIGOnlyAPIChildrenParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGMediaForIGOnlyAPIChildren(ctx context.Context, client *core.Client, id string, params GetIGMediaForIGOnlyAPIChildrenParams) (*core.Cursor[objects.Media], error) {
	var out core.Cursor[objects.Media]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "children"), params.ToParams(), &out)
	return &out, err
}

type GetIGMediaForIGOnlyAPICommentsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGMediaForIGOnlyAPICommentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGMediaForIGOnlyAPIComments(ctx context.Context, client *core.Client, id string, params GetIGMediaForIGOnlyAPICommentsParams) (*core.Cursor[objects.Comment], error) {
	var out core.Cursor[objects.Comment]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "comments"), params.ToParams(), &out)
	return &out, err
}

type CreateIGMediaForIGOnlyAPICommentsParams struct {
	Message *string     `facebook:"message"`
	Extra   core.Params `facebook:"-"`
}

func (params CreateIGMediaForIGOnlyAPICommentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Message != nil {
		out["message"] = *params.Message
	}
	return out
}

func CreateIGMediaForIGOnlyAPIComments(ctx context.Context, client *core.Client, id string, params CreateIGMediaForIGOnlyAPICommentsParams) (*objects.IGGraphComment, error) {
	var out objects.IGGraphComment
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "comments"), params.ToParams(), &out)
	return &out, err
}

type GetIGMediaForIGOnlyAPIInsightsParams struct {
	Breakdown *[]enums.MediainsightsBreakdownEnumParam `facebook:"breakdown"`
	Metric    []enums.MediainsightsMetricEnumParam     `facebook:"metric"`
	Period    *[]enums.MediainsightsPeriodEnumParam    `facebook:"period"`
	Extra     core.Params                              `facebook:"-"`
}

func (params GetIGMediaForIGOnlyAPIInsightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Breakdown != nil {
		out["breakdown"] = *params.Breakdown
	}
	out["metric"] = params.Metric
	if params.Period != nil {
		out["period"] = *params.Period
	}
	return out
}

func GetIGMediaForIGOnlyAPIInsights(ctx context.Context, client *core.Client, id string, params GetIGMediaForIGOnlyAPIInsightsParams) (*core.Cursor[objects.InsightsResult], error) {
	var out core.Cursor[objects.InsightsResult]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "insights"), params.ToParams(), &out)
	return &out, err
}

type GetIGMediaForIGOnlyAPIParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGMediaForIGOnlyAPIParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGMediaForIGOnlyAPI(ctx context.Context, client *core.Client, id string, params GetIGMediaForIGOnlyAPIParams) (*objects.Media, error) {
	var out objects.Media
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateIGMediaForIGOnlyAPIParams struct {
	CommentEnabled bool        `facebook:"comment_enabled"`
	Extra          core.Params `facebook:"-"`
}

func (params UpdateIGMediaForIGOnlyAPIParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["comment_enabled"] = params.CommentEnabled
	return out
}

func UpdateIGMediaForIGOnlyAPI(ctx context.Context, client *core.Client, id string, params UpdateIGMediaForIGOnlyAPIParams) (*objects.IGGraphMedia, error) {
	var out objects.IGGraphMedia
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
