package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetLeadGenDirectCRMIntegrationConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLeadGenDirectCRMIntegrationConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLeadGenDirectCRMIntegrationConfigBatchCall(id string, params GetLeadGenDirectCRMIntegrationConfigParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetLeadGenDirectCRMIntegrationConfigBatchRequest(id string, params GetLeadGenDirectCRMIntegrationConfigParams, options ...core.BatchOption) *core.BatchRequest[objects.LeadGenDirectCRMIntegrationConfig] {
	return core.NewBatchRequest[objects.LeadGenDirectCRMIntegrationConfig](GetLeadGenDirectCRMIntegrationConfigBatchCall(id, params, options...))
}

func DecodeGetLeadGenDirectCRMIntegrationConfigBatchResponse(response *core.BatchResponse) (*objects.LeadGenDirectCRMIntegrationConfig, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.LeadGenDirectCRMIntegrationConfig
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetLeadGenDirectCRMIntegrationConfig(ctx context.Context, client *core.Client, id string, params GetLeadGenDirectCRMIntegrationConfigParams) (*objects.LeadGenDirectCRMIntegrationConfig, error) {
	var out objects.LeadGenDirectCRMIntegrationConfig
	call := GetLeadGenDirectCRMIntegrationConfigBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
