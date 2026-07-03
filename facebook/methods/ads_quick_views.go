package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func GetAdsQuickViews(ctx context.Context, client *core.Client, id string, params GetAdsQuickViewsParams) (*objects.AdsQuickViews, error) {
	var out objects.AdsQuickViews
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
