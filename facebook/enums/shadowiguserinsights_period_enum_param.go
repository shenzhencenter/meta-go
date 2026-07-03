package enums

type ShadowiguserinsightsPeriodEnumParam string

const (
	ShadowiguserinsightsPeriodEnumParamDay            ShadowiguserinsightsPeriodEnumParam = "day"
	ShadowiguserinsightsPeriodEnumParamDaysX28        ShadowiguserinsightsPeriodEnumParam = "days_28"
	ShadowiguserinsightsPeriodEnumParamLifetime       ShadowiguserinsightsPeriodEnumParam = "lifetime"
	ShadowiguserinsightsPeriodEnumParamMonth          ShadowiguserinsightsPeriodEnumParam = "month"
	ShadowiguserinsightsPeriodEnumParamTotalOverRange ShadowiguserinsightsPeriodEnumParam = "total_over_range"
	ShadowiguserinsightsPeriodEnumParamWeek           ShadowiguserinsightsPeriodEnumParam = "week"
)

func (value ShadowiguserinsightsPeriodEnumParam) String() string {
	return string(value)
}
