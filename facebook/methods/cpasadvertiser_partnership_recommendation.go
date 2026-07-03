package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCPASAdvertiserPartnershipRecommendationParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCPASAdvertiserPartnershipRecommendationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCPASAdvertiserPartnershipRecommendation(ctx context.Context, client *core.Client, id string, params GetCPASAdvertiserPartnershipRecommendationParams) (*objects.CPASAdvertiserPartnershipRecommendation, error) {
	var out objects.CPASAdvertiserPartnershipRecommendation
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
