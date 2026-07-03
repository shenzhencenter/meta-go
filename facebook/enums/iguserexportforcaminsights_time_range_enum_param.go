package enums

type IguserexportforcaminsightsTimeRangeEnumParam string

const (
	IguserexportforcaminsightsTimeRangeEnumParamLastX14Days IguserexportforcaminsightsTimeRangeEnumParam = "LAST_14_DAYS"
	IguserexportforcaminsightsTimeRangeEnumParamLastX90Days IguserexportforcaminsightsTimeRangeEnumParam = "LAST_90_DAYS"
	IguserexportforcaminsightsTimeRangeEnumParamLifetime    IguserexportforcaminsightsTimeRangeEnumParam = "LIFETIME"
	IguserexportforcaminsightsTimeRangeEnumParamThisMonth   IguserexportforcaminsightsTimeRangeEnumParam = "THIS_MONTH"
	IguserexportforcaminsightsTimeRangeEnumParamThisWeek    IguserexportforcaminsightsTimeRangeEnumParam = "THIS_WEEK"
)

func (value IguserexportforcaminsightsTimeRangeEnumParam) String() string {
	return string(value)
}
