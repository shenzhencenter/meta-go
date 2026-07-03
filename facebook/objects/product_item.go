package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type ProductItem struct {
	AdditionalImageCdnUrls                                *[][]map[string]string                                  `json:"additional_image_cdn_urls,omitempty"`
	AdditionalImageUrls                                   *[]string                                               `json:"additional_image_urls,omitempty"`
	AdditionalVariantAttributes                           *[]map[string]string                                    `json:"additional_variant_attributes,omitempty"`
	AffiliateSellerRating                                 *float64                                                `json:"affiliate_seller_rating,omitempty"`
	AffiliateSellerRatingCount                            *int                                                    `json:"affiliate_seller_rating_count,omitempty"`
	AgeGroup                                              *enums.ProductitemAgeGroupEnum                          `json:"age_group,omitempty"`
	Applinks                                              *CatalogItemAppLinks                                    `json:"applinks,omitempty"`
	Availability                                          *enums.ProductitemAvailabilityEnum                      `json:"availability,omitempty"`
	BaseCommissionRate                                    *int                                                    `json:"base_commission_rate,omitempty"`
	Brand                                                 *string                                                 `json:"brand,omitempty"`
	BundleItems                                           *[]string                                               `json:"bundle_items,omitempty"`
	BundleRetailerIds                                     *[]core.ID                                              `json:"bundle_retailer_ids,omitempty"`
	CapabilitiesDisabledByUser                            *[]string                                               `json:"capabilities_disabled_by_user,omitempty"`
	CapabilityToReviewStatus                              *[]map[string]enums.ProductitemCapabilityToReviewStatus `json:"capability_to_review_status,omitempty"`
	Category                                              *string                                                 `json:"category,omitempty"`
	CategoryRanking                                       *int                                                    `json:"category_ranking,omitempty"`
	CategorySpecificFields                                *CatalogSubVerticalList                                 `json:"category_specific_fields,omitempty"`
	Color                                                 *string                                                 `json:"color,omitempty"`
	CommerceInsights                                      *ProductItemCommerceInsights                            `json:"commerce_insights,omitempty"`
	Condition                                             *enums.ProductitemConditionEnum                         `json:"condition,omitempty"`
	Currency                                              *string                                                 `json:"currency,omitempty"`
	CustomData                                            *[]map[string]string                                    `json:"custom_data,omitempty"`
	CustomLabelX0                                         *string                                                 `json:"custom_label_0,omitempty"`
	CustomLabelX1                                         *string                                                 `json:"custom_label_1,omitempty"`
	CustomLabelX2                                         *string                                                 `json:"custom_label_2,omitempty"`
	CustomLabelX3                                         *string                                                 `json:"custom_label_3,omitempty"`
	CustomLabelX4                                         *string                                                 `json:"custom_label_4,omitempty"`
	CustomNumberX0                                        *string                                                 `json:"custom_number_0,omitempty"`
	CustomNumberX1                                        *string                                                 `json:"custom_number_1,omitempty"`
	CustomNumberX2                                        *string                                                 `json:"custom_number_2,omitempty"`
	CustomNumberX3                                        *string                                                 `json:"custom_number_3,omitempty"`
	CustomNumberX4                                        *string                                                 `json:"custom_number_4,omitempty"`
	Description                                           *string                                                 `json:"description,omitempty"`
	Errors                                                *[]ProductItemError                                     `json:"errors,omitempty"`
	ExpirationDate                                        *string                                                 `json:"expiration_date,omitempty"`
	FbProductCategory                                     *string                                                 `json:"fb_product_category,omitempty"`
	Gender                                                *enums.ProductitemGenderEnum                            `json:"gender,omitempty"`
	GeneratedBackgroundImages                             *[]AIGeneratedProductImage                              `json:"generated_background_images,omitempty"`
	GeneratedBackgroundImagesAdUsage                      *bool                                                   `json:"generated_background_images_ad_usage,omitempty"`
	Gtin                                                  *string                                                 `json:"gtin,omitempty"`
	ID                                                    *core.ID                                                `json:"id,omitempty"`
	ImageCdnUrls                                          *[]map[string]string                                    `json:"image_cdn_urls,omitempty"`
	ImageFetchStatus                                      *enums.ProductitemImageFetchStatus                      `json:"image_fetch_status,omitempty"`
	ImageURL                                              *string                                                 `json:"image_url,omitempty"`
	Images                                                *[]string                                               `json:"images,omitempty"`
	ImporterAddress                                       *ProductItemImporterAddress                             `json:"importer_address,omitempty"`
	ImporterName                                          *string                                                 `json:"importer_name,omitempty"`
	InvalidationErrors                                    *[]ProductItemInvalidationError                         `json:"invalidation_errors,omitempty"`
	Inventory                                             *int                                                    `json:"inventory,omitempty"`
	IsBundleHero                                          *bool                                                   `json:"is_bundle_hero,omitempty"`
	LiveSpecialPrice                                      *string                                                 `json:"live_special_price,omitempty"`
	ManufacturerInfo                                      *string                                                 `json:"manufacturer_info,omitempty"`
	ManufacturerPartNumber                                *string                                                 `json:"manufacturer_part_number,omitempty"`
	MarkedForProductLaunch                                *string                                                 `json:"marked_for_product_launch,omitempty"`
	Material                                              *string                                                 `json:"material,omitempty"`
	MobileLink                                            *string                                                 `json:"mobile_link,omitempty"`
	Name                                                  *string                                                 `json:"name,omitempty"`
	OfferDisclaimer                                       *string                                                 `json:"offer_disclaimer,omitempty"`
	OfferDisclaimerURL                                    *string                                                 `json:"offer_disclaimer_url,omitempty"`
	OrderingIndex                                         *int                                                    `json:"ordering_index,omitempty"`
	OriginCountry                                         *string                                                 `json:"origin_country,omitempty"`
	OverallRanking                                        *int                                                    `json:"overall_ranking,omitempty"`
	ParentProductID                                       *core.ID                                                `json:"parent_product_id,omitempty"`
	Pattern                                               *string                                                 `json:"pattern,omitempty"`
	PostConversionSignalBasedEnforcementAppealEligibility *bool                                                   `json:"post_conversion_signal_based_enforcement_appeal_eligibility,omitempty"`
	Price                                                 *string                                                 `json:"price,omitempty"`
	ProductCatalog                                        *ProductCatalog                                         `json:"product_catalog,omitempty"`
	ProductFeed                                           *ProductFeed                                            `json:"product_feed,omitempty"`
	ProductGroup                                          *ProductGroup                                           `json:"product_group,omitempty"`
	ProductLocalInfo                                      *ProductItemLocalInfo                                   `json:"product_local_info,omitempty"`
	ProductRelationship                                   *string                                                 `json:"product_relationship,omitempty"`
	ProductType                                           *string                                                 `json:"product_type,omitempty"`
	QuantityToSellOnFacebook                              *int                                                    `json:"quantity_to_sell_on_facebook,omitempty"`
	RetailerID                                            *core.ID                                                `json:"retailer_id,omitempty"`
	RetailerProductGroupID                                *core.ID                                                `json:"retailer_product_group_id,omitempty"`
	ReviewRejectionReasons                                *[]string                                               `json:"review_rejection_reasons,omitempty"`
	ReviewStatus                                          *enums.ProductitemReviewStatus                          `json:"review_status,omitempty"`
	RichTextDescription                                   *string                                                 `json:"rich_text_description,omitempty"`
	SalePrice                                             *string                                                 `json:"sale_price,omitempty"`
	SalePriceEndDate                                      *string                                                 `json:"sale_price_end_date,omitempty"`
	SalePriceStartDate                                    *string                                                 `json:"sale_price_start_date,omitempty"`
	ShippingWeightUnit                                    *enums.ProductitemShippingWeightUnit                    `json:"shipping_weight_unit,omitempty"`
	ShippingWeightValue                                   *float64                                                `json:"shipping_weight_value,omitempty"`
	ShortDescription                                      *string                                                 `json:"short_description,omitempty"`
	Size                                                  *string                                                 `json:"size,omitempty"`
	StartDate                                             *string                                                 `json:"start_date,omitempty"`
	Status                                                *enums.ProductitemStatus                                `json:"status,omitempty"`
	Tags                                                  *[]string                                               `json:"tags,omitempty"`
	URL                                                   *string                                                 `json:"url,omitempty"`
	VendorID                                              *core.ID                                                `json:"vendor_id,omitempty"`
	VideoFetchStatus                                      *enums.ProductitemVideoFetchStatus                      `json:"video_fetch_status,omitempty"`
	Videos                                                *[]ProductItemVideoData                                 `json:"videos,omitempty"`
	Visibility                                            *enums.ProductitemVisibilityEnum                        `json:"visibility,omitempty"`
	WaComplianceCategory                                  *string                                                 `json:"wa_compliance_category,omitempty"`
}

