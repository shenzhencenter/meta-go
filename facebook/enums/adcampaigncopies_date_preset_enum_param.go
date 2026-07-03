package enums

type AdcampaigncopiesDatePresetEnumParam string

const (
	AdcampaigncopiesDatePresetEnumParamDataMaximum      AdcampaigncopiesDatePresetEnumParam = "data_maximum"
	AdcampaigncopiesDatePresetEnumParamLastX14d         AdcampaigncopiesDatePresetEnumParam = "last_14d"
	AdcampaigncopiesDatePresetEnumParamLastX28d         AdcampaigncopiesDatePresetEnumParam = "last_28d"
	AdcampaigncopiesDatePresetEnumParamLastX30d         AdcampaigncopiesDatePresetEnumParam = "last_30d"
	AdcampaigncopiesDatePresetEnumParamLastX3d          AdcampaigncopiesDatePresetEnumParam = "last_3d"
	AdcampaigncopiesDatePresetEnumParamLastX7d          AdcampaigncopiesDatePresetEnumParam = "last_7d"
	AdcampaigncopiesDatePresetEnumParamLastX90d         AdcampaigncopiesDatePresetEnumParam = "last_90d"
	AdcampaigncopiesDatePresetEnumParamLastMonth        AdcampaigncopiesDatePresetEnumParam = "last_month"
	AdcampaigncopiesDatePresetEnumParamLastQuarter      AdcampaigncopiesDatePresetEnumParam = "last_quarter"
	AdcampaigncopiesDatePresetEnumParamLastWeekMonSun   AdcampaigncopiesDatePresetEnumParam = "last_week_mon_sun"
	AdcampaigncopiesDatePresetEnumParamLastWeekSunSat   AdcampaigncopiesDatePresetEnumParam = "last_week_sun_sat"
	AdcampaigncopiesDatePresetEnumParamLastYear         AdcampaigncopiesDatePresetEnumParam = "last_year"
	AdcampaigncopiesDatePresetEnumParamMaximum          AdcampaigncopiesDatePresetEnumParam = "maximum"
	AdcampaigncopiesDatePresetEnumParamThisMonth        AdcampaigncopiesDatePresetEnumParam = "this_month"
	AdcampaigncopiesDatePresetEnumParamThisQuarter      AdcampaigncopiesDatePresetEnumParam = "this_quarter"
	AdcampaigncopiesDatePresetEnumParamThisWeekMonToday AdcampaigncopiesDatePresetEnumParam = "this_week_mon_today"
	AdcampaigncopiesDatePresetEnumParamThisWeekSunToday AdcampaigncopiesDatePresetEnumParam = "this_week_sun_today"
	AdcampaigncopiesDatePresetEnumParamThisYear         AdcampaigncopiesDatePresetEnumParam = "this_year"
	AdcampaigncopiesDatePresetEnumParamToday            AdcampaigncopiesDatePresetEnumParam = "today"
	AdcampaigncopiesDatePresetEnumParamYesterday        AdcampaigncopiesDatePresetEnumParam = "yesterday"
)

func (value AdcampaigncopiesDatePresetEnumParam) String() string {
	return string(value)
}
