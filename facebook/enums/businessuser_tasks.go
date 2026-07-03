package enums

type BusinessuserTasks string

const (
	BusinessuserTasksAdmin                   BusinessuserTasks = "ADMIN"
	BusinessuserTasksAdsRightsReviewer       BusinessuserTasks = "ADS_RIGHTS_REVIEWER"
	BusinessuserTasksDefault                 BusinessuserTasks = "DEFAULT"
	BusinessuserTasksDeveloper               BusinessuserTasks = "DEVELOPER"
	BusinessuserTasksEmployee                BusinessuserTasks = "EMPLOYEE"
	BusinessuserTasksFinanceAnalyst          BusinessuserTasks = "FINANCE_ANALYST"
	BusinessuserTasksFinanceEdit             BusinessuserTasks = "FINANCE_EDIT"
	BusinessuserTasksFinanceEditor           BusinessuserTasks = "FINANCE_EDITOR"
	BusinessuserTasksFinanceView             BusinessuserTasks = "FINANCE_VIEW"
	BusinessuserTasksManage                  BusinessuserTasks = "MANAGE"
	BusinessuserTasksPartnerCenterAdmin      BusinessuserTasks = "PARTNER_CENTER_ADMIN"
	BusinessuserTasksPartnerCenterAnalyst    BusinessuserTasks = "PARTNER_CENTER_ANALYST"
	BusinessuserTasksPartnerCenterEducation  BusinessuserTasks = "PARTNER_CENTER_EDUCATION"
	BusinessuserTasksPartnerCenterMarketing  BusinessuserTasks = "PARTNER_CENTER_MARKETING"
	BusinessuserTasksPartnerCenterOperations BusinessuserTasks = "PARTNER_CENTER_OPERATIONS"
)

func (value BusinessuserTasks) String() string {
	return string(value)
}
