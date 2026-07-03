package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetThirdPartyPartnerPanelRequestParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetThirdPartyPartnerPanelRequestParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetThirdPartyPartnerPanelRequestBatchCall(id string, params GetThirdPartyPartnerPanelRequestParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetThirdPartyPartnerPanelRequestBatchRequest(id string, params GetThirdPartyPartnerPanelRequestParams, options ...core.BatchOption) *core.BatchRequest[objects.ThirdPartyPartnerPanelRequest] {
	return core.NewBatchRequest[objects.ThirdPartyPartnerPanelRequest](GetThirdPartyPartnerPanelRequestBatchCall(id, params, options...))
}

func DecodeGetThirdPartyPartnerPanelRequestBatchResponse(response *core.BatchResponse) (*objects.ThirdPartyPartnerPanelRequest, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ThirdPartyPartnerPanelRequest
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetThirdPartyPartnerPanelRequestWithResponse(ctx context.Context, client *core.Client, id string, params GetThirdPartyPartnerPanelRequestParams) (*objects.ThirdPartyPartnerPanelRequest, *core.Response, error) {
	var out objects.ThirdPartyPartnerPanelRequest
	call := GetThirdPartyPartnerPanelRequestBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetThirdPartyPartnerPanelRequest(ctx context.Context, client *core.Client, id string, params GetThirdPartyPartnerPanelRequestParams) (*objects.ThirdPartyPartnerPanelRequest, error) {
	out, _, err := GetThirdPartyPartnerPanelRequestWithResponse(ctx, client, id, params)
	return out, err
}
