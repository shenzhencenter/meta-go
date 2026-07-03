package enums

type PageinsightsDatePresetEnumParam string

const (
	PageinsightsDatePresetEnumParamDataMaximum      PageinsightsDatePresetEnumParam = "data_maximum"
	PageinsightsDatePresetEnumParamLastX14d         PageinsightsDatePresetEnumParam = "last_14d"
	PageinsightsDatePresetEnumParamLastX28d         PageinsightsDatePresetEnumParam = "last_28d"
	PageinsightsDatePresetEnumParamLastX30d         PageinsightsDatePresetEnumParam = "last_30d"
	PageinsightsDatePresetEnumParamLastX3d          PageinsightsDatePresetEnumParam = "last_3d"
	PageinsightsDatePresetEnumParamLastX7d          PageinsightsDatePresetEnumParam = "last_7d"
	PageinsightsDatePresetEnumParamLastX90d         PageinsightsDatePresetEnumParam = "last_90d"
	PageinsightsDatePresetEnumParamLastMonth        PageinsightsDatePresetEnumParam = "last_month"
	PageinsightsDatePresetEnumParamLastQuarter      PageinsightsDatePresetEnumParam = "last_quarter"
	PageinsightsDatePresetEnumParamLastWeekMonSun   PageinsightsDatePresetEnumParam = "last_week_mon_sun"
	PageinsightsDatePresetEnumParamLastWeekSunSat   PageinsightsDatePresetEnumParam = "last_week_sun_sat"
	PageinsightsDatePresetEnumParamLastYear         PageinsightsDatePresetEnumParam = "last_year"
	PageinsightsDatePresetEnumParamMaximum          PageinsightsDatePresetEnumParam = "maximum"
	PageinsightsDatePresetEnumParamThisMonth        PageinsightsDatePresetEnumParam = "this_month"
	PageinsightsDatePresetEnumParamThisQuarter      PageinsightsDatePresetEnumParam = "this_quarter"
	PageinsightsDatePresetEnumParamThisWeekMonToday PageinsightsDatePresetEnumParam = "this_week_mon_today"
	PageinsightsDatePresetEnumParamThisWeekSunToday PageinsightsDatePresetEnumParam = "this_week_sun_today"
	PageinsightsDatePresetEnumParamThisYear         PageinsightsDatePresetEnumParam = "this_year"
	PageinsightsDatePresetEnumParamToday            PageinsightsDatePresetEnumParam = "today"
	PageinsightsDatePresetEnumParamYesterday        PageinsightsDatePresetEnumParam = "yesterday"
)

func (value PageinsightsDatePresetEnumParam) String() string {
	return string(value)
}
