package enums

type BusinesssystemUsersRoleEnumParam string

const (
	BusinesssystemUsersRoleEnumParamAdmin                   BusinesssystemUsersRoleEnumParam = "ADMIN"
	BusinesssystemUsersRoleEnumParamAdsRightsReviewer       BusinesssystemUsersRoleEnumParam = "ADS_RIGHTS_REVIEWER"
	BusinesssystemUsersRoleEnumParamDefault                 BusinesssystemUsersRoleEnumParam = "DEFAULT"
	BusinesssystemUsersRoleEnumParamDeveloper               BusinesssystemUsersRoleEnumParam = "DEVELOPER"
	BusinesssystemUsersRoleEnumParamEmployee                BusinesssystemUsersRoleEnumParam = "EMPLOYEE"
	BusinesssystemUsersRoleEnumParamFinanceAnalyst          BusinesssystemUsersRoleEnumParam = "FINANCE_ANALYST"
	BusinesssystemUsersRoleEnumParamFinanceEdit             BusinesssystemUsersRoleEnumParam = "FINANCE_EDIT"
	BusinesssystemUsersRoleEnumParamFinanceEditor           BusinesssystemUsersRoleEnumParam = "FINANCE_EDITOR"
	BusinesssystemUsersRoleEnumParamFinanceView             BusinesssystemUsersRoleEnumParam = "FINANCE_VIEW"
	BusinesssystemUsersRoleEnumParamManage                  BusinesssystemUsersRoleEnumParam = "MANAGE"
	BusinesssystemUsersRoleEnumParamPartnerCenterAdmin      BusinesssystemUsersRoleEnumParam = "PARTNER_CENTER_ADMIN"
	BusinesssystemUsersRoleEnumParamPartnerCenterAnalyst    BusinesssystemUsersRoleEnumParam = "PARTNER_CENTER_ANALYST"
	BusinesssystemUsersRoleEnumParamPartnerCenterEducation  BusinesssystemUsersRoleEnumParam = "PARTNER_CENTER_EDUCATION"
	BusinesssystemUsersRoleEnumParamPartnerCenterMarketing  BusinesssystemUsersRoleEnumParam = "PARTNER_CENTER_MARKETING"
	BusinesssystemUsersRoleEnumParamPartnerCenterOperations BusinesssystemUsersRoleEnumParam = "PARTNER_CENTER_OPERATIONS"
)

func (value BusinesssystemUsersRoleEnumParam) String() string {
	return string(value)
}
