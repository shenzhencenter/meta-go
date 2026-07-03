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

func GetSystemUserAssignedAdAccountsBatchCall(id string, params GetSystemUserAssignedAdAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "assigned_ad_accounts"), params.ToParams(), options...)
}

func NewGetSystemUserAssignedAdAccountsBatchRequest(id string, params GetSystemUserAssignedAdAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccount]](GetSystemUserAssignedAdAccountsBatchCall(id, params, options...))
}

func DecodeGetSystemUserAssignedAdAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccount], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdAccount]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetSystemUserAssignedAdAccounts(ctx context.Context, client *core.Client, id string, params GetSystemUserAssignedAdAccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	call := GetSystemUserAssignedAdAccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetSystemUserAssignedBusinessAssetGroupsBatchCall(id string, params GetSystemUserAssignedBusinessAssetGroupsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "assigned_business_asset_groups"), params.ToParams(), options...)
}

func NewGetSystemUserAssignedBusinessAssetGroupsBatchRequest(id string, params GetSystemUserAssignedBusinessAssetGroupsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BusinessAssetGroup]] {
	return core.NewBatchRequest[core.Cursor[objects.BusinessAssetGroup]](GetSystemUserAssignedBusinessAssetGroupsBatchCall(id, params, options...))
}

func DecodeGetSystemUserAssignedBusinessAssetGroupsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BusinessAssetGroup], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.BusinessAssetGroup]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetSystemUserAssignedBusinessAssetGroups(ctx context.Context, client *core.Client, id string, params GetSystemUserAssignedBusinessAssetGroupsParams) (*core.Cursor[objects.BusinessAssetGroup], error) {
	var out core.Cursor[objects.BusinessAssetGroup]
	call := GetSystemUserAssignedBusinessAssetGroupsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetSystemUserAssignedPagesBatchCall(id string, params GetSystemUserAssignedPagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "assigned_pages"), params.ToParams(), options...)
}

func NewGetSystemUserAssignedPagesBatchRequest(id string, params GetSystemUserAssignedPagesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Page]] {
	return core.NewBatchRequest[core.Cursor[objects.Page]](GetSystemUserAssignedPagesBatchCall(id, params, options...))
}

func DecodeGetSystemUserAssignedPagesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Page], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Page]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetSystemUserAssignedPages(ctx context.Context, client *core.Client, id string, params GetSystemUserAssignedPagesParams) (*core.Cursor[objects.Page], error) {
	var out core.Cursor[objects.Page]
	call := GetSystemUserAssignedPagesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetSystemUserAssignedProductCatalogsBatchCall(id string, params GetSystemUserAssignedProductCatalogsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "assigned_product_catalogs"), params.ToParams(), options...)
}

func NewGetSystemUserAssignedProductCatalogsBatchRequest(id string, params GetSystemUserAssignedProductCatalogsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProductCatalog]] {
	return core.NewBatchRequest[core.Cursor[objects.ProductCatalog]](GetSystemUserAssignedProductCatalogsBatchCall(id, params, options...))
}

func DecodeGetSystemUserAssignedProductCatalogsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProductCatalog], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ProductCatalog]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetSystemUserAssignedProductCatalogs(ctx context.Context, client *core.Client, id string, params GetSystemUserAssignedProductCatalogsParams) (*core.Cursor[objects.ProductCatalog], error) {
	var out core.Cursor[objects.ProductCatalog]
	call := GetSystemUserAssignedProductCatalogsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetSystemUserAssignedWhatsappBusinessAccountsBatchCall(id string, params GetSystemUserAssignedWhatsappBusinessAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "assigned_whatsapp_business_accounts"), params.ToParams(), options...)
}

func NewGetSystemUserAssignedWhatsappBusinessAccountsBatchRequest(id string, params GetSystemUserAssignedWhatsappBusinessAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WhatsAppBusinessAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.WhatsAppBusinessAccount]](GetSystemUserAssignedWhatsappBusinessAccountsBatchCall(id, params, options...))
}

func DecodeGetSystemUserAssignedWhatsappBusinessAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WhatsAppBusinessAccount], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.WhatsAppBusinessAccount]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetSystemUserAssignedWhatsappBusinessAccounts(ctx context.Context, client *core.Client, id string, params GetSystemUserAssignedWhatsappBusinessAccountsParams) (*core.Cursor[objects.WhatsAppBusinessAccount], error) {
	var out core.Cursor[objects.WhatsAppBusinessAccount]
	call := GetSystemUserAssignedWhatsappBusinessAccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetSystemUserBatchCall(id string, params GetSystemUserParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetSystemUserBatchRequest(id string, params GetSystemUserParams, options ...core.BatchOption) *core.BatchRequest[objects.SystemUser] {
	return core.NewBatchRequest[objects.SystemUser](GetSystemUserBatchCall(id, params, options...))
}

func DecodeGetSystemUserBatchResponse(response *core.BatchResponse) (*objects.SystemUser, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.SystemUser
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetSystemUser(ctx context.Context, client *core.Client, id string, params GetSystemUserParams) (*objects.SystemUser, error) {
	var out objects.SystemUser
	call := GetSystemUserBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
