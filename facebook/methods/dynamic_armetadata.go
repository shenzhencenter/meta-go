package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetDynamicARMetadataParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetDynamicARMetadataParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetDynamicARMetadataBatchCall(id string, params GetDynamicARMetadataParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetDynamicARMetadataBatchRequest(id string, params GetDynamicARMetadataParams, options ...core.BatchOption) *core.BatchRequest[objects.DynamicARMetadata] {
	return core.NewBatchRequest[objects.DynamicARMetadata](GetDynamicARMetadataBatchCall(id, params, options...))
}

func DecodeGetDynamicARMetadataBatchResponse(response *core.BatchResponse) (*objects.DynamicARMetadata, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.DynamicARMetadata
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetDynamicARMetadata(ctx context.Context, client *core.Client, id string, params GetDynamicARMetadataParams) (*objects.DynamicARMetadata, error) {
	var out objects.DynamicARMetadata
	call := GetDynamicARMetadataBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
