package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCRMAddressParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCRMAddressParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCRMAddressBatchCall(id string, params GetCRMAddressParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCRMAddressBatchRequest(id string, params GetCRMAddressParams, options ...core.BatchOption) *core.BatchRequest[objects.CRMAddress] {
	return core.NewBatchRequest[objects.CRMAddress](GetCRMAddressBatchCall(id, params, options...))
}

func DecodeGetCRMAddressBatchResponse(response *core.BatchResponse) (*objects.CRMAddress, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CRMAddress
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCRMAddress(ctx context.Context, client *core.Client, id string, params GetCRMAddressParams) (*objects.CRMAddress, error) {
	var out objects.CRMAddress
	call := GetCRMAddressBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
