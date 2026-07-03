package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetDynamicVideoMetadataParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetDynamicVideoMetadataParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetDynamicVideoMetadataBatchCall(id string, params GetDynamicVideoMetadataParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetDynamicVideoMetadataBatchRequest(id string, params GetDynamicVideoMetadataParams, options ...core.BatchOption) *core.BatchRequest[objects.DynamicVideoMetadata] {
	return core.NewBatchRequest[objects.DynamicVideoMetadata](GetDynamicVideoMetadataBatchCall(id, params, options...))
}

func DecodeGetDynamicVideoMetadataBatchResponse(response *core.BatchResponse) (*objects.DynamicVideoMetadata, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.DynamicVideoMetadata
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetDynamicVideoMetadata(ctx context.Context, client *core.Client, id string, params GetDynamicVideoMetadataParams) (*objects.DynamicVideoMetadata, error) {
	var out objects.DynamicVideoMetadata
	call := GetDynamicVideoMetadataBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
