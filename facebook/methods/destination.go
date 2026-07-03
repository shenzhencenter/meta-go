package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetDestinationChannelsToIntegrityStatusParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetDestinationChannelsToIntegrityStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetDestinationChannelsToIntegrityStatus(ctx context.Context, client *core.Client, id string, params GetDestinationChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
	var out core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "channels_to_integrity_status"), params.ToParams(), &out)
	return &out, err
}

type GetDestinationOverrideDetailsParams struct {
	Keys  *[]string                                      `facebook:"keys"`
	Type  *enums.DestinationoverrideDetailsTypeEnumParam `facebook:"type"`
	Extra core.Params                                    `facebook:"-"`
}

func (params GetDestinationOverrideDetailsParams) ToParams() core.Params {
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

func GetDestinationOverrideDetails(ctx context.Context, client *core.Client, id string, params GetDestinationOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], error) {
	var out core.Cursor[objects.OverrideDetails]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "override_details"), params.ToParams(), &out)
	return &out, err
}

type GetDestinationVideosMetadataParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetDestinationVideosMetadataParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetDestinationVideosMetadata(ctx context.Context, client *core.Client, id string, params GetDestinationVideosMetadataParams) (*core.Cursor[objects.DynamicVideoMetadata], error) {
	var out core.Cursor[objects.DynamicVideoMetadata]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "videos_metadata"), params.ToParams(), &out)
	return &out, err
}

type GetDestinationParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetDestinationParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetDestination(ctx context.Context, client *core.Client, id string, params GetDestinationParams) (*objects.Destination, error) {
	var out objects.Destination
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
