package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func GetVehicleOfferChannelsToIntegrityStatus(ctx context.Context, client *core.Client, id string, params GetVehicleOfferChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
	var out core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "channels_to_integrity_status"), params.ToParams(), &out)
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

func GetVehicleOfferOverrideDetails(ctx context.Context, client *core.Client, id string, params GetVehicleOfferOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], error) {
	var out core.Cursor[objects.OverrideDetails]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "override_details"), params.ToParams(), &out)
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

func GetVehicleOfferVideosMetadata(ctx context.Context, client *core.Client, id string, params GetVehicleOfferVideosMetadataParams) (*core.Cursor[objects.DynamicVideoMetadata], error) {
	var out core.Cursor[objects.DynamicVideoMetadata]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "videos_metadata"), params.ToParams(), &out)
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

func GetVehicleOffer(ctx context.Context, client *core.Client, id string, params GetVehicleOfferParams) (*objects.VehicleOffer, error) {
	var out objects.VehicleOffer
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
