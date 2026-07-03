package enums

type BusinessrolerequestTasks string

const (
	BusinessrolerequestTasksAdmin                   BusinessrolerequestTasks = "ADMIN"
	BusinessrolerequestTasksAdsRightsReviewer       BusinessrolerequestTasks = "ADS_RIGHTS_REVIEWER"
	BusinessrolerequestTasksDefault                 BusinessrolerequestTasks = "DEFAULT"
	BusinessrolerequestTasksDeveloper               BusinessrolerequestTasks = "DEVELOPER"
	BusinessrolerequestTasksEmployee                BusinessrolerequestTasks = "EMPLOYEE"
	BusinessrolerequestTasksFinanceAnalyst          BusinessrolerequestTasks = "FINANCE_ANALYST"
	BusinessrolerequestTasksFinanceEdit             BusinessrolerequestTasks = "FINANCE_EDIT"
	BusinessrolerequestTasksFinanceEditor           BusinessrolerequestTasks = "FINANCE_EDITOR"
	BusinessrolerequestTasksFinanceView             BusinessrolerequestTasks = "FINANCE_VIEW"
	BusinessrolerequestTasksManage                  BusinessrolerequestTasks = "MANAGE"
	BusinessrolerequestTasksPartnerCenterAdmin      BusinessrolerequestTasks = "PARTNER_CENTER_ADMIN"
	BusinessrolerequestTasksPartnerCenterAnalyst    BusinessrolerequestTasks = "PARTNER_CENTER_ANALYST"
	BusinessrolerequestTasksPartnerCenterEducation  BusinessrolerequestTasks = "PARTNER_CENTER_EDUCATION"
	BusinessrolerequestTasksPartnerCenterMarketing  BusinessrolerequestTasks = "PARTNER_CENTER_MARKETING"
	BusinessrolerequestTasksPartnerCenterOperations BusinessrolerequestTasks = "PARTNER_CENTER_OPERATIONS"
)

func (value BusinessrolerequestTasks) String() string {
	return string(value)
}
