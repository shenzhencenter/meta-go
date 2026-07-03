package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetExtendedCreditApplicationParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetExtendedCreditApplicationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetExtendedCreditApplicationBatchCall(id string, params GetExtendedCreditApplicationParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetExtendedCreditApplicationBatchRequest(id string, params GetExtendedCreditApplicationParams, options ...core.BatchOption) *core.BatchRequest[objects.ExtendedCreditApplication] {
	return core.NewBatchRequest[objects.ExtendedCreditApplication](GetExtendedCreditApplicationBatchCall(id, params, options...))
}

func DecodeGetExtendedCreditApplicationBatchResponse(response *core.BatchResponse) (*objects.ExtendedCreditApplication, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ExtendedCreditApplication
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetExtendedCreditApplication(ctx context.Context, client *core.Client, id string, params GetExtendedCreditApplicationParams) (*objects.ExtendedCreditApplication, error) {
	var out objects.ExtendedCreditApplication
	call := GetExtendedCreditApplicationBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
