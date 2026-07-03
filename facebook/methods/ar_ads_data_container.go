package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetArAdsDataContainerParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetArAdsDataContainerParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetArAdsDataContainer(ctx context.Context, client *core.Client, id string, params GetArAdsDataContainerParams) (*objects.ArAdsDataContainer, error) {
	var out objects.ArAdsDataContainer
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
