package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetExtendedCreditApplicationParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetExtendedCreditApplicationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetExtendedCreditApplication(ctx context.Context, client *core.Client, id string, params GetExtendedCreditApplicationParams) (*objects.ExtendedCreditApplication, error) {
	var out objects.ExtendedCreditApplication
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
