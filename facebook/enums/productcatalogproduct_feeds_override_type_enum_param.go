package enums

type ProductcatalogproductFeedsOverrideTypeEnumParam string

const (
	ProductcatalogproductFeedsOverrideTypeEnumParamBatchAPILanguageOrCountry      ProductcatalogproductFeedsOverrideTypeEnumParam = "BATCH_API_LANGUAGE_OR_COUNTRY"
	ProductcatalogproductFeedsOverrideTypeEnumParamCatalogSegmentCustomizeDefault ProductcatalogproductFeedsOverrideTypeEnumParam = "CATALOG_SEGMENT_CUSTOMIZE_DEFAULT"
	ProductcatalogproductFeedsOverrideTypeEnumParamCountry                        ProductcatalogproductFeedsOverrideTypeEnumParam = "COUNTRY"
	ProductcatalogproductFeedsOverrideTypeEnumParamLanguage                       ProductcatalogproductFeedsOverrideTypeEnumParam = "LANGUAGE"
	ProductcatalogproductFeedsOverrideTypeEnumParamLanguageAndCountry             ProductcatalogproductFeedsOverrideTypeEnumParam = "LANGUAGE_AND_COUNTRY"
	ProductcatalogproductFeedsOverrideTypeEnumParamLocal                          ProductcatalogproductFeedsOverrideTypeEnumParam = "LOCAL"
	ProductcatalogproductFeedsOverrideTypeEnumParamSmartPixelLanguageOrCountry    ProductcatalogproductFeedsOverrideTypeEnumParam = "SMART_PIXEL_LANGUAGE_OR_COUNTRY"
	ProductcatalogproductFeedsOverrideTypeEnumParamVersion                        ProductcatalogproductFeedsOverrideTypeEnumParam = "VERSION"
)

func (value ProductcatalogproductFeedsOverrideTypeEnumParam) String() string {
	return string(value)
}
