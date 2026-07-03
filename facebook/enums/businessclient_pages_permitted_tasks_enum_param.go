package enums

type BusinessclientPagesPermittedTasksEnumParam string

const (
	BusinessclientPagesPermittedTasksEnumParamAdvertise                            BusinessclientPagesPermittedTasksEnumParam = "ADVERTISE"
	BusinessclientPagesPermittedTasksEnumParamAnalyze                              BusinessclientPagesPermittedTasksEnumParam = "ANALYZE"
	BusinessclientPagesPermittedTasksEnumParamCashierRole                          BusinessclientPagesPermittedTasksEnumParam = "CASHIER_ROLE"
	BusinessclientPagesPermittedTasksEnumParamCreateContent                        BusinessclientPagesPermittedTasksEnumParam = "CREATE_CONTENT"
	BusinessclientPagesPermittedTasksEnumParamGlobalStructureManagement            BusinessclientPagesPermittedTasksEnumParam = "GLOBAL_STRUCTURE_MANAGEMENT"
	BusinessclientPagesPermittedTasksEnumParamManage                               BusinessclientPagesPermittedTasksEnumParam = "MANAGE"
	BusinessclientPagesPermittedTasksEnumParamManageJobs                           BusinessclientPagesPermittedTasksEnumParam = "MANAGE_JOBS"
	BusinessclientPagesPermittedTasksEnumParamManageLeads                          BusinessclientPagesPermittedTasksEnumParam = "MANAGE_LEADS"
	BusinessclientPagesPermittedTasksEnumParamMessaging                            BusinessclientPagesPermittedTasksEnumParam = "MESSAGING"
	BusinessclientPagesPermittedTasksEnumParamModerate                             BusinessclientPagesPermittedTasksEnumParam = "MODERATE"
	BusinessclientPagesPermittedTasksEnumParamModerateCommunity                    BusinessclientPagesPermittedTasksEnumParam = "MODERATE_COMMUNITY"
	BusinessclientPagesPermittedTasksEnumParamPagesMessaging                       BusinessclientPagesPermittedTasksEnumParam = "PAGES_MESSAGING"
	BusinessclientPagesPermittedTasksEnumParamPagesMessagingSubscriptions          BusinessclientPagesPermittedTasksEnumParam = "PAGES_MESSAGING_SUBSCRIPTIONS"
	BusinessclientPagesPermittedTasksEnumParamProfilePlusAdvertise                 BusinessclientPagesPermittedTasksEnumParam = "PROFILE_PLUS_ADVERTISE"
	BusinessclientPagesPermittedTasksEnumParamProfilePlusAnalyze                   BusinessclientPagesPermittedTasksEnumParam = "PROFILE_PLUS_ANALYZE"
	BusinessclientPagesPermittedTasksEnumParamProfilePlusCreateContent             BusinessclientPagesPermittedTasksEnumParam = "PROFILE_PLUS_CREATE_CONTENT"
	BusinessclientPagesPermittedTasksEnumParamProfilePlusCreativeManagement        BusinessclientPagesPermittedTasksEnumParam = "PROFILE_PLUS_CREATIVE_MANAGEMENT"
	BusinessclientPagesPermittedTasksEnumParamProfilePlusCreatorManagement         BusinessclientPagesPermittedTasksEnumParam = "PROFILE_PLUS_CREATOR_MANAGEMENT"
	BusinessclientPagesPermittedTasksEnumParamProfilePlusFacebookAccess            BusinessclientPagesPermittedTasksEnumParam = "PROFILE_PLUS_FACEBOOK_ACCESS"
	BusinessclientPagesPermittedTasksEnumParamProfilePlusFullControl               BusinessclientPagesPermittedTasksEnumParam = "PROFILE_PLUS_FULL_CONTROL"
	BusinessclientPagesPermittedTasksEnumParamProfilePlusGlobalStructureManagement BusinessclientPagesPermittedTasksEnumParam = "PROFILE_PLUS_GLOBAL_STRUCTURE_MANAGEMENT"
	BusinessclientPagesPermittedTasksEnumParamProfilePlusManage                    BusinessclientPagesPermittedTasksEnumParam = "PROFILE_PLUS_MANAGE"
	BusinessclientPagesPermittedTasksEnumParamProfilePlusManageLeads               BusinessclientPagesPermittedTasksEnumParam = "PROFILE_PLUS_MANAGE_LEADS"
	BusinessclientPagesPermittedTasksEnumParamProfilePlusMessaging                 BusinessclientPagesPermittedTasksEnumParam = "PROFILE_PLUS_MESSAGING"
	BusinessclientPagesPermittedTasksEnumParamProfilePlusModerate                  BusinessclientPagesPermittedTasksEnumParam = "PROFILE_PLUS_MODERATE"
	BusinessclientPagesPermittedTasksEnumParamProfilePlusModerateDelegateCommunity BusinessclientPagesPermittedTasksEnumParam = "PROFILE_PLUS_MODERATE_DELEGATE_COMMUNITY"
	BusinessclientPagesPermittedTasksEnumParamProfilePlusRevenue                   BusinessclientPagesPermittedTasksEnumParam = "PROFILE_PLUS_REVENUE"
	BusinessclientPagesPermittedTasksEnumParamReadPageMailboxes                    BusinessclientPagesPermittedTasksEnumParam = "READ_PAGE_MAILBOXES"
	BusinessclientPagesPermittedTasksEnumParamViewMonetizationInsights             BusinessclientPagesPermittedTasksEnumParam = "VIEW_MONETIZATION_INSIGHTS"
)

func (value BusinessclientPagesPermittedTasksEnumParam) String() string {
	return string(value)
}
