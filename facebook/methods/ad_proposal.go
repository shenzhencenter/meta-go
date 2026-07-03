package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdProposalParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdProposalParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdProposalBatchCall(id string, params GetAdProposalParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdProposalBatchRequest(id string, params GetAdProposalParams, options ...core.BatchOption) *core.BatchRequest[objects.AdProposal] {
	return core.NewBatchRequest[objects.AdProposal](GetAdProposalBatchCall(id, params, options...))
}

func DecodeGetAdProposalBatchResponse(response *core.BatchResponse) (*objects.AdProposal, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdProposal
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdProposalWithResponse(ctx context.Context, client *core.Client, id string, params GetAdProposalParams) (*objects.AdProposal, *core.Response, error) {
	var out objects.AdProposal
	call := GetAdProposalBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdProposal(ctx context.Context, client *core.Client, id string, params GetAdProposalParams) (*objects.AdProposal, error) {
	out, _, err := GetAdProposalWithResponse(ctx, client, id, params)
	return out, err
}
