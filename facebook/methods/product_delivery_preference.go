package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetProductDeliveryPreferenceParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductDeliveryPreferenceParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductDeliveryPreference(ctx context.Context, client *core.Client, id string, params GetProductDeliveryPreferenceParams) (*objects.ProductDeliveryPreference, error) {
	var out objects.ProductDeliveryPreference
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
