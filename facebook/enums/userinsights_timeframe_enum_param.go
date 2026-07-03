package enums

type UserinsightsTimeframeEnumParam string

const (
	UserinsightsTimeframeEnumParamLastX14Days UserinsightsTimeframeEnumParam = "last_14_days"
	UserinsightsTimeframeEnumParamLastX30Days UserinsightsTimeframeEnumParam = "last_30_days"
	UserinsightsTimeframeEnumParamLastX90Days UserinsightsTimeframeEnumParam = "last_90_days"
	UserinsightsTimeframeEnumParamPrevMonth   UserinsightsTimeframeEnumParam = "prev_month"
	UserinsightsTimeframeEnumParamThisMonth   UserinsightsTimeframeEnumParam = "this_month"
	UserinsightsTimeframeEnumParamThisWeek    UserinsightsTimeframeEnumParam = "this_week"
)

func (value UserinsightsTimeframeEnumParam) String() string {
	return string(value)
}
