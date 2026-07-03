package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdsCreationSavedStateParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsCreationSavedStateParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsCreationSavedStateBatchCall(id string, params GetAdsCreationSavedStateParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdsCreationSavedStateBatchRequest(id string, params GetAdsCreationSavedStateParams, options ...core.BatchOption) *core.BatchRequest[objects.AdsCreationSavedState] {
	return core.NewBatchRequest[objects.AdsCreationSavedState](GetAdsCreationSavedStateBatchCall(id, params, options...))
}

func DecodeGetAdsCreationSavedStateBatchResponse(response *core.BatchResponse) (*objects.AdsCreationSavedState, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdsCreationSavedState
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdsCreationSavedStateWithResponse(ctx context.Context, client *core.Client, id string, params GetAdsCreationSavedStateParams) (*objects.AdsCreationSavedState, *core.Response, error) {
	var out objects.AdsCreationSavedState
	call := GetAdsCreationSavedStateBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdsCreationSavedState(ctx context.Context, client *core.Client, id string, params GetAdsCreationSavedStateParams) (*objects.AdsCreationSavedState, error) {
	out, _, err := GetAdsCreationSavedStateWithResponse(ctx, client, id, params)
	return out, err
}
