package enums

type AdcampaigndeliveryEstimateOptimizationGoalEnumParam string

const (
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamAdvertiserSiloedValue              AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "ADVERTISER_SILOED_VALUE"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamAdRecallLift                       AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "AD_RECALL_LIFT"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamAppInstalls                        AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "APP_INSTALLS"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamAppInstallsAndOffsiteConversions   AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "APP_INSTALLS_AND_OFFSITE_CONVERSIONS"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamAutomaticObjective                 AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "AUTOMATIC_OBJECTIVE"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamConversations                      AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "CONVERSATIONS"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamDerivedEvents                      AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "DERIVED_EVENTS"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamEngagedPageViews                   AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "ENGAGED_PAGE_VIEWS"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamEngagedUsers                       AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "ENGAGED_USERS"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamEventResponses                     AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "EVENT_RESPONSES"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamImpressions                        AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "IMPRESSIONS"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamInAppValue                         AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "IN_APP_VALUE"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamLandingPageViews                   AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "LANDING_PAGE_VIEWS"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamLeadGeneration                     AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "LEAD_GENERATION"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamLinkClicks                         AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "LINK_CLICKS"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamMeaningfulCallAttempt              AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "MEANINGFUL_CALL_ATTEMPT"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamMessagingAppointmentConversion     AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "MESSAGING_APPOINTMENT_CONVERSION"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamMessagingDeepConversationAndFollow AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "MESSAGING_DEEP_CONVERSATION_AND_FOLLOW"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamMessagingPurchaseConversion        AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "MESSAGING_PURCHASE_CONVERSION"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamNone                               AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "NONE"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamOffsiteConversions                 AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "OFFSITE_CONVERSIONS"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamPageLikes                          AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "PAGE_LIKES"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamPostEngagement                     AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "POST_ENGAGEMENT"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamProfileAndPageEngagement           AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "PROFILE_AND_PAGE_ENGAGEMENT"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamProfileVisit                       AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "PROFILE_VISIT"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamQualityCall                        AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "QUALITY_CALL"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamQualityLead                        AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "QUALITY_LEAD"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamReach                              AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "REACH"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamRemindersSet                       AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "REMINDERS_SET"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamSubscribers                        AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "SUBSCRIBERS"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamThruplay                           AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "THRUPLAY"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamValue                              AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "VALUE"
	AdcampaigndeliveryEstimateOptimizationGoalEnumParamVisitInstagramProfile              AdcampaigndeliveryEstimateOptimizationGoalEnumParam = "VISIT_INSTAGRAM_PROFILE"
)

func (value AdcampaigndeliveryEstimateOptimizationGoalEnumParam) String() string {
	return string(value)
}
