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

func GetFlightChannelsToIntegrityStatus(ctx context.Context, client *core.Client, id string, params GetFlightChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
	var out core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "channels_to_integrity_status"), params.ToParams(), &out)
	return &out, err
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

func GetFlightOverrideDetails(ctx context.Context, client *core.Client, id string, params GetFlightOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], error) {
	var out core.Cursor[objects.OverrideDetails]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "override_details"), params.ToParams(), &out)
	return &out, err
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

func GetFlightVideosMetadata(ctx context.Context, client *core.Client, id string, params GetFlightVideosMetadataParams) (*core.Cursor[objects.DynamicVideoMetadata], error) {
	var out core.Cursor[objects.DynamicVideoMetadata]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "videos_metadata"), params.ToParams(), &out)
	return &out, err
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

func GetFlight(ctx context.Context, client *core.Client, id string, params GetFlightParams) (*objects.Flight, error) {
	var out objects.Flight
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
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

func UpdateFlight(ctx context.Context, client *core.Client, id string, params UpdateFlightParams) (*objects.Flight, error) {
	var out objects.Flight
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
