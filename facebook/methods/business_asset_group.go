package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type DeleteBusinessAssetGroupAssignedUsersParams struct {
	User  int         `facebook:"user"`
	Extra core.Params `facebook:"-"`
}

func (params DeleteBusinessAssetGroupAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["user"] = params.User
	return out
}

func DeleteBusinessAssetGroupAssignedUsers(ctx context.Context, client *core.Client, id string, params DeleteBusinessAssetGroupAssignedUsersParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "assigned_users"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessAssetGroupAssignedUsersParams struct {
	Business string      `facebook:"business"`
	Extra    core.Params `facebook:"-"`
}

func (params GetBusinessAssetGroupAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["business"] = params.Business
	return out
}

func GetBusinessAssetGroupAssignedUsers(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupAssignedUsersParams) (*core.Cursor[objects.AssignedUser], error) {
	var out core.Cursor[objects.AssignedUser]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "assigned_users"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessAssetGroupAssignedUsersParams struct {
	AdaccountTasks                *[]enums.BusinessassetgroupassignedUsersAdaccountTasksEnumParam                `facebook:"adaccount_tasks"`
	OfflineConversionDataSetTasks *[]enums.BusinessassetgroupassignedUsersOfflineConversionDataSetTasksEnumParam `facebook:"offline_conversion_data_set_tasks"`
	PageTasks                     *[]enums.BusinessassetgroupassignedUsersPageTasksEnumParam                     `facebook:"page_tasks"`
	PixelTasks                    *[]enums.BusinessassetgroupassignedUsersPixelTasksEnumParam                    `facebook:"pixel_tasks"`
	User                          int                                                                            `facebook:"user"`
	Extra                         core.Params                                                                    `facebook:"-"`
}

func (params CreateBusinessAssetGroupAssignedUsersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdaccountTasks != nil {
		out["adaccount_tasks"] = *params.AdaccountTasks
	}
	if params.OfflineConversionDataSetTasks != nil {
		out["offline_conversion_data_set_tasks"] = *params.OfflineConversionDataSetTasks
	}
	if params.PageTasks != nil {
		out["page_tasks"] = *params.PageTasks
	}
	if params.PixelTasks != nil {
		out["pixel_tasks"] = *params.PixelTasks
	}
	out["user"] = params.User
	return out
}

func CreateBusinessAssetGroupAssignedUsers(ctx context.Context, client *core.Client, id string, params CreateBusinessAssetGroupAssignedUsersParams) (*objects.BusinessAssetGroup, error) {
	var out objects.BusinessAssetGroup
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "assigned_users"), params.ToParams(), &out)
	return &out, err
}

type DeleteBusinessAssetGroupContainedAdaccountsParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params DeleteBusinessAssetGroupContainedAdaccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func DeleteBusinessAssetGroupContainedAdaccounts(ctx context.Context, client *core.Client, id string, params DeleteBusinessAssetGroupContainedAdaccountsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "contained_adaccounts"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessAssetGroupContainedAdaccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessAssetGroupContainedAdaccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessAssetGroupContainedAdaccounts(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupContainedAdaccountsParams) (*core.Cursor[objects.AdAccount], error) {
	var out core.Cursor[objects.AdAccount]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "contained_adaccounts"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessAssetGroupContainedAdaccountsParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params CreateBusinessAssetGroupContainedAdaccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func CreateBusinessAssetGroupContainedAdaccounts(ctx context.Context, client *core.Client, id string, params CreateBusinessAssetGroupContainedAdaccountsParams) (*objects.BusinessAssetGroup, error) {
	var out objects.BusinessAssetGroup
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "contained_adaccounts"), params.ToParams(), &out)
	return &out, err
}

type DeleteBusinessAssetGroupContainedApplicationsParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params DeleteBusinessAssetGroupContainedApplicationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func DeleteBusinessAssetGroupContainedApplications(ctx context.Context, client *core.Client, id string, params DeleteBusinessAssetGroupContainedApplicationsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "contained_applications"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessAssetGroupContainedApplicationsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessAssetGroupContainedApplicationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessAssetGroupContainedApplications(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupContainedApplicationsParams) (*core.Cursor[objects.Application], error) {
	var out core.Cursor[objects.Application]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "contained_applications"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessAssetGroupContainedApplicationsParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params CreateBusinessAssetGroupContainedApplicationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func CreateBusinessAssetGroupContainedApplications(ctx context.Context, client *core.Client, id string, params CreateBusinessAssetGroupContainedApplicationsParams) (*objects.BusinessAssetGroup, error) {
	var out objects.BusinessAssetGroup
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "contained_applications"), params.ToParams(), &out)
	return &out, err
}

type DeleteBusinessAssetGroupContainedCustomConversionsParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params DeleteBusinessAssetGroupContainedCustomConversionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func DeleteBusinessAssetGroupContainedCustomConversions(ctx context.Context, client *core.Client, id string, params DeleteBusinessAssetGroupContainedCustomConversionsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "contained_custom_conversions"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessAssetGroupContainedCustomConversionsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessAssetGroupContainedCustomConversionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessAssetGroupContainedCustomConversions(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupContainedCustomConversionsParams) (*core.Cursor[objects.CustomConversion], error) {
	var out core.Cursor[objects.CustomConversion]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "contained_custom_conversions"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessAssetGroupContainedCustomConversionsParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params CreateBusinessAssetGroupContainedCustomConversionsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func CreateBusinessAssetGroupContainedCustomConversions(ctx context.Context, client *core.Client, id string, params CreateBusinessAssetGroupContainedCustomConversionsParams) (*objects.BusinessAssetGroup, error) {
	var out objects.BusinessAssetGroup
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "contained_custom_conversions"), params.ToParams(), &out)
	return &out, err
}

