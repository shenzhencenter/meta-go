package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetStoreLocationParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetStoreLocationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetStoreLocation(ctx context.Context, client *core.Client, id string, params GetStoreLocationParams) (*objects.StoreLocation, error) {
	var out objects.StoreLocation
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
