package enums

type AdaccountadrulesHistoryActionEnumParam string

const (
	AdaccountadrulesHistoryActionEnumParamBudgetNotRedistributed               AdaccountadrulesHistoryActionEnumParam = "BUDGET_NOT_REDISTRIBUTED"
	AdaccountadrulesHistoryActionEnumParamChangedBid                           AdaccountadrulesHistoryActionEnumParam = "CHANGED_BID"
	AdaccountadrulesHistoryActionEnumParamChangedBudget                        AdaccountadrulesHistoryActionEnumParam = "CHANGED_BUDGET"
	AdaccountadrulesHistoryActionEnumParamConsolidateAscFragmentation          AdaccountadrulesHistoryActionEnumParam = "CONSOLIDATE_ASC_FRAGMENTATION"
	AdaccountadrulesHistoryActionEnumParamConsolidateFragmentation             AdaccountadrulesHistoryActionEnumParam = "CONSOLIDATE_FRAGMENTATION"
	AdaccountadrulesHistoryActionEnumParamConvertAscCpSingleInstance           AdaccountadrulesHistoryActionEnumParam = "CONVERT_ASC_CP_SINGLE_INSTANCE"
	AdaccountadrulesHistoryActionEnumParamEmail                                AdaccountadrulesHistoryActionEnumParam = "EMAIL"
	AdaccountadrulesHistoryActionEnumParamEnableAdvantageCampaignBudget        AdaccountadrulesHistoryActionEnumParam = "ENABLE_ADVANTAGE_CAMPAIGN_BUDGET"
	AdaccountadrulesHistoryActionEnumParamEnableAdvantagePlusAudience          AdaccountadrulesHistoryActionEnumParam = "ENABLE_ADVANTAGE_PLUS_AUDIENCE"
	AdaccountadrulesHistoryActionEnumParamEnableAdvantagePlusCreative          AdaccountadrulesHistoryActionEnumParam = "ENABLE_ADVANTAGE_PLUS_CREATIVE"
	AdaccountadrulesHistoryActionEnumParamEnableAdvantagePlusPlacements        AdaccountadrulesHistoryActionEnumParam = "ENABLE_ADVANTAGE_PLUS_PLACEMENTS"
	AdaccountadrulesHistoryActionEnumParamEnableAutoflow                       AdaccountadrulesHistoryActionEnumParam = "ENABLE_AUTOFLOW"
	AdaccountadrulesHistoryActionEnumParamEnableGenUncrop                      AdaccountadrulesHistoryActionEnumParam = "ENABLE_GEN_UNCROP"
	AdaccountadrulesHistoryActionEnumParamEnableLandingPageViews               AdaccountadrulesHistoryActionEnumParam = "ENABLE_LANDING_PAGE_VIEWS"
	AdaccountadrulesHistoryActionEnumParamEnableMusic                          AdaccountadrulesHistoryActionEnumParam = "ENABLE_MUSIC"
	AdaccountadrulesHistoryActionEnumParamEnablePixellessLpvOptimizationGoal   AdaccountadrulesHistoryActionEnumParam = "ENABLE_PIXELLESS_LPV_OPTIMIZATION_GOAL"
	AdaccountadrulesHistoryActionEnumParamEnableProductSetBoosting             AdaccountadrulesHistoryActionEnumParam = "ENABLE_PRODUCT_SET_BOOSTING"
	AdaccountadrulesHistoryActionEnumParamEnableReelsPlacements                AdaccountadrulesHistoryActionEnumParam = "ENABLE_REELS_PLACEMENTS"
	AdaccountadrulesHistoryActionEnumParamEnableSemanticBasedAudienceExpansion AdaccountadrulesHistoryActionEnumParam = "ENABLE_SEMANTIC_BASED_AUDIENCE_EXPANSION"
	AdaccountadrulesHistoryActionEnumParamEnableShopsAds                       AdaccountadrulesHistoryActionEnumParam = "ENABLE_SHOPS_ADS"
	AdaccountadrulesHistoryActionEnumParamEnableShopsAdsSaoff                  AdaccountadrulesHistoryActionEnumParam = "ENABLE_SHOPS_ADS_SAOFF"
	AdaccountadrulesHistoryActionEnumParamEnableWtwaUpsellInDuplication        AdaccountadrulesHistoryActionEnumParam = "ENABLE_WTWA_UPSELL_IN_DUPLICATION"
	AdaccountadrulesHistoryActionEnumParamEndpointPinged                       AdaccountadrulesHistoryActionEnumParam = "ENDPOINT_PINGED"
	AdaccountadrulesHistoryActionEnumParamError                                AdaccountadrulesHistoryActionEnumParam = "ERROR"
	AdaccountadrulesHistoryActionEnumParamFacebookNotificationSent             AdaccountadrulesHistoryActionEnumParam = "FACEBOOK_NOTIFICATION_SENT"
	AdaccountadrulesHistoryActionEnumParamMessageSent                          AdaccountadrulesHistoryActionEnumParam = "MESSAGE_SENT"
	AdaccountadrulesHistoryActionEnumParamNotChanged                           AdaccountadrulesHistoryActionEnumParam = "NOT_CHANGED"
	AdaccountadrulesHistoryActionEnumParamPaused                               AdaccountadrulesHistoryActionEnumParam = "PAUSED"
	AdaccountadrulesHistoryActionEnumParamUnpaused                             AdaccountadrulesHistoryActionEnumParam = "UNPAUSED"
)

func (value AdaccountadrulesHistoryActionEnumParam) String() string {
	return string(value)
}
