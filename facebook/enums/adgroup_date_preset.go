package enums

type AdgroupDatePreset string

const (
	AdgroupDatePresetDataMaximum      AdgroupDatePreset = "data_maximum"
	AdgroupDatePresetLastX14d         AdgroupDatePreset = "last_14d"
	AdgroupDatePresetLastX28d         AdgroupDatePreset = "last_28d"
	AdgroupDatePresetLastX30d         AdgroupDatePreset = "last_30d"
	AdgroupDatePresetLastX3d          AdgroupDatePreset = "last_3d"
	AdgroupDatePresetLastX7d          AdgroupDatePreset = "last_7d"
	AdgroupDatePresetLastX90d         AdgroupDatePreset = "last_90d"
	AdgroupDatePresetLastMonth        AdgroupDatePreset = "last_month"
	AdgroupDatePresetLastQuarter      AdgroupDatePreset = "last_quarter"
	AdgroupDatePresetLastWeekMonSun   AdgroupDatePreset = "last_week_mon_sun"
	AdgroupDatePresetLastWeekSunSat   AdgroupDatePreset = "last_week_sun_sat"
	AdgroupDatePresetLastYear         AdgroupDatePreset = "last_year"
	AdgroupDatePresetMaximum          AdgroupDatePreset = "maximum"
	AdgroupDatePresetThisMonth        AdgroupDatePreset = "this_month"
	AdgroupDatePresetThisQuarter      AdgroupDatePreset = "this_quarter"
	AdgroupDatePresetThisWeekMonToday AdgroupDatePreset = "this_week_mon_today"
	AdgroupDatePresetThisWeekSunToday AdgroupDatePreset = "this_week_sun_today"
	AdgroupDatePresetThisYear         AdgroupDatePreset = "this_year"
	AdgroupDatePresetToday            AdgroupDatePreset = "today"
	AdgroupDatePresetYesterday        AdgroupDatePreset = "yesterday"
)

func (value AdgroupDatePreset) String() string {
	return string(value)
}
