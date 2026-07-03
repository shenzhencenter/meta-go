package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdLightCampaignParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdLightCampaignParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdLightCampaign(ctx context.Context, client *core.Client, id string, params GetAdLightCampaignParams) (*objects.AdLightCampaign, error) {
	var out objects.AdLightCampaign
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
