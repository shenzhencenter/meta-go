package enums

type AdaccountadsDatePresetEnumParam string

const (
	AdaccountadsDatePresetEnumParamDataMaximum      AdaccountadsDatePresetEnumParam = "data_maximum"
	AdaccountadsDatePresetEnumParamLastX14d         AdaccountadsDatePresetEnumParam = "last_14d"
	AdaccountadsDatePresetEnumParamLastX28d         AdaccountadsDatePresetEnumParam = "last_28d"
	AdaccountadsDatePresetEnumParamLastX30d         AdaccountadsDatePresetEnumParam = "last_30d"
	AdaccountadsDatePresetEnumParamLastX3d          AdaccountadsDatePresetEnumParam = "last_3d"
	AdaccountadsDatePresetEnumParamLastX7d          AdaccountadsDatePresetEnumParam = "last_7d"
	AdaccountadsDatePresetEnumParamLastX90d         AdaccountadsDatePresetEnumParam = "last_90d"
	AdaccountadsDatePresetEnumParamLastMonth        AdaccountadsDatePresetEnumParam = "last_month"
	AdaccountadsDatePresetEnumParamLastQuarter      AdaccountadsDatePresetEnumParam = "last_quarter"
	AdaccountadsDatePresetEnumParamLastWeekMonSun   AdaccountadsDatePresetEnumParam = "last_week_mon_sun"
	AdaccountadsDatePresetEnumParamLastWeekSunSat   AdaccountadsDatePresetEnumParam = "last_week_sun_sat"
	AdaccountadsDatePresetEnumParamLastYear         AdaccountadsDatePresetEnumParam = "last_year"
	AdaccountadsDatePresetEnumParamMaximum          AdaccountadsDatePresetEnumParam = "maximum"
	AdaccountadsDatePresetEnumParamThisMonth        AdaccountadsDatePresetEnumParam = "this_month"
	AdaccountadsDatePresetEnumParamThisQuarter      AdaccountadsDatePresetEnumParam = "this_quarter"
	AdaccountadsDatePresetEnumParamThisWeekMonToday AdaccountadsDatePresetEnumParam = "this_week_mon_today"
	AdaccountadsDatePresetEnumParamThisWeekSunToday AdaccountadsDatePresetEnumParam = "this_week_sun_today"
	AdaccountadsDatePresetEnumParamThisYear         AdaccountadsDatePresetEnumParam = "this_year"
	AdaccountadsDatePresetEnumParamToday            AdaccountadsDatePresetEnumParam = "today"
	AdaccountadsDatePresetEnumParamYesterday        AdaccountadsDatePresetEnumParam = "yesterday"
)

func (value AdaccountadsDatePresetEnumParam) String() string {
	return string(value)
}
