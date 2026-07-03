package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetBusinessTagParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessTagParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessTagBatchCall(id string, params GetBusinessTagParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetBusinessTagBatchRequest(id string, params GetBusinessTagParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessTag] {
	return core.NewBatchRequest[objects.BusinessTag](GetBusinessTagBatchCall(id, params, options...))
}

func DecodeGetBusinessTagBatchResponse(response *core.BatchResponse) (*objects.BusinessTag, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessTag
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessTagWithResponse(ctx context.Context, client *core.Client, id string, params GetBusinessTagParams) (*objects.BusinessTag, *core.Response, error) {
	var out objects.BusinessTag
	call := GetBusinessTagBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetBusinessTag(ctx context.Context, client *core.Client, id string, params GetBusinessTagParams) (*objects.BusinessTag, error) {
	out, _, err := GetBusinessTagWithResponse(ctx, client, id, params)
	return out, err
}
