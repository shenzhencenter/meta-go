package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetBusinessUserAssignedAdAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessUserAssignedAdAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessUserAssignedAdAccountsBatchCall(id string, params GetBusinessUserAssignedAdAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "assigned_ad_accounts"), params.ToParams(), options...)
}

func NewGetBusinessUserAssignedAdAccountsBatchRequest(id string, params GetBusinessUserAssignedAdAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.AdAccount]](GetBusinessUserAssignedAdAccountsBatchCall(id, params, options...))
}

func DecodeGetBusinessUserAssignedAdAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdAccount], error) {
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

func GetBusinessUserAssignedAdAccounts(ctx context.Context, client *core.Client, id string, params GetBusinessUserAssignedAdAccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	call := GetBusinessUserAssignedAdAccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetBusinessUserAssignedBusinessAssetGroupsParams struct {
	ContainedAssetID *core.ID    `facebook:"contained_asset_id"`
	Extra            core.Params `facebook:"-"`
}

func (params GetBusinessUserAssignedBusinessAssetGroupsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ContainedAssetID != nil {
		out["contained_asset_id"] = *params.ContainedAssetID
	}
	return out
}

func GetBusinessUserAssignedBusinessAssetGroupsBatchCall(id string, params GetBusinessUserAssignedBusinessAssetGroupsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "assigned_business_asset_groups"), params.ToParams(), options...)
}

func NewGetBusinessUserAssignedBusinessAssetGroupsBatchRequest(id string, params GetBusinessUserAssignedBusinessAssetGroupsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.BusinessAssetGroup]] {
	return core.NewBatchRequest[core.Cursor[objects.BusinessAssetGroup]](GetBusinessUserAssignedBusinessAssetGroupsBatchCall(id, params, options...))
}

func DecodeGetBusinessUserAssignedBusinessAssetGroupsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.BusinessAssetGroup], error) {
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

func GetBusinessUserAssignedBusinessAssetGroups(ctx context.Context, client *core.Client, id string, params GetBusinessUserAssignedBusinessAssetGroupsParams) (*core.Cursor[objects.BusinessAssetGroup], error) {
	var out core.Cursor[objects.BusinessAssetGroup]
	call := GetBusinessUserAssignedBusinessAssetGroupsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetBusinessUserAssignedPagesParams struct {
	Pages *[]uint64   `facebook:"pages"`
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessUserAssignedPagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Pages != nil {
		out["pages"] = *params.Pages
	}
	return out
}

func GetBusinessUserAssignedPagesBatchCall(id string, params GetBusinessUserAssignedPagesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "assigned_pages"), params.ToParams(), options...)
}

func NewGetBusinessUserAssignedPagesBatchRequest(id string, params GetBusinessUserAssignedPagesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Page]] {
	return core.NewBatchRequest[core.Cursor[objects.Page]](GetBusinessUserAssignedPagesBatchCall(id, params, options...))
}

func DecodeGetBusinessUserAssignedPagesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Page], error) {
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

func GetBusinessUserAssignedPages(ctx context.Context, client *core.Client, id string, params GetBusinessUserAssignedPagesParams) (*core.Cursor[objects.Page], error) {
	var out core.Cursor[objects.Page]
	call := GetBusinessUserAssignedPagesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetBusinessUserAssignedProductCatalogsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessUserAssignedProductCatalogsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessUserAssignedProductCatalogsBatchCall(id string, params GetBusinessUserAssignedProductCatalogsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "assigned_product_catalogs"), params.ToParams(), options...)
}

func NewGetBusinessUserAssignedProductCatalogsBatchRequest(id string, params GetBusinessUserAssignedProductCatalogsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProductCatalog]] {
	return core.NewBatchRequest[core.Cursor[objects.ProductCatalog]](GetBusinessUserAssignedProductCatalogsBatchCall(id, params, options...))
}

func DecodeGetBusinessUserAssignedProductCatalogsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProductCatalog], error) {
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

func GetBusinessUserAssignedProductCatalogs(ctx context.Context, client *core.Client, id string, params GetBusinessUserAssignedProductCatalogsParams) (*core.Cursor[objects.ProductCatalog], error) {
	var out core.Cursor[objects.ProductCatalog]
	call := GetBusinessUserAssignedProductCatalogsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetBusinessUserAssignedWhatsappBusinessAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessUserAssignedWhatsappBusinessAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessUserAssignedWhatsappBusinessAccountsBatchCall(id string, params GetBusinessUserAssignedWhatsappBusinessAccountsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "assigned_whatsapp_business_accounts"), params.ToParams(), options...)
}

