package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetFranchiseProgramParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetFranchiseProgramParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetFranchiseProgramBatchCall(id string, params GetFranchiseProgramParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetFranchiseProgramBatchRequest(id string, params GetFranchiseProgramParams, options ...core.BatchOption) *core.BatchRequest[objects.FranchiseProgram] {
	return core.NewBatchRequest[objects.FranchiseProgram](GetFranchiseProgramBatchCall(id, params, options...))
}

func DecodeGetFranchiseProgramBatchResponse(response *core.BatchResponse) (*objects.FranchiseProgram, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.FranchiseProgram
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetFranchiseProgramWithResponse(ctx context.Context, client *core.Client, id string, params GetFranchiseProgramParams) (*objects.FranchiseProgram, *core.Response, error) {
	var out objects.FranchiseProgram
	call := GetFranchiseProgramBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetFranchiseProgram(ctx context.Context, client *core.Client, id string, params GetFranchiseProgramParams) (*objects.FranchiseProgram, error) {
	out, _, err := GetFranchiseProgramWithResponse(ctx, client, id, params)
	return out, err
}
