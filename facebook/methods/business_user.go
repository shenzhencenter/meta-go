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

func GetBusinessUserAssignedAdAccounts(ctx context.Context, client *core.Client, id string, params GetBusinessUserAssignedAdAccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "assigned_ad_accounts"), params.ToParams(), &out)
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

func GetBusinessUserAssignedBusinessAssetGroups(ctx context.Context, client *core.Client, id string, params GetBusinessUserAssignedBusinessAssetGroupsParams) (*core.Cursor[objects.BusinessAssetGroup], error) {
	var out core.Cursor[objects.BusinessAssetGroup]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "assigned_business_asset_groups"), params.ToParams(), &out)
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

func GetBusinessUserAssignedPages(ctx context.Context, client *core.Client, id string, params GetBusinessUserAssignedPagesParams) (*core.Cursor[objects.Page], error) {
	var out core.Cursor[objects.Page]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "assigned_pages"), params.ToParams(), &out)
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

func GetBusinessUserAssignedProductCatalogs(ctx context.Context, client *core.Client, id string, params GetBusinessUserAssignedProductCatalogsParams) (*core.Cursor[objects.ProductCatalog], error) {
	var out core.Cursor[objects.ProductCatalog]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "assigned_product_catalogs"), params.ToParams(), &out)
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

func GetBusinessUserAssignedWhatsappBusinessAccounts(ctx context.Context, client *core.Client, id string, params GetBusinessUserAssignedWhatsappBusinessAccountsParams) (*core.Cursor[objects.WhatsAppBusinessAccount], error) {
	var out core.Cursor[objects.WhatsAppBusinessAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "assigned_whatsapp_business_accounts"), params.ToParams(), &out)
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

func DeleteBusinessUser(ctx context.Context, client *core.Client, id string, params DeleteBusinessUserParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
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

func GetBusinessUser(ctx context.Context, client *core.Client, id string, params GetBusinessUserParams) (*objects.BusinessUser, error) {
	var out objects.BusinessUser
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
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

func UpdateBusinessUser(ctx context.Context, client *core.Client, id string, params UpdateBusinessUserParams) (*objects.BusinessUser, error) {
	var out objects.BusinessUser
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
