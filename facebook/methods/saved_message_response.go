package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetSavedMessageResponseParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetSavedMessageResponseParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetSavedMessageResponseBatchCall(id string, params GetSavedMessageResponseParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetSavedMessageResponseBatchRequest(id string, params GetSavedMessageResponseParams, options ...core.BatchOption) *core.BatchRequest[objects.SavedMessageResponse] {
	return core.NewBatchRequest[objects.SavedMessageResponse](GetSavedMessageResponseBatchCall(id, params, options...))
}

func DecodeGetSavedMessageResponseBatchResponse(response *core.BatchResponse) (*objects.SavedMessageResponse, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.SavedMessageResponse
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetSavedMessageResponse(ctx context.Context, client *core.Client, id string, params GetSavedMessageResponseParams) (*objects.SavedMessageResponse, error) {
	var out objects.SavedMessageResponse
	call := GetSavedMessageResponseBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
