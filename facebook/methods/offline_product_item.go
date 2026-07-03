package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetOfflineProductItemChannelsToIntegrityStatusParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOfflineProductItemChannelsToIntegrityStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOfflineProductItemChannelsToIntegrityStatus(ctx context.Context, client *core.Client, id string, params GetOfflineProductItemChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
	var out core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "channels_to_integrity_status"), params.ToParams(), &out)
	return &out, err
}

type GetOfflineProductItemOverrideDetailsParams struct {
	Keys  *[]string                                             `facebook:"keys"`
	Type  *enums.OfflineproductitemoverrideDetailsTypeEnumParam `facebook:"type"`
	Extra core.Params                                           `facebook:"-"`
}

func (params GetOfflineProductItemOverrideDetailsParams) ToParams() core.Params {
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

func GetOfflineProductItemOverrideDetails(ctx context.Context, client *core.Client, id string, params GetOfflineProductItemOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], error) {
	var out core.Cursor[objects.OverrideDetails]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "override_details"), params.ToParams(), &out)
	return &out, err
}

type GetOfflineProductItemParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOfflineProductItemParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOfflineProductItem(ctx context.Context, client *core.Client, id string, params GetOfflineProductItemParams) (*objects.OfflineProductItem, error) {
	var out objects.OfflineProductItem
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
