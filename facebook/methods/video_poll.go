package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetVideoPollPollOptionsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetVideoPollPollOptionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetVideoPollPollOptions(ctx context.Context, client *core.Client, id string, params GetVideoPollPollOptionsParams) (*core.Cursor[objects.VideoPollOption], error) {
	var out core.Cursor[objects.VideoPollOption]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "poll_options"), params.ToParams(), &out)
	return &out, err
}

type GetVideoPollParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetVideoPollParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetVideoPoll(ctx context.Context, client *core.Client, id string, params GetVideoPollParams) (*objects.VideoPoll, error) {
	var out objects.VideoPoll
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateVideoPollParams struct {
	Action           enums.VideopollAction `facebook:"action"`
	CloseAfterVoting *bool                 `facebook:"close_after_voting"`
	DefaultOpen      *bool                 `facebook:"default_open"`
	ShowGradient     *bool                 `facebook:"show_gradient"`
	ShowResults      *bool                 `facebook:"show_results"`
	Extra            core.Params           `facebook:"-"`
}

func (params UpdateVideoPollParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["action"] = params.Action
	if params.CloseAfterVoting != nil {
		out["close_after_voting"] = *params.CloseAfterVoting
	}
	if params.DefaultOpen != nil {
		out["default_open"] = *params.DefaultOpen
	}
	if params.ShowGradient != nil {
		out["show_gradient"] = *params.ShowGradient
	}
	if params.ShowResults != nil {
		out["show_results"] = *params.ShowResults
	}
	return out
}

func UpdateVideoPoll(ctx context.Context, client *core.Client, id string, params UpdateVideoPollParams) (*objects.VideoPoll, error) {
	var out objects.VideoPoll
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
