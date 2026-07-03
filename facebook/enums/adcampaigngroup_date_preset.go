package enums

type AdcampaigngroupDatePreset string

const (
	AdcampaigngroupDatePresetDataMaximum      AdcampaigngroupDatePreset = "data_maximum"
	AdcampaigngroupDatePresetLastX14d         AdcampaigngroupDatePreset = "last_14d"
	AdcampaigngroupDatePresetLastX28d         AdcampaigngroupDatePreset = "last_28d"
	AdcampaigngroupDatePresetLastX30d         AdcampaigngroupDatePreset = "last_30d"
	AdcampaigngroupDatePresetLastX3d          AdcampaigngroupDatePreset = "last_3d"
	AdcampaigngroupDatePresetLastX7d          AdcampaigngroupDatePreset = "last_7d"
	AdcampaigngroupDatePresetLastX90d         AdcampaigngroupDatePreset = "last_90d"
	AdcampaigngroupDatePresetLastMonth        AdcampaigngroupDatePreset = "last_month"
	AdcampaigngroupDatePresetLastQuarter      AdcampaigngroupDatePreset = "last_quarter"
	AdcampaigngroupDatePresetLastWeekMonSun   AdcampaigngroupDatePreset = "last_week_mon_sun"
	AdcampaigngroupDatePresetLastWeekSunSat   AdcampaigngroupDatePreset = "last_week_sun_sat"
	AdcampaigngroupDatePresetLastYear         AdcampaigngroupDatePreset = "last_year"
	AdcampaigngroupDatePresetMaximum          AdcampaigngroupDatePreset = "maximum"
	AdcampaigngroupDatePresetThisMonth        AdcampaigngroupDatePreset = "this_month"
	AdcampaigngroupDatePresetThisQuarter      AdcampaigngroupDatePreset = "this_quarter"
	AdcampaigngroupDatePresetThisWeekMonToday AdcampaigngroupDatePreset = "this_week_mon_today"
	AdcampaigngroupDatePresetThisWeekSunToday AdcampaigngroupDatePreset = "this_week_sun_today"
	AdcampaigngroupDatePresetThisYear         AdcampaigngroupDatePreset = "this_year"
	AdcampaigngroupDatePresetToday            AdcampaigngroupDatePreset = "today"
	AdcampaigngroupDatePresetYesterday        AdcampaigngroupDatePreset = "yesterday"
)

func (value AdcampaigngroupDatePreset) String() string {
	return string(value)
}
