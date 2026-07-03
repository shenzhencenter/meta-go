package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetUserPageOneTimeOptInTokenSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetUserPageOneTimeOptInTokenSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetUserPageOneTimeOptInTokenSettings(ctx context.Context, client *core.Client, id string, params GetUserPageOneTimeOptInTokenSettingsParams) (*objects.UserPageOneTimeOptInTokenSettings, error) {
	var out objects.UserPageOneTimeOptInTokenSettings
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
