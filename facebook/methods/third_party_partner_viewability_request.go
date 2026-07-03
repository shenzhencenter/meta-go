package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetThirdPartyPartnerViewabilityRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetThirdPartyPartnerViewabilityRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetThirdPartyPartnerViewabilityRequest(ctx context.Context, client *core.Client, id string, params GetThirdPartyPartnerViewabilityRequestParams) (*objects.ThirdPartyPartnerViewabilityRequest, error) {
	var out objects.ThirdPartyPartnerViewabilityRequest
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
