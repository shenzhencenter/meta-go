package enums

type AdspixelassignedUsersTasksEnumParam string

const (
	AdspixelassignedUsersTasksEnumParamAaAnalyze AdspixelassignedUsersTasksEnumParam = "AA_ANALYZE"
	AdspixelassignedUsersTasksEnumParamAdvertise AdspixelassignedUsersTasksEnumParam = "ADVERTISE"
	AdspixelassignedUsersTasksEnumParamAnalyze   AdspixelassignedUsersTasksEnumParam = "ANALYZE"
	AdspixelassignedUsersTasksEnumParamEdit      AdspixelassignedUsersTasksEnumParam = "EDIT"
	AdspixelassignedUsersTasksEnumParamUpload    AdspixelassignedUsersTasksEnumParam = "UPLOAD"
)

func (value AdspixelassignedUsersTasksEnumParam) String() string {
	return string(value)
}
