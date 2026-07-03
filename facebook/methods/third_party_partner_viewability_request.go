package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetThirdPartyPartnerViewabilityRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetThirdPartyPartnerViewabilityRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetThirdPartyPartnerViewabilityRequestBatchCall(id string, params GetThirdPartyPartnerViewabilityRequestParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetThirdPartyPartnerViewabilityRequestBatchRequest(id string, params GetThirdPartyPartnerViewabilityRequestParams, options ...core.BatchOption) *core.BatchRequest[objects.ThirdPartyPartnerViewabilityRequest] {
	return core.NewBatchRequest[objects.ThirdPartyPartnerViewabilityRequest](GetThirdPartyPartnerViewabilityRequestBatchCall(id, params, options...))
}

func DecodeGetThirdPartyPartnerViewabilityRequestBatchResponse(response *core.BatchResponse) (*objects.ThirdPartyPartnerViewabilityRequest, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ThirdPartyPartnerViewabilityRequest
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetThirdPartyPartnerViewabilityRequestWithResponse(ctx context.Context, client *core.Client, id string, params GetThirdPartyPartnerViewabilityRequestParams) (*objects.ThirdPartyPartnerViewabilityRequest, *core.Response, error) {
	var out objects.ThirdPartyPartnerViewabilityRequest
	call := GetThirdPartyPartnerViewabilityRequestBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetThirdPartyPartnerViewabilityRequest(ctx context.Context, client *core.Client, id string, params GetThirdPartyPartnerViewabilityRequestParams) (*objects.ThirdPartyPartnerViewabilityRequest, error) {
	out, _, err := GetThirdPartyPartnerViewabilityRequestWithResponse(ctx, client, id, params)
	return out, err
}
