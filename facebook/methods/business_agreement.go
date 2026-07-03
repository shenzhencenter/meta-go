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

func GetBusinessAgreementBatchCall(id string, params GetBusinessAgreementParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetBusinessAgreementBatchRequest(id string, params GetBusinessAgreementParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessAgreement] {
	return core.NewBatchRequest[objects.BusinessAgreement](GetBusinessAgreementBatchCall(id, params, options...))
}

func DecodeGetBusinessAgreementBatchResponse(response *core.BatchResponse) (*objects.BusinessAgreement, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessAgreement
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessAgreementWithResponse(ctx context.Context, client *core.Client, id string, params GetBusinessAgreementParams) (*objects.BusinessAgreement, *core.Response, error) {
	var out objects.BusinessAgreement
	call := GetBusinessAgreementBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetBusinessAgreement(ctx context.Context, client *core.Client, id string, params GetBusinessAgreementParams) (*objects.BusinessAgreement, error) {
	out, _, err := GetBusinessAgreementWithResponse(ctx, client, id, params)
	return out, err
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

func UpdateBusinessAgreementBatchCall(id string, params UpdateBusinessAgreementParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateBusinessAgreementBatchRequest(id string, params UpdateBusinessAgreementParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessAgreement] {
	return core.NewBatchRequest[objects.BusinessAgreement](UpdateBusinessAgreementBatchCall(id, params, options...))
}

func DecodeUpdateBusinessAgreementBatchResponse(response *core.BatchResponse) (*objects.BusinessAgreement, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessAgreement
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateBusinessAgreementWithResponse(ctx context.Context, client *core.Client, id string, params UpdateBusinessAgreementParams) (*objects.BusinessAgreement, *core.Response, error) {
	var out objects.BusinessAgreement
	call := UpdateBusinessAgreementBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateBusinessAgreement(ctx context.Context, client *core.Client, id string, params UpdateBusinessAgreementParams) (*objects.BusinessAgreement, error) {
	out, _, err := UpdateBusinessAgreementWithResponse(ctx, client, id, params)
	return out, err
}
