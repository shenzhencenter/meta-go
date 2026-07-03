package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetPartnerAccountLinkingParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPartnerAccountLinkingParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPartnerAccountLinkingBatchCall(id string, params GetPartnerAccountLinkingParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetPartnerAccountLinkingBatchRequest(id string, params GetPartnerAccountLinkingParams, options ...core.BatchOption) *core.BatchRequest[objects.PartnerAccountLinking] {
	return core.NewBatchRequest[objects.PartnerAccountLinking](GetPartnerAccountLinkingBatchCall(id, params, options...))
}

func DecodeGetPartnerAccountLinkingBatchResponse(response *core.BatchResponse) (*objects.PartnerAccountLinking, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PartnerAccountLinking
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPartnerAccountLinkingWithResponse(ctx context.Context, client *core.Client, id string, params GetPartnerAccountLinkingParams) (*objects.PartnerAccountLinking, *core.Response, error) {
	var out objects.PartnerAccountLinking
	call := GetPartnerAccountLinkingBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPartnerAccountLinking(ctx context.Context, client *core.Client, id string, params GetPartnerAccountLinkingParams) (*objects.PartnerAccountLinking, error) {
	out, _, err := GetPartnerAccountLinkingWithResponse(ctx, client, id, params)
	return out, err
}
