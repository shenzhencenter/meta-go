package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCopyrightOwnershipTransferParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCopyrightOwnershipTransferParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCopyrightOwnershipTransferBatchCall(id string, params GetCopyrightOwnershipTransferParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCopyrightOwnershipTransferBatchRequest(id string, params GetCopyrightOwnershipTransferParams, options ...core.BatchOption) *core.BatchRequest[objects.CopyrightOwnershipTransfer] {
	return core.NewBatchRequest[objects.CopyrightOwnershipTransfer](GetCopyrightOwnershipTransferBatchCall(id, params, options...))
}

func DecodeGetCopyrightOwnershipTransferBatchResponse(response *core.BatchResponse) (*objects.CopyrightOwnershipTransfer, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CopyrightOwnershipTransfer
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCopyrightOwnershipTransfer(ctx context.Context, client *core.Client, id string, params GetCopyrightOwnershipTransferParams) (*objects.CopyrightOwnershipTransfer, error) {
	var out objects.CopyrightOwnershipTransfer
	call := GetCopyrightOwnershipTransferBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
