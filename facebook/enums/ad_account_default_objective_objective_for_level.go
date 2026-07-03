package enums

type AdaccountdefaultobjectiveObjectiveForLevel string

const (
	AdaccountdefaultobjectiveObjectiveForLevelAppInstalls         AdaccountdefaultobjectiveObjectiveForLevel = "APP_INSTALLS"
	AdaccountdefaultobjectiveObjectiveForLevelBrandAwareness      AdaccountdefaultobjectiveObjectiveForLevel = "BRAND_AWARENESS"
	AdaccountdefaultobjectiveObjectiveForLevelEventResponses      AdaccountdefaultobjectiveObjectiveForLevel = "EVENT_RESPONSES"
	AdaccountdefaultobjectiveObjectiveForLevelLeadGeneration      AdaccountdefaultobjectiveObjectiveForLevel = "LEAD_GENERATION"
	AdaccountdefaultobjectiveObjectiveForLevelLinkClicks          AdaccountdefaultobjectiveObjectiveForLevel = "LINK_CLICKS"
	AdaccountdefaultobjectiveObjectiveForLevelLocalAwareness      AdaccountdefaultobjectiveObjectiveForLevel = "LOCAL_AWARENESS"
	AdaccountdefaultobjectiveObjectiveForLevelMessages            AdaccountdefaultobjectiveObjectiveForLevel = "MESSAGES"
	AdaccountdefaultobjectiveObjectiveForLevelOfferClaims         AdaccountdefaultobjectiveObjectiveForLevel = "OFFER_CLAIMS"
	AdaccountdefaultobjectiveObjectiveForLevelOutcomeAppPromotion AdaccountdefaultobjectiveObjectiveForLevel = "OUTCOME_APP_PROMOTION"
	AdaccountdefaultobjectiveObjectiveForLevelOutcomeAwareness    AdaccountdefaultobjectiveObjectiveForLevel = "OUTCOME_AWARENESS"
	AdaccountdefaultobjectiveObjectiveForLevelOutcomeEngagement   AdaccountdefaultobjectiveObjectiveForLevel = "OUTCOME_ENGAGEMENT"
	AdaccountdefaultobjectiveObjectiveForLevelOutcomeLeads        AdaccountdefaultobjectiveObjectiveForLevel = "OUTCOME_LEADS"
	AdaccountdefaultobjectiveObjectiveForLevelOutcomeSales        AdaccountdefaultobjectiveObjectiveForLevel = "OUTCOME_SALES"
	AdaccountdefaultobjectiveObjectiveForLevelOutcomeTraffic      AdaccountdefaultobjectiveObjectiveForLevel = "OUTCOME_TRAFFIC"
	AdaccountdefaultobjectiveObjectiveForLevelPageLikes           AdaccountdefaultobjectiveObjectiveForLevel = "PAGE_LIKES"
	AdaccountdefaultobjectiveObjectiveForLevelPostEngagement      AdaccountdefaultobjectiveObjectiveForLevel = "POST_ENGAGEMENT"
	AdaccountdefaultobjectiveObjectiveForLevelProductCatalogSales AdaccountdefaultobjectiveObjectiveForLevel = "PRODUCT_CATALOG_SALES"
	AdaccountdefaultobjectiveObjectiveForLevelReach               AdaccountdefaultobjectiveObjectiveForLevel = "REACH"
	AdaccountdefaultobjectiveObjectiveForLevelStoreVisits         AdaccountdefaultobjectiveObjectiveForLevel = "STORE_VISITS"
	AdaccountdefaultobjectiveObjectiveForLevelVideoViews          AdaccountdefaultobjectiveObjectiveForLevel = "VIDEO_VIEWS"
	AdaccountdefaultobjectiveObjectiveForLevelWebsiteConversions  AdaccountdefaultobjectiveObjectiveForLevel = "WEBSITE_CONVERSIONS"
)

func (value AdaccountdefaultobjectiveObjectiveForLevel) String() string {
	return string(value)
}
