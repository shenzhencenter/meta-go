package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
	"time"
)

type GetCommerceMerchantSettingsCommerceOrdersParams struct {
	Filters       *[]enums.CommercemerchantsettingscommerceOrdersFiltersEnumParam `facebook:"filters"`
	State         *[]enums.CommercemerchantsettingscommerceOrdersStateEnumParam   `facebook:"state"`
	UpdatedAfter  *time.Time                                                      `facebook:"updated_after"`
	UpdatedBefore *time.Time                                                      `facebook:"updated_before"`
	Extra         core.Params                                                     `facebook:"-"`
}

func (params GetCommerceMerchantSettingsCommerceOrdersParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Filters != nil {
		out["filters"] = *params.Filters
	}
	if params.State != nil {
		out["state"] = *params.State
	}
	if params.UpdatedAfter != nil {
		out["updated_after"] = *params.UpdatedAfter
	}
	if params.UpdatedBefore != nil {
		out["updated_before"] = *params.UpdatedBefore
	}
	return out
}

func GetCommerceMerchantSettingsCommerceOrders(ctx context.Context, client *core.Client, id string, params GetCommerceMerchantSettingsCommerceOrdersParams) (*core.Cursor[objects.CommerceOrder], error) {
	var out core.Cursor[objects.CommerceOrder]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "commerce_orders"), params.ToParams(), &out)
	return &out, err
}

type GetCommerceMerchantSettingsCommercePayoutsParams struct {
	EndTime   *time.Time  `facebook:"end_time"`
	StartTime *time.Time  `facebook:"start_time"`
	Extra     core.Params `facebook:"-"`
}

func (params GetCommerceMerchantSettingsCommercePayoutsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.EndTime != nil {
		out["end_time"] = *params.EndTime
	}
	if params.StartTime != nil {
		out["start_time"] = *params.StartTime
	}
	return out
}

func GetCommerceMerchantSettingsCommercePayouts(ctx context.Context, client *core.Client, id string, params GetCommerceMerchantSettingsCommercePayoutsParams) (*core.Cursor[objects.CommercePayout], error) {
	var out core.Cursor[objects.CommercePayout]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "commerce_payouts"), params.ToParams(), &out)
	return &out, err
}

type GetCommerceMerchantSettingsOrderManagementAppsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCommerceMerchantSettingsOrderManagementAppsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCommerceMerchantSettingsOrderManagementApps(ctx context.Context, client *core.Client, id string, params GetCommerceMerchantSettingsOrderManagementAppsParams) (*core.Cursor[objects.Application], error) {
	var out core.Cursor[objects.Application]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "order_management_apps"), params.ToParams(), &out)
	return &out, err
}

type GetCommerceMerchantSettingsProductCatalogsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCommerceMerchantSettingsProductCatalogsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCommerceMerchantSettingsProductCatalogs(ctx context.Context, client *core.Client, id string, params GetCommerceMerchantSettingsProductCatalogsParams) (*core.Cursor[objects.ProductCatalog], error) {
	var out core.Cursor[objects.ProductCatalog]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "product_catalogs"), params.ToParams(), &out)
	return &out, err
}

type GetCommerceMerchantSettingsReturnsParams struct {
	EndTimeCreated   *time.Time                                                `facebook:"end_time_created"`
	MerchantReturnID *core.ID                                                  `facebook:"merchant_return_id"`
	StartTimeCreated *time.Time                                                `facebook:"start_time_created"`
	Statuses         *[]enums.CommercemerchantsettingsreturnsStatusesEnumParam `facebook:"statuses"`
	Extra            core.Params                                               `facebook:"-"`
}

func (params GetCommerceMerchantSettingsReturnsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.EndTimeCreated != nil {
		out["end_time_created"] = *params.EndTimeCreated
	}
	if params.MerchantReturnID != nil {
		out["merchant_return_id"] = *params.MerchantReturnID
	}
	if params.StartTimeCreated != nil {
		out["start_time_created"] = *params.StartTimeCreated
	}
	if params.Statuses != nil {
		out["statuses"] = *params.Statuses
	}
	return out
}

func GetCommerceMerchantSettingsReturns(ctx context.Context, client *core.Client, id string, params GetCommerceMerchantSettingsReturnsParams) (*core.Cursor[objects.CommerceReturn], error) {
	var out core.Cursor[objects.CommerceReturn]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "returns"), params.ToParams(), &out)
	return &out, err
}

