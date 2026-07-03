package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdSavedKeywordsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdSavedKeywordsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdSavedKeywordsBatchCall(id string, params GetAdSavedKeywordsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdSavedKeywordsBatchRequest(id string, params GetAdSavedKeywordsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdSavedKeywords] {
	return core.NewBatchRequest[objects.AdSavedKeywords](GetAdSavedKeywordsBatchCall(id, params, options...))
}

func DecodeGetAdSavedKeywordsBatchResponse(response *core.BatchResponse) (*objects.AdSavedKeywords, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdSavedKeywords
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdSavedKeywords(ctx context.Context, client *core.Client, id string, params GetAdSavedKeywordsParams) (*objects.AdSavedKeywords, error) {
	var out objects.AdSavedKeywords
	call := GetAdSavedKeywordsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
