package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdsStartYourDayWidgetParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsStartYourDayWidgetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsStartYourDayWidgetBatchCall(id string, params GetAdsStartYourDayWidgetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdsStartYourDayWidgetBatchRequest(id string, params GetAdsStartYourDayWidgetParams, options ...core.BatchOption) *core.BatchRequest[objects.AdsStartYourDayWidget] {
	return core.NewBatchRequest[objects.AdsStartYourDayWidget](GetAdsStartYourDayWidgetBatchCall(id, params, options...))
}

func DecodeGetAdsStartYourDayWidgetBatchResponse(response *core.BatchResponse) (*objects.AdsStartYourDayWidget, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdsStartYourDayWidget
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdsStartYourDayWidget(ctx context.Context, client *core.Client, id string, params GetAdsStartYourDayWidgetParams) (*objects.AdsStartYourDayWidget, error) {
	var out objects.AdsStartYourDayWidget
	call := GetAdsStartYourDayWidgetBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
