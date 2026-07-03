package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetChinaBusinessOnboardingVettingRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetChinaBusinessOnboardingVettingRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetChinaBusinessOnboardingVettingRequestBatchCall(id string, params GetChinaBusinessOnboardingVettingRequestParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetChinaBusinessOnboardingVettingRequestBatchRequest(id string, params GetChinaBusinessOnboardingVettingRequestParams, options ...core.BatchOption) *core.BatchRequest[objects.ChinaBusinessOnboardingVettingRequest] {
	return core.NewBatchRequest[objects.ChinaBusinessOnboardingVettingRequest](GetChinaBusinessOnboardingVettingRequestBatchCall(id, params, options...))
}

func DecodeGetChinaBusinessOnboardingVettingRequestBatchResponse(response *core.BatchResponse) (*objects.ChinaBusinessOnboardingVettingRequest, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ChinaBusinessOnboardingVettingRequest
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetChinaBusinessOnboardingVettingRequestWithResponse(ctx context.Context, client *core.Client, id string, params GetChinaBusinessOnboardingVettingRequestParams) (*objects.ChinaBusinessOnboardingVettingRequest, *core.Response, error) {
	var out objects.ChinaBusinessOnboardingVettingRequest
	call := GetChinaBusinessOnboardingVettingRequestBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetChinaBusinessOnboardingVettingRequest(ctx context.Context, client *core.Client, id string, params GetChinaBusinessOnboardingVettingRequestParams) (*objects.ChinaBusinessOnboardingVettingRequest, error) {
	out, _, err := GetChinaBusinessOnboardingVettingRequestWithResponse(ctx, client, id, params)
	return out, err
}
