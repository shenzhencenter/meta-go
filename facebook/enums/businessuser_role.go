package enums

type BusinessuserRole string

const (
	BusinessuserRoleAdmin                   BusinessuserRole = "ADMIN"
	BusinessuserRoleAdsRightsReviewer       BusinessuserRole = "ADS_RIGHTS_REVIEWER"
	BusinessuserRoleDefault                 BusinessuserRole = "DEFAULT"
	BusinessuserRoleDeveloper               BusinessuserRole = "DEVELOPER"
	BusinessuserRoleEmployee                BusinessuserRole = "EMPLOYEE"
	BusinessuserRoleFinanceAnalyst          BusinessuserRole = "FINANCE_ANALYST"
	BusinessuserRoleFinanceEdit             BusinessuserRole = "FINANCE_EDIT"
	BusinessuserRoleFinanceEditor           BusinessuserRole = "FINANCE_EDITOR"
	BusinessuserRoleFinanceView             BusinessuserRole = "FINANCE_VIEW"
	BusinessuserRoleManage                  BusinessuserRole = "MANAGE"
	BusinessuserRolePartnerCenterAdmin      BusinessuserRole = "PARTNER_CENTER_ADMIN"
	BusinessuserRolePartnerCenterAnalyst    BusinessuserRole = "PARTNER_CENTER_ANALYST"
	BusinessuserRolePartnerCenterEducation  BusinessuserRole = "PARTNER_CENTER_EDUCATION"
	BusinessuserRolePartnerCenterMarketing  BusinessuserRole = "PARTNER_CENTER_MARKETING"
	BusinessuserRolePartnerCenterOperations BusinessuserRole = "PARTNER_CENTER_OPERATIONS"
)

func (value BusinessuserRole) String() string {
	return string(value)
}
