package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetSavedAudienceParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetSavedAudienceParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetSavedAudienceBatchCall(id string, params GetSavedAudienceParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetSavedAudienceBatchRequest(id string, params GetSavedAudienceParams, options ...core.BatchOption) *core.BatchRequest[objects.SavedAudience] {
	return core.NewBatchRequest[objects.SavedAudience](GetSavedAudienceBatchCall(id, params, options...))
}

func DecodeGetSavedAudienceBatchResponse(response *core.BatchResponse) (*objects.SavedAudience, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.SavedAudience
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetSavedAudienceWithResponse(ctx context.Context, client *core.Client, id string, params GetSavedAudienceParams) (*objects.SavedAudience, *core.Response, error) {
	var out objects.SavedAudience
	call := GetSavedAudienceBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetSavedAudience(ctx context.Context, client *core.Client, id string, params GetSavedAudienceParams) (*objects.SavedAudience, error) {
	out, _, err := GetSavedAudienceWithResponse(ctx, client, id, params)
	return out, err
}
