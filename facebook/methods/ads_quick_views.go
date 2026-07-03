package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdsQuickViewsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsQuickViewsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsQuickViewsBatchCall(id string, params GetAdsQuickViewsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdsQuickViewsBatchRequest(id string, params GetAdsQuickViewsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdsQuickViews] {
	return core.NewBatchRequest[objects.AdsQuickViews](GetAdsQuickViewsBatchCall(id, params, options...))
}

func DecodeGetAdsQuickViewsBatchResponse(response *core.BatchResponse) (*objects.AdsQuickViews, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdsQuickViews
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdsQuickViews(ctx context.Context, client *core.Client, id string, params GetAdsQuickViewsParams) (*objects.AdsQuickViews, error) {
	var out objects.AdsQuickViews
	call := GetAdsQuickViewsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
