package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetOrganizationParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOrganizationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOrganizationBatchCall(id string, params GetOrganizationParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetOrganizationBatchRequest(id string, params GetOrganizationParams, options ...core.BatchOption) *core.BatchRequest[objects.Organization] {
	return core.NewBatchRequest[objects.Organization](GetOrganizationBatchCall(id, params, options...))
}

func DecodeGetOrganizationBatchResponse(response *core.BatchResponse) (*objects.Organization, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Organization
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOrganizationWithResponse(ctx context.Context, client *core.Client, id string, params GetOrganizationParams) (*objects.Organization, *core.Response, error) {
	var out objects.Organization
	call := GetOrganizationBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetOrganization(ctx context.Context, client *core.Client, id string, params GetOrganizationParams) (*objects.Organization, error) {
	out, _, err := GetOrganizationWithResponse(ctx, client, id, params)
	return out, err
}
