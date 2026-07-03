package enums

type AdgroupcopiesDatePresetEnumParam string

const (
	AdgroupcopiesDatePresetEnumParamDataMaximum      AdgroupcopiesDatePresetEnumParam = "data_maximum"
	AdgroupcopiesDatePresetEnumParamLastX14d         AdgroupcopiesDatePresetEnumParam = "last_14d"
	AdgroupcopiesDatePresetEnumParamLastX28d         AdgroupcopiesDatePresetEnumParam = "last_28d"
	AdgroupcopiesDatePresetEnumParamLastX30d         AdgroupcopiesDatePresetEnumParam = "last_30d"
	AdgroupcopiesDatePresetEnumParamLastX3d          AdgroupcopiesDatePresetEnumParam = "last_3d"
	AdgroupcopiesDatePresetEnumParamLastX7d          AdgroupcopiesDatePresetEnumParam = "last_7d"
	AdgroupcopiesDatePresetEnumParamLastX90d         AdgroupcopiesDatePresetEnumParam = "last_90d"
	AdgroupcopiesDatePresetEnumParamLastMonth        AdgroupcopiesDatePresetEnumParam = "last_month"
	AdgroupcopiesDatePresetEnumParamLastQuarter      AdgroupcopiesDatePresetEnumParam = "last_quarter"
	AdgroupcopiesDatePresetEnumParamLastWeekMonSun   AdgroupcopiesDatePresetEnumParam = "last_week_mon_sun"
	AdgroupcopiesDatePresetEnumParamLastWeekSunSat   AdgroupcopiesDatePresetEnumParam = "last_week_sun_sat"
	AdgroupcopiesDatePresetEnumParamLastYear         AdgroupcopiesDatePresetEnumParam = "last_year"
	AdgroupcopiesDatePresetEnumParamMaximum          AdgroupcopiesDatePresetEnumParam = "maximum"
	AdgroupcopiesDatePresetEnumParamThisMonth        AdgroupcopiesDatePresetEnumParam = "this_month"
	AdgroupcopiesDatePresetEnumParamThisQuarter      AdgroupcopiesDatePresetEnumParam = "this_quarter"
	AdgroupcopiesDatePresetEnumParamThisWeekMonToday AdgroupcopiesDatePresetEnumParam = "this_week_mon_today"
	AdgroupcopiesDatePresetEnumParamThisWeekSunToday AdgroupcopiesDatePresetEnumParam = "this_week_sun_today"
	AdgroupcopiesDatePresetEnumParamThisYear         AdgroupcopiesDatePresetEnumParam = "this_year"
	AdgroupcopiesDatePresetEnumParamToday            AdgroupcopiesDatePresetEnumParam = "today"
	AdgroupcopiesDatePresetEnumParamYesterday        AdgroupcopiesDatePresetEnumParam = "yesterday"
)

func (value AdgroupcopiesDatePresetEnumParam) String() string {
	return string(value)
}
