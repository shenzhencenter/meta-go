package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetHomeListingChannelsToIntegrityStatusParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetHomeListingChannelsToIntegrityStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetHomeListingChannelsToIntegrityStatus(ctx context.Context, client *core.Client, id string, params GetHomeListingChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
	var out core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "channels_to_integrity_status"), params.ToParams(), &out)
	return &out, err
}

type GetHomeListingOverrideDetailsParams struct {
	Keys  *[]string                                      `facebook:"keys"`
	Type  *enums.HomelistingoverrideDetailsTypeEnumParam `facebook:"type"`
	Extra core.Params                                    `facebook:"-"`
}

func (params GetHomeListingOverrideDetailsParams) ToParams() core.Params {
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

func GetHomeListingOverrideDetails(ctx context.Context, client *core.Client, id string, params GetHomeListingOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], error) {
	var out core.Cursor[objects.OverrideDetails]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "override_details"), params.ToParams(), &out)
	return &out, err
}

type GetHomeListingVideosMetadataParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetHomeListingVideosMetadataParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetHomeListingVideosMetadata(ctx context.Context, client *core.Client, id string, params GetHomeListingVideosMetadataParams) (*core.Cursor[objects.DynamicVideoMetadata], error) {
	var out core.Cursor[objects.DynamicVideoMetadata]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "videos_metadata"), params.ToParams(), &out)
	return &out, err
}

type DeleteHomeListingParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteHomeListingParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteHomeListing(ctx context.Context, client *core.Client, id string, params DeleteHomeListingParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetHomeListingParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetHomeListingParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetHomeListing(ctx context.Context, client *core.Client, id string, params GetHomeListingParams) (*objects.HomeListing, error) {
	var out objects.HomeListing
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateHomeListingParams struct {
	Address      *map[string]interface{}   `facebook:"address"`
	Availability *string                   `facebook:"availability"`
	Currency     *string                   `facebook:"currency"`
	Description  *string                   `facebook:"description"`
	Images       *[]map[string]interface{} `facebook:"images"`
	ListingType  *string                   `facebook:"listing_type"`
	Name         *string                   `facebook:"name"`
	NumBaths     *float64                  `facebook:"num_baths"`
	NumBeds      *float64                  `facebook:"num_beds"`
	NumUnits     *float64                  `facebook:"num_units"`
	Price        *float64                  `facebook:"price"`
	PropertyType *string                   `facebook:"property_type"`
	URL          *string                   `facebook:"url"`
	YearBuilt    *uint64                   `facebook:"year_built"`
	Extra        core.Params               `facebook:"-"`
}

func (params UpdateHomeListingParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Address != nil {
		out["address"] = *params.Address
	}
	if params.Availability != nil {
		out["availability"] = *params.Availability
	}
	if params.Currency != nil {
		out["currency"] = *params.Currency
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.Images != nil {
		out["images"] = *params.Images
	}
	if params.ListingType != nil {
		out["listing_type"] = *params.ListingType
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.NumBaths != nil {
		out["num_baths"] = *params.NumBaths
	}
	if params.NumBeds != nil {
		out["num_beds"] = *params.NumBeds
	}
	if params.NumUnits != nil {
		out["num_units"] = *params.NumUnits
	}
	if params.Price != nil {
		out["price"] = *params.Price
	}
	if params.PropertyType != nil {
		out["property_type"] = *params.PropertyType
	}
	if params.URL != nil {
		out["url"] = *params.URL
	}
	if params.YearBuilt != nil {
		out["year_built"] = *params.YearBuilt
	}
	return out
}

func UpdateHomeListing(ctx context.Context, client *core.Client, id string, params UpdateHomeListingParams) (*objects.HomeListing, error) {
	var out objects.HomeListing
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
