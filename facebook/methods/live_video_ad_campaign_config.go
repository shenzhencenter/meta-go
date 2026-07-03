package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetLiveVideoAdCampaignConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLiveVideoAdCampaignConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLiveVideoAdCampaignConfig(ctx context.Context, client *core.Client, id string, params GetLiveVideoAdCampaignConfigParams) (*objects.LiveVideoAdCampaignConfig, error) {
	var out objects.LiveVideoAdCampaignConfig
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
