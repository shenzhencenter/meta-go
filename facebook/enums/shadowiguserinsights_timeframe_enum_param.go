package enums

type ShadowiguserinsightsTimeframeEnumParam string

const (
	ShadowiguserinsightsTimeframeEnumParamLastX14Days ShadowiguserinsightsTimeframeEnumParam = "last_14_days"
	ShadowiguserinsightsTimeframeEnumParamLastX30Days ShadowiguserinsightsTimeframeEnumParam = "last_30_days"
	ShadowiguserinsightsTimeframeEnumParamLastX90Days ShadowiguserinsightsTimeframeEnumParam = "last_90_days"
	ShadowiguserinsightsTimeframeEnumParamPrevMonth   ShadowiguserinsightsTimeframeEnumParam = "prev_month"
	ShadowiguserinsightsTimeframeEnumParamThisMonth   ShadowiguserinsightsTimeframeEnumParam = "this_month"
	ShadowiguserinsightsTimeframeEnumParamThisWeek    ShadowiguserinsightsTimeframeEnumParam = "this_week"
)

func (value ShadowiguserinsightsTimeframeEnumParam) String() string {
	return string(value)
}
