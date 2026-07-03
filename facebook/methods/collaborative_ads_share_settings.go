package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetCollaborativeAdsShareSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCollaborativeAdsShareSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCollaborativeAdsShareSettings(ctx context.Context, client *core.Client, id string, params GetCollaborativeAdsShareSettingsParams) (*objects.CollaborativeAdsShareSettings, error) {
	var out objects.CollaborativeAdsShareSettings
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
