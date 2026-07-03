package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func GetHotelChannelsToIntegrityStatus(ctx context.Context, client *core.Client, id string, params GetHotelChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
	var out core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "channels_to_integrity_status"), params.ToParams(), &out)
	return &out, err
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

func GetHotelHotelRooms(ctx context.Context, client *core.Client, id string, params GetHotelHotelRoomsParams) (*core.Cursor[objects.HotelRoom], error) {
	var out core.Cursor[objects.HotelRoom]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "hotel_rooms"), params.ToParams(), &out)
	return &out, err
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

func GetHotelOverrideDetails(ctx context.Context, client *core.Client, id string, params GetHotelOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], error) {
	var out core.Cursor[objects.OverrideDetails]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "override_details"), params.ToParams(), &out)
	return &out, err
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

func GetHotelVideosMetadata(ctx context.Context, client *core.Client, id string, params GetHotelVideosMetadataParams) (*core.Cursor[objects.DynamicVideoMetadata], error) {
	var out core.Cursor[objects.DynamicVideoMetadata]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "videos_metadata"), params.ToParams(), &out)
	return &out, err
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

func DeleteHotel(ctx context.Context, client *core.Client, id string, params DeleteHotelParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
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

func GetHotel(ctx context.Context, client *core.Client, id string, params GetHotelParams) (*objects.Hotel, error) {
	var out objects.Hotel
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
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

func UpdateHotel(ctx context.Context, client *core.Client, id string, params UpdateHotelParams) (*objects.Hotel, error) {
	var out objects.Hotel
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
