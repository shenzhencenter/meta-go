package enums

type BusinessrolerequestRole string

const (
	BusinessrolerequestRoleAdmin                   BusinessrolerequestRole = "ADMIN"
	BusinessrolerequestRoleAdsRightsReviewer       BusinessrolerequestRole = "ADS_RIGHTS_REVIEWER"
	BusinessrolerequestRoleDefault                 BusinessrolerequestRole = "DEFAULT"
	BusinessrolerequestRoleDeveloper               BusinessrolerequestRole = "DEVELOPER"
	BusinessrolerequestRoleEmployee                BusinessrolerequestRole = "EMPLOYEE"
	BusinessrolerequestRoleFinanceAnalyst          BusinessrolerequestRole = "FINANCE_ANALYST"
	BusinessrolerequestRoleFinanceEdit             BusinessrolerequestRole = "FINANCE_EDIT"
	BusinessrolerequestRoleFinanceEditor           BusinessrolerequestRole = "FINANCE_EDITOR"
	BusinessrolerequestRoleFinanceView             BusinessrolerequestRole = "FINANCE_VIEW"
	BusinessrolerequestRoleManage                  BusinessrolerequestRole = "MANAGE"
	BusinessrolerequestRolePartnerCenterAdmin      BusinessrolerequestRole = "PARTNER_CENTER_ADMIN"
	BusinessrolerequestRolePartnerCenterAnalyst    BusinessrolerequestRole = "PARTNER_CENTER_ANALYST"
	BusinessrolerequestRolePartnerCenterEducation  BusinessrolerequestRole = "PARTNER_CENTER_EDUCATION"
	BusinessrolerequestRolePartnerCenterMarketing  BusinessrolerequestRole = "PARTNER_CENTER_MARKETING"
	BusinessrolerequestRolePartnerCenterOperations BusinessrolerequestRole = "PARTNER_CENTER_OPERATIONS"
)

func (value BusinessrolerequestRole) String() string {
	return string(value)
}
