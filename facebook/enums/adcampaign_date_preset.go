package enums

type AdcampaignDatePreset string

const (
	AdcampaignDatePresetDataMaximum      AdcampaignDatePreset = "data_maximum"
	AdcampaignDatePresetLastX14d         AdcampaignDatePreset = "last_14d"
	AdcampaignDatePresetLastX28d         AdcampaignDatePreset = "last_28d"
	AdcampaignDatePresetLastX30d         AdcampaignDatePreset = "last_30d"
	AdcampaignDatePresetLastX3d          AdcampaignDatePreset = "last_3d"
	AdcampaignDatePresetLastX7d          AdcampaignDatePreset = "last_7d"
	AdcampaignDatePresetLastX90d         AdcampaignDatePreset = "last_90d"
	AdcampaignDatePresetLastMonth        AdcampaignDatePreset = "last_month"
	AdcampaignDatePresetLastQuarter      AdcampaignDatePreset = "last_quarter"
	AdcampaignDatePresetLastWeekMonSun   AdcampaignDatePreset = "last_week_mon_sun"
	AdcampaignDatePresetLastWeekSunSat   AdcampaignDatePreset = "last_week_sun_sat"
	AdcampaignDatePresetLastYear         AdcampaignDatePreset = "last_year"
	AdcampaignDatePresetMaximum          AdcampaignDatePreset = "maximum"
	AdcampaignDatePresetThisMonth        AdcampaignDatePreset = "this_month"
	AdcampaignDatePresetThisQuarter      AdcampaignDatePreset = "this_quarter"
	AdcampaignDatePresetThisWeekMonToday AdcampaignDatePreset = "this_week_mon_today"
	AdcampaignDatePresetThisWeekSunToday AdcampaignDatePreset = "this_week_sun_today"
	AdcampaignDatePresetThisYear         AdcampaignDatePreset = "this_year"
	AdcampaignDatePresetToday            AdcampaignDatePreset = "today"
	AdcampaignDatePresetYesterday        AdcampaignDatePreset = "yesterday"
)

func (value AdcampaignDatePreset) String() string {
	return string(value)
}
