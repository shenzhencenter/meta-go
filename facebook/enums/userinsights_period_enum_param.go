package enums

type UserinsightsPeriodEnumParam string

const (
	UserinsightsPeriodEnumParamDay            UserinsightsPeriodEnumParam = "day"
	UserinsightsPeriodEnumParamDaysX28        UserinsightsPeriodEnumParam = "days_28"
	UserinsightsPeriodEnumParamLifetime       UserinsightsPeriodEnumParam = "lifetime"
	UserinsightsPeriodEnumParamMonth          UserinsightsPeriodEnumParam = "month"
	UserinsightsPeriodEnumParamTotalOverRange UserinsightsPeriodEnumParam = "total_over_range"
	UserinsightsPeriodEnumParamWeek           UserinsightsPeriodEnumParam = "week"
)

func (value UserinsightsPeriodEnumParam) String() string {
	return string(value)
}
