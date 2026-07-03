package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetThirdPartyPartnerLiftRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetThirdPartyPartnerLiftRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetThirdPartyPartnerLiftRequest(ctx context.Context, client *core.Client, id string, params GetThirdPartyPartnerLiftRequestParams) (*objects.ThirdPartyPartnerLiftRequest, error) {
	var out objects.ThirdPartyPartnerLiftRequest
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