var ProductItemFields = struct {
	AdditionalImageCdnUrls                                string
	AdditionalImageUrls                                   string
	AdditionalVariantAttributes                           string
	AffiliateSellerRating                                 string
	AffiliateSellerRatingCount                            string
	AgeGroup                                              string
	Applinks                                              string
	Availability                                          string
	BaseCommissionRate                                    string
	Brand                                                 string
	BundleItems                                           string
	BundleRetailerIds                                     string
	CapabilitiesDisabledByUser                            string
	CapabilityToReviewStatus                              string
	Category                                              string
	CategoryRanking                                       string
	CategorySpecificFields                                string
	Color                                                 string
	CommerceInsights                                      string
	Condition                                             string
	Currency                                              string
	CustomData                                            string
	CustomLabelX0                                         string
	CustomLabelX1                                         string
	CustomLabelX2                                         string
	CustomLabelX3                                         string
	CustomLabelX4                                         string
	CustomNumberX0                                        string
	CustomNumberX1                                        string
	CustomNumberX2                                        string
	CustomNumberX3                                        string
	CustomNumberX4                                        string
	Description                                           string
	Errors                                                string
	ExpirationDate                                        string
	FbProductCategory                                     string
	Gender                                                string
	GeneratedBackgroundImages                             string
	GeneratedBackgroundImagesAdUsage                      string
	Gtin                                                  string
	ID                                                    string
	ImageCdnUrls                                          string
	ImageFetchStatus                                      string
	ImageURL                                              string
	Images                                                string
	ImporterAddress                                       string
	ImporterName                                          string
	InvalidationErrors                                    string
	Inventory                                             string
	IsBundleHero                                          string
	LiveSpecialPrice                                      string
	ManufacturerInfo                                      string
	ManufacturerPartNumber                                string
	MarkedForProductLaunch                                string
	Material                                              string
	MobileLink                                            string
	Name                                                  string
	OfferDisclaimer                                       string
	OfferDisclaimerURL                                    string
	OrderingIndex                                         string
	OriginCountry                                         string
	OverallRanking                                        string
	ParentProductID                                       string
	Pattern                                               string
	PostConversionSignalBasedEnforcementAppealEligibility string
	Price                                                 string
	ProductCatalog                                        string
	ProductFeed                                           string
	ProductGroup                                          string
	ProductLocalInfo                                      string
	ProductRelationship                                   string
	ProductType                                           string
	QuantityToSellOnFacebook                              string
	RetailerID                                            string
	RetailerProductGroupID                                string
	ReviewRejectionReasons                                string
	ReviewStatus                                          string
	RichTextDescription                                   string
	SalePrice                                             string
	SalePriceEndDate                                      string
	SalePriceStartDate                                    string
	ShippingWeightUnit                                    string
	ShippingWeightValue                                   string
	ShortDescription                                      string
	Size                                                  string
	StartDate                                             string
	Status                                                string
	Tags                                                  string
	URL                                                   string
	VendorID                                              string
	VideoFetchStatus                                      string
	Videos                                                string
	Visibility                                            string
	WaComplianceCategory                                  string
}{
	AdditionalImageCdnUrls:           "additional_image_cdn_urls",
	AdditionalImageUrls:              "additional_image_urls",
	AdditionalVariantAttributes:      "additional_variant_attributes",
	AffiliateSellerRating:            "affiliate_seller_rating",
	AffiliateSellerRatingCount:       "affiliate_seller_rating_count",
	AgeGroup:                         "age_group",
	Applinks:                         "applinks",
	Availability:                     "availability",
	BaseCommissionRate:               "base_commission_rate",
	Brand:                            "brand",
	BundleItems:                      "bundle_items",
	BundleRetailerIds:                "bundle_retailer_ids",
	CapabilitiesDisabledByUser:       "capabilities_disabled_by_user",
	CapabilityToReviewStatus:         "capability_to_review_status",
	Category:                         "category",
	CategoryRanking:                  "category_ranking",
	CategorySpecificFields:           "category_specific_fields",
	Color:                            "color",
	CommerceInsights:                 "commerce_insights",
	Condition:                        "condition",
	Currency:                         "currency",
	CustomData:                       "custom_data",
	CustomLabelX0:                    "custom_label_0",
	CustomLabelX1:                    "custom_label_1",
	CustomLabelX2:                    "custom_label_2",
	CustomLabelX3:                    "custom_label_3",
	CustomLabelX4:                    "custom_label_4",
	CustomNumberX0:                   "custom_number_0",
	CustomNumberX1:                   "custom_number_1",
	CustomNumberX2:                   "custom_number_2",
	CustomNumberX3:                   "custom_number_3",
	CustomNumberX4:                   "custom_number_4",
	Description:                      "description",
	Errors:                           "errors",
	ExpirationDate:                   "expiration_date",
	FbProductCategory:                "fb_product_category",
	Gender:                           "gender",
	GeneratedBackgroundImages:        "generated_background_images",
	GeneratedBackgroundImagesAdUsage: "generated_background_images_ad_usage",
	Gtin:                             "gtin",
	ID:                               "id",
	ImageCdnUrls:                     "image_cdn_urls",
	ImageFetchStatus:                 "image_fetch_status",
	ImageURL:                         "image_url",
	Images:                           "images",
	ImporterAddress:                  "importer_address",
	ImporterName:                     "importer_name",
	InvalidationErrors:               "invalidation_errors",
	Inventory:                        "inventory",
	IsBundleHero:                     "is_bundle_hero",
	LiveSpecialPrice:                 "live_special_price",
	ManufacturerInfo:                 "manufacturer_info",
	ManufacturerPartNumber:           "manufacturer_part_number",
	MarkedForProductLaunch:           "marked_for_product_launch",
	Material:                         "material",
	MobileLink:                       "mobile_link",
	Name:                             "name",
	OfferDisclaimer:                  "offer_disclaimer",
	OfferDisclaimerURL:               "offer_disclaimer_url",
	OrderingIndex:                    "ordering_index",
	OriginCountry:                    "origin_country",
	OverallRanking:                   "overall_ranking",
	ParentProductID:                  "parent_product_id",
	Pattern:                          "pattern",
	PostConversionSignalBasedEnforcementAppealEligibility: "post_conversion_signal_based_enforcement_appeal_eligibility",
	Price:                    "price",
	ProductCatalog:           "product_catalog",
	ProductFeed:              "product_feed",
	ProductGroup:             "product_group",
	ProductLocalInfo:         "product_local_info",
	ProductRelationship:      "product_relationship",
	ProductType:              "product_type",
	QuantityToSellOnFacebook: "quantity_to_sell_on_facebook",
	RetailerID:               "retailer_id",
	RetailerProductGroupID:   "retailer_product_group_id",
	ReviewRejectionReasons:   "review_rejection_reasons",
	ReviewStatus:             "review_status",
	RichTextDescription:      "rich_text_description",
	SalePrice:                "sale_price",
	SalePriceEndDate:         "sale_price_end_date",
	SalePriceStartDate:       "sale_price_start_date",
	ShippingWeightUnit:       "shipping_weight_unit",
	ShippingWeightValue:      "shipping_weight_value",
	ShortDescription:         "short_description",
	Size:                     "size",
	StartDate:                "start_date",
	Status:                   "status",
	Tags:                     "tags",
	URL:                      "url",
	VendorID:                 "vendor_id",
	VideoFetchStatus:         "video_fetch_status",
	Videos:                   "videos",
	Visibility:               "visibility",
	WaComplianceCategory:     "wa_compliance_category",
}
