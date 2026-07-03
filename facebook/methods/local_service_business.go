package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetLocalServiceBusinessChannelsToIntegrityStatusParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLocalServiceBusinessChannelsToIntegrityStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLocalServiceBusinessChannelsToIntegrityStatus(ctx context.Context, client *core.Client, id string, params GetLocalServiceBusinessChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
	var out core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "channels_to_integrity_status"), params.ToParams(), &out)
	return &out, err
}

type GetLocalServiceBusinessOverrideDetailsParams struct {
	Keys  *[]string                                               `facebook:"keys"`
	Type  *enums.LocalservicebusinessoverrideDetailsTypeEnumParam `facebook:"type"`
	Extra core.Params                                             `facebook:"-"`
}

func (params GetLocalServiceBusinessOverrideDetailsParams) ToParams() core.Params {
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

func GetLocalServiceBusinessOverrideDetails(ctx context.Context, client *core.Client, id string, params GetLocalServiceBusinessOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], error) {
	var out core.Cursor[objects.OverrideDetails]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "override_details"), params.ToParams(), &out)
	return &out, err
}

type GetLocalServiceBusinessParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetLocalServiceBusinessParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetLocalServiceBusiness(ctx context.Context, client *core.Client, id string, params GetLocalServiceBusinessParams) (*objects.LocalServiceBusiness, error) {
	var out objects.LocalServiceBusiness
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
