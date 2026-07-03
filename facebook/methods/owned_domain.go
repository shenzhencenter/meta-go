package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetOwnedDomainParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOwnedDomainParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOwnedDomainBatchCall(id string, params GetOwnedDomainParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetOwnedDomainBatchRequest(id string, params GetOwnedDomainParams, options ...core.BatchOption) *core.BatchRequest[objects.OwnedDomain] {
	return core.NewBatchRequest[objects.OwnedDomain](GetOwnedDomainBatchCall(id, params, options...))
}

func DecodeGetOwnedDomainBatchResponse(response *core.BatchResponse) (*objects.OwnedDomain, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.OwnedDomain
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOwnedDomainWithResponse(ctx context.Context, client *core.Client, id string, params GetOwnedDomainParams) (*objects.OwnedDomain, *core.Response, error) {
	var out objects.OwnedDomain
	call := GetOwnedDomainBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetOwnedDomain(ctx context.Context, client *core.Client, id string, params GetOwnedDomainParams) (*objects.OwnedDomain, error) {
	out, _, err := GetOwnedDomainWithResponse(ctx, client, id, params)
	return out, err
}
