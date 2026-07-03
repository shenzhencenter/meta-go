package enums

type AdaccounttargetingsearchObjectiveEnumParam string

const (
	AdaccounttargetingsearchObjectiveEnumParamAppInstalls         AdaccounttargetingsearchObjectiveEnumParam = "APP_INSTALLS"
	AdaccounttargetingsearchObjectiveEnumParamBrandAwareness      AdaccounttargetingsearchObjectiveEnumParam = "BRAND_AWARENESS"
	AdaccounttargetingsearchObjectiveEnumParamConversions         AdaccounttargetingsearchObjectiveEnumParam = "CONVERSIONS"
	AdaccounttargetingsearchObjectiveEnumParamEventResponses      AdaccounttargetingsearchObjectiveEnumParam = "EVENT_RESPONSES"
	AdaccounttargetingsearchObjectiveEnumParamLeadGeneration      AdaccounttargetingsearchObjectiveEnumParam = "LEAD_GENERATION"
	AdaccounttargetingsearchObjectiveEnumParamLinkClicks          AdaccounttargetingsearchObjectiveEnumParam = "LINK_CLICKS"
	AdaccounttargetingsearchObjectiveEnumParamLocalAwareness      AdaccounttargetingsearchObjectiveEnumParam = "LOCAL_AWARENESS"
	AdaccounttargetingsearchObjectiveEnumParamMessages            AdaccounttargetingsearchObjectiveEnumParam = "MESSAGES"
	AdaccounttargetingsearchObjectiveEnumParamOfferClaims         AdaccounttargetingsearchObjectiveEnumParam = "OFFER_CLAIMS"
	AdaccounttargetingsearchObjectiveEnumParamOutcomeAppPromotion AdaccounttargetingsearchObjectiveEnumParam = "OUTCOME_APP_PROMOTION"
	AdaccounttargetingsearchObjectiveEnumParamOutcomeAwareness    AdaccounttargetingsearchObjectiveEnumParam = "OUTCOME_AWARENESS"
	AdaccounttargetingsearchObjectiveEnumParamOutcomeEngagement   AdaccounttargetingsearchObjectiveEnumParam = "OUTCOME_ENGAGEMENT"
	AdaccounttargetingsearchObjectiveEnumParamOutcomeLeads        AdaccounttargetingsearchObjectiveEnumParam = "OUTCOME_LEADS"
	AdaccounttargetingsearchObjectiveEnumParamOutcomeSales        AdaccounttargetingsearchObjectiveEnumParam = "OUTCOME_SALES"
	AdaccounttargetingsearchObjectiveEnumParamOutcomeTraffic      AdaccounttargetingsearchObjectiveEnumParam = "OUTCOME_TRAFFIC"
	AdaccounttargetingsearchObjectiveEnumParamPageLikes           AdaccounttargetingsearchObjectiveEnumParam = "PAGE_LIKES"
	AdaccounttargetingsearchObjectiveEnumParamPostEngagement      AdaccounttargetingsearchObjectiveEnumParam = "POST_ENGAGEMENT"
	AdaccounttargetingsearchObjectiveEnumParamProductCatalogSales AdaccounttargetingsearchObjectiveEnumParam = "PRODUCT_CATALOG_SALES"
	AdaccounttargetingsearchObjectiveEnumParamReach               AdaccounttargetingsearchObjectiveEnumParam = "REACH"
	AdaccounttargetingsearchObjectiveEnumParamStoreVisits         AdaccounttargetingsearchObjectiveEnumParam = "STORE_VISITS"
	AdaccounttargetingsearchObjectiveEnumParamVideoViews          AdaccounttargetingsearchObjectiveEnumParam = "VIDEO_VIEWS"
)

func (value AdaccounttargetingsearchObjectiveEnumParam) String() string {
	return string(value)
}
