package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetManagementSiteLinkParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetManagementSiteLinkParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetManagementSiteLink(ctx context.Context, client *core.Client, id string, params GetManagementSiteLinkParams) (*objects.ManagementSiteLink, error) {
	var out objects.ManagementSiteLink
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
