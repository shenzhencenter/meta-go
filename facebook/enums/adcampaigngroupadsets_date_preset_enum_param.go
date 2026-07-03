package enums

type AdcampaigngroupadsetsDatePresetEnumParam string

const (
	AdcampaigngroupadsetsDatePresetEnumParamDataMaximum      AdcampaigngroupadsetsDatePresetEnumParam = "data_maximum"
	AdcampaigngroupadsetsDatePresetEnumParamLastX14d         AdcampaigngroupadsetsDatePresetEnumParam = "last_14d"
	AdcampaigngroupadsetsDatePresetEnumParamLastX28d         AdcampaigngroupadsetsDatePresetEnumParam = "last_28d"
	AdcampaigngroupadsetsDatePresetEnumParamLastX30d         AdcampaigngroupadsetsDatePresetEnumParam = "last_30d"
	AdcampaigngroupadsetsDatePresetEnumParamLastX3d          AdcampaigngroupadsetsDatePresetEnumParam = "last_3d"
	AdcampaigngroupadsetsDatePresetEnumParamLastX7d          AdcampaigngroupadsetsDatePresetEnumParam = "last_7d"
	AdcampaigngroupadsetsDatePresetEnumParamLastX90d         AdcampaigngroupadsetsDatePresetEnumParam = "last_90d"
	AdcampaigngroupadsetsDatePresetEnumParamLastMonth        AdcampaigngroupadsetsDatePresetEnumParam = "last_month"
	AdcampaigngroupadsetsDatePresetEnumParamLastQuarter      AdcampaigngroupadsetsDatePresetEnumParam = "last_quarter"
	AdcampaigngroupadsetsDatePresetEnumParamLastWeekMonSun   AdcampaigngroupadsetsDatePresetEnumParam = "last_week_mon_sun"
	AdcampaigngroupadsetsDatePresetEnumParamLastWeekSunSat   AdcampaigngroupadsetsDatePresetEnumParam = "last_week_sun_sat"
	AdcampaigngroupadsetsDatePresetEnumParamLastYear         AdcampaigngroupadsetsDatePresetEnumParam = "last_year"
	AdcampaigngroupadsetsDatePresetEnumParamMaximum          AdcampaigngroupadsetsDatePresetEnumParam = "maximum"
	AdcampaigngroupadsetsDatePresetEnumParamThisMonth        AdcampaigngroupadsetsDatePresetEnumParam = "this_month"
	AdcampaigngroupadsetsDatePresetEnumParamThisQuarter      AdcampaigngroupadsetsDatePresetEnumParam = "this_quarter"
	AdcampaigngroupadsetsDatePresetEnumParamThisWeekMonToday AdcampaigngroupadsetsDatePresetEnumParam = "this_week_mon_today"
	AdcampaigngroupadsetsDatePresetEnumParamThisWeekSunToday AdcampaigngroupadsetsDatePresetEnumParam = "this_week_sun_today"
	AdcampaigngroupadsetsDatePresetEnumParamThisYear         AdcampaigngroupadsetsDatePresetEnumParam = "this_year"
	AdcampaigngroupadsetsDatePresetEnumParamToday            AdcampaigngroupadsetsDatePresetEnumParam = "today"
	AdcampaigngroupadsetsDatePresetEnumParamYesterday        AdcampaigngroupadsetsDatePresetEnumParam = "yesterday"
)

func (value AdcampaigngroupadsetsDatePresetEnumParam) String() string {
	return string(value)
}
