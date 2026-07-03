package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetHotelChannelsToIntegrityStatusParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetHotelChannelsToIntegrityStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetHotelChannelsToIntegrityStatusBatchCall(id string, params GetHotelChannelsToIntegrityStatusParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "channels_to_integrity_status"), params.ToParams(), options...)
}

func NewGetHotelChannelsToIntegrityStatusBatchRequest(id string, params GetHotelChannelsToIntegrityStatusParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]] {
	return core.NewBatchRequest[core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]](GetHotelChannelsToIntegrityStatusBatchCall(id, params, options...))
}

func DecodeGetHotelChannelsToIntegrityStatusBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
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

func GetHotelChannelsToIntegrityStatusWithResponse(ctx context.Context, client *core.Client, id string, params GetHotelChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], *core.Response, error) {
	var out core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]
	call := GetHotelChannelsToIntegrityStatusBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetHotelChannelsToIntegrityStatus(ctx context.Context, client *core.Client, id string, params GetHotelChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
	out, _, err := GetHotelChannelsToIntegrityStatusWithResponse(ctx, client, id, params)
	return out, err
}

type GetHotelHotelRoomsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetHotelHotelRoomsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetHotelHotelRoomsBatchCall(id string, params GetHotelHotelRoomsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "hotel_rooms"), params.ToParams(), options...)
}

func NewGetHotelHotelRoomsBatchRequest(id string, params GetHotelHotelRoomsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.HotelRoom]] {
	return core.NewBatchRequest[core.Cursor[objects.HotelRoom]](GetHotelHotelRoomsBatchCall(id, params, options...))
}

func DecodeGetHotelHotelRoomsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.HotelRoom], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.HotelRoom]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetHotelHotelRoomsWithResponse(ctx context.Context, client *core.Client, id string, params GetHotelHotelRoomsParams) (*core.Cursor[objects.HotelRoom], *core.Response, error) {
	var out core.Cursor[objects.HotelRoom]
	call := GetHotelHotelRoomsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetHotelHotelRooms(ctx context.Context, client *core.Client, id string, params GetHotelHotelRoomsParams) (*core.Cursor[objects.HotelRoom], error) {
	out, _, err := GetHotelHotelRoomsWithResponse(ctx, client, id, params)
	return out, err
}

type GetHotelOverrideDetailsParams struct {
	Keys  *[]string                                `facebook:"keys"`
	Type  *enums.HoteloverrideDetailsTypeEnumParam `facebook:"type"`
	Extra core.Params                              `facebook:"-"`
}

func (params GetHotelOverrideDetailsParams) ToParams() core.Params {
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

func GetHotelOverrideDetailsBatchCall(id string, params GetHotelOverrideDetailsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "override_details"), params.ToParams(), options...)
}

func NewGetHotelOverrideDetailsBatchRequest(id string, params GetHotelOverrideDetailsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.OverrideDetails]] {
	return core.NewBatchRequest[core.Cursor[objects.OverrideDetails]](GetHotelOverrideDetailsBatchCall(id, params, options...))
}

func DecodeGetHotelOverrideDetailsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.OverrideDetails], error) {
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

func GetHotelOverrideDetailsWithResponse(ctx context.Context, client *core.Client, id string, params GetHotelOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], *core.Response, error) {
	var out core.Cursor[objects.OverrideDetails]
	call := GetHotelOverrideDetailsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetHotelOverrideDetails(ctx context.Context, client *core.Client, id string, params GetHotelOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], error) {
	out, _, err := GetHotelOverrideDetailsWithResponse(ctx, client, id, params)
	return out, err
}

type GetHotelVideosMetadataParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetHotelVideosMetadataParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetHotelVideosMetadataBatchCall(id string, params GetHotelVideosMetadataParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "videos_metadata"), params.ToParams(), options...)
}

func NewGetHotelVideosMetadataBatchRequest(id string, params GetHotelVideosMetadataParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.DynamicVideoMetadata]] {
	return core.NewBatchRequest[core.Cursor[objects.DynamicVideoMetadata]](GetHotelVideosMetadataBatchCall(id, params, options...))
}

func DecodeGetHotelVideosMetadataBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.DynamicVideoMetadata], error) {
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

func GetHotelVideosMetadataWithResponse(ctx context.Context, client *core.Client, id string, params GetHotelVideosMetadataParams) (*core.Cursor[objects.DynamicVideoMetadata], *core.Response, error) {
	var out core.Cursor[objects.DynamicVideoMetadata]
	call := GetHotelVideosMetadataBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetHotelVideosMetadata(ctx context.Context, client *core.Client, id string, params GetHotelVideosMetadataParams) (*core.Cursor[objects.DynamicVideoMetadata], error) {
	out, _, err := GetHotelVideosMetadataWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteHotelParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteHotelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteHotelBatchCall(id string, params DeleteHotelParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteHotelBatchRequest(id string, params DeleteHotelParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteHotelBatchCall(id, params, options...))
}

func DecodeDeleteHotelBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func DeleteHotelWithResponse(ctx context.Context, client *core.Client, id string, params DeleteHotelParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteHotelBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteHotel(ctx context.Context, client *core.Client, id string, params DeleteHotelParams) (*map[string]interface{}, error) {
	out, _, err := DeleteHotelWithResponse(ctx, client, id, params)
	return out, err
}

type GetHotelParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetHotelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetHotelBatchCall(id string, params GetHotelParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetHotelBatchRequest(id string, params GetHotelParams, options ...core.BatchOption) *core.BatchRequest[objects.Hotel] {
	return core.NewBatchRequest[objects.Hotel](GetHotelBatchCall(id, params, options...))
}

func DecodeGetHotelBatchResponse(response *core.BatchResponse) (*objects.Hotel, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Hotel
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetHotelWithResponse(ctx context.Context, client *core.Client, id string, params GetHotelParams) (*objects.Hotel, *core.Response, error) {
	var out objects.Hotel
	call := GetHotelBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetHotel(ctx context.Context, client *core.Client, id string, params GetHotelParams) (*objects.Hotel, error) {
	out, _, err := GetHotelWithResponse(ctx, client, id, params)
	return out, err
}

type UpdateHotelParams struct {
	Address      *map[string]interface{}   `facebook:"address"`
	Applinks     *map[string]interface{}   `facebook:"applinks"`
	BasePrice    *uint64                   `facebook:"base_price"`
	Brand        *string                   `facebook:"brand"`
	Currency     *string                   `facebook:"currency"`
	Description  *string                   `facebook:"description"`
	GuestRatings *[]map[string]interface{} `facebook:"guest_ratings"`
	Images       *[]map[string]interface{} `facebook:"images"`
	Name         *string                   `facebook:"name"`
	Phone        *string                   `facebook:"phone"`
	StarRating   *float64                  `facebook:"star_rating"`
	URL          *string                   `facebook:"url"`
	Extra        core.Params               `facebook:"-"`
}

func (params UpdateHotelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Address != nil {
		out["address"] = *params.Address
	}
	if params.Applinks != nil {
		out["applinks"] = *params.Applinks
	}
	if params.BasePrice != nil {
		out["base_price"] = *params.BasePrice
	}
	if params.Brand != nil {
		out["brand"] = *params.Brand
	}
	if params.Currency != nil {
		out["currency"] = *params.Currency
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.GuestRatings != nil {
		out["guest_ratings"] = *params.GuestRatings
	}
	if params.Images != nil {
		out["images"] = *params.Images
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.Phone != nil {
		out["phone"] = *params.Phone
	}
	if params.StarRating != nil {
		out["star_rating"] = *params.StarRating
	}
	if params.URL != nil {
		out["url"] = *params.URL
	}
	return out
}

func UpdateHotelBatchCall(id string, params UpdateHotelParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateHotelBatchRequest(id string, params UpdateHotelParams, options ...core.BatchOption) *core.BatchRequest[objects.Hotel] {
	return core.NewBatchRequest[objects.Hotel](UpdateHotelBatchCall(id, params, options...))
}

func DecodeUpdateHotelBatchResponse(response *core.BatchResponse) (*objects.Hotel, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Hotel
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateHotelWithResponse(ctx context.Context, client *core.Client, id string, params UpdateHotelParams) (*objects.Hotel, *core.Response, error) {
	var out objects.Hotel
	call := UpdateHotelBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateHotel(ctx context.Context, client *core.Client, id string, params UpdateHotelParams) (*objects.Hotel, error) {
	out, _, err := UpdateHotelWithResponse(ctx, client, id, params)
	return out, err
}
