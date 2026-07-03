package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetWearableDevicePublicKeyParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWearableDevicePublicKeyParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWearableDevicePublicKey(ctx context.Context, client *core.Client, id string, params GetWearableDevicePublicKeyParams) (*objects.WearableDevicePublicKey, error) {
	var out objects.WearableDevicePublicKey
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
