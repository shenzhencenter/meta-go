package enums

type AdcampaigngroupadsDatePresetEnumParam string

const (
	AdcampaigngroupadsDatePresetEnumParamDataMaximum      AdcampaigngroupadsDatePresetEnumParam = "data_maximum"
	AdcampaigngroupadsDatePresetEnumParamLastX14d         AdcampaigngroupadsDatePresetEnumParam = "last_14d"
	AdcampaigngroupadsDatePresetEnumParamLastX28d         AdcampaigngroupadsDatePresetEnumParam = "last_28d"
	AdcampaigngroupadsDatePresetEnumParamLastX30d         AdcampaigngroupadsDatePresetEnumParam = "last_30d"
	AdcampaigngroupadsDatePresetEnumParamLastX3d          AdcampaigngroupadsDatePresetEnumParam = "last_3d"
	AdcampaigngroupadsDatePresetEnumParamLastX7d          AdcampaigngroupadsDatePresetEnumParam = "last_7d"
	AdcampaigngroupadsDatePresetEnumParamLastX90d         AdcampaigngroupadsDatePresetEnumParam = "last_90d"
	AdcampaigngroupadsDatePresetEnumParamLastMonth        AdcampaigngroupadsDatePresetEnumParam = "last_month"
	AdcampaigngroupadsDatePresetEnumParamLastQuarter      AdcampaigngroupadsDatePresetEnumParam = "last_quarter"
	AdcampaigngroupadsDatePresetEnumParamLastWeekMonSun   AdcampaigngroupadsDatePresetEnumParam = "last_week_mon_sun"
	AdcampaigngroupadsDatePresetEnumParamLastWeekSunSat   AdcampaigngroupadsDatePresetEnumParam = "last_week_sun_sat"
	AdcampaigngroupadsDatePresetEnumParamLastYear         AdcampaigngroupadsDatePresetEnumParam = "last_year"
	AdcampaigngroupadsDatePresetEnumParamMaximum          AdcampaigngroupadsDatePresetEnumParam = "maximum"
	AdcampaigngroupadsDatePresetEnumParamThisMonth        AdcampaigngroupadsDatePresetEnumParam = "this_month"
	AdcampaigngroupadsDatePresetEnumParamThisQuarter      AdcampaigngroupadsDatePresetEnumParam = "this_quarter"
	AdcampaigngroupadsDatePresetEnumParamThisWeekMonToday AdcampaigngroupadsDatePresetEnumParam = "this_week_mon_today"
	AdcampaigngroupadsDatePresetEnumParamThisWeekSunToday AdcampaigngroupadsDatePresetEnumParam = "this_week_sun_today"
	AdcampaigngroupadsDatePresetEnumParamThisYear         AdcampaigngroupadsDatePresetEnumParam = "this_year"
	AdcampaigngroupadsDatePresetEnumParamToday            AdcampaigngroupadsDatePresetEnumParam = "today"
	AdcampaigngroupadsDatePresetEnumParamYesterday        AdcampaigngroupadsDatePresetEnumParam = "yesterday"
)

func (value AdcampaigngroupadsDatePresetEnumParam) String() string {
	return string(value)
}
