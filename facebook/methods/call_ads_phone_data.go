package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetCallAdsPhoneDataParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCallAdsPhoneDataParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCallAdsPhoneData(ctx context.Context, client *core.Client, id string, params GetCallAdsPhoneDataParams) (*objects.CallAdsPhoneData, error) {
	var out objects.CallAdsPhoneData
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
