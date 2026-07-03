package enums

type AdrulehistoryActionEnumParam string

const (
	AdrulehistoryActionEnumParamBudgetNotRedistributed               AdrulehistoryActionEnumParam = "BUDGET_NOT_REDISTRIBUTED"
	AdrulehistoryActionEnumParamChangedBid                           AdrulehistoryActionEnumParam = "CHANGED_BID"
	AdrulehistoryActionEnumParamChangedBudget                        AdrulehistoryActionEnumParam = "CHANGED_BUDGET"
	AdrulehistoryActionEnumParamConsolidateAscFragmentation          AdrulehistoryActionEnumParam = "CONSOLIDATE_ASC_FRAGMENTATION"
	AdrulehistoryActionEnumParamConsolidateFragmentation             AdrulehistoryActionEnumParam = "CONSOLIDATE_FRAGMENTATION"
	AdrulehistoryActionEnumParamConvertAscCpSingleInstance           AdrulehistoryActionEnumParam = "CONVERT_ASC_CP_SINGLE_INSTANCE"
	AdrulehistoryActionEnumParamEmail                                AdrulehistoryActionEnumParam = "EMAIL"
	AdrulehistoryActionEnumParamEnableAdvantageCampaignBudget        AdrulehistoryActionEnumParam = "ENABLE_ADVANTAGE_CAMPAIGN_BUDGET"
	AdrulehistoryActionEnumParamEnableAdvantagePlusAudience          AdrulehistoryActionEnumParam = "ENABLE_ADVANTAGE_PLUS_AUDIENCE"
	AdrulehistoryActionEnumParamEnableAdvantagePlusCreative          AdrulehistoryActionEnumParam = "ENABLE_ADVANTAGE_PLUS_CREATIVE"
	AdrulehistoryActionEnumParamEnableAdvantagePlusPlacements        AdrulehistoryActionEnumParam = "ENABLE_ADVANTAGE_PLUS_PLACEMENTS"
	AdrulehistoryActionEnumParamEnableAutoflow                       AdrulehistoryActionEnumParam = "ENABLE_AUTOFLOW"
	AdrulehistoryActionEnumParamEnableGenUncrop                      AdrulehistoryActionEnumParam = "ENABLE_GEN_UNCROP"
	AdrulehistoryActionEnumParamEnableLandingPageViews               AdrulehistoryActionEnumParam = "ENABLE_LANDING_PAGE_VIEWS"
	AdrulehistoryActionEnumParamEnableMusic                          AdrulehistoryActionEnumParam = "ENABLE_MUSIC"
	AdrulehistoryActionEnumParamEnablePixellessLpvOptimizationGoal   AdrulehistoryActionEnumParam = "ENABLE_PIXELLESS_LPV_OPTIMIZATION_GOAL"
	AdrulehistoryActionEnumParamEnableProductSetBoosting             AdrulehistoryActionEnumParam = "ENABLE_PRODUCT_SET_BOOSTING"
	AdrulehistoryActionEnumParamEnableReelsPlacements                AdrulehistoryActionEnumParam = "ENABLE_REELS_PLACEMENTS"
	AdrulehistoryActionEnumParamEnableSemanticBasedAudienceExpansion AdrulehistoryActionEnumParam = "ENABLE_SEMANTIC_BASED_AUDIENCE_EXPANSION"
	AdrulehistoryActionEnumParamEnableShopsAds                       AdrulehistoryActionEnumParam = "ENABLE_SHOPS_ADS"
	AdrulehistoryActionEnumParamEnableShopsAdsSaoff                  AdrulehistoryActionEnumParam = "ENABLE_SHOPS_ADS_SAOFF"
	AdrulehistoryActionEnumParamEnableWtwaUpsellInDuplication        AdrulehistoryActionEnumParam = "ENABLE_WTWA_UPSELL_IN_DUPLICATION"
	AdrulehistoryActionEnumParamEndpointPinged                       AdrulehistoryActionEnumParam = "ENDPOINT_PINGED"
	AdrulehistoryActionEnumParamError                                AdrulehistoryActionEnumParam = "ERROR"
	AdrulehistoryActionEnumParamFacebookNotificationSent             AdrulehistoryActionEnumParam = "FACEBOOK_NOTIFICATION_SENT"
	AdrulehistoryActionEnumParamMessageSent                          AdrulehistoryActionEnumParam = "MESSAGE_SENT"
	AdrulehistoryActionEnumParamNotChanged                           AdrulehistoryActionEnumParam = "NOT_CHANGED"
	AdrulehistoryActionEnumParamPaused                               AdrulehistoryActionEnumParam = "PAUSED"
	AdrulehistoryActionEnumParamUnpaused                             AdrulehistoryActionEnumParam = "UNPAUSED"
)

func (value AdrulehistoryActionEnumParam) String() string {
	return string(value)
}
