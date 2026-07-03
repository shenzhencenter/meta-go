package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type DeleteAdCampaignGendeleteParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteAdCampaignGendeleteParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteAdCampaignGendelete(ctx context.Context, client *core.Client, id string, params DeleteAdCampaignGendeleteParams) (*objects.AdCampaignDelete, error) {
	var out objects.AdCampaignDelete
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, ""), params.ToParams(), &out)
	return &out, err
}
