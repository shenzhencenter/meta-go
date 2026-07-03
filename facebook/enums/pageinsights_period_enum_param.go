package enums

type PageinsightsPeriodEnumParam string

const (
	PageinsightsPeriodEnumParamDay            PageinsightsPeriodEnumParam = "day"
	PageinsightsPeriodEnumParamDaysX28        PageinsightsPeriodEnumParam = "days_28"
	PageinsightsPeriodEnumParamLifetime       PageinsightsPeriodEnumParam = "lifetime"
	PageinsightsPeriodEnumParamMonth          PageinsightsPeriodEnumParam = "month"
	PageinsightsPeriodEnumParamTotalOverRange PageinsightsPeriodEnumParam = "total_over_range"
	PageinsightsPeriodEnumParamWeek           PageinsightsPeriodEnumParam = "week"
)

func (value PageinsightsPeriodEnumParam) String() string {
	return string(value)
}
