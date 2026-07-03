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

func GetCPASAdvertiserPartnershipRecommendationBatchCall(id string, params GetCPASAdvertiserPartnershipRecommendationParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCPASAdvertiserPartnershipRecommendationBatchRequest(id string, params GetCPASAdvertiserPartnershipRecommendationParams, options ...core.BatchOption) *core.BatchRequest[objects.CPASAdvertiserPartnershipRecommendation] {
	return core.NewBatchRequest[objects.CPASAdvertiserPartnershipRecommendation](GetCPASAdvertiserPartnershipRecommendationBatchCall(id, params, options...))
}

func DecodeGetCPASAdvertiserPartnershipRecommendationBatchResponse(response *core.BatchResponse) (*objects.CPASAdvertiserPartnershipRecommendation, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CPASAdvertiserPartnershipRecommendation
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCPASAdvertiserPartnershipRecommendation(ctx context.Context, client *core.Client, id string, params GetCPASAdvertiserPartnershipRecommendationParams) (*objects.CPASAdvertiserPartnershipRecommendation, error) {
	var out objects.CPASAdvertiserPartnershipRecommendation
	call := GetCPASAdvertiserPartnershipRecommendationBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
