package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdsPixelCapabilityOverrideParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsPixelCapabilityOverrideParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsPixelCapabilityOverrideBatchCall(id string, params GetAdsPixelCapabilityOverrideParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdsPixelCapabilityOverrideBatchRequest(id string, params GetAdsPixelCapabilityOverrideParams, options ...core.BatchOption) *core.BatchRequest[objects.AdsPixelCapabilityOverride] {
	return core.NewBatchRequest[objects.AdsPixelCapabilityOverride](GetAdsPixelCapabilityOverrideBatchCall(id, params, options...))
}

func DecodeGetAdsPixelCapabilityOverrideBatchResponse(response *core.BatchResponse) (*objects.AdsPixelCapabilityOverride, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdsPixelCapabilityOverride
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdsPixelCapabilityOverride(ctx context.Context, client *core.Client, id string, params GetAdsPixelCapabilityOverrideParams) (*objects.AdsPixelCapabilityOverride, error) {
	var out objects.AdsPixelCapabilityOverride
	call := GetAdsPixelCapabilityOverrideBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
