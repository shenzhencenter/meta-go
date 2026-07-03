package enums

type AdspixelstatsAggregationEnumParam string

const (
	AdspixelstatsAggregationEnumParamBrowserType            AdspixelstatsAggregationEnumParam = "browser_type"
	AdspixelstatsAggregationEnumParamCustomDataField        AdspixelstatsAggregationEnumParam = "custom_data_field"
	AdspixelstatsAggregationEnumParamDeviceOs               AdspixelstatsAggregationEnumParam = "device_os"
	AdspixelstatsAggregationEnumParamDeviceType             AdspixelstatsAggregationEnumParam = "device_type"
	AdspixelstatsAggregationEnumParamEvent                  AdspixelstatsAggregationEnumParam = "event"
	AdspixelstatsAggregationEnumParamEventDetectionMethod   AdspixelstatsAggregationEnumParam = "event_detection_method"
	AdspixelstatsAggregationEnumParamEventProcessingResults AdspixelstatsAggregationEnumParam = "event_processing_results"
	AdspixelstatsAggregationEnumParamEventSource            AdspixelstatsAggregationEnumParam = "event_source"
	AdspixelstatsAggregationEnumParamEventTotalCounts       AdspixelstatsAggregationEnumParam = "event_total_counts"
	AdspixelstatsAggregationEnumParamEventValueCount        AdspixelstatsAggregationEnumParam = "event_value_count"
	AdspixelstatsAggregationEnumParamHadPii                 AdspixelstatsAggregationEnumParam = "had_pii"
	AdspixelstatsAggregationEnumParamHost                   AdspixelstatsAggregationEnumParam = "host"
	AdspixelstatsAggregationEnumParamMatchKeys              AdspixelstatsAggregationEnumParam = "match_keys"
	AdspixelstatsAggregationEnumParamPixelFire              AdspixelstatsAggregationEnumParam = "pixel_fire"
	AdspixelstatsAggregationEnumParamURL                    AdspixelstatsAggregationEnumParam = "url"
	AdspixelstatsAggregationEnumParamURLByRule              AdspixelstatsAggregationEnumParam = "url_by_rule"
)

func (value AdspixelstatsAggregationEnumParam) String() string {
	return string(value)
}
