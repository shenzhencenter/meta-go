package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetALMAdAccountInfoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetALMAdAccountInfoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetALMAdAccountInfo(ctx context.Context, client *core.Client, id string, params GetALMAdAccountInfoParams) (*objects.ALMAdAccountInfo, error) {
	var out objects.ALMAdAccountInfo
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