func NewGetBusinessUserAssignedWhatsappBusinessAccountsBatchRequest(id string, params GetBusinessUserAssignedWhatsappBusinessAccountsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.WhatsAppBusinessAccount]] {
	return core.NewBatchRequest[core.Cursor[objects.WhatsAppBusinessAccount]](GetBusinessUserAssignedWhatsappBusinessAccountsBatchCall(id, params, options...))
}

func DecodeGetBusinessUserAssignedWhatsappBusinessAccountsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.WhatsAppBusinessAccount], error) {
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

func GetBusinessUserAssignedWhatsappBusinessAccounts(ctx context.Context, client *core.Client, id string, params GetBusinessUserAssignedWhatsappBusinessAccountsParams) (*core.Cursor[objects.WhatsAppBusinessAccount], error) {
	var out core.Cursor[objects.WhatsAppBusinessAccount]
	call := GetBusinessUserAssignedWhatsappBusinessAccountsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeleteBusinessUserParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteBusinessUserParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteBusinessUserBatchCall(id string, params DeleteBusinessUserParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteBusinessUserBatchRequest(id string, params DeleteBusinessUserParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteBusinessUserBatchCall(id, params, options...))
}

func DecodeDeleteBusinessUserBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func DeleteBusinessUser(ctx context.Context, client *core.Client, id string, params DeleteBusinessUserParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteBusinessUserBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetBusinessUserParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessUserParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessUserBatchCall(id string, params GetBusinessUserParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetBusinessUserBatchRequest(id string, params GetBusinessUserParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessUser] {
	return core.NewBatchRequest[objects.BusinessUser](GetBusinessUserBatchCall(id, params, options...))
}

func DecodeGetBusinessUserBatchResponse(response *core.BatchResponse) (*objects.BusinessUser, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessUser
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetBusinessUser(ctx context.Context, client *core.Client, id string, params GetBusinessUserParams) (*objects.BusinessUser, error) {
	var out objects.BusinessUser
	call := GetBusinessUserBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type UpdateBusinessUserParams struct {
	ClearPendingEmail     *bool                      `facebook:"clear_pending_email"`
	Email                 *string                    `facebook:"email"`
	FirstName             *string                    `facebook:"first_name"`
	LastName              *string                    `facebook:"last_name"`
	PendingEmail          *string                    `facebook:"pending_email"`
	Role                  *enums.BusinessuserRole    `facebook:"role"`
	SkipVerificationEmail *bool                      `facebook:"skip_verification_email"`
	Tasks                 *[]enums.BusinessuserTasks `facebook:"tasks"`
	Title                 *string                    `facebook:"title"`
	Extra                 core.Params                `facebook:"-"`
}

func (params UpdateBusinessUserParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ClearPendingEmail != nil {
		out["clear_pending_email"] = *params.ClearPendingEmail
	}
	if params.Email != nil {
		out["email"] = *params.Email
	}
	if params.FirstName != nil {
		out["first_name"] = *params.FirstName
	}
	if params.LastName != nil {
		out["last_name"] = *params.LastName
	}
	if params.PendingEmail != nil {
		out["pending_email"] = *params.PendingEmail
	}
	if params.Role != nil {
		out["role"] = *params.Role
	}
	if params.SkipVerificationEmail != nil {
		out["skip_verification_email"] = *params.SkipVerificationEmail
	}
	if params.Tasks != nil {
		out["tasks"] = *params.Tasks
	}
	if params.Title != nil {
		out["title"] = *params.Title
	}
	return out
}

func UpdateBusinessUserBatchCall(id string, params UpdateBusinessUserParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateBusinessUserBatchRequest(id string, params UpdateBusinessUserParams, options ...core.BatchOption) *core.BatchRequest[objects.BusinessUser] {
	return core.NewBatchRequest[objects.BusinessUser](UpdateBusinessUserBatchCall(id, params, options...))
}

func DecodeUpdateBusinessUserBatchResponse(response *core.BatchResponse) (*objects.BusinessUser, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.BusinessUser
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateBusinessUser(ctx context.Context, client *core.Client, id string, params UpdateBusinessUserParams) (*objects.BusinessUser, error) {
	var out objects.BusinessUser
	call := UpdateBusinessUserBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
