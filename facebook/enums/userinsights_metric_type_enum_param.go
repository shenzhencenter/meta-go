package enums

type UserinsightsMetricTypeEnumParam string

const (
	UserinsightsMetricTypeEnumParamDefault    UserinsightsMetricTypeEnumParam = "default"
	UserinsightsMetricTypeEnumParamTimeSeries UserinsightsMetricTypeEnumParam = "time_series"
	UserinsightsMetricTypeEnumParamTotalValue UserinsightsMetricTypeEnumParam = "total_value"
)

func (value UserinsightsMetricTypeEnumParam) String() string {
	return string(value)
}
