package enums

type HomelistinggetCapabilities string

const (
	HomelistinggetCapabilitiesB2cMarketplace           HomelistinggetCapabilities = "B2C_MARKETPLACE"
	HomelistinggetCapabilitiesBizMsgAiAgent            HomelistinggetCapabilities = "BIZ_MSG_AI_AGENT"
	HomelistinggetCapabilitiesBusinessInboxInMessenger HomelistinggetCapabilities = "BUSINESS_INBOX_IN_MESSENGER"
	HomelistinggetCapabilitiesBuyOnFacebook            HomelistinggetCapabilities = "BUY_ON_FACEBOOK"
	HomelistinggetCapabilitiesC2cMarketplace           HomelistinggetCapabilities = "C2C_MARKETPLACE"
	HomelistinggetCapabilitiesCpasParentCatalog        HomelistinggetCapabilities = "CPAS_PARENT_CATALOG"
	HomelistinggetCapabilitiesDa                       HomelistinggetCapabilities = "DA"
	HomelistinggetCapabilitiesDailyDeals               HomelistinggetCapabilities = "DAILY_DEALS"
	HomelistinggetCapabilitiesDailyDealsLegacy         HomelistinggetCapabilities = "DAILY_DEALS_LEGACY"
	HomelistinggetCapabilitiesEvent                    HomelistinggetCapabilities = "EVENT"
	HomelistinggetCapabilitiesEventDeprecated          HomelistinggetCapabilities = "EVENT_DEPRECATED"
	HomelistinggetCapabilitiesGroups                   HomelistinggetCapabilities = "GROUPS"
	HomelistinggetCapabilitiesIgOnsiteShopping         HomelistinggetCapabilities = "IG_ONSITE_SHOPPING"
	HomelistinggetCapabilitiesIgProductTagging         HomelistinggetCapabilities = "IG_PRODUCT_TAGGING"
	HomelistinggetCapabilitiesLdp                      HomelistinggetCapabilities = "LDP"
	HomelistinggetCapabilitiesMarketplace              HomelistinggetCapabilities = "MARKETPLACE"
	HomelistinggetCapabilitiesMarketplaceAdsDeprecated HomelistinggetCapabilities = "MARKETPLACE_ADS_DEPRECATED"
	HomelistinggetCapabilitiesMarketplaceHomeRentals   HomelistinggetCapabilities = "MARKETPLACE_HOME_RENTALS"
	HomelistinggetCapabilitiesMarketplaceHomeSales     HomelistinggetCapabilities = "MARKETPLACE_HOME_SALES"
	HomelistinggetCapabilitiesMarketplaceMotors        HomelistinggetCapabilities = "MARKETPLACE_MOTORS"
	HomelistinggetCapabilitiesMarketplaceShops         HomelistinggetCapabilities = "MARKETPLACE_SHOPS"
	HomelistinggetCapabilitiesMiniShops                HomelistinggetCapabilities = "MINI_SHOPS"
	HomelistinggetCapabilitiesNeighborhoods            HomelistinggetCapabilities = "NEIGHBORHOODS"
	HomelistinggetCapabilitiesOfflineConversions       HomelistinggetCapabilities = "OFFLINE_CONVERSIONS"
	HomelistinggetCapabilitiesProfile                  HomelistinggetCapabilities = "PROFILE"
	HomelistinggetCapabilitiesService                  HomelistinggetCapabilities = "SERVICE"
	HomelistinggetCapabilitiesShops                    HomelistinggetCapabilities = "SHOPS"
	HomelistinggetCapabilitiesTestCapability           HomelistinggetCapabilities = "TEST_CAPABILITY"
	HomelistinggetCapabilitiesUniversalCheckout        HomelistinggetCapabilities = "UNIVERSAL_CHECKOUT"
	HomelistinggetCapabilitiesUsMarketplace            HomelistinggetCapabilities = "US_MARKETPLACE"
	HomelistinggetCapabilitiesWhatsapp                 HomelistinggetCapabilities = "WHATSAPP"
	HomelistinggetCapabilitiesWhatsappMarketingMessage HomelistinggetCapabilities = "WHATSAPP_MARKETING_MESSAGE"
)

func (value HomelistinggetCapabilities) String() string {
	return string(value)
}