type DeleteBusinessAssetGroupContainedInstagramAccountsParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params DeleteBusinessAssetGroupContainedInstagramAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func DeleteBusinessAssetGroupContainedInstagramAccounts(ctx context.Context, client *core.Client, id string, params DeleteBusinessAssetGroupContainedInstagramAccountsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "contained_instagram_accounts"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessAssetGroupContainedInstagramAccountsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessAssetGroupContainedInstagramAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessAssetGroupContainedInstagramAccounts(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupContainedInstagramAccountsParams) (*core.Cursor[objects.IGUser], error) {
	var out core.Cursor[objects.IGUser]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "contained_instagram_accounts"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessAssetGroupContainedInstagramAccountsParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params CreateBusinessAssetGroupContainedInstagramAccountsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func CreateBusinessAssetGroupContainedInstagramAccounts(ctx context.Context, client *core.Client, id string, params CreateBusinessAssetGroupContainedInstagramAccountsParams) (*objects.BusinessAssetGroup, error) {
	var out objects.BusinessAssetGroup
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "contained_instagram_accounts"), params.ToParams(), &out)
	return &out, err
}

type DeleteBusinessAssetGroupContainedPagesParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params DeleteBusinessAssetGroupContainedPagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func DeleteBusinessAssetGroupContainedPages(ctx context.Context, client *core.Client, id string, params DeleteBusinessAssetGroupContainedPagesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "contained_pages"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessAssetGroupContainedPagesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessAssetGroupContainedPagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessAssetGroupContainedPages(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupContainedPagesParams) (*core.Cursor[objects.Page], error) {
	var out core.Cursor[objects.Page]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "contained_pages"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessAssetGroupContainedPagesParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params CreateBusinessAssetGroupContainedPagesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func CreateBusinessAssetGroupContainedPages(ctx context.Context, client *core.Client, id string, params CreateBusinessAssetGroupContainedPagesParams) (*objects.BusinessAssetGroup, error) {
	var out objects.BusinessAssetGroup
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "contained_pages"), params.ToParams(), &out)
	return &out, err
}

type DeleteBusinessAssetGroupContainedPixelsParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params DeleteBusinessAssetGroupContainedPixelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func DeleteBusinessAssetGroupContainedPixels(ctx context.Context, client *core.Client, id string, params DeleteBusinessAssetGroupContainedPixelsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "contained_pixels"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessAssetGroupContainedPixelsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessAssetGroupContainedPixelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessAssetGroupContainedPixels(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupContainedPixelsParams) (*core.Cursor[objects.AdsPixel], error) {
	var out core.Cursor[objects.AdsPixel]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "contained_pixels"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessAssetGroupContainedPixelsParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params CreateBusinessAssetGroupContainedPixelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func CreateBusinessAssetGroupContainedPixels(ctx context.Context, client *core.Client, id string, params CreateBusinessAssetGroupContainedPixelsParams) (*objects.BusinessAssetGroup, error) {
	var out objects.BusinessAssetGroup
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "contained_pixels"), params.ToParams(), &out)
	return &out, err
}

type DeleteBusinessAssetGroupContainedProductCatalogsParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params DeleteBusinessAssetGroupContainedProductCatalogsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func DeleteBusinessAssetGroupContainedProductCatalogs(ctx context.Context, client *core.Client, id string, params DeleteBusinessAssetGroupContainedProductCatalogsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "contained_product_catalogs"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessAssetGroupContainedProductCatalogsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessAssetGroupContainedProductCatalogsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessAssetGroupContainedProductCatalogs(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupContainedProductCatalogsParams) (*core.Cursor[objects.ProductCatalog], error) {
	var out core.Cursor[objects.ProductCatalog]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "contained_product_catalogs"), params.ToParams(), &out)
	return &out, err
}

type CreateBusinessAssetGroupContainedProductCatalogsParams struct {
	AssetID core.ID     `facebook:"asset_id"`
	Extra   core.Params `facebook:"-"`
}

func (params CreateBusinessAssetGroupContainedProductCatalogsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["asset_id"] = params.AssetID
	return out
}

func CreateBusinessAssetGroupContainedProductCatalogs(ctx context.Context, client *core.Client, id string, params CreateBusinessAssetGroupContainedProductCatalogsParams) (*objects.BusinessAssetGroup, error) {
	var out objects.BusinessAssetGroup
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "contained_product_catalogs"), params.ToParams(), &out)
	return &out, err
}

type GetBusinessAssetGroupParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetBusinessAssetGroupParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetBusinessAssetGroup(ctx context.Context, client *core.Client, id string, params GetBusinessAssetGroupParams) (*objects.BusinessAssetGroup, error) {
	var out objects.BusinessAssetGroup
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateBusinessAssetGroupParams struct {
	Name  *string     `facebook:"name"`
	Extra core.Params `facebook:"-"`
}

func (params UpdateBusinessAssetGroupParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	return out
}

func UpdateBusinessAssetGroup(ctx context.Context, client *core.Client, id string, params UpdateBusinessAssetGroupParams) (*objects.BusinessAssetGroup, error) {
	var out objects.BusinessAssetGroup
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
