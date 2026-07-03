package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetThirdPartyPartnerLiftRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetThirdPartyPartnerLiftRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetThirdPartyPartnerLiftRequestBatchCall(id string, params GetThirdPartyPartnerLiftRequestParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetThirdPartyPartnerLiftRequestBatchRequest(id string, params GetThirdPartyPartnerLiftRequestParams, options ...core.BatchOption) *core.BatchRequest[objects.ThirdPartyPartnerLiftRequest] {
	return core.NewBatchRequest[objects.ThirdPartyPartnerLiftRequest](GetThirdPartyPartnerLiftRequestBatchCall(id, params, options...))
}

func DecodeGetThirdPartyPartnerLiftRequestBatchResponse(response *core.BatchResponse) (*objects.ThirdPartyPartnerLiftRequest, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ThirdPartyPartnerLiftRequest
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetThirdPartyPartnerLiftRequestWithResponse(ctx context.Context, client *core.Client, id string, params GetThirdPartyPartnerLiftRequestParams) (*objects.ThirdPartyPartnerLiftRequest, *core.Response, error) {
	var out objects.ThirdPartyPartnerLiftRequest
	call := GetThirdPartyPartnerLiftRequestBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetThirdPartyPartnerLiftRequest(ctx context.Context, client *core.Client, id string, params GetThirdPartyPartnerLiftRequestParams) (*objects.ThirdPartyPartnerLiftRequest, error) {
	out, _, err := GetThirdPartyPartnerLiftRequestWithResponse(ctx, client, id, params)
	return out, err
}
