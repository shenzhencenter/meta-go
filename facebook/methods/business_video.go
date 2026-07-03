package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetBusinessVideoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessVideoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessVideoBatchCall(id string, params GetBusinessVideoParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetBusinessVideoBatchRequest(id string, params GetBusinessVideoParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessVideo] {
	return core.NewBatchRequest[objects.BusinessVideo](GetBusinessVideoBatchCall(id, params, options...))
}

func DecodeGetBusinessVideoBatchResponse(response *core.BatchResponse) (*objects.BusinessVideo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessVideo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessVideo(ctx context.Context, client *core.Client, id string, params GetBusinessVideoParams) (*objects.BusinessVideo, error) {
	var out objects.BusinessVideo
	call := GetBusinessVideoBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
