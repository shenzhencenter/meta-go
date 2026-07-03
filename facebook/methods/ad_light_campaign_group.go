package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdLightCampaignGroupParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdLightCampaignGroupParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdLightCampaignGroup(ctx context.Context, client *core.Client, id string, params GetAdLightCampaignGroupParams) (*objects.AdLightCampaignGroup, error) {
	var out objects.AdLightCampaignGroup
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
