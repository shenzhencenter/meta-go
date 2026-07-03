package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetLeadGenDataDraftParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLeadGenDataDraftParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLeadGenDataDraftBatchCall(id string, params GetLeadGenDataDraftParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetLeadGenDataDraftBatchRequest(id string, params GetLeadGenDataDraftParams, options ...core.BatchOption) *core.BatchRequest[objects.LeadGenDataDraft] {
	return core.NewBatchRequest[objects.LeadGenDataDraft](GetLeadGenDataDraftBatchCall(id, params, options...))
}

func DecodeGetLeadGenDataDraftBatchResponse(response *core.BatchResponse) (*objects.LeadGenDataDraft, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.LeadGenDataDraft
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetLeadGenDataDraft(ctx context.Context, client *core.Client, id string, params GetLeadGenDataDraftParams) (*objects.LeadGenDataDraft, error) {
	var out objects.LeadGenDataDraft
	call := GetLeadGenDataDraftBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
