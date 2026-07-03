package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetThirdPartyPartnerPanelScheduledParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetThirdPartyPartnerPanelScheduledParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetThirdPartyPartnerPanelScheduled(ctx context.Context, client *core.Client, id string, params GetThirdPartyPartnerPanelScheduledParams) (*objects.ThirdPartyPartnerPanelScheduled, error) {
	var out objects.ThirdPartyPartnerPanelScheduled
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
