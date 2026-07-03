package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetTransactableItemChannelsToIntegrityStatusParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetTransactableItemChannelsToIntegrityStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetTransactableItemChannelsToIntegrityStatus(ctx context.Context, client *core.Client, id string, params GetTransactableItemChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
	var out core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "channels_to_integrity_status"), params.ToParams(), &out)
	return &out, err
}

type GetTransactableItemOverrideDetailsParams struct {
	Keys  *[]string                                           `facebook:"keys"`
	Type  *enums.TransactableitemoverrideDetailsTypeEnumParam `facebook:"type"`
	Extra core.Params                                         `facebook:"-"`
}

func (params GetTransactableItemOverrideDetailsParams) ToParams() core.Params {
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

func GetTransactableItemOverrideDetails(ctx context.Context, client *core.Client, id string, params GetTransactableItemOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], error) {
	var out core.Cursor[objects.OverrideDetails]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "override_details"), params.ToParams(), &out)
	return &out, err
}

type GetTransactableItemParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetTransactableItemParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetTransactableItem(ctx context.Context, client *core.Client, id string, params GetTransactableItemParams) (*objects.TransactableItem, error) {
	var out objects.TransactableItem
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
