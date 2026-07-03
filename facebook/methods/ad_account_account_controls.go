package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type CreateAdAccountAccountControlsAccountControlsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params CreateAdAccountAccountControlsAccountControlsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func CreateAdAccountAccountControlsAccountControlsBatchCall(id string, params CreateAdAccountAccountControlsAccountControlsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "account_controls"), params.ToParams(), options...)
}

func NewCreateAdAccountAccountControlsAccountControlsBatchRequest(id string, params CreateAdAccountAccountControlsAccountControlsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdAccountAccountControlsPost] {
	return core.NewBatchRequest[objects.AdAccountAccountControlsPost](CreateAdAccountAccountControlsAccountControlsBatchCall(id, params, options...))
}

func DecodeCreateAdAccountAccountControlsAccountControlsBatchResponse(response *core.BatchResponse) (*objects.AdAccountAccountControlsPost, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdAccountAccountControlsPost
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAccountAccountControlsAccountControls(ctx context.Context, client *core.Client, id string, params CreateAdAccountAccountControlsAccountControlsParams) (*objects.AdAccountAccountControlsPost, error) {
	var out objects.AdAccountAccountControlsPost
	call := CreateAdAccountAccountControlsAccountControlsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
