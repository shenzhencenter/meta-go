package enums

type AdcampaigninsightsDatePresetEnumParam string

const (
	AdcampaigninsightsDatePresetEnumParamDataMaximum      AdcampaigninsightsDatePresetEnumParam = "data_maximum"
	AdcampaigninsightsDatePresetEnumParamLastX14d         AdcampaigninsightsDatePresetEnumParam = "last_14d"
	AdcampaigninsightsDatePresetEnumParamLastX28d         AdcampaigninsightsDatePresetEnumParam = "last_28d"
	AdcampaigninsightsDatePresetEnumParamLastX30d         AdcampaigninsightsDatePresetEnumParam = "last_30d"
	AdcampaigninsightsDatePresetEnumParamLastX3d          AdcampaigninsightsDatePresetEnumParam = "last_3d"
	AdcampaigninsightsDatePresetEnumParamLastX7d          AdcampaigninsightsDatePresetEnumParam = "last_7d"
	AdcampaigninsightsDatePresetEnumParamLastX90d         AdcampaigninsightsDatePresetEnumParam = "last_90d"
	AdcampaigninsightsDatePresetEnumParamLastMonth        AdcampaigninsightsDatePresetEnumParam = "last_month"
	AdcampaigninsightsDatePresetEnumParamLastQuarter      AdcampaigninsightsDatePresetEnumParam = "last_quarter"
	AdcampaigninsightsDatePresetEnumParamLastWeekMonSun   AdcampaigninsightsDatePresetEnumParam = "last_week_mon_sun"
	AdcampaigninsightsDatePresetEnumParamLastWeekSunSat   AdcampaigninsightsDatePresetEnumParam = "last_week_sun_sat"
	AdcampaigninsightsDatePresetEnumParamLastYear         AdcampaigninsightsDatePresetEnumParam = "last_year"
	AdcampaigninsightsDatePresetEnumParamMaximum          AdcampaigninsightsDatePresetEnumParam = "maximum"
	AdcampaigninsightsDatePresetEnumParamThisMonth        AdcampaigninsightsDatePresetEnumParam = "this_month"
	AdcampaigninsightsDatePresetEnumParamThisQuarter      AdcampaigninsightsDatePresetEnumParam = "this_quarter"
	AdcampaigninsightsDatePresetEnumParamThisWeekMonToday AdcampaigninsightsDatePresetEnumParam = "this_week_mon_today"
	AdcampaigninsightsDatePresetEnumParamThisWeekSunToday AdcampaigninsightsDatePresetEnumParam = "this_week_sun_today"
	AdcampaigninsightsDatePresetEnumParamThisYear         AdcampaigninsightsDatePresetEnumParam = "this_year"
	AdcampaigninsightsDatePresetEnumParamToday            AdcampaigninsightsDatePresetEnumParam = "today"
	AdcampaigninsightsDatePresetEnumParamYesterday        AdcampaigninsightsDatePresetEnumParam = "yesterday"
)

func (value AdcampaigninsightsDatePresetEnumParam) String() string {
	return string(value)
}
