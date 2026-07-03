package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetAdAccountAgencyFeeConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAccountAgencyFeeConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAccountAgencyFeeConfig(ctx context.Context, client *core.Client, id string, params GetAdAccountAgencyFeeConfigParams) (*objects.AdAccountAgencyFeeConfig, error) {
	var out objects.AdAccountAgencyFeeConfig
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
