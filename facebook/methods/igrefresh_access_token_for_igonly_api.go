package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetIGRefreshAccessTokenForIGOnlyAPIParams struct {
	AccessToken string      `facebook:"access_token"`
	GrantType   string      `facebook:"grant_type"`
	Extra       core.Params `facebook:"-"`
}

func (params GetIGRefreshAccessTokenForIGOnlyAPIParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["access_token"] = params.AccessToken
	out["grant_type"] = params.GrantType
	return out
}

func GetIGRefreshAccessTokenForIGOnlyAPI(ctx context.Context, client *core.Client, id string, params GetIGRefreshAccessTokenForIGOnlyAPIParams) (*objects.RefreshAccessToken, error) {
	var out objects.RefreshAccessToken
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
