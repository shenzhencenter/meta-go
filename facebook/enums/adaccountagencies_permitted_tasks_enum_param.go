package enums

type AdaccountagenciesPermittedTasksEnumParam string

const (
	AdaccountagenciesPermittedTasksEnumParamAaAnalyze AdaccountagenciesPermittedTasksEnumParam = "AA_ANALYZE"
	AdaccountagenciesPermittedTasksEnumParamAdvertise AdaccountagenciesPermittedTasksEnumParam = "ADVERTISE"
	AdaccountagenciesPermittedTasksEnumParamAnalyze   AdaccountagenciesPermittedTasksEnumParam = "ANALYZE"
	AdaccountagenciesPermittedTasksEnumParamDraft     AdaccountagenciesPermittedTasksEnumParam = "DRAFT"
	AdaccountagenciesPermittedTasksEnumParamManage    AdaccountagenciesPermittedTasksEnumParam = "MANAGE"
)

func (value AdaccountagenciesPermittedTasksEnumParam) String() string {
	return string(value)
}
