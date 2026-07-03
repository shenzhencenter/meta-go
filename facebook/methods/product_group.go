package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
	"time"
)

type GetProductGroupProductsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductGroupProductsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductGroupProducts(ctx context.Context, client *core.Client, id string, params GetProductGroupProductsParams) (*core.Cursor[objects.ProductItem], error) {
	var out core.Cursor[objects.ProductItem]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "products"), params.ToParams(), &out)
	return &out, err
}

type CreateProductGroupProductsParams struct {
	AdditionalImageUrls         *[]string                                                  `facebook:"additional_image_urls"`
	AdditionalVariantAttributes *map[string]interface{}                                    `facebook:"additional_variant_attributes"`
	AgeGroup                    *enums.ProductgroupproductsAgeGroupEnumParam               `facebook:"age_group"`
	AndroidAppName              *string                                                    `facebook:"android_app_name"`
	AndroidClass                *string                                                    `facebook:"android_class"`
	AndroidPackage              *string                                                    `facebook:"android_package"`
	AndroidURL                  *string                                                    `facebook:"android_url"`
	Availability                *enums.ProductgroupproductsAvailabilityEnumParam           `facebook:"availability"`
	Brand                       *string                                                    `facebook:"brand"`
	Category                    *string                                                    `facebook:"category"`
	CheckoutURL                 *string                                                    `facebook:"checkout_url"`
	Color                       *string                                                    `facebook:"color"`
	CommerceTaxCategory         *enums.ProductgroupproductsCommerceTaxCategoryEnumParam    `facebook:"commerce_tax_category"`
	Condition                   *enums.ProductgroupproductsConditionEnumParam              `facebook:"condition"`
	Currency                    string                                                     `facebook:"currency"`
	CustomData                  *map[string]interface{}                                    `facebook:"custom_data"`
	CustomLabelX0               *string                                                    `facebook:"custom_label_0"`
	CustomLabelX1               *string                                                    `facebook:"custom_label_1"`
	CustomLabelX2               *string                                                    `facebook:"custom_label_2"`
	CustomLabelX3               *string                                                    `facebook:"custom_label_3"`
	CustomLabelX4               *string                                                    `facebook:"custom_label_4"`
	CustomNumberX0              *uint64                                                    `facebook:"custom_number_0"`
	CustomNumberX1              *uint64                                                    `facebook:"custom_number_1"`
	CustomNumberX2              *uint64                                                    `facebook:"custom_number_2"`
	CustomNumberX3              *uint64                                                    `facebook:"custom_number_3"`
	CustomNumberX4              *uint64                                                    `facebook:"custom_number_4"`
	Description                 *string                                                    `facebook:"description"`
	ExpirationDate              *string                                                    `facebook:"expiration_date"`
	FbProductCategory           *string                                                    `facebook:"fb_product_category"`
	Gender                      *enums.ProductgroupproductsGenderEnumParam                 `facebook:"gender"`
	Gtin                        *string                                                    `facebook:"gtin"`
	ImageURL                    string                                                     `facebook:"image_url"`
	Inventory                   *uint64                                                    `facebook:"inventory"`
	IosAppName                  *string                                                    `facebook:"ios_app_name"`
	IosAppStoreID               *core.ID                                                   `facebook:"ios_app_store_id"`
	IosURL                      *string                                                    `facebook:"ios_url"`
	IpadAppName                 *string                                                    `facebook:"ipad_app_name"`
	IpadAppStoreID              *core.ID                                                   `facebook:"ipad_app_store_id"`
	IpadURL                     *string                                                    `facebook:"ipad_url"`
	IphoneAppName               *string                                                    `facebook:"iphone_app_name"`
	IphoneAppStoreID            *core.ID                                                   `facebook:"iphone_app_store_id"`
	IphoneURL                   *string                                                    `facebook:"iphone_url"`
	LaunchDate                  *string                                                    `facebook:"launch_date"`
	LiveSpecialPrice            *string                                                    `facebook:"live_special_price"`
	ManufacturerPartNumber      *string                                                    `facebook:"manufacturer_part_number"`
	MarkedForProductLaunch      *enums.ProductgroupproductsMarkedForProductLaunchEnumParam `facebook:"marked_for_product_launch"`
	Material                    *string                                                    `facebook:"material"`
	MobileLink                  *string                                                    `facebook:"mobile_link"`
	Name                        string                                                     `facebook:"name"`
	OrderingIndex               *uint64                                                    `facebook:"ordering_index"`
	Pattern                     *string                                                    `facebook:"pattern"`
	Price                       uint64                                                     `facebook:"price"`
	ProductPriorityX0           *float64                                                   `facebook:"product_priority_0"`
	ProductPriorityX1           *float64                                                   `facebook:"product_priority_1"`
	ProductPriorityX2           *float64                                                   `facebook:"product_priority_2"`
	ProductPriorityX3           *float64                                                   `facebook:"product_priority_3"`
	ProductPriorityX4           *float64                                                   `facebook:"product_priority_4"`
	ProductType                 *string                                                    `facebook:"product_type"`
	QuantityToSellOnFacebook    *uint64                                                    `facebook:"quantity_to_sell_on_facebook"`
	RetailerID                  core.ID                                                    `facebook:"retailer_id"`
	ReturnPolicyDays            *uint64                                                    `facebook:"return_policy_days"`
	RichTextDescription         *string                                                    `facebook:"rich_text_description"`
	SalePrice                   *uint64                                                    `facebook:"sale_price"`
	SalePriceEndDate            *time.Time                                                 `facebook:"sale_price_end_date"`
	SalePriceStartDate          *time.Time                                                 `facebook:"sale_price_start_date"`
	ShortDescription            *string                                                    `facebook:"short_description"`
	Size                        *string                                                    `facebook:"size"`
	StartDate                   *string                                                    `facebook:"start_date"`
	URL                         *string                                                    `facebook:"url"`
	Visibility                  *enums.ProductgroupproductsVisibilityEnumParam             `facebook:"visibility"`
	WindowsPhoneAppID           *core.ID                                                   `facebook:"windows_phone_app_id"`
	WindowsPhoneAppName         *string                                                    `facebook:"windows_phone_app_name"`
	WindowsPhoneURL             *string                                                    `facebook:"windows_phone_url"`
	Extra                       core.Params                                                `facebook:"-"`
}

