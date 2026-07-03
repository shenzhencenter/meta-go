package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCreditCardParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCreditCardParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCreditCard(ctx context.Context, client *core.Client, id string, params GetCreditCardParams) (*objects.CreditCard, error) {
	var out objects.CreditCard
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
