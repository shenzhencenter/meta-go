package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetArAdsDataContainerParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetArAdsDataContainerParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetArAdsDataContainerBatchCall(id string, params GetArAdsDataContainerParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetArAdsDataContainerBatchRequest(id string, params GetArAdsDataContainerParams, options ...core.BatchOption) *core.BatchRequest[objects.ArAdsDataContainer] {
	return core.NewBatchRequest[objects.ArAdsDataContainer](GetArAdsDataContainerBatchCall(id, params, options...))
}

func DecodeGetArAdsDataContainerBatchResponse(response *core.BatchResponse) (*objects.ArAdsDataContainer, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ArAdsDataContainer
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetArAdsDataContainer(ctx context.Context, client *core.Client, id string, params GetArAdsDataContainerParams) (*objects.ArAdsDataContainer, error) {
	var out objects.ArAdsDataContainer
	call := GetArAdsDataContainerBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