type GetCommerceMerchantSettingsSetupStatusParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCommerceMerchantSettingsSetupStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCommerceMerchantSettingsSetupStatus(ctx context.Context, client *core.Client, id string, params GetCommerceMerchantSettingsSetupStatusParams) (*core.Cursor[objects.CommerceMerchantSettingsSetupStatus], error) {
	var out core.Cursor[objects.CommerceMerchantSettingsSetupStatus]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "setup_status"), params.ToParams(), &out)
	return &out, err
}

type CreateCommerceMerchantSettingsShippingProfilesParams struct {
	HandlingTime             *map[string]interface{}  `facebook:"handling_time"`
	IsDefault                *bool                    `facebook:"is_default"`
	IsDefaultShippingProfile *bool                    `facebook:"is_default_shipping_profile"`
	Name                     string                   `facebook:"name"`
	ReferenceID              *core.ID                 `facebook:"reference_id"`
	ShippingDestinations     []map[string]interface{} `facebook:"shipping_destinations"`
	Extra                    core.Params              `facebook:"-"`
}

func (params CreateCommerceMerchantSettingsShippingProfilesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.HandlingTime != nil {
		out["handling_time"] = *params.HandlingTime
	}
	if params.IsDefault != nil {
		out["is_default"] = *params.IsDefault
	}
	if params.IsDefaultShippingProfile != nil {
		out["is_default_shipping_profile"] = *params.IsDefaultShippingProfile
	}
	out["name"] = params.Name
	if params.ReferenceID != nil {
		out["reference_id"] = *params.ReferenceID
	}
	out["shipping_destinations"] = params.ShippingDestinations
	return out
}

func CreateCommerceMerchantSettingsShippingProfiles(ctx context.Context, client *core.Client, id string, params CreateCommerceMerchantSettingsShippingProfilesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "shipping_profiles"), params.ToParams(), &out)
	return &out, err
}

type GetCommerceMerchantSettingsShopsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCommerceMerchantSettingsShopsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCommerceMerchantSettingsShops(ctx context.Context, client *core.Client, id string, params GetCommerceMerchantSettingsShopsParams) (*core.Cursor[objects.Shop], error) {
	var out core.Cursor[objects.Shop]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "shops"), params.ToParams(), &out)
	return &out, err
}

type GetCommerceMerchantSettingsTaxSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCommerceMerchantSettingsTaxSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCommerceMerchantSettingsTaxSettings(ctx context.Context, client *core.Client, id string, params GetCommerceMerchantSettingsTaxSettingsParams) (*core.Cursor[objects.CommerceMerchantTaxSettings], error) {
	var out core.Cursor[objects.CommerceMerchantTaxSettings]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "tax_settings"), params.ToParams(), &out)
	return &out, err
}

type GetCommerceMerchantSettingsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCommerceMerchantSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCommerceMerchantSettings(ctx context.Context, client *core.Client, id string, params GetCommerceMerchantSettingsParams) (*objects.CommerceMerchantSettings, error) {
	var out objects.CommerceMerchantSettings
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateCommerceMerchantSettingsParams struct {
	CheckoutConfig         *map[string]interface{}                       `facebook:"checkout_config"`
	KoreaFtcListing        *string                                       `facebook:"korea_ftc_listing"`
	MerchantStatus         *enums.CommercemerchantsettingsMerchantStatus `facebook:"merchant_status"`
	PrivacyPolicyLocalized *map[string]interface{}                       `facebook:"privacy_policy_localized"`
	Extra                  core.Params                                   `facebook:"-"`
}

func (params UpdateCommerceMerchantSettingsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CheckoutConfig != nil {
		out["checkout_config"] = *params.CheckoutConfig
	}
	if params.KoreaFtcListing != nil {
		out["korea_ftc_listing"] = *params.KoreaFtcListing
	}
	if params.MerchantStatus != nil {
		out["merchant_status"] = *params.MerchantStatus
	}
	if params.PrivacyPolicyLocalized != nil {
		out["privacy_policy_localized"] = *params.PrivacyPolicyLocalized
	}
	return out
}

func UpdateCommerceMerchantSettings(ctx context.Context, client *core.Client, id string, params UpdateCommerceMerchantSettingsParams) (*objects.CommerceMerchantSettings, error) {
	var out objects.CommerceMerchantSettings
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
