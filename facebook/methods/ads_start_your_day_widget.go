package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func GetAdsStartYourDayWidget(ctx context.Context, client *core.Client, id string, params GetAdsStartYourDayWidgetParams) (*objects.AdsStartYourDayWidget, error) {
	var out objects.AdsStartYourDayWidget
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
