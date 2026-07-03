package enums

type AdaccountusersettingsSydCampaignTrendsObjective string

const (
	AdaccountusersettingsSydCampaignTrendsObjectiveAppInstalls         AdaccountusersettingsSydCampaignTrendsObjective = "APP_INSTALLS"
	AdaccountusersettingsSydCampaignTrendsObjectiveBrandAwareness      AdaccountusersettingsSydCampaignTrendsObjective = "BRAND_AWARENESS"
	AdaccountusersettingsSydCampaignTrendsObjectiveEventResponses      AdaccountusersettingsSydCampaignTrendsObjective = "EVENT_RESPONSES"
	AdaccountusersettingsSydCampaignTrendsObjectiveLeadGeneration      AdaccountusersettingsSydCampaignTrendsObjective = "LEAD_GENERATION"
	AdaccountusersettingsSydCampaignTrendsObjectiveLinkClicks          AdaccountusersettingsSydCampaignTrendsObjective = "LINK_CLICKS"
	AdaccountusersettingsSydCampaignTrendsObjectiveLocalAwareness      AdaccountusersettingsSydCampaignTrendsObjective = "LOCAL_AWARENESS"
	AdaccountusersettingsSydCampaignTrendsObjectiveMessages            AdaccountusersettingsSydCampaignTrendsObjective = "MESSAGES"
	AdaccountusersettingsSydCampaignTrendsObjectiveOfferClaims         AdaccountusersettingsSydCampaignTrendsObjective = "OFFER_CLAIMS"
	AdaccountusersettingsSydCampaignTrendsObjectiveOutcomeAppPromotion AdaccountusersettingsSydCampaignTrendsObjective = "OUTCOME_APP_PROMOTION"
	AdaccountusersettingsSydCampaignTrendsObjectiveOutcomeAwareness    AdaccountusersettingsSydCampaignTrendsObjective = "OUTCOME_AWARENESS"
	AdaccountusersettingsSydCampaignTrendsObjectiveOutcomeEngagement   AdaccountusersettingsSydCampaignTrendsObjective = "OUTCOME_ENGAGEMENT"
	AdaccountusersettingsSydCampaignTrendsObjectiveOutcomeLeads        AdaccountusersettingsSydCampaignTrendsObjective = "OUTCOME_LEADS"
	AdaccountusersettingsSydCampaignTrendsObjectiveOutcomeSales        AdaccountusersettingsSydCampaignTrendsObjective = "OUTCOME_SALES"
	AdaccountusersettingsSydCampaignTrendsObjectiveOutcomeTraffic      AdaccountusersettingsSydCampaignTrendsObjective = "OUTCOME_TRAFFIC"
	AdaccountusersettingsSydCampaignTrendsObjectivePageLikes           AdaccountusersettingsSydCampaignTrendsObjective = "PAGE_LIKES"
	AdaccountusersettingsSydCampaignTrendsObjectivePostEngagement      AdaccountusersettingsSydCampaignTrendsObjective = "POST_ENGAGEMENT"
	AdaccountusersettingsSydCampaignTrendsObjectiveProductCatalogSales AdaccountusersettingsSydCampaignTrendsObjective = "PRODUCT_CATALOG_SALES"
	AdaccountusersettingsSydCampaignTrendsObjectiveReach               AdaccountusersettingsSydCampaignTrendsObjective = "REACH"
	AdaccountusersettingsSydCampaignTrendsObjectiveStoreVisits         AdaccountusersettingsSydCampaignTrendsObjective = "STORE_VISITS"
	AdaccountusersettingsSydCampaignTrendsObjectiveVideoViews          AdaccountusersettingsSydCampaignTrendsObjective = "VIDEO_VIEWS"
	AdaccountusersettingsSydCampaignTrendsObjectiveWebsiteConversions  AdaccountusersettingsSydCampaignTrendsObjective = "WEBSITE_CONVERSIONS"
)

func (value AdaccountusersettingsSydCampaignTrendsObjective) String() string {
	return string(value)
}
