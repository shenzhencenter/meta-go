package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func GetVehicleChannelsToIntegrityStatus(ctx context.Context, client *core.Client, id string, params GetVehicleChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
	var out core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "channels_to_integrity_status"), params.ToParams(), &out)
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

func GetVehicleOverrideDetails(ctx context.Context, client *core.Client, id string, params GetVehicleOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], error) {
	var out core.Cursor[objects.OverrideDetails]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "override_details"), params.ToParams(), &out)
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

func GetVehicleVideosMetadata(ctx context.Context, client *core.Client, id string, params GetVehicleVideosMetadataParams) (*core.Cursor[objects.DynamicVideoMetadata], error) {
	var out core.Cursor[objects.DynamicVideoMetadata]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "videos_metadata"), params.ToParams(), &out)
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

func GetVehicle(ctx context.Context, client *core.Client, id string, params GetVehicleParams) (*objects.Vehicle, error) {
	var out objects.Vehicle
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
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

func UpdateVehicle(ctx context.Context, client *core.Client, id string, params UpdateVehicleParams) (*objects.Vehicle, error) {
	var out objects.Vehicle
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
