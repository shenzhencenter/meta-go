package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAutomotiveModelChannelsToIntegrityStatusParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAutomotiveModelChannelsToIntegrityStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAutomotiveModelChannelsToIntegrityStatusBatchCall(id string, params GetAutomotiveModelChannelsToIntegrityStatusParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "channels_to_integrity_status"), params.ToParams(), options...)
}

func NewGetAutomotiveModelChannelsToIntegrityStatusBatchRequest(id string, params GetAutomotiveModelChannelsToIntegrityStatusParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]] {
	return core.NewBatchRequest[core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]](GetAutomotiveModelChannelsToIntegrityStatusBatchCall(id, params, options...))
}

func DecodeGetAutomotiveModelChannelsToIntegrityStatusBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
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

func GetAutomotiveModelChannelsToIntegrityStatusWithResponse(ctx context.Context, client *core.Client, id string, params GetAutomotiveModelChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], *core.Response, error) {
	var out core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]
	call := GetAutomotiveModelChannelsToIntegrityStatusBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAutomotiveModelChannelsToIntegrityStatus(ctx context.Context, client *core.Client, id string, params GetAutomotiveModelChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
	out, _, err := GetAutomotiveModelChannelsToIntegrityStatusWithResponse(ctx, client, id, params)
	return out, err
}

type GetAutomotiveModelOverrideDetailsParams struct {
	Keys  *[]string                                          `facebook:"keys"`
	Type  *enums.AutomotivemodeloverrideDetailsTypeEnumParam `facebook:"type"`
	Extra core.Params                                        `facebook:"-"`
}

func (params GetAutomotiveModelOverrideDetailsParams) ToParams() core.Params {
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

func GetAutomotiveModelOverrideDetailsBatchCall(id string, params GetAutomotiveModelOverrideDetailsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "override_details"), params.ToParams(), options...)
}

func NewGetAutomotiveModelOverrideDetailsBatchRequest(id string, params GetAutomotiveModelOverrideDetailsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.OverrideDetails]] {
	return core.NewBatchRequest[core.Cursor[objects.OverrideDetails]](GetAutomotiveModelOverrideDetailsBatchCall(id, params, options...))
}

func DecodeGetAutomotiveModelOverrideDetailsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.OverrideDetails], error) {
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

func GetAutomotiveModelOverrideDetailsWithResponse(ctx context.Context, client *core.Client, id string, params GetAutomotiveModelOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], *core.Response, error) {
	var out core.Cursor[objects.OverrideDetails]
	call := GetAutomotiveModelOverrideDetailsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAutomotiveModelOverrideDetails(ctx context.Context, client *core.Client, id string, params GetAutomotiveModelOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], error) {
	out, _, err := GetAutomotiveModelOverrideDetailsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAutomotiveModelVideosMetadataParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAutomotiveModelVideosMetadataParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAutomotiveModelVideosMetadataBatchCall(id string, params GetAutomotiveModelVideosMetadataParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "videos_metadata"), params.ToParams(), options...)
}

func NewGetAutomotiveModelVideosMetadataBatchRequest(id string, params GetAutomotiveModelVideosMetadataParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.DynamicVideoMetadata]] {
	return core.NewBatchRequest[core.Cursor[objects.DynamicVideoMetadata]](GetAutomotiveModelVideosMetadataBatchCall(id, params, options...))
}

func DecodeGetAutomotiveModelVideosMetadataBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.DynamicVideoMetadata], error) {
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

func GetAutomotiveModelVideosMetadataWithResponse(ctx context.Context, client *core.Client, id string, params GetAutomotiveModelVideosMetadataParams) (*core.Cursor[objects.DynamicVideoMetadata], *core.Response, error) {
	var out core.Cursor[objects.DynamicVideoMetadata]
	call := GetAutomotiveModelVideosMetadataBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAutomotiveModelVideosMetadata(ctx context.Context, client *core.Client, id string, params GetAutomotiveModelVideosMetadataParams) (*core.Cursor[objects.DynamicVideoMetadata], error) {
	out, _, err := GetAutomotiveModelVideosMetadataWithResponse(ctx, client, id, params)
	return out, err
}

type GetAutomotiveModelParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAutomotiveModelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAutomotiveModelBatchCall(id string, params GetAutomotiveModelParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAutomotiveModelBatchRequest(id string, params GetAutomotiveModelParams, options ...core.BatchOption) *core.BatchRequest[objects.AutomotiveModel] {
	return core.NewBatchRequest[objects.AutomotiveModel](GetAutomotiveModelBatchCall(id, params, options...))
}

func DecodeGetAutomotiveModelBatchResponse(response *core.BatchResponse) (*objects.AutomotiveModel, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AutomotiveModel
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAutomotiveModelWithResponse(ctx context.Context, client *core.Client, id string, params GetAutomotiveModelParams) (*objects.AutomotiveModel, *core.Response, error) {
	var out objects.AutomotiveModel
	call := GetAutomotiveModelBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAutomotiveModel(ctx context.Context, client *core.Client, id string, params GetAutomotiveModelParams) (*objects.AutomotiveModel, error) {
	out, _, err := GetAutomotiveModelWithResponse(ctx, client, id, params)
	return out, err
}
