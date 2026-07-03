package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetFBImageCopyrightMatchParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetFBImageCopyrightMatchParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetFBImageCopyrightMatchBatchCall(id string, params GetFBImageCopyrightMatchParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetFBImageCopyrightMatchBatchRequest(id string, params GetFBImageCopyrightMatchParams, options ...core.BatchOption) *core.BatchRequest[objects.FBImageCopyrightMatch] {
	return core.NewBatchRequest[objects.FBImageCopyrightMatch](GetFBImageCopyrightMatchBatchCall(id, params, options...))
}

func DecodeGetFBImageCopyrightMatchBatchResponse(response *core.BatchResponse) (*objects.FBImageCopyrightMatch, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.FBImageCopyrightMatch
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetFBImageCopyrightMatch(ctx context.Context, client *core.Client, id string, params GetFBImageCopyrightMatchParams) (*objects.FBImageCopyrightMatch, error) {
	var out objects.FBImageCopyrightMatch
	call := GetFBImageCopyrightMatchBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
