package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetCPASMerchantConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCPASMerchantConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCPASMerchantConfig(ctx context.Context, client *core.Client, id string, params GetCPASMerchantConfigParams) (*objects.CPASMerchantConfig, error) {
	var out objects.CPASMerchantConfig
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
