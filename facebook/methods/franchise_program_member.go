package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetFranchiseProgramMemberParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetFranchiseProgramMemberParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetFranchiseProgramMemberBatchCall(id string, params GetFranchiseProgramMemberParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetFranchiseProgramMemberBatchRequest(id string, params GetFranchiseProgramMemberParams, options ...core.BatchOption) *core.BatchRequest[objects.FranchiseProgramMember] {
	return core.NewBatchRequest[objects.FranchiseProgramMember](GetFranchiseProgramMemberBatchCall(id, params, options...))
}

func DecodeGetFranchiseProgramMemberBatchResponse(response *core.BatchResponse) (*objects.FranchiseProgramMember, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.FranchiseProgramMember
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetFranchiseProgramMemberWithResponse(ctx context.Context, client *core.Client, id string, params GetFranchiseProgramMemberParams) (*objects.FranchiseProgramMember, *core.Response, error) {
	var out objects.FranchiseProgramMember
	call := GetFranchiseProgramMemberBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetFranchiseProgramMember(ctx context.Context, client *core.Client, id string, params GetFranchiseProgramMemberParams) (*objects.FranchiseProgramMember, error) {
	out, _, err := GetFranchiseProgramMemberWithResponse(ctx, client, id, params)
	return out, err
}
