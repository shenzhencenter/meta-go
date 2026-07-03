package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetImageReferenceMatchParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetImageReferenceMatchParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetImageReferenceMatchBatchCall(id string, params GetImageReferenceMatchParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetImageReferenceMatchBatchRequest(id string, params GetImageReferenceMatchParams, options ...core.BatchOption) *core.BatchRequest[objects.ImageReferenceMatch] {
	return core.NewBatchRequest[objects.ImageReferenceMatch](GetImageReferenceMatchBatchCall(id, params, options...))
}

func DecodeGetImageReferenceMatchBatchResponse(response *core.BatchResponse) (*objects.ImageReferenceMatch, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ImageReferenceMatch
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetImageReferenceMatchWithResponse(ctx context.Context, client *core.Client, id string, params GetImageReferenceMatchParams) (*objects.ImageReferenceMatch, *core.Response, error) {
	var out objects.ImageReferenceMatch
	call := GetImageReferenceMatchBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetImageReferenceMatch(ctx context.Context, client *core.Client, id string, params GetImageReferenceMatchParams) (*objects.ImageReferenceMatch, error) {
	out, _, err := GetImageReferenceMatchWithResponse(ctx, client, id, params)
	return out, err
}
