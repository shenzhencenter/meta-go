package enums

type PagepostinsightsPeriodEnumParam string

const (
	PagepostinsightsPeriodEnumParamDay            PagepostinsightsPeriodEnumParam = "day"
	PagepostinsightsPeriodEnumParamDaysX28        PagepostinsightsPeriodEnumParam = "days_28"
	PagepostinsightsPeriodEnumParamLifetime       PagepostinsightsPeriodEnumParam = "lifetime"
	PagepostinsightsPeriodEnumParamMonth          PagepostinsightsPeriodEnumParam = "month"
	PagepostinsightsPeriodEnumParamTotalOverRange PagepostinsightsPeriodEnumParam = "total_over_range"
	PagepostinsightsPeriodEnumParamWeek           PagepostinsightsPeriodEnumParam = "week"
)

func (value PagepostinsightsPeriodEnumParam) String() string {
	return string(value)
}
