package enums

type IguserexportforcaminsightsPeriodEnumParam string

const (
	IguserexportforcaminsightsPeriodEnumParamDay     IguserexportforcaminsightsPeriodEnumParam = "DAY"
	IguserexportforcaminsightsPeriodEnumParamOverall IguserexportforcaminsightsPeriodEnumParam = "OVERALL"
)

func (value IguserexportforcaminsightsPeriodEnumParam) String() string {
	return string(value)
}
