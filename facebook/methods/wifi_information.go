package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetWifiInformationParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWifiInformationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWifiInformation(ctx context.Context, client *core.Client, id string, params GetWifiInformationParams) (*objects.WifiInformation, error) {
	var out objects.WifiInformation
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
