package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetBusinessFranchiseConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessFranchiseConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessFranchiseConfigBatchCall(id string, params GetBusinessFranchiseConfigParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetBusinessFranchiseConfigBatchRequest(id string, params GetBusinessFranchiseConfigParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessFranchiseConfig] {
	return core.NewBatchRequest[objects.BusinessFranchiseConfig](GetBusinessFranchiseConfigBatchCall(id, params, options...))
}

func DecodeGetBusinessFranchiseConfigBatchResponse(response *core.BatchResponse) (*objects.BusinessFranchiseConfig, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessFranchiseConfig
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessFranchiseConfig(ctx context.Context, client *core.Client, id string, params GetBusinessFranchiseConfigParams) (*objects.BusinessFranchiseConfig, error) {
	var out objects.BusinessFranchiseConfig
	call := GetBusinessFranchiseConfigBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
