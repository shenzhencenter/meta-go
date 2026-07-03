package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetExternalMerchantSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetExternalMerchantSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetExternalMerchantSettings(ctx context.Context, client *core.Client, id string, params GetExternalMerchantSettingsParams) (*objects.ExternalMerchantSettings, error) {
	var out objects.ExternalMerchantSettings
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
