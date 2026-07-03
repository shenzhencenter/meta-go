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

func GetBusinessAssetSharingAgreementBatchCall(id string, params GetBusinessAssetSharingAgreementParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetBusinessAssetSharingAgreementBatchRequest(id string, params GetBusinessAssetSharingAgreementParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessAssetSharingAgreement] {
	return core.NewBatchRequest[objects.BusinessAssetSharingAgreement](GetBusinessAssetSharingAgreementBatchCall(id, params, options...))
}

func DecodeGetBusinessAssetSharingAgreementBatchResponse(response *core.BatchResponse) (*objects.BusinessAssetSharingAgreement, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessAssetSharingAgreement
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessAssetSharingAgreement(ctx context.Context, client *core.Client, id string, params GetBusinessAssetSharingAgreementParams) (*objects.BusinessAssetSharingAgreement, error) {
	var out objects.BusinessAssetSharingAgreement
	call := GetBusinessAssetSharingAgreementBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func UpdateBusinessAssetSharingAgreementBatchCall(id string, params UpdateBusinessAssetSharingAgreementParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateBusinessAssetSharingAgreementBatchRequest(id string, params UpdateBusinessAssetSharingAgreementParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessAssetSharingAgreement] {
	return core.NewBatchRequest[objects.BusinessAssetSharingAgreement](UpdateBusinessAssetSharingAgreementBatchCall(id, params, options...))
}

func DecodeUpdateBusinessAssetSharingAgreementBatchResponse(response *core.BatchResponse) (*objects.BusinessAssetSharingAgreement, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessAssetSharingAgreement
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateBusinessAssetSharingAgreement(ctx context.Context, client *core.Client, id string, params UpdateBusinessAssetSharingAgreementParams) (*objects.BusinessAssetSharingAgreement, error) {
	var out objects.BusinessAssetSharingAgreement
	call := UpdateBusinessAssetSharingAgreementBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
