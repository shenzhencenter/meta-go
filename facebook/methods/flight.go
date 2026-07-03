package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetFlightChannelsToIntegrityStatusParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetFlightChannelsToIntegrityStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetFlightChannelsToIntegrityStatusBatchCall(id string, params GetFlightChannelsToIntegrityStatusParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "channels_to_integrity_status"), params.ToParams(), options...)
}

func NewGetFlightChannelsToIntegrityStatusBatchRequest(id string, params GetFlightChannelsToIntegrityStatusParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]] {
	return core.NewBatchRequest[core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]](GetFlightChannelsToIntegrityStatusBatchCall(id, params, options...))
}

func DecodeGetFlightChannelsToIntegrityStatusBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
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

func GetFlightChannelsToIntegrityStatusWithResponse(ctx context.Context, client *core.Client, id string, params GetFlightChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], *core.Response, error) {
	var out core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]
	call := GetFlightChannelsToIntegrityStatusBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetFlightChannelsToIntegrityStatus(ctx context.Context, client *core.Client, id string, params GetFlightChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
	out, _, err := GetFlightChannelsToIntegrityStatusWithResponse(ctx, client, id, params)
	return out, err
}

type GetFlightOverrideDetailsParams struct {
	Keys  *[]string                                 `facebook:"keys"`
	Type  *enums.FlightoverrideDetailsTypeEnumParam `facebook:"type"`
	Extra core.Params                               `facebook:"-"`
}

func (params GetFlightOverrideDetailsParams) ToParams() core.Params {
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

func GetFlightOverrideDetailsBatchCall(id string, params GetFlightOverrideDetailsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "override_details"), params.ToParams(), options...)
}

func NewGetFlightOverrideDetailsBatchRequest(id string, params GetFlightOverrideDetailsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.OverrideDetails]] {
	return core.NewBatchRequest[core.Cursor[objects.OverrideDetails]](GetFlightOverrideDetailsBatchCall(id, params, options...))
}

func DecodeGetFlightOverrideDetailsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.OverrideDetails], error) {
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

func GetFlightOverrideDetailsWithResponse(ctx context.Context, client *core.Client, id string, params GetFlightOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], *core.Response, error) {
	var out core.Cursor[objects.OverrideDetails]
	call := GetFlightOverrideDetailsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetFlightOverrideDetails(ctx context.Context, client *core.Client, id string, params GetFlightOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], error) {
	out, _, err := GetFlightOverrideDetailsWithResponse(ctx, client, id, params)
	return out, err
}

type GetFlightVideosMetadataParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetFlightVideosMetadataParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetFlightVideosMetadataBatchCall(id string, params GetFlightVideosMetadataParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "videos_metadata"), params.ToParams(), options...)
}

func NewGetFlightVideosMetadataBatchRequest(id string, params GetFlightVideosMetadataParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.DynamicVideoMetadata]] {
	return core.NewBatchRequest[core.Cursor[objects.DynamicVideoMetadata]](GetFlightVideosMetadataBatchCall(id, params, options...))
}

func DecodeGetFlightVideosMetadataBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.DynamicVideoMetadata], error) {
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

func GetFlightVideosMetadataWithResponse(ctx context.Context, client *core.Client, id string, params GetFlightVideosMetadataParams) (*core.Cursor[objects.DynamicVideoMetadata], *core.Response, error) {
	var out core.Cursor[objects.DynamicVideoMetadata]
	call := GetFlightVideosMetadataBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetFlightVideosMetadata(ctx context.Context, client *core.Client, id string, params GetFlightVideosMetadataParams) (*core.Cursor[objects.DynamicVideoMetadata], error) {
	out, _, err := GetFlightVideosMetadataWithResponse(ctx, client, id, params)
	return out, err
}

type GetFlightParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetFlightParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetFlightBatchCall(id string, params GetFlightParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetFlightBatchRequest(id string, params GetFlightParams, options ...core.BatchOption) *core.BatchRequest[objects.Flight] {
	return core.NewBatchRequest[objects.Flight](GetFlightBatchCall(id, params, options...))
}

func DecodeGetFlightBatchResponse(response *core.BatchResponse) (*objects.Flight, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Flight
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetFlightWithResponse(ctx context.Context, client *core.Client, id string, params GetFlightParams) (*objects.Flight, *core.Response, error) {
	var out objects.Flight
	call := GetFlightBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetFlight(ctx context.Context, client *core.Client, id string, params GetFlightParams) (*objects.Flight, error) {
	out, _, err := GetFlightWithResponse(ctx, client, id, params)
	return out, err
}

type UpdateFlightParams struct {
	Currency           *string                   `facebook:"currency"`
	Description        *string                   `facebook:"description"`
	DestinationAirport *string                   `facebook:"destination_airport"`
	DestinationCity    *string                   `facebook:"destination_city"`
	Images             *[]map[string]interface{} `facebook:"images"`
	OriginAirport      *string                   `facebook:"origin_airport"`
	OriginCity         *string                   `facebook:"origin_city"`
	Price              *uint64                   `facebook:"price"`
	URL                *string                   `facebook:"url"`
	Extra              core.Params               `facebook:"-"`
}

func (params UpdateFlightParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Currency != nil {
		out["currency"] = *params.Currency
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.DestinationAirport != nil {
		out["destination_airport"] = *params.DestinationAirport
	}
	if params.DestinationCity != nil {
		out["destination_city"] = *params.DestinationCity
	}
	if params.Images != nil {
		out["images"] = *params.Images
	}
	if params.OriginAirport != nil {
		out["origin_airport"] = *params.OriginAirport
	}
	if params.OriginCity != nil {
		out["origin_city"] = *params.OriginCity
	}
	if params.Price != nil {
		out["price"] = *params.Price
	}
	if params.URL != nil {
		out["url"] = *params.URL
	}
	return out
}

func UpdateFlightBatchCall(id string, params UpdateFlightParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateFlightBatchRequest(id string, params UpdateFlightParams, options ...core.BatchOption) *core.BatchRequest[objects.Flight] {
	return core.NewBatchRequest[objects.Flight](UpdateFlightBatchCall(id, params, options...))
}

func DecodeUpdateFlightBatchResponse(response *core.BatchResponse) (*objects.Flight, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Flight
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateFlightWithResponse(ctx context.Context, client *core.Client, id string, params UpdateFlightParams) (*objects.Flight, *core.Response, error) {
	var out objects.Flight
	call := UpdateFlightBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateFlight(ctx context.Context, client *core.Client, id string, params UpdateFlightParams) (*objects.Flight, error) {
	out, _, err := UpdateFlightWithResponse(ctx, client, id, params)
	return out, err
}
