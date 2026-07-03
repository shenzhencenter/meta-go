package enums

type PhotoinsightsPeriodEnumParam string

const (
	PhotoinsightsPeriodEnumParamDay            PhotoinsightsPeriodEnumParam = "day"
	PhotoinsightsPeriodEnumParamDaysX28        PhotoinsightsPeriodEnumParam = "days_28"
	PhotoinsightsPeriodEnumParamLifetime       PhotoinsightsPeriodEnumParam = "lifetime"
	PhotoinsightsPeriodEnumParamMonth          PhotoinsightsPeriodEnumParam = "month"
	PhotoinsightsPeriodEnumParamTotalOverRange PhotoinsightsPeriodEnumParam = "total_over_range"
	PhotoinsightsPeriodEnumParamWeek           PhotoinsightsPeriodEnumParam = "week"
)

func (value PhotoinsightsPeriodEnumParam) String() string {
	return string(value)
}
