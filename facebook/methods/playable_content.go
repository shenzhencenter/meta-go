package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetPlayableContentParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPlayableContentParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPlayableContentBatchCall(id string, params GetPlayableContentParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetPlayableContentBatchRequest(id string, params GetPlayableContentParams, options ...core.BatchOption) *core.BatchRequest[objects.PlayableContent] {
	return core.NewBatchRequest[objects.PlayableContent](GetPlayableContentBatchCall(id, params, options...))
}

func DecodeGetPlayableContentBatchResponse(response *core.BatchResponse) (*objects.PlayableContent, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PlayableContent
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPlayableContent(ctx context.Context, client *core.Client, id string, params GetPlayableContentParams) (*objects.PlayableContent, error) {
	var out objects.PlayableContent
	call := GetPlayableContentBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
