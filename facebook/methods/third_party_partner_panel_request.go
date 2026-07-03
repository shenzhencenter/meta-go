package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetThirdPartyPartnerPanelRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetThirdPartyPartnerPanelRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetThirdPartyPartnerPanelRequest(ctx context.Context, client *core.Client, id string, params GetThirdPartyPartnerPanelRequestParams) (*objects.ThirdPartyPartnerPanelRequest, error) {
	var out objects.ThirdPartyPartnerPanelRequest
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
