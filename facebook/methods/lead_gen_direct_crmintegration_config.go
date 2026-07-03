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

func GetLeadGenDirectCRMIntegrationConfig(ctx context.Context, client *core.Client, id string, params GetLeadGenDirectCRMIntegrationConfigParams) (*objects.LeadGenDirectCRMIntegrationConfig, error) {
	var out objects.LeadGenDirectCRMIntegrationConfig
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
