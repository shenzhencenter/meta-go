package enums

type PageagenciesPermittedTasksEnumParam string

const (
	PageagenciesPermittedTasksEnumParamAdvertise                            PageagenciesPermittedTasksEnumParam = "ADVERTISE"
	PageagenciesPermittedTasksEnumParamAnalyze                              PageagenciesPermittedTasksEnumParam = "ANALYZE"
	PageagenciesPermittedTasksEnumParamCashierRole                          PageagenciesPermittedTasksEnumParam = "CASHIER_ROLE"
	PageagenciesPermittedTasksEnumParamCreateContent                        PageagenciesPermittedTasksEnumParam = "CREATE_CONTENT"
	PageagenciesPermittedTasksEnumParamGlobalStructureManagement            PageagenciesPermittedTasksEnumParam = "GLOBAL_STRUCTURE_MANAGEMENT"
	PageagenciesPermittedTasksEnumParamManage                               PageagenciesPermittedTasksEnumParam = "MANAGE"
	PageagenciesPermittedTasksEnumParamManageJobs                           PageagenciesPermittedTasksEnumParam = "MANAGE_JOBS"
	PageagenciesPermittedTasksEnumParamManageLeads                          PageagenciesPermittedTasksEnumParam = "MANAGE_LEADS"
	PageagenciesPermittedTasksEnumParamMessaging                            PageagenciesPermittedTasksEnumParam = "MESSAGING"
	PageagenciesPermittedTasksEnumParamModerate                             PageagenciesPermittedTasksEnumParam = "MODERATE"
	PageagenciesPermittedTasksEnumParamModerateCommunity                    PageagenciesPermittedTasksEnumParam = "MODERATE_COMMUNITY"
	PageagenciesPermittedTasksEnumParamPagesMessaging                       PageagenciesPermittedTasksEnumParam = "PAGES_MESSAGING"
	PageagenciesPermittedTasksEnumParamPagesMessagingSubscriptions          PageagenciesPermittedTasksEnumParam = "PAGES_MESSAGING_SUBSCRIPTIONS"
	PageagenciesPermittedTasksEnumParamProfilePlusAdvertise                 PageagenciesPermittedTasksEnumParam = "PROFILE_PLUS_ADVERTISE"
	PageagenciesPermittedTasksEnumParamProfilePlusAnalyze                   PageagenciesPermittedTasksEnumParam = "PROFILE_PLUS_ANALYZE"
	PageagenciesPermittedTasksEnumParamProfilePlusCreateContent             PageagenciesPermittedTasksEnumParam = "PROFILE_PLUS_CREATE_CONTENT"
	PageagenciesPermittedTasksEnumParamProfilePlusCreativeManagement        PageagenciesPermittedTasksEnumParam = "PROFILE_PLUS_CREATIVE_MANAGEMENT"
	PageagenciesPermittedTasksEnumParamProfilePlusCreatorManagement         PageagenciesPermittedTasksEnumParam = "PROFILE_PLUS_CREATOR_MANAGEMENT"
	PageagenciesPermittedTasksEnumParamProfilePlusFacebookAccess            PageagenciesPermittedTasksEnumParam = "PROFILE_PLUS_FACEBOOK_ACCESS"
	PageagenciesPermittedTasksEnumParamProfilePlusFullControl               PageagenciesPermittedTasksEnumParam = "PROFILE_PLUS_FULL_CONTROL"
	PageagenciesPermittedTasksEnumParamProfilePlusGlobalStructureManagement PageagenciesPermittedTasksEnumParam = "PROFILE_PLUS_GLOBAL_STRUCTURE_MANAGEMENT"
	PageagenciesPermittedTasksEnumParamProfilePlusManage                    PageagenciesPermittedTasksEnumParam = "PROFILE_PLUS_MANAGE"
	PageagenciesPermittedTasksEnumParamProfilePlusManageLeads               PageagenciesPermittedTasksEnumParam = "PROFILE_PLUS_MANAGE_LEADS"
	PageagenciesPermittedTasksEnumParamProfilePlusMessaging                 PageagenciesPermittedTasksEnumParam = "PROFILE_PLUS_MESSAGING"
	PageagenciesPermittedTasksEnumParamProfilePlusModerate                  PageagenciesPermittedTasksEnumParam = "PROFILE_PLUS_MODERATE"
	PageagenciesPermittedTasksEnumParamProfilePlusModerateDelegateCommunity PageagenciesPermittedTasksEnumParam = "PROFILE_PLUS_MODERATE_DELEGATE_COMMUNITY"
	PageagenciesPermittedTasksEnumParamProfilePlusRevenue                   PageagenciesPermittedTasksEnumParam = "PROFILE_PLUS_REVENUE"
	PageagenciesPermittedTasksEnumParamReadPageMailboxes                    PageagenciesPermittedTasksEnumParam = "READ_PAGE_MAILBOXES"
	PageagenciesPermittedTasksEnumParamViewMonetizationInsights             PageagenciesPermittedTasksEnumParam = "VIEW_MONETIZATION_INSIGHTS"
)

func (value PageagenciesPermittedTasksEnumParam) String() string {
	return string(value)
}
