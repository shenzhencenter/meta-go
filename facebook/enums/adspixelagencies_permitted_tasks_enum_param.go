package enums

type AdspixelagenciesPermittedTasksEnumParam string

const (
	AdspixelagenciesPermittedTasksEnumParamAdvertise AdspixelagenciesPermittedTasksEnumParam = "ADVERTISE"
	AdspixelagenciesPermittedTasksEnumParamAnalyze   AdspixelagenciesPermittedTasksEnumParam = "ANALYZE"
	AdspixelagenciesPermittedTasksEnumParamEdit      AdspixelagenciesPermittedTasksEnumParam = "EDIT"
	AdspixelagenciesPermittedTasksEnumParamUpload    AdspixelagenciesPermittedTasksEnumParam = "UPLOAD"
)

func (value AdspixelagenciesPermittedTasksEnumParam) String() string {
	return string(value)
}
