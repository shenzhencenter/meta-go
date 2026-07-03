package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetPartnerIntegrationLinkedParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPartnerIntegrationLinkedParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPartnerIntegrationLinkedBatchCall(id string, params GetPartnerIntegrationLinkedParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetPartnerIntegrationLinkedBatchRequest(id string, params GetPartnerIntegrationLinkedParams, options ...core.BatchOption) *core.BatchRequest[objects.PartnerIntegrationLinked] {
	return core.NewBatchRequest[objects.PartnerIntegrationLinked](GetPartnerIntegrationLinkedBatchCall(id, params, options...))
}

func DecodeGetPartnerIntegrationLinkedBatchResponse(response *core.BatchResponse) (*objects.PartnerIntegrationLinked, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PartnerIntegrationLinked
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPartnerIntegrationLinkedWithResponse(ctx context.Context, client *core.Client, id string, params GetPartnerIntegrationLinkedParams) (*objects.PartnerIntegrationLinked, *core.Response, error) {
	var out objects.PartnerIntegrationLinked
	call := GetPartnerIntegrationLinkedBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPartnerIntegrationLinked(ctx context.Context, client *core.Client, id string, params GetPartnerIntegrationLinkedParams) (*objects.PartnerIntegrationLinked, error) {
	out, _, err := GetPartnerIntegrationLinkedWithResponse(ctx, client, id, params)
	return out, err
}
