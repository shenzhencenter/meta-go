package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAutomotiveModelChannelsToIntegrityStatusParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAutomotiveModelChannelsToIntegrityStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAutomotiveModelChannelsToIntegrityStatus(ctx context.Context, client *core.Client, id string, params GetAutomotiveModelChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
	var out core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "channels_to_integrity_status"), params.ToParams(), &out)
	return &out, err
}

type GetAutomotiveModelOverrideDetailsParams struct {
	Keys  *[]string                                          `facebook:"keys"`
	Type  *enums.AutomotivemodeloverrideDetailsTypeEnumParam `facebook:"type"`
	Extra core.Params                                        `facebook:"-"`
}

func (params GetAutomotiveModelOverrideDetailsParams) ToParams() core.Params {
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

func GetAutomotiveModelOverrideDetails(ctx context.Context, client *core.Client, id string, params GetAutomotiveModelOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], error) {
	var out core.Cursor[objects.OverrideDetails]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "override_details"), params.ToParams(), &out)
	return &out, err
}

type GetAutomotiveModelVideosMetadataParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAutomotiveModelVideosMetadataParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAutomotiveModelVideosMetadata(ctx context.Context, client *core.Client, id string, params GetAutomotiveModelVideosMetadataParams) (*core.Cursor[objects.DynamicVideoMetadata], error) {
	var out core.Cursor[objects.DynamicVideoMetadata]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "videos_metadata"), params.ToParams(), &out)
	return &out, err
}

type GetAutomotiveModelParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAutomotiveModelParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAutomotiveModel(ctx context.Context, client *core.Client, id string, params GetAutomotiveModelParams) (*objects.AutomotiveModel, error) {
	var out objects.AutomotiveModel
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
