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

func GetCommerceMerchantSettingsCommerceOrdersBatchCall(id string, params GetCommerceMerchantSettingsCommerceOrdersParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "commerce_orders"), params.ToParams(), options...)
}

func NewGetCommerceMerchantSettingsCommerceOrdersBatchRequest(id string, params GetCommerceMerchantSettingsCommerceOrdersParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CommerceOrder]] {
	return core.NewBatchRequest[core.Cursor[objects.CommerceOrder]](GetCommerceMerchantSettingsCommerceOrdersBatchCall(id, params, options...))
}

func DecodeGetCommerceMerchantSettingsCommerceOrdersBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CommerceOrder], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CommerceOrder]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCommerceMerchantSettingsCommerceOrders(ctx context.Context, client *core.Client, id string, params GetCommerceMerchantSettingsCommerceOrdersParams) (*core.Cursor[objects.CommerceOrder], error) {
	var out core.Cursor[objects.CommerceOrder]
	call := GetCommerceMerchantSettingsCommerceOrdersBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetCommerceMerchantSettingsCommercePayoutsBatchCall(id string, params GetCommerceMerchantSettingsCommercePayoutsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "commerce_payouts"), params.ToParams(), options...)
}

func NewGetCommerceMerchantSettingsCommercePayoutsBatchRequest(id string, params GetCommerceMerchantSettingsCommercePayoutsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CommercePayout]] {
	return core.NewBatchRequest[core.Cursor[objects.CommercePayout]](GetCommerceMerchantSettingsCommercePayoutsBatchCall(id, params, options...))
}

func DecodeGetCommerceMerchantSettingsCommercePayoutsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CommercePayout], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CommercePayout]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCommerceMerchantSettingsCommercePayouts(ctx context.Context, client *core.Client, id string, params GetCommerceMerchantSettingsCommercePayoutsParams) (*core.Cursor[objects.CommercePayout], error) {
	var out core.Cursor[objects.CommercePayout]
	call := GetCommerceMerchantSettingsCommercePayoutsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetCommerceMerchantSettingsOrderManagementAppsBatchCall(id string, params GetCommerceMerchantSettingsOrderManagementAppsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "order_management_apps"), params.ToParams(), options...)
}

func NewGetCommerceMerchantSettingsOrderManagementAppsBatchRequest(id string, params GetCommerceMerchantSettingsOrderManagementAppsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Application]] {
	return core.NewBatchRequest[core.Cursor[objects.Application]](GetCommerceMerchantSettingsOrderManagementAppsBatchCall(id, params, options...))
}

func DecodeGetCommerceMerchantSettingsOrderManagementAppsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Application], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Application]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCommerceMerchantSettingsOrderManagementApps(ctx context.Context, client *core.Client, id string, params GetCommerceMerchantSettingsOrderManagementAppsParams) (*core.Cursor[objects.Application], error) {
	var out core.Cursor[objects.Application]
	call := GetCommerceMerchantSettingsOrderManagementAppsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetCommerceMerchantSettingsProductCatalogsBatchCall(id string, params GetCommerceMerchantSettingsProductCatalogsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "product_catalogs"), params.ToParams(), options...)
}

func NewGetCommerceMerchantSettingsProductCatalogsBatchRequest(id string, params GetCommerceMerchantSettingsProductCatalogsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProductCatalog]] {
	return core.NewBatchRequest[core.Cursor[objects.ProductCatalog]](GetCommerceMerchantSettingsProductCatalogsBatchCall(id, params, options...))
}

func DecodeGetCommerceMerchantSettingsProductCatalogsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProductCatalog], error) {
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

func GetCommerceMerchantSettingsProductCatalogs(ctx context.Context, client *core.Client, id string, params GetCommerceMerchantSettingsProductCatalogsParams) (*core.Cursor[objects.ProductCatalog], error) {
	var out core.Cursor[objects.ProductCatalog]
	call := GetCommerceMerchantSettingsProductCatalogsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetCommerceMerchantSettingsReturnsBatchCall(id string, params GetCommerceMerchantSettingsReturnsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "returns"), params.ToParams(), options...)
}

func NewGetCommerceMerchantSettingsReturnsBatchRequest(id string, params GetCommerceMerchantSettingsReturnsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CommerceReturn]] {
	return core.NewBatchRequest[core.Cursor[objects.CommerceReturn]](GetCommerceMerchantSettingsReturnsBatchCall(id, params, options...))
}

func DecodeGetCommerceMerchantSettingsReturnsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CommerceReturn], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CommerceReturn]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCommerceMerchantSettingsReturns(ctx context.Context, client *core.Client, id string, params GetCommerceMerchantSettingsReturnsParams) (*core.Cursor[objects.CommerceReturn], error) {
	var out core.Cursor[objects.CommerceReturn]
	call := GetCommerceMerchantSettingsReturnsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetCommerceMerchantSettingsSetupStatusBatchCall(id string, params GetCommerceMerchantSettingsSetupStatusParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "setup_status"), params.ToParams(), options...)
}

