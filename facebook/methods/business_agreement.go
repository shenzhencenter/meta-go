package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetBusinessAgreementParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessAgreementParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessAgreement(ctx context.Context, client *core.Client, id string, params GetBusinessAgreementParams) (*objects.BusinessAgreement, error) {
	var out objects.BusinessAgreement
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateBusinessAgreementParams struct {
	AssetID       *core.ID                              `facebook:"asset_id"`
	RequestStatus *enums.BusinessagreementRequestStatus `facebook:"request_status"`
	Extra         core.Params                           `facebook:"-"`
}

func (params UpdateBusinessAgreementParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AssetID != nil {
		out["asset_id"] = *params.AssetID
	}
	if params.RequestStatus != nil {
		out["request_status"] = *params.RequestStatus
	}
	return out
}

func UpdateBusinessAgreement(ctx context.Context, client *core.Client, id string, params UpdateBusinessAgreementParams) (*objects.BusinessAgreement, error) {
	var out objects.BusinessAgreement
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
