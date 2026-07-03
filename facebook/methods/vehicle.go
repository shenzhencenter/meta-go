package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetVehicleChannelsToIntegrityStatusParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetVehicleChannelsToIntegrityStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetVehicleChannelsToIntegrityStatusBatchCall(id string, params GetVehicleChannelsToIntegrityStatusParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "channels_to_integrity_status"), params.ToParams(), options...)
}

func NewGetVehicleChannelsToIntegrityStatusBatchRequest(id string, params GetVehicleChannelsToIntegrityStatusParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]] {
	return core.NewBatchRequest[core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]](GetVehicleChannelsToIntegrityStatusBatchCall(id, params, options...))
}

func DecodeGetVehicleChannelsToIntegrityStatusBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
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

func GetVehicleChannelsToIntegrityStatus(ctx context.Context, client *core.Client, id string, params GetVehicleChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
	var out core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]
	call := GetVehicleChannelsToIntegrityStatusBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetVehicleOverrideDetailsParams struct {
	Keys  *[]string                                  `facebook:"keys"`
	Type  *enums.VehicleoverrideDetailsTypeEnumParam `facebook:"type"`
	Extra core.Params                                `facebook:"-"`
}

func (params GetVehicleOverrideDetailsParams) ToParams() core.Params {
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

func GetVehicleOverrideDetailsBatchCall(id string, params GetVehicleOverrideDetailsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "override_details"), params.ToParams(), options...)
}

func NewGetVehicleOverrideDetailsBatchRequest(id string, params GetVehicleOverrideDetailsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.OverrideDetails]] {
	return core.NewBatchRequest[core.Cursor[objects.OverrideDetails]](GetVehicleOverrideDetailsBatchCall(id, params, options...))
}

func DecodeGetVehicleOverrideDetailsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.OverrideDetails], error) {
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

func GetVehicleOverrideDetails(ctx context.Context, client *core.Client, id string, params GetVehicleOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], error) {
	var out core.Cursor[objects.OverrideDetails]
	call := GetVehicleOverrideDetailsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetVehicleVideosMetadataParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetVehicleVideosMetadataParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetVehicleVideosMetadataBatchCall(id string, params GetVehicleVideosMetadataParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "videos_metadata"), params.ToParams(), options...)
}

func NewGetVehicleVideosMetadataBatchRequest(id string, params GetVehicleVideosMetadataParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.DynamicVideoMetadata]] {
	return core.NewBatchRequest[core.Cursor[objects.DynamicVideoMetadata]](GetVehicleVideosMetadataBatchCall(id, params, options...))
}

func DecodeGetVehicleVideosMetadataBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.DynamicVideoMetadata], error) {
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

func GetVehicleVideosMetadata(ctx context.Context, client *core.Client, id string, params GetVehicleVideosMetadataParams) (*core.Cursor[objects.DynamicVideoMetadata], error) {
	var out core.Cursor[objects.DynamicVideoMetadata]
	call := GetVehicleVideosMetadataBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetVehicleParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetVehicleParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetVehicleBatchCall(id string, params GetVehicleParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetVehicleBatchRequest(id string, params GetVehicleParams, options ...core.BatchOption) *core.BatchRequest[objects.Vehicle] {
	return core.NewBatchRequest[objects.Vehicle](GetVehicleBatchCall(id, params, options...))
}

func DecodeGetVehicleBatchResponse(response *core.BatchResponse) (*objects.Vehicle, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Vehicle
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetVehicle(ctx context.Context, client *core.Client, id string, params GetVehicleParams) (*objects.Vehicle, error) {
	var out objects.Vehicle
	call := GetVehicleBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type UpdateVehicleParams struct {
	Address        *map[string]interface{}      `facebook:"address"`
	Applinks       *map[string]interface{}      `facebook:"applinks"`
	Availability   *enums.VehicleAvailability   `facebook:"availability"`
	BodyStyle      *enums.VehicleBodyStyle      `facebook:"body_style"`
	Condition      *enums.VehicleCondition      `facebook:"condition"`
	Currency       *string                      `facebook:"currency"`
	DateFirstOnLot *string                      `facebook:"date_first_on_lot"`
	DealerID       *core.ID                     `facebook:"dealer_id"`
	DealerName     *string                      `facebook:"dealer_name"`
	DealerPhone    *string                      `facebook:"dealer_phone"`
	Description    *string                      `facebook:"description"`
	Drivetrain     *enums.VehicleDrivetrain     `facebook:"drivetrain"`
	ExteriorColor  *string                      `facebook:"exterior_color"`
	FbPageID       *core.ID                     `facebook:"fb_page_id"`
	FuelType       *enums.VehicleFuelType       `facebook:"fuel_type"`
	Images         *[]map[string]interface{}    `facebook:"images"`
	InteriorColor  *string                      `facebook:"interior_color"`
	Make           *string                      `facebook:"make"`
	Mileage        *map[string]interface{}      `facebook:"mileage"`
	Model          *string                      `facebook:"model"`
	Price          *uint64                      `facebook:"price"`
	StateOfVehicle *enums.VehicleStateOfVehicle `facebook:"state_of_vehicle"`
	Title          *string                      `facebook:"title"`
	Transmission   *enums.VehicleTransmission   `facebook:"transmission"`
	Trim           *string                      `facebook:"trim"`
	URL            *string                      `facebook:"url"`
	VehicleType    *enums.VehicleVehicleType    `facebook:"vehicle_type"`
	Vin            *string                      `facebook:"vin"`
	Year           *uint64                      `facebook:"year"`
	Extra          core.Params                  `facebook:"-"`
}

func (params UpdateVehicleParams) ToParams() core.Params {
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
	if params.Availability != nil {
		out["availability"] = *params.Availability
	}
	if params.BodyStyle != nil {
		out["body_style"] = *params.BodyStyle
	}
	if params.Condition != nil {
		out["condition"] = *params.Condition
	}
	if params.Currency != nil {
		out["currency"] = *params.Currency
	}
	if params.DateFirstOnLot != nil {
		out["date_first_on_lot"] = *params.DateFirstOnLot
	}
	if params.DealerID != nil {
		out["dealer_id"] = *params.DealerID
	}
	if params.DealerName != nil {
		out["dealer_name"] = *params.DealerName
	}
	if params.DealerPhone != nil {
		out["dealer_phone"] = *params.DealerPhone
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.Drivetrain != nil {
		out["drivetrain"] = *params.Drivetrain
	}
	if params.ExteriorColor != nil {
		out["exterior_color"] = *params.ExteriorColor
	}
	if params.FbPageID != nil {
		out["fb_page_id"] = *params.FbPageID
	}
	if params.FuelType != nil {
		out["fuel_type"] = *params.FuelType
	}
	if params.Images != nil {
		out["images"] = *params.Images
	}
	if params.InteriorColor != nil {
		out["interior_color"] = *params.InteriorColor
	}
	if params.Make != nil {
		out["make"] = *params.Make
	}
	if params.Mileage != nil {
		out["mileage"] = *params.Mileage
	}
	if params.Model != nil {
		out["model"] = *params.Model
	}
	if params.Price != nil {
		out["price"] = *params.Price
	}
	if params.StateOfVehicle != nil {
		out["state_of_vehicle"] = *params.StateOfVehicle
	}
	if params.Title != nil {
		out["title"] = *params.Title
	}
	if params.Transmission != nil {
		out["transmission"] = *params.Transmission
	}
	if params.Trim != nil {
		out["trim"] = *params.Trim
	}
	if params.URL != nil {
		out["url"] = *params.URL
	}
	if params.VehicleType != nil {
		out["vehicle_type"] = *params.VehicleType
	}
	if params.Vin != nil {
		out["vin"] = *params.Vin
	}
	if params.Year != nil {
		out["year"] = *params.Year
	}
	return out
}

func UpdateVehicleBatchCall(id string, params UpdateVehicleParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateVehicleBatchRequest(id string, params UpdateVehicleParams, options ...core.BatchOption) *core.BatchRequest[objects.Vehicle] {
	return core.NewBatchRequest[objects.Vehicle](UpdateVehicleBatchCall(id, params, options...))
}

func DecodeUpdateVehicleBatchResponse(response *core.BatchResponse) (*objects.Vehicle, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Vehicle
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateVehicle(ctx context.Context, client *core.Client, id string, params UpdateVehicleParams) (*objects.Vehicle, error) {
	var out objects.Vehicle
	call := UpdateVehicleBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
