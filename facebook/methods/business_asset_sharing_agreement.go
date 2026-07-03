package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetBusinessAssetSharingAgreementParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessAssetSharingAgreementParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessAssetSharingAgreement(ctx context.Context, client *core.Client, id string, params GetBusinessAssetSharingAgreementParams) (*objects.BusinessAssetSharingAgreement, error) {
	var out objects.BusinessAssetSharingAgreement
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateBusinessAssetSharingAgreementParams struct {
	RequestResponse *string     `facebook:"request_response"`
	Extra           core.Params `facebook:"-"`
}

func (params UpdateBusinessAssetSharingAgreementParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.RequestResponse != nil {
		out["request_response"] = *params.RequestResponse
	}
	return out
}

func UpdateBusinessAssetSharingAgreement(ctx context.Context, client *core.Client, id string, params UpdateBusinessAssetSharingAgreementParams) (*objects.BusinessAssetSharingAgreement, error) {
	var out objects.BusinessAssetSharingAgreement
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
