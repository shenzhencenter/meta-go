package enums

type AdaccountadsetsOptimizationGoalEnumParam string

const (
	AdaccountadsetsOptimizationGoalEnumParamAdvertiserSiloedValue              AdaccountadsetsOptimizationGoalEnumParam = "ADVERTISER_SILOED_VALUE"
	AdaccountadsetsOptimizationGoalEnumParamAdRecallLift                       AdaccountadsetsOptimizationGoalEnumParam = "AD_RECALL_LIFT"
	AdaccountadsetsOptimizationGoalEnumParamAppInstalls                        AdaccountadsetsOptimizationGoalEnumParam = "APP_INSTALLS"
	AdaccountadsetsOptimizationGoalEnumParamAppInstallsAndOffsiteConversions   AdaccountadsetsOptimizationGoalEnumParam = "APP_INSTALLS_AND_OFFSITE_CONVERSIONS"
	AdaccountadsetsOptimizationGoalEnumParamAutomaticObjective                 AdaccountadsetsOptimizationGoalEnumParam = "AUTOMATIC_OBJECTIVE"
	AdaccountadsetsOptimizationGoalEnumParamConversations                      AdaccountadsetsOptimizationGoalEnumParam = "CONVERSATIONS"
	AdaccountadsetsOptimizationGoalEnumParamDerivedEvents                      AdaccountadsetsOptimizationGoalEnumParam = "DERIVED_EVENTS"
	AdaccountadsetsOptimizationGoalEnumParamEngagedPageViews                   AdaccountadsetsOptimizationGoalEnumParam = "ENGAGED_PAGE_VIEWS"
	AdaccountadsetsOptimizationGoalEnumParamEngagedUsers                       AdaccountadsetsOptimizationGoalEnumParam = "ENGAGED_USERS"
	AdaccountadsetsOptimizationGoalEnumParamEventResponses                     AdaccountadsetsOptimizationGoalEnumParam = "EVENT_RESPONSES"
	AdaccountadsetsOptimizationGoalEnumParamImpressions                        AdaccountadsetsOptimizationGoalEnumParam = "IMPRESSIONS"
	AdaccountadsetsOptimizationGoalEnumParamInAppValue                         AdaccountadsetsOptimizationGoalEnumParam = "IN_APP_VALUE"
	AdaccountadsetsOptimizationGoalEnumParamLandingPageViews                   AdaccountadsetsOptimizationGoalEnumParam = "LANDING_PAGE_VIEWS"
	AdaccountadsetsOptimizationGoalEnumParamLeadGeneration                     AdaccountadsetsOptimizationGoalEnumParam = "LEAD_GENERATION"
	AdaccountadsetsOptimizationGoalEnumParamLinkClicks                         AdaccountadsetsOptimizationGoalEnumParam = "LINK_CLICKS"
	AdaccountadsetsOptimizationGoalEnumParamMeaningfulCallAttempt              AdaccountadsetsOptimizationGoalEnumParam = "MEANINGFUL_CALL_ATTEMPT"
	AdaccountadsetsOptimizationGoalEnumParamMessagingAppointmentConversion     AdaccountadsetsOptimizationGoalEnumParam = "MESSAGING_APPOINTMENT_CONVERSION"
	AdaccountadsetsOptimizationGoalEnumParamMessagingDeepConversationAndFollow AdaccountadsetsOptimizationGoalEnumParam = "MESSAGING_DEEP_CONVERSATION_AND_FOLLOW"
	AdaccountadsetsOptimizationGoalEnumParamMessagingPurchaseConversion        AdaccountadsetsOptimizationGoalEnumParam = "MESSAGING_PURCHASE_CONVERSION"
	AdaccountadsetsOptimizationGoalEnumParamNone                               AdaccountadsetsOptimizationGoalEnumParam = "NONE"
	AdaccountadsetsOptimizationGoalEnumParamOffsiteConversions                 AdaccountadsetsOptimizationGoalEnumParam = "OFFSITE_CONVERSIONS"
	AdaccountadsetsOptimizationGoalEnumParamPageLikes                          AdaccountadsetsOptimizationGoalEnumParam = "PAGE_LIKES"
	AdaccountadsetsOptimizationGoalEnumParamPostEngagement                     AdaccountadsetsOptimizationGoalEnumParam = "POST_ENGAGEMENT"
	AdaccountadsetsOptimizationGoalEnumParamProfileAndPageEngagement           AdaccountadsetsOptimizationGoalEnumParam = "PROFILE_AND_PAGE_ENGAGEMENT"
	AdaccountadsetsOptimizationGoalEnumParamProfileVisit                       AdaccountadsetsOptimizationGoalEnumParam = "PROFILE_VISIT"
	AdaccountadsetsOptimizationGoalEnumParamQualityCall                        AdaccountadsetsOptimizationGoalEnumParam = "QUALITY_CALL"
	AdaccountadsetsOptimizationGoalEnumParamQualityLead                        AdaccountadsetsOptimizationGoalEnumParam = "QUALITY_LEAD"
	AdaccountadsetsOptimizationGoalEnumParamReach                              AdaccountadsetsOptimizationGoalEnumParam = "REACH"
	AdaccountadsetsOptimizationGoalEnumParamRemindersSet                       AdaccountadsetsOptimizationGoalEnumParam = "REMINDERS_SET"
	AdaccountadsetsOptimizationGoalEnumParamSubscribers                        AdaccountadsetsOptimizationGoalEnumParam = "SUBSCRIBERS"
	AdaccountadsetsOptimizationGoalEnumParamThruplay                           AdaccountadsetsOptimizationGoalEnumParam = "THRUPLAY"
	AdaccountadsetsOptimizationGoalEnumParamValue                              AdaccountadsetsOptimizationGoalEnumParam = "VALUE"
	AdaccountadsetsOptimizationGoalEnumParamVisitInstagramProfile              AdaccountadsetsOptimizationGoalEnumParam = "VISIT_INSTAGRAM_PROFILE"
)

func (value AdaccountadsetsOptimizationGoalEnumParam) String() string {
	return string(value)
}
