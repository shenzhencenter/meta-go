package enums

type MediainsightsPeriodEnumParam string

const (
	MediainsightsPeriodEnumParamDay            MediainsightsPeriodEnumParam = "day"
	MediainsightsPeriodEnumParamDaysX28        MediainsightsPeriodEnumParam = "days_28"
	MediainsightsPeriodEnumParamLifetime       MediainsightsPeriodEnumParam = "lifetime"
	MediainsightsPeriodEnumParamMonth          MediainsightsPeriodEnumParam = "month"
	MediainsightsPeriodEnumParamTotalOverRange MediainsightsPeriodEnumParam = "total_over_range"
	MediainsightsPeriodEnumParamWeek           MediainsightsPeriodEnumParam = "week"
)

func (value MediainsightsPeriodEnumParam) String() string {
	return string(value)
}
