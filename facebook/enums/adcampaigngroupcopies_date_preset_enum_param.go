package enums

type AdcampaigngroupcopiesDatePresetEnumParam string

const (
	AdcampaigngroupcopiesDatePresetEnumParamDataMaximum      AdcampaigngroupcopiesDatePresetEnumParam = "data_maximum"
	AdcampaigngroupcopiesDatePresetEnumParamLastX14d         AdcampaigngroupcopiesDatePresetEnumParam = "last_14d"
	AdcampaigngroupcopiesDatePresetEnumParamLastX28d         AdcampaigngroupcopiesDatePresetEnumParam = "last_28d"
	AdcampaigngroupcopiesDatePresetEnumParamLastX30d         AdcampaigngroupcopiesDatePresetEnumParam = "last_30d"
	AdcampaigngroupcopiesDatePresetEnumParamLastX3d          AdcampaigngroupcopiesDatePresetEnumParam = "last_3d"
	AdcampaigngroupcopiesDatePresetEnumParamLastX7d          AdcampaigngroupcopiesDatePresetEnumParam = "last_7d"
	AdcampaigngroupcopiesDatePresetEnumParamLastX90d         AdcampaigngroupcopiesDatePresetEnumParam = "last_90d"
	AdcampaigngroupcopiesDatePresetEnumParamLastMonth        AdcampaigngroupcopiesDatePresetEnumParam = "last_month"
	AdcampaigngroupcopiesDatePresetEnumParamLastQuarter      AdcampaigngroupcopiesDatePresetEnumParam = "last_quarter"
	AdcampaigngroupcopiesDatePresetEnumParamLastWeekMonSun   AdcampaigngroupcopiesDatePresetEnumParam = "last_week_mon_sun"
	AdcampaigngroupcopiesDatePresetEnumParamLastWeekSunSat   AdcampaigngroupcopiesDatePresetEnumParam = "last_week_sun_sat"
	AdcampaigngroupcopiesDatePresetEnumParamLastYear         AdcampaigngroupcopiesDatePresetEnumParam = "last_year"
	AdcampaigngroupcopiesDatePresetEnumParamMaximum          AdcampaigngroupcopiesDatePresetEnumParam = "maximum"
	AdcampaigngroupcopiesDatePresetEnumParamThisMonth        AdcampaigngroupcopiesDatePresetEnumParam = "this_month"
	AdcampaigngroupcopiesDatePresetEnumParamThisQuarter      AdcampaigngroupcopiesDatePresetEnumParam = "this_quarter"
	AdcampaigngroupcopiesDatePresetEnumParamThisWeekMonToday AdcampaigngroupcopiesDatePresetEnumParam = "this_week_mon_today"
	AdcampaigngroupcopiesDatePresetEnumParamThisWeekSunToday AdcampaigngroupcopiesDatePresetEnumParam = "this_week_sun_today"
	AdcampaigngroupcopiesDatePresetEnumParamThisYear         AdcampaigngroupcopiesDatePresetEnumParam = "this_year"
	AdcampaigngroupcopiesDatePresetEnumParamToday            AdcampaigngroupcopiesDatePresetEnumParam = "today"
	AdcampaigngroupcopiesDatePresetEnumParamYesterday        AdcampaigngroupcopiesDatePresetEnumParam = "yesterday"
)

func (value AdcampaigngroupcopiesDatePresetEnumParam) String() string {
	return string(value)
}
