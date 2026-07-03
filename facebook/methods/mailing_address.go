package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetMailingAddressParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetMailingAddressParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetMailingAddress(ctx context.Context, client *core.Client, id string, params GetMailingAddressParams) (*objects.MailingAddress, error) {
	var out objects.MailingAddress
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
