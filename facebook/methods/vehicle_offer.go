package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetVehicleOfferChannelsToIntegrityStatusParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetVehicleOfferChannelsToIntegrityStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetVehicleOfferChannelsToIntegrityStatusBatchCall(id string, params GetVehicleOfferChannelsToIntegrityStatusParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "channels_to_integrity_status"), params.ToParams(), options...)
}

func NewGetVehicleOfferChannelsToIntegrityStatusBatchRequest(id string, params GetVehicleOfferChannelsToIntegrityStatusParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]] {
	return core.NewBatchRequest[core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]](GetVehicleOfferChannelsToIntegrityStatusBatchCall(id, params, options...))
}

func DecodeGetVehicleOfferChannelsToIntegrityStatusBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
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

func GetVehicleOfferChannelsToIntegrityStatus(ctx context.Context, client *core.Client, id string, params GetVehicleOfferChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
	var out core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]
	call := GetVehicleOfferChannelsToIntegrityStatusBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetVehicleOfferOverrideDetailsParams struct {
	Keys  *[]string                                       `facebook:"keys"`
	Type  *enums.VehicleofferoverrideDetailsTypeEnumParam `facebook:"type"`
	Extra core.Params                                     `facebook:"-"`
}

func (params GetVehicleOfferOverrideDetailsParams) ToParams() core.Params {
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

func GetVehicleOfferOverrideDetailsBatchCall(id string, params GetVehicleOfferOverrideDetailsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "override_details"), params.ToParams(), options...)
}

func NewGetVehicleOfferOverrideDetailsBatchRequest(id string, params GetVehicleOfferOverrideDetailsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.OverrideDetails]] {
	return core.NewBatchRequest[core.Cursor[objects.OverrideDetails]](GetVehicleOfferOverrideDetailsBatchCall(id, params, options...))
}

func DecodeGetVehicleOfferOverrideDetailsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.OverrideDetails], error) {
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

func GetVehicleOfferOverrideDetails(ctx context.Context, client *core.Client, id string, params GetVehicleOfferOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], error) {
	var out core.Cursor[objects.OverrideDetails]
	call := GetVehicleOfferOverrideDetailsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetVehicleOfferVideosMetadataParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetVehicleOfferVideosMetadataParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetVehicleOfferVideosMetadataBatchCall(id string, params GetVehicleOfferVideosMetadataParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "videos_metadata"), params.ToParams(), options...)
}

func NewGetVehicleOfferVideosMetadataBatchRequest(id string, params GetVehicleOfferVideosMetadataParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.DynamicVideoMetadata]] {
	return core.NewBatchRequest[core.Cursor[objects.DynamicVideoMetadata]](GetVehicleOfferVideosMetadataBatchCall(id, params, options...))
}

func DecodeGetVehicleOfferVideosMetadataBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.DynamicVideoMetadata], error) {
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

func GetVehicleOfferVideosMetadata(ctx context.Context, client *core.Client, id string, params GetVehicleOfferVideosMetadataParams) (*core.Cursor[objects.DynamicVideoMetadata], error) {
	var out core.Cursor[objects.DynamicVideoMetadata]
	call := GetVehicleOfferVideosMetadataBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetVehicleOfferParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetVehicleOfferParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetVehicleOfferBatchCall(id string, params GetVehicleOfferParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetVehicleOfferBatchRequest(id string, params GetVehicleOfferParams, options ...core.BatchOption) *core.BatchRequest[objects.VehicleOffer] {
	return core.NewBatchRequest[objects.VehicleOffer](GetVehicleOfferBatchCall(id, params, options...))
}

func DecodeGetVehicleOfferBatchResponse(response *core.BatchResponse) (*objects.VehicleOffer, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.VehicleOffer
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetVehicleOffer(ctx context.Context, client *core.Client, id string, params GetVehicleOfferParams) (*objects.VehicleOffer, error) {
	var out objects.VehicleOffer
	call := GetVehicleOfferBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
