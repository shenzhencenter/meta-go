package enums

type AdaccountinsightsDatePresetEnumParam string

const (
	AdaccountinsightsDatePresetEnumParamDataMaximum      AdaccountinsightsDatePresetEnumParam = "data_maximum"
	AdaccountinsightsDatePresetEnumParamLastX14d         AdaccountinsightsDatePresetEnumParam = "last_14d"
	AdaccountinsightsDatePresetEnumParamLastX28d         AdaccountinsightsDatePresetEnumParam = "last_28d"
	AdaccountinsightsDatePresetEnumParamLastX30d         AdaccountinsightsDatePresetEnumParam = "last_30d"
	AdaccountinsightsDatePresetEnumParamLastX3d          AdaccountinsightsDatePresetEnumParam = "last_3d"
	AdaccountinsightsDatePresetEnumParamLastX7d          AdaccountinsightsDatePresetEnumParam = "last_7d"
	AdaccountinsightsDatePresetEnumParamLastX90d         AdaccountinsightsDatePresetEnumParam = "last_90d"
	AdaccountinsightsDatePresetEnumParamLastMonth        AdaccountinsightsDatePresetEnumParam = "last_month"
	AdaccountinsightsDatePresetEnumParamLastQuarter      AdaccountinsightsDatePresetEnumParam = "last_quarter"
	AdaccountinsightsDatePresetEnumParamLastWeekMonSun   AdaccountinsightsDatePresetEnumParam = "last_week_mon_sun"
	AdaccountinsightsDatePresetEnumParamLastWeekSunSat   AdaccountinsightsDatePresetEnumParam = "last_week_sun_sat"
	AdaccountinsightsDatePresetEnumParamLastYear         AdaccountinsightsDatePresetEnumParam = "last_year"
	AdaccountinsightsDatePresetEnumParamMaximum          AdaccountinsightsDatePresetEnumParam = "maximum"
	AdaccountinsightsDatePresetEnumParamThisMonth        AdaccountinsightsDatePresetEnumParam = "this_month"
	AdaccountinsightsDatePresetEnumParamThisQuarter      AdaccountinsightsDatePresetEnumParam = "this_quarter"
	AdaccountinsightsDatePresetEnumParamThisWeekMonToday AdaccountinsightsDatePresetEnumParam = "this_week_mon_today"
	AdaccountinsightsDatePresetEnumParamThisWeekSunToday AdaccountinsightsDatePresetEnumParam = "this_week_sun_today"
	AdaccountinsightsDatePresetEnumParamThisYear         AdaccountinsightsDatePresetEnumParam = "this_year"
	AdaccountinsightsDatePresetEnumParamToday            AdaccountinsightsDatePresetEnumParam = "today"
	AdaccountinsightsDatePresetEnumParamYesterday        AdaccountinsightsDatePresetEnumParam = "yesterday"
)

func (value AdaccountinsightsDatePresetEnumParam) String() string {
	return string(value)
}
