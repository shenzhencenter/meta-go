package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetEventTicketTierParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetEventTicketTierParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetEventTicketTier(ctx context.Context, client *core.Client, id string, params GetEventTicketTierParams) (*objects.EventTicketTier, error) {
	var out objects.EventTicketTier
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
