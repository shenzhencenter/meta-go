package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
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

func GetBusinessObjectTransferOwnershipAgreementBatchCall(id string, params GetBusinessObjectTransferOwnershipAgreementParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetBusinessObjectTransferOwnershipAgreementBatchRequest(id string, params GetBusinessObjectTransferOwnershipAgreementParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessObjectTransferOwnershipAgreement] {
	return core.NewBatchRequest[objects.BusinessObjectTransferOwnershipAgreement](GetBusinessObjectTransferOwnershipAgreementBatchCall(id, params, options...))
}

func DecodeGetBusinessObjectTransferOwnershipAgreementBatchResponse(response *core.BatchResponse) (*objects.BusinessObjectTransferOwnershipAgreement, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessObjectTransferOwnershipAgreement
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessObjectTransferOwnershipAgreement(ctx context.Context, client *core.Client, id string, params GetBusinessObjectTransferOwnershipAgreementParams) (*objects.BusinessObjectTransferOwnershipAgreement, error) {
	var out objects.BusinessObjectTransferOwnershipAgreement
	call := GetBusinessObjectTransferOwnershipAgreementBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
