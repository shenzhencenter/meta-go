package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetOrganizationParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOrganizationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOrganization(ctx context.Context, client *core.Client, id string, params GetOrganizationParams) (*objects.Organization, error) {
	var out objects.Organization
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
