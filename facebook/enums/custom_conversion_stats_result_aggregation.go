package enums

type CustomconversionstatsresultAggregation string

const (
	CustomconversionstatsresultAggregationCount              CustomconversionstatsresultAggregation = "count"
	CustomconversionstatsresultAggregationDeviceType         CustomconversionstatsresultAggregation = "device_type"
	CustomconversionstatsresultAggregationHost               CustomconversionstatsresultAggregation = "host"
	CustomconversionstatsresultAggregationPixelFire          CustomconversionstatsresultAggregation = "pixel_fire"
	CustomconversionstatsresultAggregationUnmatchedCount     CustomconversionstatsresultAggregation = "unmatched_count"
	CustomconversionstatsresultAggregationUnmatchedUsdAmount CustomconversionstatsresultAggregation = "unmatched_usd_amount"
	CustomconversionstatsresultAggregationURL                CustomconversionstatsresultAggregation = "url"
	CustomconversionstatsresultAggregationUsdAmount          CustomconversionstatsresultAggregation = "usd_amount"
)

func (value CustomconversionstatsresultAggregation) String() string {
	return string(value)
}