func (params CreateProductGroupProductsParams) ToParams() core.Params {
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
	out["currency"] = params.Currency
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
	out["image_url"] = params.ImageURL
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
	out["name"] = params.Name
	if params.OrderingIndex != nil {
		out["ordering_index"] = *params.OrderingIndex
	}
	if params.Pattern != nil {
		out["pattern"] = *params.Pattern
	}
	out["price"] = params.Price
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
	out["retailer_id"] = params.RetailerID
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

func CreateProductGroupProducts(ctx context.Context, client *core.Client, id string, params CreateProductGroupProductsParams) (*objects.ProductItem, error) {
	var out objects.ProductItem
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "products"), params.ToParams(), &out)
	return &out, err
}

type DeleteProductGroupParams struct {
	DeletionMethod *enums.ProductgroupDeletionMethod `facebook:"deletion_method"`
	Extra          core.Params                       `facebook:"-"`
}

func (params DeleteProductGroupParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.DeletionMethod != nil {
		out["deletion_method"] = *params.DeletionMethod
	}
	return out
}

func DeleteProductGroup(ctx context.Context, client *core.Client, id string, params DeleteProductGroupParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetProductGroupParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductGroupParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductGroup(ctx context.Context, client *core.Client, id string, params GetProductGroupParams) (*objects.ProductGroup, error) {
	var out objects.ProductGroup
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateProductGroupParams struct {
	DefaultProductID *core.ID                  `facebook:"default_product_id"`
	Variants         *[]map[string]interface{} `facebook:"variants"`
	Extra            core.Params               `facebook:"-"`
}

func (params UpdateProductGroupParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.DefaultProductID != nil {
		out["default_product_id"] = *params.DefaultProductID
	}
	if params.Variants != nil {
		out["variants"] = *params.Variants
	}
	return out
}

func UpdateProductGroup(ctx context.Context, client *core.Client, id string, params UpdateProductGroupParams) (*objects.ProductGroup, error) {
	var out objects.ProductGroup
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
