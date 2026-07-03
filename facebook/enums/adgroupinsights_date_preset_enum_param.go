package enums

type AdgroupinsightsDatePresetEnumParam string

const (
	AdgroupinsightsDatePresetEnumParamDataMaximum      AdgroupinsightsDatePresetEnumParam = "data_maximum"
	AdgroupinsightsDatePresetEnumParamLastX14d         AdgroupinsightsDatePresetEnumParam = "last_14d"
	AdgroupinsightsDatePresetEnumParamLastX28d         AdgroupinsightsDatePresetEnumParam = "last_28d"
	AdgroupinsightsDatePresetEnumParamLastX30d         AdgroupinsightsDatePresetEnumParam = "last_30d"
	AdgroupinsightsDatePresetEnumParamLastX3d          AdgroupinsightsDatePresetEnumParam = "last_3d"
	AdgroupinsightsDatePresetEnumParamLastX7d          AdgroupinsightsDatePresetEnumParam = "last_7d"
	AdgroupinsightsDatePresetEnumParamLastX90d         AdgroupinsightsDatePresetEnumParam = "last_90d"
	AdgroupinsightsDatePresetEnumParamLastMonth        AdgroupinsightsDatePresetEnumParam = "last_month"
	AdgroupinsightsDatePresetEnumParamLastQuarter      AdgroupinsightsDatePresetEnumParam = "last_quarter"
	AdgroupinsightsDatePresetEnumParamLastWeekMonSun   AdgroupinsightsDatePresetEnumParam = "last_week_mon_sun"
	AdgroupinsightsDatePresetEnumParamLastWeekSunSat   AdgroupinsightsDatePresetEnumParam = "last_week_sun_sat"
	AdgroupinsightsDatePresetEnumParamLastYear         AdgroupinsightsDatePresetEnumParam = "last_year"
	AdgroupinsightsDatePresetEnumParamMaximum          AdgroupinsightsDatePresetEnumParam = "maximum"
	AdgroupinsightsDatePresetEnumParamThisMonth        AdgroupinsightsDatePresetEnumParam = "this_month"
	AdgroupinsightsDatePresetEnumParamThisQuarter      AdgroupinsightsDatePresetEnumParam = "this_quarter"
	AdgroupinsightsDatePresetEnumParamThisWeekMonToday AdgroupinsightsDatePresetEnumParam = "this_week_mon_today"
	AdgroupinsightsDatePresetEnumParamThisWeekSunToday AdgroupinsightsDatePresetEnumParam = "this_week_sun_today"
	AdgroupinsightsDatePresetEnumParamThisYear         AdgroupinsightsDatePresetEnumParam = "this_year"
	AdgroupinsightsDatePresetEnumParamToday            AdgroupinsightsDatePresetEnumParam = "today"
	AdgroupinsightsDatePresetEnumParamYesterday        AdgroupinsightsDatePresetEnumParam = "yesterday"
)

func (value AdgroupinsightsDatePresetEnumParam) String() string {
	return string(value)
}
