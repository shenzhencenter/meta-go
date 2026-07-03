package enums

type AdcampaignOptimizationGoal string

const (
	AdcampaignOptimizationGoalAdvertiserSiloedValue              AdcampaignOptimizationGoal = "ADVERTISER_SILOED_VALUE"
	AdcampaignOptimizationGoalAdRecallLift                       AdcampaignOptimizationGoal = "AD_RECALL_LIFT"
	AdcampaignOptimizationGoalAppInstalls                        AdcampaignOptimizationGoal = "APP_INSTALLS"
	AdcampaignOptimizationGoalAppInstallsAndOffsiteConversions   AdcampaignOptimizationGoal = "APP_INSTALLS_AND_OFFSITE_CONVERSIONS"
	AdcampaignOptimizationGoalAutomaticObjective                 AdcampaignOptimizationGoal = "AUTOMATIC_OBJECTIVE"
	AdcampaignOptimizationGoalConversations                      AdcampaignOptimizationGoal = "CONVERSATIONS"
	AdcampaignOptimizationGoalDerivedEvents                      AdcampaignOptimizationGoal = "DERIVED_EVENTS"
	AdcampaignOptimizationGoalEngagedPageViews                   AdcampaignOptimizationGoal = "ENGAGED_PAGE_VIEWS"
	AdcampaignOptimizationGoalEngagedUsers                       AdcampaignOptimizationGoal = "ENGAGED_USERS"
	AdcampaignOptimizationGoalEventResponses                     AdcampaignOptimizationGoal = "EVENT_RESPONSES"
	AdcampaignOptimizationGoalImpressions                        AdcampaignOptimizationGoal = "IMPRESSIONS"
	AdcampaignOptimizationGoalInAppValue                         AdcampaignOptimizationGoal = "IN_APP_VALUE"
	AdcampaignOptimizationGoalLandingPageViews                   AdcampaignOptimizationGoal = "LANDING_PAGE_VIEWS"
	AdcampaignOptimizationGoalLeadGeneration                     AdcampaignOptimizationGoal = "LEAD_GENERATION"
	AdcampaignOptimizationGoalLinkClicks                         AdcampaignOptimizationGoal = "LINK_CLICKS"
	AdcampaignOptimizationGoalMeaningfulCallAttempt              AdcampaignOptimizationGoal = "MEANINGFUL_CALL_ATTEMPT"
	AdcampaignOptimizationGoalMessagingAppointmentConversion     AdcampaignOptimizationGoal = "MESSAGING_APPOINTMENT_CONVERSION"
	AdcampaignOptimizationGoalMessagingDeepConversationAndFollow AdcampaignOptimizationGoal = "MESSAGING_DEEP_CONVERSATION_AND_FOLLOW"
	AdcampaignOptimizationGoalMessagingPurchaseConversion        AdcampaignOptimizationGoal = "MESSAGING_PURCHASE_CONVERSION"
	AdcampaignOptimizationGoalNone                               AdcampaignOptimizationGoal = "NONE"
	AdcampaignOptimizationGoalOffsiteConversions                 AdcampaignOptimizationGoal = "OFFSITE_CONVERSIONS"
	AdcampaignOptimizationGoalPageLikes                          AdcampaignOptimizationGoal = "PAGE_LIKES"
	AdcampaignOptimizationGoalPostEngagement                     AdcampaignOptimizationGoal = "POST_ENGAGEMENT"
	AdcampaignOptimizationGoalProfileAndPageEngagement           AdcampaignOptimizationGoal = "PROFILE_AND_PAGE_ENGAGEMENT"
	AdcampaignOptimizationGoalProfileVisit                       AdcampaignOptimizationGoal = "PROFILE_VISIT"
	AdcampaignOptimizationGoalQualityCall                        AdcampaignOptimizationGoal = "QUALITY_CALL"
	AdcampaignOptimizationGoalQualityLead                        AdcampaignOptimizationGoal = "QUALITY_LEAD"
	AdcampaignOptimizationGoalReach                              AdcampaignOptimizationGoal = "REACH"
	AdcampaignOptimizationGoalRemindersSet                       AdcampaignOptimizationGoal = "REMINDERS_SET"
	AdcampaignOptimizationGoalSubscribers                        AdcampaignOptimizationGoal = "SUBSCRIBERS"
	AdcampaignOptimizationGoalThruplay                           AdcampaignOptimizationGoal = "THRUPLAY"
	AdcampaignOptimizationGoalValue                              AdcampaignOptimizationGoal = "VALUE"
	AdcampaignOptimizationGoalVisitInstagramProfile              AdcampaignOptimizationGoal = "VISIT_INSTAGRAM_PROFILE"
)

func (value AdcampaignOptimizationGoal) String() string {
	return string(value)
}
