package enums

type ShadowiguserinsightsMetricTypeEnumParam string

const (
	ShadowiguserinsightsMetricTypeEnumParamDefault    ShadowiguserinsightsMetricTypeEnumParam = "default"
	ShadowiguserinsightsMetricTypeEnumParamTimeSeries ShadowiguserinsightsMetricTypeEnumParam = "time_series"
	ShadowiguserinsightsMetricTypeEnumParamTotalValue ShadowiguserinsightsMetricTypeEnumParam = "total_value"
)

func (value ShadowiguserinsightsMetricTypeEnumParam) String() string {
	return string(value)
}
