package enums

type PagepostinsightsDatePresetEnumParam string

const (
	PagepostinsightsDatePresetEnumParamDataMaximum      PagepostinsightsDatePresetEnumParam = "data_maximum"
	PagepostinsightsDatePresetEnumParamLastX14d         PagepostinsightsDatePresetEnumParam = "last_14d"
	PagepostinsightsDatePresetEnumParamLastX28d         PagepostinsightsDatePresetEnumParam = "last_28d"
	PagepostinsightsDatePresetEnumParamLastX30d         PagepostinsightsDatePresetEnumParam = "last_30d"
	PagepostinsightsDatePresetEnumParamLastX3d          PagepostinsightsDatePresetEnumParam = "last_3d"
	PagepostinsightsDatePresetEnumParamLastX7d          PagepostinsightsDatePresetEnumParam = "last_7d"
	PagepostinsightsDatePresetEnumParamLastX90d         PagepostinsightsDatePresetEnumParam = "last_90d"
	PagepostinsightsDatePresetEnumParamLastMonth        PagepostinsightsDatePresetEnumParam = "last_month"
	PagepostinsightsDatePresetEnumParamLastQuarter      PagepostinsightsDatePresetEnumParam = "last_quarter"
	PagepostinsightsDatePresetEnumParamLastWeekMonSun   PagepostinsightsDatePresetEnumParam = "last_week_mon_sun"
	PagepostinsightsDatePresetEnumParamLastWeekSunSat   PagepostinsightsDatePresetEnumParam = "last_week_sun_sat"
	PagepostinsightsDatePresetEnumParamLastYear         PagepostinsightsDatePresetEnumParam = "last_year"
	PagepostinsightsDatePresetEnumParamMaximum          PagepostinsightsDatePresetEnumParam = "maximum"
	PagepostinsightsDatePresetEnumParamThisMonth        PagepostinsightsDatePresetEnumParam = "this_month"
	PagepostinsightsDatePresetEnumParamThisQuarter      PagepostinsightsDatePresetEnumParam = "this_quarter"
	PagepostinsightsDatePresetEnumParamThisWeekMonToday PagepostinsightsDatePresetEnumParam = "this_week_mon_today"
	PagepostinsightsDatePresetEnumParamThisWeekSunToday PagepostinsightsDatePresetEnumParam = "this_week_sun_today"
	PagepostinsightsDatePresetEnumParamThisYear         PagepostinsightsDatePresetEnumParam = "this_year"
	PagepostinsightsDatePresetEnumParamToday            PagepostinsightsDatePresetEnumParam = "today"
	PagepostinsightsDatePresetEnumParamYesterday        PagepostinsightsDatePresetEnumParam = "yesterday"
)

func (value PagepostinsightsDatePresetEnumParam) String() string {
	return string(value)
}
