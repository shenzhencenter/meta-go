package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdsConversionGoalConversionEventsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsConversionGoalConversionEventsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsConversionGoalConversionEventsBatchCall(id string, params GetAdsConversionGoalConversionEventsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "conversion_events"), params.ToParams(), options...)
}

func NewGetAdsConversionGoalConversionEventsBatchRequest(id string, params GetAdsConversionGoalConversionEventsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdsSingleChannelConversionEvent]] {
	return core.NewBatchRequest[core.Cursor[objects.AdsSingleChannelConversionEvent]](GetAdsConversionGoalConversionEventsBatchCall(id, params, options...))
}

func DecodeGetAdsConversionGoalConversionEventsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdsSingleChannelConversionEvent], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdsSingleChannelConversionEvent]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdsConversionGoalConversionEventsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdsConversionGoalConversionEventsParams) (*core.Cursor[objects.AdsSingleChannelConversionEvent], *core.Response, error) {
	var out core.Cursor[objects.AdsSingleChannelConversionEvent]
	call := GetAdsConversionGoalConversionEventsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdsConversionGoalConversionEvents(ctx context.Context, client *core.Client, id string, params GetAdsConversionGoalConversionEventsParams) (*core.Cursor[objects.AdsSingleChannelConversionEvent], error) {
	out, _, err := GetAdsConversionGoalConversionEventsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdsConversionGoalParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsConversionGoalParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsConversionGoalBatchCall(id string, params GetAdsConversionGoalParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdsConversionGoalBatchRequest(id string, params GetAdsConversionGoalParams, options ...core.BatchOption) *core.BatchRequest[objects.AdsConversionGoal] {
	return core.NewBatchRequest[objects.AdsConversionGoal](GetAdsConversionGoalBatchCall(id, params, options...))
}

func DecodeGetAdsConversionGoalBatchResponse(response *core.BatchResponse) (*objects.AdsConversionGoal, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdsConversionGoal
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdsConversionGoalWithResponse(ctx context.Context, client *core.Client, id string, params GetAdsConversionGoalParams) (*objects.AdsConversionGoal, *core.Response, error) {
	var out objects.AdsConversionGoal
	call := GetAdsConversionGoalBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdsConversionGoal(ctx context.Context, client *core.Client, id string, params GetAdsConversionGoalParams) (*objects.AdsConversionGoal, error) {
	out, _, err := GetAdsConversionGoalWithResponse(ctx, client, id, params)
	return out, err
}
