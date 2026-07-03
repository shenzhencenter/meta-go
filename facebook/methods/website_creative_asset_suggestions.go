package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetWebsiteCreativeAssetSuggestionsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWebsiteCreativeAssetSuggestionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWebsiteCreativeAssetSuggestionsBatchCall(id string, params GetWebsiteCreativeAssetSuggestionsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetWebsiteCreativeAssetSuggestionsBatchRequest(id string, params GetWebsiteCreativeAssetSuggestionsParams, options ...core.BatchOption) *core.BatchRequest[objects.WebsiteCreativeAssetSuggestions] {
	return core.NewBatchRequest[objects.WebsiteCreativeAssetSuggestions](GetWebsiteCreativeAssetSuggestionsBatchCall(id, params, options...))
}

func DecodeGetWebsiteCreativeAssetSuggestionsBatchResponse(response *core.BatchResponse) (*objects.WebsiteCreativeAssetSuggestions, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WebsiteCreativeAssetSuggestions
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWebsiteCreativeAssetSuggestionsWithResponse(ctx context.Context, client *core.Client, id string, params GetWebsiteCreativeAssetSuggestionsParams) (*objects.WebsiteCreativeAssetSuggestions, *core.Response, error) {
	var out objects.WebsiteCreativeAssetSuggestions
	call := GetWebsiteCreativeAssetSuggestionsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetWebsiteCreativeAssetSuggestions(ctx context.Context, client *core.Client, id string, params GetWebsiteCreativeAssetSuggestionsParams) (*objects.WebsiteCreativeAssetSuggestions, error) {
	out, _, err := GetWebsiteCreativeAssetSuggestionsWithResponse(ctx, client, id, params)
	return out, err
}
