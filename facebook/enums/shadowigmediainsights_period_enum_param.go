package enums

type ShadowigmediainsightsPeriodEnumParam string

const (
	ShadowigmediainsightsPeriodEnumParamDay            ShadowigmediainsightsPeriodEnumParam = "day"
	ShadowigmediainsightsPeriodEnumParamDaysX28        ShadowigmediainsightsPeriodEnumParam = "days_28"
	ShadowigmediainsightsPeriodEnumParamLifetime       ShadowigmediainsightsPeriodEnumParam = "lifetime"
	ShadowigmediainsightsPeriodEnumParamMonth          ShadowigmediainsightsPeriodEnumParam = "month"
	ShadowigmediainsightsPeriodEnumParamTotalOverRange ShadowigmediainsightsPeriodEnumParam = "total_over_range"
	ShadowigmediainsightsPeriodEnumParamWeek           ShadowigmediainsightsPeriodEnumParam = "week"
)

func (value ShadowigmediainsightsPeriodEnumParam) String() string {
	return string(value)
}
