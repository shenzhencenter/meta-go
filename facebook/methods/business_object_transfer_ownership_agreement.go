package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetBusinessObjectTransferOwnershipAgreementParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessObjectTransferOwnershipAgreementParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessObjectTransferOwnershipAgreement(ctx context.Context, client *core.Client, id string, params GetBusinessObjectTransferOwnershipAgreementParams) (*objects.BusinessObjectTransferOwnershipAgreement, error) {
	var out objects.BusinessObjectTransferOwnershipAgreement
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
