package enums

type AdaccountadsetsDatePresetEnumParam string

const (
	AdaccountadsetsDatePresetEnumParamDataMaximum      AdaccountadsetsDatePresetEnumParam = "DATA_MAXIMUM"
	AdaccountadsetsDatePresetEnumParamLastX14D         AdaccountadsetsDatePresetEnumParam = "LAST_14D"
	AdaccountadsetsDatePresetEnumParamLastX28D         AdaccountadsetsDatePresetEnumParam = "LAST_28D"
	AdaccountadsetsDatePresetEnumParamLastX30D         AdaccountadsetsDatePresetEnumParam = "LAST_30D"
	AdaccountadsetsDatePresetEnumParamLastX3D          AdaccountadsetsDatePresetEnumParam = "LAST_3D"
	AdaccountadsetsDatePresetEnumParamLastX7D          AdaccountadsetsDatePresetEnumParam = "LAST_7D"
	AdaccountadsetsDatePresetEnumParamLastX90D         AdaccountadsetsDatePresetEnumParam = "LAST_90D"
	AdaccountadsetsDatePresetEnumParamLastMonth        AdaccountadsetsDatePresetEnumParam = "LAST_MONTH"
	AdaccountadsetsDatePresetEnumParamLastQuarter      AdaccountadsetsDatePresetEnumParam = "LAST_QUARTER"
	AdaccountadsetsDatePresetEnumParamLastWeekMonSun   AdaccountadsetsDatePresetEnumParam = "LAST_WEEK_MON_SUN"
	AdaccountadsetsDatePresetEnumParamLastWeekSunSat   AdaccountadsetsDatePresetEnumParam = "LAST_WEEK_SUN_SAT"
	AdaccountadsetsDatePresetEnumParamLastYear         AdaccountadsetsDatePresetEnumParam = "LAST_YEAR"
	AdaccountadsetsDatePresetEnumParamMaximum          AdaccountadsetsDatePresetEnumParam = "MAXIMUM"
	AdaccountadsetsDatePresetEnumParamThisMonth        AdaccountadsetsDatePresetEnumParam = "THIS_MONTH"
	AdaccountadsetsDatePresetEnumParamThisQuarter      AdaccountadsetsDatePresetEnumParam = "THIS_QUARTER"
	AdaccountadsetsDatePresetEnumParamThisWeekMonToday AdaccountadsetsDatePresetEnumParam = "THIS_WEEK_MON_TODAY"
	AdaccountadsetsDatePresetEnumParamThisWeekSunToday AdaccountadsetsDatePresetEnumParam = "THIS_WEEK_SUN_TODAY"
	AdaccountadsetsDatePresetEnumParamThisYear         AdaccountadsetsDatePresetEnumParam = "THIS_YEAR"
	AdaccountadsetsDatePresetEnumParamToday            AdaccountadsetsDatePresetEnumParam = "TODAY"
	AdaccountadsetsDatePresetEnumParamYesterday        AdaccountadsetsDatePresetEnumParam = "YESTERDAY"
)

func (value AdaccountadsetsDatePresetEnumParam) String() string {
	return string(value)
}
