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

func GetVideoPollPollOptionsBatchCall(id string, params GetVideoPollPollOptionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "poll_options"), params.ToParams(), options...)
}

func NewGetVideoPollPollOptionsBatchRequest(id string, params GetVideoPollPollOptionsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.VideoPollOption]] {
	return core.NewBatchRequest[core.Cursor[objects.VideoPollOption]](GetVideoPollPollOptionsBatchCall(id, params, options...))
}

func DecodeGetVideoPollPollOptionsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.VideoPollOption], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.VideoPollOption]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetVideoPollPollOptions(ctx context.Context, client *core.Client, id string, params GetVideoPollPollOptionsParams) (*core.Cursor[objects.VideoPollOption], error) {
	var out core.Cursor[objects.VideoPollOption]
	call := GetVideoPollPollOptionsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetVideoPollBatchCall(id string, params GetVideoPollParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetVideoPollBatchRequest(id string, params GetVideoPollParams, options ...core.BatchOption) *core.BatchRequest[objects.VideoPoll] {
	return core.NewBatchRequest[objects.VideoPoll](GetVideoPollBatchCall(id, params, options...))
}

func DecodeGetVideoPollBatchResponse(response *core.BatchResponse) (*objects.VideoPoll, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.VideoPoll
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetVideoPoll(ctx context.Context, client *core.Client, id string, params GetVideoPollParams) (*objects.VideoPoll, error) {
	var out objects.VideoPoll
	call := GetVideoPollBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func UpdateVideoPollBatchCall(id string, params UpdateVideoPollParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateVideoPollBatchRequest(id string, params UpdateVideoPollParams, options ...core.BatchOption) *core.BatchRequest[objects.VideoPoll] {
	return core.NewBatchRequest[objects.VideoPoll](UpdateVideoPollBatchCall(id, params, options...))
}

func DecodeUpdateVideoPollBatchResponse(response *core.BatchResponse) (*objects.VideoPoll, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.VideoPoll
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateVideoPoll(ctx context.Context, client *core.Client, id string, params UpdateVideoPollParams) (*objects.VideoPoll, error) {
	var out objects.VideoPoll
	call := UpdateVideoPollBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
