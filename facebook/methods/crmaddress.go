package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetCRMAddressParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCRMAddressParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCRMAddress(ctx context.Context, client *core.Client, id string, params GetCRMAddressParams) (*objects.CRMAddress, error) {
	var out objects.CRMAddress
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
