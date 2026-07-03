package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetMessengerAdsPartialAutomatedStepListStepsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetMessengerAdsPartialAutomatedStepListStepsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetMessengerAdsPartialAutomatedStepListStepsBatchCall(id string, params GetMessengerAdsPartialAutomatedStepListStepsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "steps"), params.ToParams(), options...)
}

func NewGetMessengerAdsPartialAutomatedStepListStepsBatchRequest(id string, params GetMessengerAdsPartialAutomatedStepListStepsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.MessengerAdsPartialAutomatedStep]] {
	return core.NewBatchRequest[core.Cursor[objects.MessengerAdsPartialAutomatedStep]](GetMessengerAdsPartialAutomatedStepListStepsBatchCall(id, params, options...))
}

func DecodeGetMessengerAdsPartialAutomatedStepListStepsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.MessengerAdsPartialAutomatedStep], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.MessengerAdsPartialAutomatedStep]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetMessengerAdsPartialAutomatedStepListSteps(ctx context.Context, client *core.Client, id string, params GetMessengerAdsPartialAutomatedStepListStepsParams) (*core.Cursor[objects.MessengerAdsPartialAutomatedStep], error) {
	var out core.Cursor[objects.MessengerAdsPartialAutomatedStep]
	call := GetMessengerAdsPartialAutomatedStepListStepsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetMessengerAdsPartialAutomatedStepListParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetMessengerAdsPartialAutomatedStepListParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetMessengerAdsPartialAutomatedStepListBatchCall(id string, params GetMessengerAdsPartialAutomatedStepListParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetMessengerAdsPartialAutomatedStepListBatchRequest(id string, params GetMessengerAdsPartialAutomatedStepListParams, options ...core.BatchOption) *core.BatchRequest[objects.MessengerAdsPartialAutomatedStepList] {
	return core.NewBatchRequest[objects.MessengerAdsPartialAutomatedStepList](GetMessengerAdsPartialAutomatedStepListBatchCall(id, params, options...))
}

func DecodeGetMessengerAdsPartialAutomatedStepListBatchResponse(response *core.BatchResponse) (*objects.MessengerAdsPartialAutomatedStepList, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.MessengerAdsPartialAutomatedStepList
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetMessengerAdsPartialAutomatedStepList(ctx context.Context, client *core.Client, id string, params GetMessengerAdsPartialAutomatedStepListParams) (*objects.MessengerAdsPartialAutomatedStepList, error) {
	var out objects.MessengerAdsPartialAutomatedStepList
	call := GetMessengerAdsPartialAutomatedStepListBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
