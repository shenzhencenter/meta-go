package enums

type CustomconversionstatsAggregationEnumParam string

const (
	CustomconversionstatsAggregationEnumParamCount              CustomconversionstatsAggregationEnumParam = "count"
	CustomconversionstatsAggregationEnumParamDeviceType         CustomconversionstatsAggregationEnumParam = "device_type"
	CustomconversionstatsAggregationEnumParamHost               CustomconversionstatsAggregationEnumParam = "host"
	CustomconversionstatsAggregationEnumParamPixelFire          CustomconversionstatsAggregationEnumParam = "pixel_fire"
	CustomconversionstatsAggregationEnumParamUnmatchedCount     CustomconversionstatsAggregationEnumParam = "unmatched_count"
	CustomconversionstatsAggregationEnumParamUnmatchedUsdAmount CustomconversionstatsAggregationEnumParam = "unmatched_usd_amount"
	CustomconversionstatsAggregationEnumParamURL                CustomconversionstatsAggregationEnumParam = "url"
	CustomconversionstatsAggregationEnumParamUsdAmount          CustomconversionstatsAggregationEnumParam = "usd_amount"
)

func (value CustomconversionstatsAggregationEnumParam) String() string {
	return string(value)
}
