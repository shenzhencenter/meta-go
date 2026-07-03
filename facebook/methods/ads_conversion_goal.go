package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func GetAdsConversionGoalConversionEvents(ctx context.Context, client *core.Client, id string, params GetAdsConversionGoalConversionEventsParams) (*core.Cursor[objects.AdsSingleChannelConversionEvent], error) {
	var out core.Cursor[objects.AdsSingleChannelConversionEvent]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "conversion_events"), params.ToParams(), &out)
	return &out, err
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

func GetAdsConversionGoal(ctx context.Context, client *core.Client, id string, params GetAdsConversionGoalParams) (*objects.AdsConversionGoal, error) {
	var out objects.AdsConversionGoal
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
