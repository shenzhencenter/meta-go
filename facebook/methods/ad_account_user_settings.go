package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetAdAccountUserSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountUserSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountUserSettings(ctx context.Context, client *core.Client, id string, params GetAdAccountUserSettingsParams) (*objects.AdAccountUserSettings, error) {
	var out objects.AdAccountUserSettings
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
