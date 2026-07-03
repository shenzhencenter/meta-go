package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetAdProposalParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdProposalParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdProposal(ctx context.Context, client *core.Client, id string, params GetAdProposalParams) (*objects.AdProposal, error) {
	var out objects.AdProposal
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
