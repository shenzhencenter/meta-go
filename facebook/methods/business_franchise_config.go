package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetBusinessFranchiseConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessFranchiseConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessFranchiseConfig(ctx context.Context, client *core.Client, id string, params GetBusinessFranchiseConfigParams) (*objects.BusinessFranchiseConfig, error) {
	var out objects.BusinessFranchiseConfig
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
