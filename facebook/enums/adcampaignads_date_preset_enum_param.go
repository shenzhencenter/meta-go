package enums

type AdcampaignadsDatePresetEnumParam string

const (
	AdcampaignadsDatePresetEnumParamDataMaximum      AdcampaignadsDatePresetEnumParam = "data_maximum"
	AdcampaignadsDatePresetEnumParamLastX14d         AdcampaignadsDatePresetEnumParam = "last_14d"
	AdcampaignadsDatePresetEnumParamLastX28d         AdcampaignadsDatePresetEnumParam = "last_28d"
	AdcampaignadsDatePresetEnumParamLastX30d         AdcampaignadsDatePresetEnumParam = "last_30d"
	AdcampaignadsDatePresetEnumParamLastX3d          AdcampaignadsDatePresetEnumParam = "last_3d"
	AdcampaignadsDatePresetEnumParamLastX7d          AdcampaignadsDatePresetEnumParam = "last_7d"
	AdcampaignadsDatePresetEnumParamLastX90d         AdcampaignadsDatePresetEnumParam = "last_90d"
	AdcampaignadsDatePresetEnumParamLastMonth        AdcampaignadsDatePresetEnumParam = "last_month"
	AdcampaignadsDatePresetEnumParamLastQuarter      AdcampaignadsDatePresetEnumParam = "last_quarter"
	AdcampaignadsDatePresetEnumParamLastWeekMonSun   AdcampaignadsDatePresetEnumParam = "last_week_mon_sun"
	AdcampaignadsDatePresetEnumParamLastWeekSunSat   AdcampaignadsDatePresetEnumParam = "last_week_sun_sat"
	AdcampaignadsDatePresetEnumParamLastYear         AdcampaignadsDatePresetEnumParam = "last_year"
	AdcampaignadsDatePresetEnumParamMaximum          AdcampaignadsDatePresetEnumParam = "maximum"
	AdcampaignadsDatePresetEnumParamThisMonth        AdcampaignadsDatePresetEnumParam = "this_month"
	AdcampaignadsDatePresetEnumParamThisQuarter      AdcampaignadsDatePresetEnumParam = "this_quarter"
	AdcampaignadsDatePresetEnumParamThisWeekMonToday AdcampaignadsDatePresetEnumParam = "this_week_mon_today"
	AdcampaignadsDatePresetEnumParamThisWeekSunToday AdcampaignadsDatePresetEnumParam = "this_week_sun_today"
	AdcampaignadsDatePresetEnumParamThisYear         AdcampaignadsDatePresetEnumParam = "this_year"
	AdcampaignadsDatePresetEnumParamToday            AdcampaignadsDatePresetEnumParam = "today"
	AdcampaignadsDatePresetEnumParamYesterday        AdcampaignadsDatePresetEnumParam = "yesterday"
)

func (value AdcampaignadsDatePresetEnumParam) String() string {
	return string(value)
}
