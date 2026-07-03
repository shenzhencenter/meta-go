package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetFAMEKumoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetFAMEKumoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetFAMEKumo(ctx context.Context, client *core.Client, id string, params GetFAMEKumoParams) (*objects.FAMEKumo, error) {
	var out objects.FAMEKumo
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
