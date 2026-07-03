package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func GetAdsPixelCapabilityOverride(ctx context.Context, client *core.Client, id string, params GetAdsPixelCapabilityOverrideParams) (*objects.AdsPixelCapabilityOverride, error) {
	var out objects.AdsPixelCapabilityOverride
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
