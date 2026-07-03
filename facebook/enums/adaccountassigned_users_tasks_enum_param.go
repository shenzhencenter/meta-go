package enums

type AdaccountassignedUsersTasksEnumParam string

const (
	AdaccountassignedUsersTasksEnumParamAaAnalyze AdaccountassignedUsersTasksEnumParam = "AA_ANALYZE"
	AdaccountassignedUsersTasksEnumParamAdvertise AdaccountassignedUsersTasksEnumParam = "ADVERTISE"
	AdaccountassignedUsersTasksEnumParamAnalyze   AdaccountassignedUsersTasksEnumParam = "ANALYZE"
	AdaccountassignedUsersTasksEnumParamDraft     AdaccountassignedUsersTasksEnumParam = "DRAFT"
	AdaccountassignedUsersTasksEnumParamManage    AdaccountassignedUsersTasksEnumParam = "MANAGE"
)

func (value AdaccountassignedUsersTasksEnumParam) String() string {
	return string(value)
}
