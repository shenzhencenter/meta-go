package enums

type PostinsightsDatePresetEnumParam string

const (
	PostinsightsDatePresetEnumParamDataMaximum      PostinsightsDatePresetEnumParam = "data_maximum"
	PostinsightsDatePresetEnumParamLastX14d         PostinsightsDatePresetEnumParam = "last_14d"
	PostinsightsDatePresetEnumParamLastX28d         PostinsightsDatePresetEnumParam = "last_28d"
	PostinsightsDatePresetEnumParamLastX30d         PostinsightsDatePresetEnumParam = "last_30d"
	PostinsightsDatePresetEnumParamLastX3d          PostinsightsDatePresetEnumParam = "last_3d"
	PostinsightsDatePresetEnumParamLastX7d          PostinsightsDatePresetEnumParam = "last_7d"
	PostinsightsDatePresetEnumParamLastX90d         PostinsightsDatePresetEnumParam = "last_90d"
	PostinsightsDatePresetEnumParamLastMonth        PostinsightsDatePresetEnumParam = "last_month"
	PostinsightsDatePresetEnumParamLastQuarter      PostinsightsDatePresetEnumParam = "last_quarter"
	PostinsightsDatePresetEnumParamLastWeekMonSun   PostinsightsDatePresetEnumParam = "last_week_mon_sun"
	PostinsightsDatePresetEnumParamLastWeekSunSat   PostinsightsDatePresetEnumParam = "last_week_sun_sat"
	PostinsightsDatePresetEnumParamLastYear         PostinsightsDatePresetEnumParam = "last_year"
	PostinsightsDatePresetEnumParamMaximum          PostinsightsDatePresetEnumParam = "maximum"
	PostinsightsDatePresetEnumParamThisMonth        PostinsightsDatePresetEnumParam = "this_month"
	PostinsightsDatePresetEnumParamThisQuarter      PostinsightsDatePresetEnumParam = "this_quarter"
	PostinsightsDatePresetEnumParamThisWeekMonToday PostinsightsDatePresetEnumParam = "this_week_mon_today"
	PostinsightsDatePresetEnumParamThisWeekSunToday PostinsightsDatePresetEnumParam = "this_week_sun_today"
	PostinsightsDatePresetEnumParamThisYear         PostinsightsDatePresetEnumParam = "this_year"
	PostinsightsDatePresetEnumParamToday            PostinsightsDatePresetEnumParam = "today"
	PostinsightsDatePresetEnumParamYesterday        PostinsightsDatePresetEnumParam = "yesterday"
)

func (value PostinsightsDatePresetEnumParam) String() string {
	return string(value)
}
