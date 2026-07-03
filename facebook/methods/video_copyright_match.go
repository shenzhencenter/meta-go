package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetVideoCopyrightMatchParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetVideoCopyrightMatchParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetVideoCopyrightMatchBatchCall(id string, params GetVideoCopyrightMatchParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetVideoCopyrightMatchBatchRequest(id string, params GetVideoCopyrightMatchParams, options ...core.BatchOption) *core.BatchRequest[objects.VideoCopyrightMatch] {
	return core.NewBatchRequest[objects.VideoCopyrightMatch](GetVideoCopyrightMatchBatchCall(id, params, options...))
}

func DecodeGetVideoCopyrightMatchBatchResponse(response *core.BatchResponse) (*objects.VideoCopyrightMatch, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.VideoCopyrightMatch
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetVideoCopyrightMatchWithResponse(ctx context.Context, client *core.Client, id string, params GetVideoCopyrightMatchParams) (*objects.VideoCopyrightMatch, *core.Response, error) {
	var out objects.VideoCopyrightMatch
	call := GetVideoCopyrightMatchBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetVideoCopyrightMatch(ctx context.Context, client *core.Client, id string, params GetVideoCopyrightMatchParams) (*objects.VideoCopyrightMatch, error) {
	out, _, err := GetVideoCopyrightMatchWithResponse(ctx, client, id, params)
	return out, err
}
