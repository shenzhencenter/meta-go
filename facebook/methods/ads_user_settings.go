package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetAdsUserSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsUserSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsUserSettings(ctx context.Context, client *core.Client, id string, params GetAdsUserSettingsParams) (*objects.AdsUserSettings, error) {
	var out objects.AdsUserSettings
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
