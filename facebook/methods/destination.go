package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetDestinationChannelsToIntegrityStatusParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetDestinationChannelsToIntegrityStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetDestinationChannelsToIntegrityStatusBatchCall(id string, params GetDestinationChannelsToIntegrityStatusParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "channels_to_integrity_status"), params.ToParams(), options...)
}

func NewGetDestinationChannelsToIntegrityStatusBatchRequest(id string, params GetDestinationChannelsToIntegrityStatusParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]] {
	return core.NewBatchRequest[core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]](GetDestinationChannelsToIntegrityStatusBatchCall(id, params, options...))
}

func DecodeGetDestinationChannelsToIntegrityStatusBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetDestinationChannelsToIntegrityStatus(ctx context.Context, client *core.Client, id string, params GetDestinationChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
	var out core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]
	call := GetDestinationChannelsToIntegrityStatusBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetDestinationOverrideDetailsParams struct {
	Keys  *[]string                                      `facebook:"keys"`
	Type  *enums.DestinationoverrideDetailsTypeEnumParam `facebook:"type"`
	Extra core.Params                                    `facebook:"-"`
}

func (params GetDestinationOverrideDetailsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Keys != nil {
		out["keys"] = *params.Keys
	}
	if params.Type != nil {
		out["type"] = *params.Type
	}
	return out
}

func GetDestinationOverrideDetailsBatchCall(id string, params GetDestinationOverrideDetailsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "override_details"), params.ToParams(), options...)
}

func NewGetDestinationOverrideDetailsBatchRequest(id string, params GetDestinationOverrideDetailsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.OverrideDetails]] {
	return core.NewBatchRequest[core.Cursor[objects.OverrideDetails]](GetDestinationOverrideDetailsBatchCall(id, params, options...))
}

func DecodeGetDestinationOverrideDetailsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.OverrideDetails], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.OverrideDetails]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetDestinationOverrideDetails(ctx context.Context, client *core.Client, id string, params GetDestinationOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], error) {
	var out core.Cursor[objects.OverrideDetails]
	call := GetDestinationOverrideDetailsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetDestinationVideosMetadataParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetDestinationVideosMetadataParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetDestinationVideosMetadataBatchCall(id string, params GetDestinationVideosMetadataParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "videos_metadata"), params.ToParams(), options...)
}

func NewGetDestinationVideosMetadataBatchRequest(id string, params GetDestinationVideosMetadataParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.DynamicVideoMetadata]] {
	return core.NewBatchRequest[core.Cursor[objects.DynamicVideoMetadata]](GetDestinationVideosMetadataBatchCall(id, params, options...))
}

func DecodeGetDestinationVideosMetadataBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.DynamicVideoMetadata], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.DynamicVideoMetadata]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetDestinationVideosMetadata(ctx context.Context, client *core.Client, id string, params GetDestinationVideosMetadataParams) (*core.Cursor[objects.DynamicVideoMetadata], error) {
	var out core.Cursor[objects.DynamicVideoMetadata]
	call := GetDestinationVideosMetadataBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetDestinationParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetDestinationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetDestinationBatchCall(id string, params GetDestinationParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetDestinationBatchRequest(id string, params GetDestinationParams, options ...core.BatchOption) *core.BatchRequest[objects.Destination] {
	return core.NewBatchRequest[objects.Destination](GetDestinationBatchCall(id, params, options...))
}

func DecodeGetDestinationBatchResponse(response *core.BatchResponse) (*objects.Destination, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Destination
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetDestination(ctx context.Context, client *core.Client, id string, params GetDestinationParams) (*objects.Destination, error) {
	var out objects.Destination
	call := GetDestinationBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
