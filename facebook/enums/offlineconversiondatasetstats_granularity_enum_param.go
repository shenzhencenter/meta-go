package enums

type OfflineconversiondatasetstatsGranularityEnumParam string

const (
	OfflineconversiondatasetstatsGranularityEnumParamDaily     OfflineconversiondatasetstatsGranularityEnumParam = "daily"
	OfflineconversiondatasetstatsGranularityEnumParamHourly    OfflineconversiondatasetstatsGranularityEnumParam = "hourly"
	OfflineconversiondatasetstatsGranularityEnumParamSixHourly OfflineconversiondatasetstatsGranularityEnumParam = "six_hourly"
)

func (value OfflineconversiondatasetstatsGranularityEnumParam) String() string {
	return string(value)
}
