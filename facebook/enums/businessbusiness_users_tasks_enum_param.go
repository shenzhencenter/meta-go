package enums

type BusinessbusinessUsersTasksEnumParam string

const (
	BusinessbusinessUsersTasksEnumParamAdmin                   BusinessbusinessUsersTasksEnumParam = "ADMIN"
	BusinessbusinessUsersTasksEnumParamAdsRightsReviewer       BusinessbusinessUsersTasksEnumParam = "ADS_RIGHTS_REVIEWER"
	BusinessbusinessUsersTasksEnumParamDefault                 BusinessbusinessUsersTasksEnumParam = "DEFAULT"
	BusinessbusinessUsersTasksEnumParamDeveloper               BusinessbusinessUsersTasksEnumParam = "DEVELOPER"
	BusinessbusinessUsersTasksEnumParamEmployee                BusinessbusinessUsersTasksEnumParam = "EMPLOYEE"
	BusinessbusinessUsersTasksEnumParamFinanceAnalyst          BusinessbusinessUsersTasksEnumParam = "FINANCE_ANALYST"
	BusinessbusinessUsersTasksEnumParamFinanceEdit             BusinessbusinessUsersTasksEnumParam = "FINANCE_EDIT"
	BusinessbusinessUsersTasksEnumParamFinanceEditor           BusinessbusinessUsersTasksEnumParam = "FINANCE_EDITOR"
	BusinessbusinessUsersTasksEnumParamFinanceView             BusinessbusinessUsersTasksEnumParam = "FINANCE_VIEW"
	BusinessbusinessUsersTasksEnumParamManage                  BusinessbusinessUsersTasksEnumParam = "MANAGE"
	BusinessbusinessUsersTasksEnumParamPartnerCenterAdmin      BusinessbusinessUsersTasksEnumParam = "PARTNER_CENTER_ADMIN"
	BusinessbusinessUsersTasksEnumParamPartnerCenterAnalyst    BusinessbusinessUsersTasksEnumParam = "PARTNER_CENTER_ANALYST"
	BusinessbusinessUsersTasksEnumParamPartnerCenterEducation  BusinessbusinessUsersTasksEnumParam = "PARTNER_CENTER_EDUCATION"
	BusinessbusinessUsersTasksEnumParamPartnerCenterMarketing  BusinessbusinessUsersTasksEnumParam = "PARTNER_CENTER_MARKETING"
	BusinessbusinessUsersTasksEnumParamPartnerCenterOperations BusinessbusinessUsersTasksEnumParam = "PARTNER_CENTER_OPERATIONS"
)

func (value BusinessbusinessUsersTasksEnumParam) String() string {
	return string(value)
}
