package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetPartnerIntegrationLinkedParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPartnerIntegrationLinkedParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPartnerIntegrationLinked(ctx context.Context, client *core.Client, id string, params GetPartnerIntegrationLinkedParams) (*objects.PartnerIntegrationLinked, error) {
	var out objects.PartnerIntegrationLinked
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
