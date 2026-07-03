package enums

type BusinessbusinessUsersRoleEnumParam string

const (
	BusinessbusinessUsersRoleEnumParamAdmin                   BusinessbusinessUsersRoleEnumParam = "ADMIN"
	BusinessbusinessUsersRoleEnumParamAdsRightsReviewer       BusinessbusinessUsersRoleEnumParam = "ADS_RIGHTS_REVIEWER"
	BusinessbusinessUsersRoleEnumParamDefault                 BusinessbusinessUsersRoleEnumParam = "DEFAULT"
	BusinessbusinessUsersRoleEnumParamDeveloper               BusinessbusinessUsersRoleEnumParam = "DEVELOPER"
	BusinessbusinessUsersRoleEnumParamEmployee                BusinessbusinessUsersRoleEnumParam = "EMPLOYEE"
	BusinessbusinessUsersRoleEnumParamFinanceAnalyst          BusinessbusinessUsersRoleEnumParam = "FINANCE_ANALYST"
	BusinessbusinessUsersRoleEnumParamFinanceEdit             BusinessbusinessUsersRoleEnumParam = "FINANCE_EDIT"
	BusinessbusinessUsersRoleEnumParamFinanceEditor           BusinessbusinessUsersRoleEnumParam = "FINANCE_EDITOR"
	BusinessbusinessUsersRoleEnumParamFinanceView             BusinessbusinessUsersRoleEnumParam = "FINANCE_VIEW"
	BusinessbusinessUsersRoleEnumParamManage                  BusinessbusinessUsersRoleEnumParam = "MANAGE"
	BusinessbusinessUsersRoleEnumParamPartnerCenterAdmin      BusinessbusinessUsersRoleEnumParam = "PARTNER_CENTER_ADMIN"
	BusinessbusinessUsersRoleEnumParamPartnerCenterAnalyst    BusinessbusinessUsersRoleEnumParam = "PARTNER_CENTER_ANALYST"
	BusinessbusinessUsersRoleEnumParamPartnerCenterEducation  BusinessbusinessUsersRoleEnumParam = "PARTNER_CENTER_EDUCATION"
	BusinessbusinessUsersRoleEnumParamPartnerCenterMarketing  BusinessbusinessUsersRoleEnumParam = "PARTNER_CENTER_MARKETING"
	BusinessbusinessUsersRoleEnumParamPartnerCenterOperations BusinessbusinessUsersRoleEnumParam = "PARTNER_CENTER_OPERATIONS"
)

func (value BusinessbusinessUsersRoleEnumParam) String() string {
	return string(value)
}
