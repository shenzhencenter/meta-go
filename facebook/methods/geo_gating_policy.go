package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetGeoGatingPolicyParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetGeoGatingPolicyParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetGeoGatingPolicyBatchCall(id string, params GetGeoGatingPolicyParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetGeoGatingPolicyBatchRequest(id string, params GetGeoGatingPolicyParams, options ...core.BatchOption) *core.BatchRequest[objects.GeoGatingPolicy] {
	return core.NewBatchRequest[objects.GeoGatingPolicy](GetGeoGatingPolicyBatchCall(id, params, options...))
}

func DecodeGetGeoGatingPolicyBatchResponse(response *core.BatchResponse) (*objects.GeoGatingPolicy, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.GeoGatingPolicy
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetGeoGatingPolicyWithResponse(ctx context.Context, client *core.Client, id string, params GetGeoGatingPolicyParams) (*objects.GeoGatingPolicy, *core.Response, error) {
	var out objects.GeoGatingPolicy
	call := GetGeoGatingPolicyBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetGeoGatingPolicy(ctx context.Context, client *core.Client, id string, params GetGeoGatingPolicyParams) (*objects.GeoGatingPolicy, error) {
	out, _, err := GetGeoGatingPolicyWithResponse(ctx, client, id, params)
	return out, err
}
