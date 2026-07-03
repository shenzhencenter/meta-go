package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetIGAccessTokenForIGOnlyAPIParams struct {
	AccessToken  string      `facebook:"access_token"`
	ClientSecret string      `facebook:"client_secret"`
	GrantType    string      `facebook:"grant_type"`
	Extra        core.Params `facebook:"-"`
}

func (params GetIGAccessTokenForIGOnlyAPIParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["access_token"] = params.AccessToken
	out["client_secret"] = params.ClientSecret
	out["grant_type"] = params.GrantType
	return out
}

func GetIGAccessTokenForIGOnlyAPI(ctx context.Context, client *core.Client, id string, params GetIGAccessTokenForIGOnlyAPIParams) (*objects.AccessToken, error) {
	var out objects.AccessToken
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
