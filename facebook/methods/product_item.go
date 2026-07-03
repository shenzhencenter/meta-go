package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
	"time"
)

type GetProductItemChannelsToIntegrityStatusParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductItemChannelsToIntegrityStatusParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductItemChannelsToIntegrityStatusBatchCall(id string, params GetProductItemChannelsToIntegrityStatusParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "channels_to_integrity_status"), params.ToParams(), options...)
}

func NewGetProductItemChannelsToIntegrityStatusBatchRequest(id string, params GetProductItemChannelsToIntegrityStatusParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]] {
	return core.NewBatchRequest[core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]](GetProductItemChannelsToIntegrityStatusBatchCall(id, params, options...))
}

func DecodeGetProductItemChannelsToIntegrityStatusBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductItemChannelsToIntegrityStatus(ctx context.Context, client *core.Client, id string, params GetProductItemChannelsToIntegrityStatusParams) (*core.Cursor[objects.CatalogItemChannelsToIntegrityStatus], error) {
	var out core.Cursor[objects.CatalogItemChannelsToIntegrityStatus]
	call := GetProductItemChannelsToIntegrityStatusBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetProductItemOverrideDetailsParams struct {
	Keys  *[]string                                      `facebook:"keys"`
	Type  *enums.ProductitemoverrideDetailsTypeEnumParam `facebook:"type"`
	Extra core.Params                                    `facebook:"-"`
}

func (params GetProductItemOverrideDetailsParams) ToParams() core.Params {
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

func GetProductItemOverrideDetailsBatchCall(id string, params GetProductItemOverrideDetailsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "override_details"), params.ToParams(), options...)
}

func NewGetProductItemOverrideDetailsBatchRequest(id string, params GetProductItemOverrideDetailsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.OverrideDetails]] {
	return core.NewBatchRequest[core.Cursor[objects.OverrideDetails]](GetProductItemOverrideDetailsBatchCall(id, params, options...))
}

func DecodeGetProductItemOverrideDetailsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.OverrideDetails], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.OverrideDetails]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductItemOverrideDetails(ctx context.Context, client *core.Client, id string, params GetProductItemOverrideDetailsParams) (*core.Cursor[objects.OverrideDetails], error) {
	var out core.Cursor[objects.OverrideDetails]
	call := GetProductItemOverrideDetailsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetProductItemProductSetsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductItemProductSetsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductItemProductSetsBatchCall(id string, params GetProductItemProductSetsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "product_sets"), params.ToParams(), options...)
}

func NewGetProductItemProductSetsBatchRequest(id string, params GetProductItemProductSetsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProductSet]] {
	return core.NewBatchRequest[core.Cursor[objects.ProductSet]](GetProductItemProductSetsBatchCall(id, params, options...))
}

func DecodeGetProductItemProductSetsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProductSet], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ProductSet]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductItemProductSets(ctx context.Context, client *core.Client, id string, params GetProductItemProductSetsParams) (*core.Cursor[objects.ProductSet], error) {
	var out core.Cursor[objects.ProductSet]
	call := GetProductItemProductSetsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetProductItemVideosMetadataParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductItemVideosMetadataParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductItemVideosMetadataBatchCall(id string, params GetProductItemVideosMetadataParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "videos_metadata"), params.ToParams(), options...)
}

func NewGetProductItemVideosMetadataBatchRequest(id string, params GetProductItemVideosMetadataParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.DynamicVideoMetadata]] {
	return core.NewBatchRequest[core.Cursor[objects.DynamicVideoMetadata]](GetProductItemVideosMetadataBatchCall(id, params, options...))
}

func DecodeGetProductItemVideosMetadataBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.DynamicVideoMetadata], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.DynamicVideoMetadata]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductItemVideosMetadata(ctx context.Context, client *core.Client, id string, params GetProductItemVideosMetadataParams) (*core.Cursor[objects.DynamicVideoMetadata], error) {
	var out core.Cursor[objects.DynamicVideoMetadata]
	call := GetProductItemVideosMetadataBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type DeleteProductItemParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteProductItemParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteProductItemBatchCall(id string, params DeleteProductItemParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteProductItemBatchRequest(id string, params DeleteProductItemParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteProductItemBatchCall(id, params, options...))
}

func DecodeDeleteProductItemBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteProductItem(ctx context.Context, client *core.Client, id string, params DeleteProductItemParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	call := DeleteProductItemBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetProductItemParams struct {
	CatalogID        *core.ID    `facebook:"catalog_id"`
	ImageHeight      *uint64     `facebook:"image_height"`
	ImageWidth       *uint64     `facebook:"image_width"`
	OverrideCountry  *string     `facebook:"override_country"`
	OverrideLanguage *string     `facebook:"override_language"`
	Extra            core.Params `facebook:"-"`
}

func (params GetProductItemParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CatalogID != nil {
		out["catalog_id"] = *params.CatalogID
	}
	if params.ImageHeight != nil {
		out["image_height"] = *params.ImageHeight
	}
	if params.ImageWidth != nil {
		out["image_width"] = *params.ImageWidth
	}
	if params.OverrideCountry != nil {
		out["override_country"] = *params.OverrideCountry
	}
	if params.OverrideLanguage != nil {
		out["override_language"] = *params.OverrideLanguage
	}
	return out
}

func GetProductItemBatchCall(id string, params GetProductItemParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetProductItemBatchRequest(id string, params GetProductItemParams, options ...core.BatchOption) *core.BatchRequest[objects.ProductItem] {
	return core.NewBatchRequest[objects.ProductItem](GetProductItemBatchCall(id, params, options...))
}

func DecodeGetProductItemBatchResponse(response *core.BatchResponse) (*objects.ProductItem, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ProductItem
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductItem(ctx context.Context, client *core.Client, id string, params GetProductItemParams) (*objects.ProductItem, error) {
	var out objects.ProductItem
	call := GetProductItemBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type UpdateProductItemParams struct {
	AdditionalImageUrls         *[]string                                `facebook:"additional_image_urls"`
	AdditionalVariantAttributes *map[string]interface{}                  `facebook:"additional_variant_attributes"`
	AgeGroup                    *enums.ProductitemAgeGroup               `facebook:"age_group"`
	AndroidAppName              *string                                  `facebook:"android_app_name"`
	AndroidClass                *string                                  `facebook:"android_class"`
	AndroidPackage              *string                                  `facebook:"android_package"`
	AndroidURL                  *string                                  `facebook:"android_url"`
	Availability                *enums.ProductitemAvailability           `facebook:"availability"`
	Brand                       *string                                  `facebook:"brand"`
	Category                    *string                                  `facebook:"category"`
	CategorySpecificFields      *map[string]interface{}                  `facebook:"category_specific_fields"`
	CheckoutURL                 *string                                  `facebook:"checkout_url"`
	Color                       *string                                  `facebook:"color"`
	CommerceTaxCategory         *enums.ProductitemCommerceTaxCategory    `facebook:"commerce_tax_category"`
	Condition                   *enums.ProductitemCondition              `facebook:"condition"`
	Currency                    *string                                  `facebook:"currency"`
	CustomData                  *map[string]interface{}                  `facebook:"custom_data"`
	CustomLabelX0               *string                                  `facebook:"custom_label_0"`
	CustomLabelX1               *string                                  `facebook:"custom_label_1"`
	CustomLabelX2               *string                                  `facebook:"custom_label_2"`
	CustomLabelX3               *string                                  `facebook:"custom_label_3"`
	CustomLabelX4               *string                                  `facebook:"custom_label_4"`
	CustomNumberX0              *uint64                                  `facebook:"custom_number_0"`
	CustomNumberX1              *uint64                                  `facebook:"custom_number_1"`
	CustomNumberX2              *uint64                                  `facebook:"custom_number_2"`
	CustomNumberX3              *uint64                                  `facebook:"custom_number_3"`
	CustomNumberX4              *uint64                                  `facebook:"custom_number_4"`
	Description                 *string                                  `facebook:"description"`
	ExpirationDate              *string                                  `facebook:"expiration_date"`
	FbProductCategory           *string                                  `facebook:"fb_product_category"`
	Gender                      *enums.ProductitemGender                 `facebook:"gender"`
	Gtin                        *string                                  `facebook:"gtin"`
	ImageURL                    *string                                  `facebook:"image_url"`
	ImporterAddress             *map[string]interface{}                  `facebook:"importer_address"`
	ImporterName                *string                                  `facebook:"importer_name"`
	Inventory                   *uint64                                  `facebook:"inventory"`
	IosAppName                  *string                                  `facebook:"ios_app_name"`
	IosAppStoreID               *core.ID                                 `facebook:"ios_app_store_id"`
	IosURL                      *string                                  `facebook:"ios_url"`
	IpadAppName                 *string                                  `facebook:"ipad_app_name"`
	IpadAppStoreID              *core.ID                                 `facebook:"ipad_app_store_id"`
	IpadURL                     *string                                  `facebook:"ipad_url"`
	IphoneAppName               *string                                  `facebook:"iphone_app_name"`
	IphoneAppStoreID            *core.ID                                 `facebook:"iphone_app_store_id"`
	IphoneURL                   *string                                  `facebook:"iphone_url"`
	LaunchDate                  *string                                  `facebook:"launch_date"`
	LiveSpecialPrice            *string                                  `facebook:"live_special_price"`
	ManufacturerInfo            *string                                  `facebook:"manufacturer_info"`
	ManufacturerPartNumber      *string                                  `facebook:"manufacturer_part_number"`
	MarkedForProductLaunch      *enums.ProductitemMarkedForProductLaunch `facebook:"marked_for_product_launch"`
	Material                    *string                                  `facebook:"material"`
	MobileLink                  *string                                  `facebook:"mobile_link"`
	Name                        *string                                  `facebook:"name"`
	OrderingIndex               *uint64                                  `facebook:"ordering_index"`
	OriginCountry               *enums.ProductitemOriginCountry          `facebook:"origin_country"`
	Pattern                     *string                                  `facebook:"pattern"`
	Price                       *uint64                                  `facebook:"price"`
	ProductPriorityX0           *float64                                 `facebook:"product_priority_0"`
	ProductPriorityX1           *float64                                 `facebook:"product_priority_1"`
	ProductPriorityX2           *float64                                 `facebook:"product_priority_2"`
	ProductPriorityX3           *float64                                 `facebook:"product_priority_3"`
	ProductPriorityX4           *float64                                 `facebook:"product_priority_4"`
	ProductType                 *string                                  `facebook:"product_type"`
	QuantityToSellOnFacebook    *uint64                                  `facebook:"quantity_to_sell_on_facebook"`
	RetailerID                  *core.ID                                 `facebook:"retailer_id"`
	ReturnPolicyDays            *uint64                                  `facebook:"return_policy_days"`
	RichTextDescription         *string                                  `facebook:"rich_text_description"`
	SalePrice                   *uint64                                  `facebook:"sale_price"`
	SalePriceEndDate            *time.Time                               `facebook:"sale_price_end_date"`
	SalePriceStartDate          *time.Time                               `facebook:"sale_price_start_date"`
	ShortDescription            *string                                  `facebook:"short_description"`
	Size                        *string                                  `facebook:"size"`
	StartDate                   *string                                  `facebook:"start_date"`
	URL                         *string                                  `facebook:"url"`
	Visibility                  *enums.ProductitemVisibility             `facebook:"visibility"`
	WaComplianceCategory        *enums.ProductitemWaComplianceCategory   `facebook:"wa_compliance_category"`
	WindowsPhoneAppID           *core.ID                                 `facebook:"windows_phone_app_id"`
	WindowsPhoneAppName         *string                                  `facebook:"windows_phone_app_name"`
	WindowsPhoneURL             *string                                  `facebook:"windows_phone_url"`
	Extra                       core.Params                              `facebook:"-"`
}

func (params UpdateProductItemParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdditionalImageUrls != nil {
		out["additional_image_urls"] = *params.AdditionalImageUrls
	}
	if params.AdditionalVariantAttributes != nil {
		out["additional_variant_attributes"] = *params.AdditionalVariantAttributes
	}
	if params.AgeGroup != nil {
		out["age_group"] = *params.AgeGroup
	}
	if params.AndroidAppName != nil {
		out["android_app_name"] = *params.AndroidAppName
	}
	if params.AndroidClass != nil {
		out["android_class"] = *params.AndroidClass
	}
	if params.AndroidPackage != nil {
		out["android_package"] = *params.AndroidPackage
	}
	if params.AndroidURL != nil {
		out["android_url"] = *params.AndroidURL
	}
	if params.Availability != nil {
		out["availability"] = *params.Availability
	}
	if params.Brand != nil {
		out["brand"] = *params.Brand
	}
	if params.Category != nil {
		out["category"] = *params.Category
	}
	if params.CategorySpecificFields != nil {
		out["category_specific_fields"] = *params.CategorySpecificFields
	}
	if params.CheckoutURL != nil {
		out["checkout_url"] = *params.CheckoutURL
	}
	if params.Color != nil {
		out["color"] = *params.Color
	}
	if params.CommerceTaxCategory != nil {
		out["commerce_tax_category"] = *params.CommerceTaxCategory
	}
	if params.Condition != nil {
		out["condition"] = *params.Condition
	}
	if params.Currency != nil {
		out["currency"] = *params.Currency
	}
	if params.CustomData != nil {
		out["custom_data"] = *params.CustomData
	}
	if params.CustomLabelX0 != nil {
		out["custom_label_0"] = *params.CustomLabelX0
	}
	if params.CustomLabelX1 != nil {
		out["custom_label_1"] = *params.CustomLabelX1
	}
	if params.CustomLabelX2 != nil {
		out["custom_label_2"] = *params.CustomLabelX2
	}
	if params.CustomLabelX3 != nil {
		out["custom_label_3"] = *params.CustomLabelX3
	}
	if params.CustomLabelX4 != nil {
		out["custom_label_4"] = *params.CustomLabelX4
	}
	if params.CustomNumberX0 != nil {
		out["custom_number_0"] = *params.CustomNumberX0
	}
	if params.CustomNumberX1 != nil {
		out["custom_number_1"] = *params.CustomNumberX1
	}
	if params.CustomNumberX2 != nil {
		out["custom_number_2"] = *params.CustomNumberX2
	}
	if params.CustomNumberX3 != nil {
		out["custom_number_3"] = *params.CustomNumberX3
	}
	if params.CustomNumberX4 != nil {
		out["custom_number_4"] = *params.CustomNumberX4
	}
	if params.Description != nil {
		out["description"] = *params.Description
	}
	if params.ExpirationDate != nil {
		out["expiration_date"] = *params.ExpirationDate
	}
	if params.FbProductCategory != nil {
		out["fb_product_category"] = *params.FbProductCategory
	}
	if params.Gender != nil {
		out["gender"] = *params.Gender
	}
	if params.Gtin != nil {
		out["gtin"] = *params.Gtin
	}
	if params.ImageURL != nil {
		out["image_url"] = *params.ImageURL
	}
	if params.ImporterAddress != nil {
		out["importer_address"] = *params.ImporterAddress
	}
	if params.ImporterName != nil {
		out["importer_name"] = *params.ImporterName
	}
	if params.Inventory != nil {
		out["inventory"] = *params.Inventory
	}
	if params.IosAppName != nil {
		out["ios_app_name"] = *params.IosAppName
	}
	if params.IosAppStoreID != nil {
		out["ios_app_store_id"] = *params.IosAppStoreID
	}
	if params.IosURL != nil {
		out["ios_url"] = *params.IosURL
	}
	if params.IpadAppName != nil {
		out["ipad_app_name"] = *params.IpadAppName
	}
	if params.IpadAppStoreID != nil {
		out["ipad_app_store_id"] = *params.IpadAppStoreID
	}
	if params.IpadURL != nil {
		out["ipad_url"] = *params.IpadURL
	}
	if params.IphoneAppName != nil {
		out["iphone_app_name"] = *params.IphoneAppName
	}
	if params.IphoneAppStoreID != nil {
		out["iphone_app_store_id"] = *params.IphoneAppStoreID
	}
	if params.IphoneURL != nil {
		out["iphone_url"] = *params.IphoneURL
	}
	if params.LaunchDate != nil {
		out["launch_date"] = *params.LaunchDate
	}
	if params.LiveSpecialPrice != nil {
		out["live_special_price"] = *params.LiveSpecialPrice
	}
	if params.ManufacturerInfo != nil {
		out["manufacturer_info"] = *params.ManufacturerInfo
	}
	if params.ManufacturerPartNumber != nil {
		out["manufacturer_part_number"] = *params.ManufacturerPartNumber
	}
	if params.MarkedForProductLaunch != nil {
		out["marked_for_product_launch"] = *params.MarkedForProductLaunch
	}
	if params.Material != nil {
		out["material"] = *params.Material
	}
	if params.MobileLink != nil {
		out["mobile_link"] = *params.MobileLink
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.OrderingIndex != nil {
		out["ordering_index"] = *params.OrderingIndex
	}
	if params.OriginCountry != nil {
		out["origin_country"] = *params.OriginCountry
	}
	if params.Pattern != nil {
		out["pattern"] = *params.Pattern
	}
	if params.Price != nil {
		out["price"] = *params.Price
	}
	if params.ProductPriorityX0 != nil {
		out["product_priority_0"] = *params.ProductPriorityX0
	}
	if params.ProductPriorityX1 != nil {
		out["product_priority_1"] = *params.ProductPriorityX1
	}
	if params.ProductPriorityX2 != nil {
		out["product_priority_2"] = *params.ProductPriorityX2
	}
	if params.ProductPriorityX3 != nil {
		out["product_priority_3"] = *params.ProductPriorityX3
	}
	if params.ProductPriorityX4 != nil {
		out["product_priority_4"] = *params.ProductPriorityX4
	}
	if params.ProductType != nil {
		out["product_type"] = *params.ProductType
	}
	if params.QuantityToSellOnFacebook != nil {
		out["quantity_to_sell_on_facebook"] = *params.QuantityToSellOnFacebook
	}
	if params.RetailerID != nil {
		out["retailer_id"] = *params.RetailerID
	}
	if params.ReturnPolicyDays != nil {
		out["return_policy_days"] = *params.ReturnPolicyDays
	}
	if params.RichTextDescription != nil {
		out["rich_text_description"] = *params.RichTextDescription
	}
	if params.SalePrice != nil {
		out["sale_price"] = *params.SalePrice
	}
	if params.SalePriceEndDate != nil {
		out["sale_price_end_date"] = *params.SalePriceEndDate
	}
	if params.SalePriceStartDate != nil {
		out["sale_price_start_date"] = *params.SalePriceStartDate
	}
	if params.ShortDescription != nil {
		out["short_description"] = *params.ShortDescription
	}
	if params.Size != nil {
		out["size"] = *params.Size
	}
	if params.StartDate != nil {
		out["start_date"] = *params.StartDate
	}
	if params.URL != nil {
		out["url"] = *params.URL
	}
	if params.Visibility != nil {
		out["visibility"] = *params.Visibility
	}
	if params.WaComplianceCategory != nil {
		out["wa_compliance_category"] = *params.WaComplianceCategory
	}
	if params.WindowsPhoneAppID != nil {
		out["windows_phone_app_id"] = *params.WindowsPhoneAppID
	}
	if params.WindowsPhoneAppName != nil {
		out["windows_phone_app_name"] = *params.WindowsPhoneAppName
	}
	if params.WindowsPhoneURL != nil {
		out["windows_phone_url"] = *params.WindowsPhoneURL
	}
	return out
}

func UpdateProductItemBatchCall(id string, params UpdateProductItemParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateProductItemBatchRequest(id string, params UpdateProductItemParams, options ...core.BatchOption) *core.BatchRequest[objects.ProductItem] {
	return core.NewBatchRequest[objects.ProductItem](UpdateProductItemBatchCall(id, params, options...))
}

func DecodeUpdateProductItemBatchResponse(response *core.BatchResponse) (*objects.ProductItem, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ProductItem
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateProductItem(ctx context.Context, client *core.Client, id string, params UpdateProductItemParams) (*objects.ProductItem, error) {
	var out objects.ProductItem
	call := UpdateProductItemBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
