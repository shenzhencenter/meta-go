package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetOfflineTermsOfServiceParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOfflineTermsOfServiceParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOfflineTermsOfService(ctx context.Context, client *core.Client, id string, params GetOfflineTermsOfServiceParams) (*objects.OfflineTermsOfService, error) {
	var out objects.OfflineTermsOfService
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
