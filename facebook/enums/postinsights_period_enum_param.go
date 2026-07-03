package enums

type PostinsightsPeriodEnumParam string

const (
	PostinsightsPeriodEnumParamDay            PostinsightsPeriodEnumParam = "day"
	PostinsightsPeriodEnumParamDaysX28        PostinsightsPeriodEnumParam = "days_28"
	PostinsightsPeriodEnumParamLifetime       PostinsightsPeriodEnumParam = "lifetime"
	PostinsightsPeriodEnumParamMonth          PostinsightsPeriodEnumParam = "month"
	PostinsightsPeriodEnumParamTotalOverRange PostinsightsPeriodEnumParam = "total_over_range"
	PostinsightsPeriodEnumParamWeek           PostinsightsPeriodEnumParam = "week"
)

func (value PostinsightsPeriodEnumParam) String() string {
	return string(value)
}
