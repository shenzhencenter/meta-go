package enums

type ProductcatalogdiagnosticsAffectedChannelsEnumParam string

const (
	ProductcatalogdiagnosticsAffectedChannelsEnumParamB2cMarketplace           ProductcatalogdiagnosticsAffectedChannelsEnumParam = "b2c_marketplace"
	ProductcatalogdiagnosticsAffectedChannelsEnumParamC2cMarketplace           ProductcatalogdiagnosticsAffectedChannelsEnumParam = "c2c_marketplace"
	ProductcatalogdiagnosticsAffectedChannelsEnumParamDa                       ProductcatalogdiagnosticsAffectedChannelsEnumParam = "da"
	ProductcatalogdiagnosticsAffectedChannelsEnumParamDailyDeals               ProductcatalogdiagnosticsAffectedChannelsEnumParam = "daily_deals"
	ProductcatalogdiagnosticsAffectedChannelsEnumParamDailyDealsLegacy         ProductcatalogdiagnosticsAffectedChannelsEnumParam = "daily_deals_legacy"
	ProductcatalogdiagnosticsAffectedChannelsEnumParamIgProductTagging         ProductcatalogdiagnosticsAffectedChannelsEnumParam = "ig_product_tagging"
	ProductcatalogdiagnosticsAffectedChannelsEnumParamMarketplace              ProductcatalogdiagnosticsAffectedChannelsEnumParam = "marketplace"
	ProductcatalogdiagnosticsAffectedChannelsEnumParamMarketplaceAdsDeprecated ProductcatalogdiagnosticsAffectedChannelsEnumParam = "marketplace_ads_deprecated"
	ProductcatalogdiagnosticsAffectedChannelsEnumParamMarketplaceShops         ProductcatalogdiagnosticsAffectedChannelsEnumParam = "marketplace_shops"
	ProductcatalogdiagnosticsAffectedChannelsEnumParamMiniShops                ProductcatalogdiagnosticsAffectedChannelsEnumParam = "mini_shops"
	ProductcatalogdiagnosticsAffectedChannelsEnumParamOfflineConversions       ProductcatalogdiagnosticsAffectedChannelsEnumParam = "offline_conversions"
	ProductcatalogdiagnosticsAffectedChannelsEnumParamShops                    ProductcatalogdiagnosticsAffectedChannelsEnumParam = "shops"
	ProductcatalogdiagnosticsAffectedChannelsEnumParamUniversalCheckout        ProductcatalogdiagnosticsAffectedChannelsEnumParam = "universal_checkout"
	ProductcatalogdiagnosticsAffectedChannelsEnumParamWhatsapp                 ProductcatalogdiagnosticsAffectedChannelsEnumParam = "whatsapp"
)

func (value ProductcatalogdiagnosticsAffectedChannelsEnumParam) String() string {
	return string(value)
}
