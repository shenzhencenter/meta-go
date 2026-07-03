package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetChinaBusinessOnboardingVettingRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetChinaBusinessOnboardingVettingRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetChinaBusinessOnboardingVettingRequest(ctx context.Context, client *core.Client, id string, params GetChinaBusinessOnboardingVettingRequestParams) (*objects.ChinaBusinessOnboardingVettingRequest, error) {
	var out objects.ChinaBusinessOnboardingVettingRequest
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
