package enums

type AdsetOptimizationGoal string

const (
	AdsetOptimizationGoalAdvertiserSiloedValue              AdsetOptimizationGoal = "ADVERTISER_SILOED_VALUE"
	AdsetOptimizationGoalAdRecallLift                       AdsetOptimizationGoal = "AD_RECALL_LIFT"
	AdsetOptimizationGoalAppInstalls                        AdsetOptimizationGoal = "APP_INSTALLS"
	AdsetOptimizationGoalAppInstallsAndOffsiteConversions   AdsetOptimizationGoal = "APP_INSTALLS_AND_OFFSITE_CONVERSIONS"
	AdsetOptimizationGoalAutomaticObjective                 AdsetOptimizationGoal = "AUTOMATIC_OBJECTIVE"
	AdsetOptimizationGoalConversations                      AdsetOptimizationGoal = "CONVERSATIONS"
	AdsetOptimizationGoalDerivedEvents                      AdsetOptimizationGoal = "DERIVED_EVENTS"
	AdsetOptimizationGoalEngagedPageViews                   AdsetOptimizationGoal = "ENGAGED_PAGE_VIEWS"
	AdsetOptimizationGoalEngagedUsers                       AdsetOptimizationGoal = "ENGAGED_USERS"
	AdsetOptimizationGoalEventResponses                     AdsetOptimizationGoal = "EVENT_RESPONSES"
	AdsetOptimizationGoalImpressions                        AdsetOptimizationGoal = "IMPRESSIONS"
	AdsetOptimizationGoalInAppValue                         AdsetOptimizationGoal = "IN_APP_VALUE"
	AdsetOptimizationGoalLandingPageViews                   AdsetOptimizationGoal = "LANDING_PAGE_VIEWS"
	AdsetOptimizationGoalLeadGeneration                     AdsetOptimizationGoal = "LEAD_GENERATION"
	AdsetOptimizationGoalLinkClicks                         AdsetOptimizationGoal = "LINK_CLICKS"
	AdsetOptimizationGoalMeaningfulCallAttempt              AdsetOptimizationGoal = "MEANINGFUL_CALL_ATTEMPT"
	AdsetOptimizationGoalMessagingAppointmentConversion     AdsetOptimizationGoal = "MESSAGING_APPOINTMENT_CONVERSION"
	AdsetOptimizationGoalMessagingDeepConversationAndFollow AdsetOptimizationGoal = "MESSAGING_DEEP_CONVERSATION_AND_FOLLOW"
	AdsetOptimizationGoalMessagingPurchaseConversion        AdsetOptimizationGoal = "MESSAGING_PURCHASE_CONVERSION"
	AdsetOptimizationGoalNone                               AdsetOptimizationGoal = "NONE"
	AdsetOptimizationGoalOffsiteConversions                 AdsetOptimizationGoal = "OFFSITE_CONVERSIONS"
	AdsetOptimizationGoalPageLikes                          AdsetOptimizationGoal = "PAGE_LIKES"
	AdsetOptimizationGoalPostEngagement                     AdsetOptimizationGoal = "POST_ENGAGEMENT"
	AdsetOptimizationGoalProfileAndPageEngagement           AdsetOptimizationGoal = "PROFILE_AND_PAGE_ENGAGEMENT"
	AdsetOptimizationGoalProfileVisit                       AdsetOptimizationGoal = "PROFILE_VISIT"
	AdsetOptimizationGoalQualityCall                        AdsetOptimizationGoal = "QUALITY_CALL"
	AdsetOptimizationGoalQualityLead                        AdsetOptimizationGoal = "QUALITY_LEAD"
	AdsetOptimizationGoalReach                              AdsetOptimizationGoal = "REACH"
	AdsetOptimizationGoalRemindersSet                       AdsetOptimizationGoal = "REMINDERS_SET"
	AdsetOptimizationGoalSubscribers                        AdsetOptimizationGoal = "SUBSCRIBERS"
	AdsetOptimizationGoalThruplay                           AdsetOptimizationGoal = "THRUPLAY"
	AdsetOptimizationGoalValue                              AdsetOptimizationGoal = "VALUE"
	AdsetOptimizationGoalVisitInstagramProfile              AdsetOptimizationGoal = "VISIT_INSTAGRAM_PROFILE"
)

func (value AdsetOptimizationGoal) String() string {
	return string(value)
}
