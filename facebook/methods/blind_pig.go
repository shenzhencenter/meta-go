package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetBlindPigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBlindPigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBlindPig(ctx context.Context, client *core.Client, id string, params GetBlindPigParams) (*objects.BlindPig, error) {
	var out objects.BlindPig
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
