package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type DeleteAdCampaignGroupGendeleteParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteAdCampaignGroupGendeleteParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteAdCampaignGroupGendelete(ctx context.Context, client *core.Client, id string, params DeleteAdCampaignGroupGendeleteParams) (*objects.AdCampaignGroupDelete, error) {
	var out objects.AdCampaignGroupDelete
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, ""), params.ToParams(), &out)
	return &out, err
}
