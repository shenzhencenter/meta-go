package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCTXPartnerAppWelcomeMessageFlowParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCTXPartnerAppWelcomeMessageFlowParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCTXPartnerAppWelcomeMessageFlow(ctx context.Context, client *core.Client, id string, params GetCTXPartnerAppWelcomeMessageFlowParams) (*objects.CTXPartnerAppWelcomeMessageFlow, error) {
	var out objects.CTXPartnerAppWelcomeMessageFlow
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
