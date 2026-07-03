package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetBCPCampaignParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBCPCampaignParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBCPCampaign(ctx context.Context, client *core.Client, id string, params GetBCPCampaignParams) (*objects.BCPCampaign, error) {
	var out objects.BCPCampaign
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
