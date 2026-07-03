package enums

type PhotoinsightsDatePresetEnumParam string

const (
	PhotoinsightsDatePresetEnumParamDataMaximum      PhotoinsightsDatePresetEnumParam = "data_maximum"
	PhotoinsightsDatePresetEnumParamLastX14d         PhotoinsightsDatePresetEnumParam = "last_14d"
	PhotoinsightsDatePresetEnumParamLastX28d         PhotoinsightsDatePresetEnumParam = "last_28d"
	PhotoinsightsDatePresetEnumParamLastX30d         PhotoinsightsDatePresetEnumParam = "last_30d"
	PhotoinsightsDatePresetEnumParamLastX3d          PhotoinsightsDatePresetEnumParam = "last_3d"
	PhotoinsightsDatePresetEnumParamLastX7d          PhotoinsightsDatePresetEnumParam = "last_7d"
	PhotoinsightsDatePresetEnumParamLastX90d         PhotoinsightsDatePresetEnumParam = "last_90d"
	PhotoinsightsDatePresetEnumParamLastMonth        PhotoinsightsDatePresetEnumParam = "last_month"
	PhotoinsightsDatePresetEnumParamLastQuarter      PhotoinsightsDatePresetEnumParam = "last_quarter"
	PhotoinsightsDatePresetEnumParamLastWeekMonSun   PhotoinsightsDatePresetEnumParam = "last_week_mon_sun"
	PhotoinsightsDatePresetEnumParamLastWeekSunSat   PhotoinsightsDatePresetEnumParam = "last_week_sun_sat"
	PhotoinsightsDatePresetEnumParamLastYear         PhotoinsightsDatePresetEnumParam = "last_year"
	PhotoinsightsDatePresetEnumParamMaximum          PhotoinsightsDatePresetEnumParam = "maximum"
	PhotoinsightsDatePresetEnumParamThisMonth        PhotoinsightsDatePresetEnumParam = "this_month"
	PhotoinsightsDatePresetEnumParamThisQuarter      PhotoinsightsDatePresetEnumParam = "this_quarter"
	PhotoinsightsDatePresetEnumParamThisWeekMonToday PhotoinsightsDatePresetEnumParam = "this_week_mon_today"
	PhotoinsightsDatePresetEnumParamThisWeekSunToday PhotoinsightsDatePresetEnumParam = "this_week_sun_today"
	PhotoinsightsDatePresetEnumParamThisYear         PhotoinsightsDatePresetEnumParam = "this_year"
	PhotoinsightsDatePresetEnumParamToday            PhotoinsightsDatePresetEnumParam = "today"
	PhotoinsightsDatePresetEnumParamYesterday        PhotoinsightsDatePresetEnumParam = "yesterday"
)

func (value PhotoinsightsDatePresetEnumParam) String() string {
	return string(value)
}
