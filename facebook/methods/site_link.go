package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetSiteLinkParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetSiteLinkParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetSiteLink(ctx context.Context, client *core.Client, id string, params GetSiteLinkParams) (*objects.SiteLink, error) {
	var out objects.SiteLink
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
