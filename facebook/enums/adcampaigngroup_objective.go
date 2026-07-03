package enums

type AdcampaigngroupObjective string

const (
	AdcampaigngroupObjectiveAppInstalls         AdcampaigngroupObjective = "APP_INSTALLS"
	AdcampaigngroupObjectiveBrandAwareness      AdcampaigngroupObjective = "BRAND_AWARENESS"
	AdcampaigngroupObjectiveConversions         AdcampaigngroupObjective = "CONVERSIONS"
	AdcampaigngroupObjectiveEventResponses      AdcampaigngroupObjective = "EVENT_RESPONSES"
	AdcampaigngroupObjectiveLeadGeneration      AdcampaigngroupObjective = "LEAD_GENERATION"
	AdcampaigngroupObjectiveLinkClicks          AdcampaigngroupObjective = "LINK_CLICKS"
	AdcampaigngroupObjectiveLocalAwareness      AdcampaigngroupObjective = "LOCAL_AWARENESS"
	AdcampaigngroupObjectiveMessages            AdcampaigngroupObjective = "MESSAGES"
	AdcampaigngroupObjectiveOfferClaims         AdcampaigngroupObjective = "OFFER_CLAIMS"
	AdcampaigngroupObjectiveOutcomeAppPromotion AdcampaigngroupObjective = "OUTCOME_APP_PROMOTION"
	AdcampaigngroupObjectiveOutcomeAwareness    AdcampaigngroupObjective = "OUTCOME_AWARENESS"
	AdcampaigngroupObjectiveOutcomeEngagement   AdcampaigngroupObjective = "OUTCOME_ENGAGEMENT"
	AdcampaigngroupObjectiveOutcomeLeads        AdcampaigngroupObjective = "OUTCOME_LEADS"
	AdcampaigngroupObjectiveOutcomeSales        AdcampaigngroupObjective = "OUTCOME_SALES"
	AdcampaigngroupObjectiveOutcomeTraffic      AdcampaigngroupObjective = "OUTCOME_TRAFFIC"
	AdcampaigngroupObjectivePageLikes           AdcampaigngroupObjective = "PAGE_LIKES"
	AdcampaigngroupObjectivePostEngagement      AdcampaigngroupObjective = "POST_ENGAGEMENT"
	AdcampaigngroupObjectiveProductCatalogSales AdcampaigngroupObjective = "PRODUCT_CATALOG_SALES"
	AdcampaigngroupObjectiveReach               AdcampaigngroupObjective = "REACH"
	AdcampaigngroupObjectiveStoreVisits         AdcampaigngroupObjective = "STORE_VISITS"
	AdcampaigngroupObjectiveVideoViews          AdcampaigngroupObjective = "VIDEO_VIEWS"
)

func (value AdcampaigngroupObjective) String() string {
	return string(value)
}
