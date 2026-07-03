package enums

type ProductcatalogdiagnosticgroupAffectedChannels string

const (
	ProductcatalogdiagnosticgroupAffectedChannelsB2cMarketplace           ProductcatalogdiagnosticgroupAffectedChannels = "b2c_marketplace"
	ProductcatalogdiagnosticgroupAffectedChannelsC2cMarketplace           ProductcatalogdiagnosticgroupAffectedChannels = "c2c_marketplace"
	ProductcatalogdiagnosticgroupAffectedChannelsDa                       ProductcatalogdiagnosticgroupAffectedChannels = "da"
	ProductcatalogdiagnosticgroupAffectedChannelsDailyDeals               ProductcatalogdiagnosticgroupAffectedChannels = "daily_deals"
	ProductcatalogdiagnosticgroupAffectedChannelsDailyDealsLegacy         ProductcatalogdiagnosticgroupAffectedChannels = "daily_deals_legacy"
	ProductcatalogdiagnosticgroupAffectedChannelsIgProductTagging         ProductcatalogdiagnosticgroupAffectedChannels = "ig_product_tagging"
	ProductcatalogdiagnosticgroupAffectedChannelsMarketplace              ProductcatalogdiagnosticgroupAffectedChannels = "marketplace"
	ProductcatalogdiagnosticgroupAffectedChannelsMarketplaceAdsDeprecated ProductcatalogdiagnosticgroupAffectedChannels = "marketplace_ads_deprecated"
	ProductcatalogdiagnosticgroupAffectedChannelsMarketplaceShops         ProductcatalogdiagnosticgroupAffectedChannels = "marketplace_shops"
	ProductcatalogdiagnosticgroupAffectedChannelsMiniShops                ProductcatalogdiagnosticgroupAffectedChannels = "mini_shops"
	ProductcatalogdiagnosticgroupAffectedChannelsOfflineConversions       ProductcatalogdiagnosticgroupAffectedChannels = "offline_conversions"
	ProductcatalogdiagnosticgroupAffectedChannelsShops                    ProductcatalogdiagnosticgroupAffectedChannels = "shops"
	ProductcatalogdiagnosticgroupAffectedChannelsUniversalCheckout        ProductcatalogdiagnosticgroupAffectedChannels = "universal_checkout"
	ProductcatalogdiagnosticgroupAffectedChannelsWhatsapp                 ProductcatalogdiagnosticgroupAffectedChannels = "whatsapp"
)

func (value ProductcatalogdiagnosticgroupAffectedChannels) String() string {
	return string(value)
}
