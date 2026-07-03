package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetSystemUserAssignedAdAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetSystemUserAssignedAdAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetSystemUserAssignedAdAccounts(ctx context.Context, client *core.Client, id string, params GetSystemUserAssignedAdAccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "assigned_ad_accounts"), params.ToParams(), &out)
	return &out, err
}

type GetSystemUserAssignedBusinessAssetGroupsParams struct {
	ContainedAssetID *core.ID    `facebook:"contained_asset_id"`
	Extra            core.Params `facebook:"-"`
}

func (params GetSystemUserAssignedBusinessAssetGroupsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ContainedAssetID != nil {
		out["contained_asset_id"] = *params.ContainedAssetID
	}
	return out
}

func GetSystemUserAssignedBusinessAssetGroups(ctx context.Context, client *core.Client, id string, params GetSystemUserAssignedBusinessAssetGroupsParams) (*core.Cursor[objects.BusinessAssetGroup], error) {
	var out core.Cursor[objects.BusinessAssetGroup]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "assigned_business_asset_groups"), params.ToParams(), &out)
	return &out, err
}

type GetSystemUserAssignedPagesParams struct {
	Pages *[]uint64   `facebook:"pages"`
	Extra core.Params `facebook:"-"`
}

func (params GetSystemUserAssignedPagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Pages != nil {
		out["pages"] = *params.Pages
	}
	return out
}

func GetSystemUserAssignedPages(ctx context.Context, client *core.Client, id string, params GetSystemUserAssignedPagesParams) (*core.Cursor[objects.Page], error) {
	var out core.Cursor[objects.Page]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "assigned_pages"), params.ToParams(), &out)
	return &out, err
}

type GetSystemUserAssignedProductCatalogsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetSystemUserAssignedProductCatalogsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetSystemUserAssignedProductCatalogs(ctx context.Context, client *core.Client, id string, params GetSystemUserAssignedProductCatalogsParams) (*core.Cursor[objects.ProductCatalog], error) {
	var out core.Cursor[objects.ProductCatalog]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "assigned_product_catalogs"), params.ToParams(), &out)
	return &out, err
}

type GetSystemUserAssignedWhatsappBusinessAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetSystemUserAssignedWhatsappBusinessAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetSystemUserAssignedWhatsappBusinessAccounts(ctx context.Context, client *core.Client, id string, params GetSystemUserAssignedWhatsappBusinessAccountsParams) (*core.Cursor[objects.WhatsAppBusinessAccount], error) {
	var out core.Cursor[objects.WhatsAppBusinessAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "assigned_whatsapp_business_accounts"), params.ToParams(), &out)
	return &out, err
}

type GetSystemUserParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetSystemUserParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetSystemUser(ctx context.Context, client *core.Client, id string, params GetSystemUserParams) (*objects.SystemUser, error) {
	var out objects.SystemUser
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
