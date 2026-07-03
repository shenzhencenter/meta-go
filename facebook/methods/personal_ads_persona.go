package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetPersonalAdsPersonaParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPersonalAdsPersonaParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPersonalAdsPersona(ctx context.Context, client *core.Client, id string, params GetPersonalAdsPersonaParams) (*objects.PersonalAdsPersona, error) {
	var out objects.PersonalAdsPersona
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
