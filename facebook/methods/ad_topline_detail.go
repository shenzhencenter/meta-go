package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdToplineDetailParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdToplineDetailParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdToplineDetailBatchCall(id string, params GetAdToplineDetailParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdToplineDetailBatchRequest(id string, params GetAdToplineDetailParams, options ...core.BatchOption) *core.BatchRequest[objects.AdToplineDetail] {
	return core.NewBatchRequest[objects.AdToplineDetail](GetAdToplineDetailBatchCall(id, params, options...))
}

func DecodeGetAdToplineDetailBatchResponse(response *core.BatchResponse) (*objects.AdToplineDetail, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdToplineDetail
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdToplineDetail(ctx context.Context, client *core.Client, id string, params GetAdToplineDetailParams) (*objects.AdToplineDetail, error) {
	var out objects.AdToplineDetail
	call := GetAdToplineDetailBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
