package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdDraftParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdDraftParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdDraftBatchCall(id string, params GetAdDraftParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdDraftBatchRequest(id string, params GetAdDraftParams, options ...core.BatchOption) *core.BatchRequest[objects.AdDraft] {
	return core.NewBatchRequest[objects.AdDraft](GetAdDraftBatchCall(id, params, options...))
}

func DecodeGetAdDraftBatchResponse(response *core.BatchResponse) (*objects.AdDraft, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdDraft
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdDraft(ctx context.Context, client *core.Client, id string, params GetAdDraftParams) (*objects.AdDraft, error) {
	var out objects.AdDraft
	call := GetAdDraftBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
