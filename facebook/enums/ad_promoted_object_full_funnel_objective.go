package enums

type AdpromotedobjectFullFunnelObjective string

const (
	AdpromotedobjectFullFunnelObjectiveAppInstalls         AdpromotedobjectFullFunnelObjective = "APP_INSTALLS"
	AdpromotedobjectFullFunnelObjectiveBrandAwareness      AdpromotedobjectFullFunnelObjective = "BRAND_AWARENESS"
	AdpromotedobjectFullFunnelObjectiveEventResponses      AdpromotedobjectFullFunnelObjective = "EVENT_RESPONSES"
	AdpromotedobjectFullFunnelObjectiveLeadGeneration      AdpromotedobjectFullFunnelObjective = "LEAD_GENERATION"
	AdpromotedobjectFullFunnelObjectiveLinkClicks          AdpromotedobjectFullFunnelObjective = "LINK_CLICKS"
	AdpromotedobjectFullFunnelObjectiveLocalAwareness      AdpromotedobjectFullFunnelObjective = "LOCAL_AWARENESS"
	AdpromotedobjectFullFunnelObjectiveMessages            AdpromotedobjectFullFunnelObjective = "MESSAGES"
	AdpromotedobjectFullFunnelObjectiveOfferClaims         AdpromotedobjectFullFunnelObjective = "OFFER_CLAIMS"
	AdpromotedobjectFullFunnelObjectiveOutcomeAppPromotion AdpromotedobjectFullFunnelObjective = "OUTCOME_APP_PROMOTION"
	AdpromotedobjectFullFunnelObjectiveOutcomeAwareness    AdpromotedobjectFullFunnelObjective = "OUTCOME_AWARENESS"
	AdpromotedobjectFullFunnelObjectiveOutcomeEngagement   AdpromotedobjectFullFunnelObjective = "OUTCOME_ENGAGEMENT"
	AdpromotedobjectFullFunnelObjectiveOutcomeLeads        AdpromotedobjectFullFunnelObjective = "OUTCOME_LEADS"
	AdpromotedobjectFullFunnelObjectiveOutcomeSales        AdpromotedobjectFullFunnelObjective = "OUTCOME_SALES"
	AdpromotedobjectFullFunnelObjectiveOutcomeTraffic      AdpromotedobjectFullFunnelObjective = "OUTCOME_TRAFFIC"
	AdpromotedobjectFullFunnelObjectivePageLikes           AdpromotedobjectFullFunnelObjective = "PAGE_LIKES"
	AdpromotedobjectFullFunnelObjectivePostEngagement      AdpromotedobjectFullFunnelObjective = "POST_ENGAGEMENT"
	AdpromotedobjectFullFunnelObjectiveProductCatalogSales AdpromotedobjectFullFunnelObjective = "PRODUCT_CATALOG_SALES"
	AdpromotedobjectFullFunnelObjectiveReach               AdpromotedobjectFullFunnelObjective = "REACH"
	AdpromotedobjectFullFunnelObjectiveStoreVisits         AdpromotedobjectFullFunnelObjective = "STORE_VISITS"
	AdpromotedobjectFullFunnelObjectiveVideoViews          AdpromotedobjectFullFunnelObjective = "VIDEO_VIEWS"
	AdpromotedobjectFullFunnelObjectiveWebsiteConversions  AdpromotedobjectFullFunnelObjective = "WEBSITE_CONVERSIONS"
)

func (value AdpromotedobjectFullFunnelObjective) String() string {
	return string(value)
}