func NewGetCommerceMerchantSettingsSetupStatusBatchRequest(id string, params GetCommerceMerchantSettingsSetupStatusParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CommerceMerchantSettingsSetupStatus]] {
	return core.NewBatchRequest[core.Cursor[objects.CommerceMerchantSettingsSetupStatus]](GetCommerceMerchantSettingsSetupStatusBatchCall(id, params, options...))
}

func DecodeGetCommerceMerchantSettingsSetupStatusBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CommerceMerchantSettingsSetupStatus], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CommerceMerchantSettingsSetupStatus]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCommerceMerchantSettingsSetupStatus(ctx context.Context, client *core.Client, id string, params GetCommerceMerchantSettingsSetupStatusParams) (*core.Cursor[objects.CommerceMerchantSettingsSetupStatus], error) {
	var out core.Cursor[objects.CommerceMerchantSettingsSetupStatus]
	call := GetCommerceMerchantSettingsSetupStatusBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateCommerceMerchantSettingsShippingProfilesBatchCall(id string, params CreateCommerceMerchantSettingsShippingProfilesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "shipping_profiles"), params.ToParams(), options...)
}

func NewCreateCommerceMerchantSettingsShippingProfilesBatchRequest(id string, params CreateCommerceMerchantSettingsShippingProfilesParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](CreateCommerceMerchantSettingsShippingProfilesBatchCall(id, params, options...))
}

func DecodeCreateCommerceMerchantSettingsShippingProfilesBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func CreateCommerceMerchantSettingsShippingProfiles(ctx context.Context, client *core.Client, id string, params CreateCommerceMerchantSettingsShippingProfilesParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := CreateCommerceMerchantSettingsShippingProfilesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetCommerceMerchantSettingsShopsBatchCall(id string, params GetCommerceMerchantSettingsShopsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "shops"), params.ToParams(), options...)
}

func NewGetCommerceMerchantSettingsShopsBatchRequest(id string, params GetCommerceMerchantSettingsShopsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Shop]] {
	return core.NewBatchRequest[core.Cursor[objects.Shop]](GetCommerceMerchantSettingsShopsBatchCall(id, params, options...))
}

func DecodeGetCommerceMerchantSettingsShopsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Shop], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Shop]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCommerceMerchantSettingsShops(ctx context.Context, client *core.Client, id string, params GetCommerceMerchantSettingsShopsParams) (*core.Cursor[objects.Shop], error) {
	var out core.Cursor[objects.Shop]
	call := GetCommerceMerchantSettingsShopsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetCommerceMerchantSettingsTaxSettingsBatchCall(id string, params GetCommerceMerchantSettingsTaxSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "tax_settings"), params.ToParams(), options...)
}

func NewGetCommerceMerchantSettingsTaxSettingsBatchRequest(id string, params GetCommerceMerchantSettingsTaxSettingsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CommerceMerchantTaxSettings]] {
	return core.NewBatchRequest[core.Cursor[objects.CommerceMerchantTaxSettings]](GetCommerceMerchantSettingsTaxSettingsBatchCall(id, params, options...))
}

func DecodeGetCommerceMerchantSettingsTaxSettingsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CommerceMerchantTaxSettings], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CommerceMerchantTaxSettings]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCommerceMerchantSettingsTaxSettings(ctx context.Context, client *core.Client, id string, params GetCommerceMerchantSettingsTaxSettingsParams) (*core.Cursor[objects.CommerceMerchantTaxSettings], error) {
	var out core.Cursor[objects.CommerceMerchantTaxSettings]
	call := GetCommerceMerchantSettingsTaxSettingsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetCommerceMerchantSettingsBatchCall(id string, params GetCommerceMerchantSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCommerceMerchantSettingsBatchRequest(id string, params GetCommerceMerchantSettingsParams, options ...core.BatchOption) *core.BatchRequest[objects.CommerceMerchantSettings] {
	return core.NewBatchRequest[objects.CommerceMerchantSettings](GetCommerceMerchantSettingsBatchCall(id, params, options...))
}

func DecodeGetCommerceMerchantSettingsBatchResponse(response *core.BatchResponse) (*objects.CommerceMerchantSettings, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CommerceMerchantSettings
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCommerceMerchantSettings(ctx context.Context, client *core.Client, id string, params GetCommerceMerchantSettingsParams) (*objects.CommerceMerchantSettings, error) {
	var out objects.CommerceMerchantSettings
	call := GetCommerceMerchantSettingsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func UpdateCommerceMerchantSettingsBatchCall(id string, params UpdateCommerceMerchantSettingsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateCommerceMerchantSettingsBatchRequest(id string, params UpdateCommerceMerchantSettingsParams, options ...core.BatchOption) *core.BatchRequest[objects.CommerceMerchantSettings] {
	return core.NewBatchRequest[objects.CommerceMerchantSettings](UpdateCommerceMerchantSettingsBatchCall(id, params, options...))
}

func DecodeUpdateCommerceMerchantSettingsBatchResponse(response *core.BatchResponse) (*objects.CommerceMerchantSettings, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CommerceMerchantSettings
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateCommerceMerchantSettings(ctx context.Context, client *core.Client, id string, params UpdateCommerceMerchantSettingsParams) (*objects.CommerceMerchantSettings, error) {
	var out objects.CommerceMerchantSettings
	call := UpdateCommerceMerchantSettingsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
