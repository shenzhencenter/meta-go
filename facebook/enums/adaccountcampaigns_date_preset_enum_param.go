package enums

type AdaccountcampaignsDatePresetEnumParam string

const (
	AdaccountcampaignsDatePresetEnumParamDataMaximum      AdaccountcampaignsDatePresetEnumParam = "data_maximum"
	AdaccountcampaignsDatePresetEnumParamLastX14d         AdaccountcampaignsDatePresetEnumParam = "last_14d"
	AdaccountcampaignsDatePresetEnumParamLastX28d         AdaccountcampaignsDatePresetEnumParam = "last_28d"
	AdaccountcampaignsDatePresetEnumParamLastX30d         AdaccountcampaignsDatePresetEnumParam = "last_30d"
	AdaccountcampaignsDatePresetEnumParamLastX3d          AdaccountcampaignsDatePresetEnumParam = "last_3d"
	AdaccountcampaignsDatePresetEnumParamLastX7d          AdaccountcampaignsDatePresetEnumParam = "last_7d"
	AdaccountcampaignsDatePresetEnumParamLastX90d         AdaccountcampaignsDatePresetEnumParam = "last_90d"
	AdaccountcampaignsDatePresetEnumParamLastMonth        AdaccountcampaignsDatePresetEnumParam = "last_month"
	AdaccountcampaignsDatePresetEnumParamLastQuarter      AdaccountcampaignsDatePresetEnumParam = "last_quarter"
	AdaccountcampaignsDatePresetEnumParamLastWeekMonSun   AdaccountcampaignsDatePresetEnumParam = "last_week_mon_sun"
	AdaccountcampaignsDatePresetEnumParamLastWeekSunSat   AdaccountcampaignsDatePresetEnumParam = "last_week_sun_sat"
	AdaccountcampaignsDatePresetEnumParamLastYear         AdaccountcampaignsDatePresetEnumParam = "last_year"
	AdaccountcampaignsDatePresetEnumParamMaximum          AdaccountcampaignsDatePresetEnumParam = "maximum"
	AdaccountcampaignsDatePresetEnumParamThisMonth        AdaccountcampaignsDatePresetEnumParam = "this_month"
	AdaccountcampaignsDatePresetEnumParamThisQuarter      AdaccountcampaignsDatePresetEnumParam = "this_quarter"
	AdaccountcampaignsDatePresetEnumParamThisWeekMonToday AdaccountcampaignsDatePresetEnumParam = "this_week_mon_today"
	AdaccountcampaignsDatePresetEnumParamThisWeekSunToday AdaccountcampaignsDatePresetEnumParam = "this_week_sun_today"
	AdaccountcampaignsDatePresetEnumParamThisYear         AdaccountcampaignsDatePresetEnumParam = "this_year"
	AdaccountcampaignsDatePresetEnumParamToday            AdaccountcampaignsDatePresetEnumParam = "today"
	AdaccountcampaignsDatePresetEnumParamYesterday        AdaccountcampaignsDatePresetEnumParam = "yesterday"
)

func (value AdaccountcampaignsDatePresetEnumParam) String() string {
	return string(value)
}
