package enums

type AdcampaigngroupinsightsDatePresetEnumParam string

const (
	AdcampaigngroupinsightsDatePresetEnumParamDataMaximum      AdcampaigngroupinsightsDatePresetEnumParam = "data_maximum"
	AdcampaigngroupinsightsDatePresetEnumParamLastX14d         AdcampaigngroupinsightsDatePresetEnumParam = "last_14d"
	AdcampaigngroupinsightsDatePresetEnumParamLastX28d         AdcampaigngroupinsightsDatePresetEnumParam = "last_28d"
	AdcampaigngroupinsightsDatePresetEnumParamLastX30d         AdcampaigngroupinsightsDatePresetEnumParam = "last_30d"
	AdcampaigngroupinsightsDatePresetEnumParamLastX3d          AdcampaigngroupinsightsDatePresetEnumParam = "last_3d"
	AdcampaigngroupinsightsDatePresetEnumParamLastX7d          AdcampaigngroupinsightsDatePresetEnumParam = "last_7d"
	AdcampaigngroupinsightsDatePresetEnumParamLastX90d         AdcampaigngroupinsightsDatePresetEnumParam = "last_90d"
	AdcampaigngroupinsightsDatePresetEnumParamLastMonth        AdcampaigngroupinsightsDatePresetEnumParam = "last_month"
	AdcampaigngroupinsightsDatePresetEnumParamLastQuarter      AdcampaigngroupinsightsDatePresetEnumParam = "last_quarter"
	AdcampaigngroupinsightsDatePresetEnumParamLastWeekMonSun   AdcampaigngroupinsightsDatePresetEnumParam = "last_week_mon_sun"
	AdcampaigngroupinsightsDatePresetEnumParamLastWeekSunSat   AdcampaigngroupinsightsDatePresetEnumParam = "last_week_sun_sat"
	AdcampaigngroupinsightsDatePresetEnumParamLastYear         AdcampaigngroupinsightsDatePresetEnumParam = "last_year"
	AdcampaigngroupinsightsDatePresetEnumParamMaximum          AdcampaigngroupinsightsDatePresetEnumParam = "maximum"
	AdcampaigngroupinsightsDatePresetEnumParamThisMonth        AdcampaigngroupinsightsDatePresetEnumParam = "this_month"
	AdcampaigngroupinsightsDatePresetEnumParamThisQuarter      AdcampaigngroupinsightsDatePresetEnumParam = "this_quarter"
	AdcampaigngroupinsightsDatePresetEnumParamThisWeekMonToday AdcampaigngroupinsightsDatePresetEnumParam = "this_week_mon_today"
	AdcampaigngroupinsightsDatePresetEnumParamThisWeekSunToday AdcampaigngroupinsightsDatePresetEnumParam = "this_week_sun_today"
	AdcampaigngroupinsightsDatePresetEnumParamThisYear         AdcampaigngroupinsightsDatePresetEnumParam = "this_year"
	AdcampaigngroupinsightsDatePresetEnumParamToday            AdcampaigngroupinsightsDatePresetEnumParam = "today"
	AdcampaigngroupinsightsDatePresetEnumParamYesterday        AdcampaigngroupinsightsDatePresetEnumParam = "yesterday"
)

func (value AdcampaigngroupinsightsDatePresetEnumParam) String() string {
	return string(value)
}
