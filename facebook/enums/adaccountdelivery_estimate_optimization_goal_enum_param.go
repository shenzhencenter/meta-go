package enums

type AdaccountdeliveryEstimateOptimizationGoalEnumParam string

const (
	AdaccountdeliveryEstimateOptimizationGoalEnumParamAdvertiserSiloedValue              AdaccountdeliveryEstimateOptimizationGoalEnumParam = "ADVERTISER_SILOED_VALUE"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamAdRecallLift                       AdaccountdeliveryEstimateOptimizationGoalEnumParam = "AD_RECALL_LIFT"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamAppInstalls                        AdaccountdeliveryEstimateOptimizationGoalEnumParam = "APP_INSTALLS"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamAppInstallsAndOffsiteConversions   AdaccountdeliveryEstimateOptimizationGoalEnumParam = "APP_INSTALLS_AND_OFFSITE_CONVERSIONS"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamAutomaticObjective                 AdaccountdeliveryEstimateOptimizationGoalEnumParam = "AUTOMATIC_OBJECTIVE"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamConversations                      AdaccountdeliveryEstimateOptimizationGoalEnumParam = "CONVERSATIONS"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamDerivedEvents                      AdaccountdeliveryEstimateOptimizationGoalEnumParam = "DERIVED_EVENTS"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamEngagedPageViews                   AdaccountdeliveryEstimateOptimizationGoalEnumParam = "ENGAGED_PAGE_VIEWS"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamEngagedUsers                       AdaccountdeliveryEstimateOptimizationGoalEnumParam = "ENGAGED_USERS"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamEventResponses                     AdaccountdeliveryEstimateOptimizationGoalEnumParam = "EVENT_RESPONSES"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamImpressions                        AdaccountdeliveryEstimateOptimizationGoalEnumParam = "IMPRESSIONS"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamInAppValue                         AdaccountdeliveryEstimateOptimizationGoalEnumParam = "IN_APP_VALUE"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamLandingPageViews                   AdaccountdeliveryEstimateOptimizationGoalEnumParam = "LANDING_PAGE_VIEWS"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamLeadGeneration                     AdaccountdeliveryEstimateOptimizationGoalEnumParam = "LEAD_GENERATION"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamLinkClicks                         AdaccountdeliveryEstimateOptimizationGoalEnumParam = "LINK_CLICKS"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamMeaningfulCallAttempt              AdaccountdeliveryEstimateOptimizationGoalEnumParam = "MEANINGFUL_CALL_ATTEMPT"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamMessagingAppointmentConversion     AdaccountdeliveryEstimateOptimizationGoalEnumParam = "MESSAGING_APPOINTMENT_CONVERSION"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamMessagingDeepConversationAndFollow AdaccountdeliveryEstimateOptimizationGoalEnumParam = "MESSAGING_DEEP_CONVERSATION_AND_FOLLOW"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamMessagingPurchaseConversion        AdaccountdeliveryEstimateOptimizationGoalEnumParam = "MESSAGING_PURCHASE_CONVERSION"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamNone                               AdaccountdeliveryEstimateOptimizationGoalEnumParam = "NONE"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamOffsiteConversions                 AdaccountdeliveryEstimateOptimizationGoalEnumParam = "OFFSITE_CONVERSIONS"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamPageLikes                          AdaccountdeliveryEstimateOptimizationGoalEnumParam = "PAGE_LIKES"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamPostEngagement                     AdaccountdeliveryEstimateOptimizationGoalEnumParam = "POST_ENGAGEMENT"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamProfileAndPageEngagement           AdaccountdeliveryEstimateOptimizationGoalEnumParam = "PROFILE_AND_PAGE_ENGAGEMENT"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamProfileVisit                       AdaccountdeliveryEstimateOptimizationGoalEnumParam = "PROFILE_VISIT"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamQualityCall                        AdaccountdeliveryEstimateOptimizationGoalEnumParam = "QUALITY_CALL"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamQualityLead                        AdaccountdeliveryEstimateOptimizationGoalEnumParam = "QUALITY_LEAD"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamReach                              AdaccountdeliveryEstimateOptimizationGoalEnumParam = "REACH"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamRemindersSet                       AdaccountdeliveryEstimateOptimizationGoalEnumParam = "REMINDERS_SET"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamSubscribers                        AdaccountdeliveryEstimateOptimizationGoalEnumParam = "SUBSCRIBERS"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamThruplay                           AdaccountdeliveryEstimateOptimizationGoalEnumParam = "THRUPLAY"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamValue                              AdaccountdeliveryEstimateOptimizationGoalEnumParam = "VALUE"
	AdaccountdeliveryEstimateOptimizationGoalEnumParamVisitInstagramProfile              AdaccountdeliveryEstimateOptimizationGoalEnumParam = "VISIT_INSTAGRAM_PROFILE"
)

func (value AdaccountdeliveryEstimateOptimizationGoalEnumParam) String() string {
	return string(value)
}
