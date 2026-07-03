package enums

type AdaccountcampaignsObjectiveEnumParam string

const (
	AdaccountcampaignsObjectiveEnumParamAppInstalls         AdaccountcampaignsObjectiveEnumParam = "APP_INSTALLS"
	AdaccountcampaignsObjectiveEnumParamBrandAwareness      AdaccountcampaignsObjectiveEnumParam = "BRAND_AWARENESS"
	AdaccountcampaignsObjectiveEnumParamConversions         AdaccountcampaignsObjectiveEnumParam = "CONVERSIONS"
	AdaccountcampaignsObjectiveEnumParamEventResponses      AdaccountcampaignsObjectiveEnumParam = "EVENT_RESPONSES"
	AdaccountcampaignsObjectiveEnumParamLeadGeneration      AdaccountcampaignsObjectiveEnumParam = "LEAD_GENERATION"
	AdaccountcampaignsObjectiveEnumParamLinkClicks          AdaccountcampaignsObjectiveEnumParam = "LINK_CLICKS"
	AdaccountcampaignsObjectiveEnumParamLocalAwareness      AdaccountcampaignsObjectiveEnumParam = "LOCAL_AWARENESS"
	AdaccountcampaignsObjectiveEnumParamMessages            AdaccountcampaignsObjectiveEnumParam = "MESSAGES"
	AdaccountcampaignsObjectiveEnumParamOfferClaims         AdaccountcampaignsObjectiveEnumParam = "OFFER_CLAIMS"
	AdaccountcampaignsObjectiveEnumParamOutcomeAppPromotion AdaccountcampaignsObjectiveEnumParam = "OUTCOME_APP_PROMOTION"
	AdaccountcampaignsObjectiveEnumParamOutcomeAwareness    AdaccountcampaignsObjectiveEnumParam = "OUTCOME_AWARENESS"
	AdaccountcampaignsObjectiveEnumParamOutcomeEngagement   AdaccountcampaignsObjectiveEnumParam = "OUTCOME_ENGAGEMENT"
	AdaccountcampaignsObjectiveEnumParamOutcomeLeads        AdaccountcampaignsObjectiveEnumParam = "OUTCOME_LEADS"
	AdaccountcampaignsObjectiveEnumParamOutcomeSales        AdaccountcampaignsObjectiveEnumParam = "OUTCOME_SALES"
	AdaccountcampaignsObjectiveEnumParamOutcomeTraffic      AdaccountcampaignsObjectiveEnumParam = "OUTCOME_TRAFFIC"
	AdaccountcampaignsObjectiveEnumParamPageLikes           AdaccountcampaignsObjectiveEnumParam = "PAGE_LIKES"
	AdaccountcampaignsObjectiveEnumParamPostEngagement      AdaccountcampaignsObjectiveEnumParam = "POST_ENGAGEMENT"
	AdaccountcampaignsObjectiveEnumParamProductCatalogSales AdaccountcampaignsObjectiveEnumParam = "PRODUCT_CATALOG_SALES"
	AdaccountcampaignsObjectiveEnumParamReach               AdaccountcampaignsObjectiveEnumParam = "REACH"
	AdaccountcampaignsObjectiveEnumParamStoreVisits         AdaccountcampaignsObjectiveEnumParam = "STORE_VISITS"
	AdaccountcampaignsObjectiveEnumParamVideoViews          AdaccountcampaignsObjectiveEnumParam = "VIDEO_VIEWS"
)

func (value AdaccountcampaignsObjectiveEnumParam) String() string {
	return string(value)
}
