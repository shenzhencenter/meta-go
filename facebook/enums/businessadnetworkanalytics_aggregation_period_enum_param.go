package enums

type BusinessadnetworkanalyticsAggregationPeriodEnumParam string

const (
	BusinessadnetworkanalyticsAggregationPeriodEnumParamDay   BusinessadnetworkanalyticsAggregationPeriodEnumParam = "DAY"
	BusinessadnetworkanalyticsAggregationPeriodEnumParamTotal BusinessadnetworkanalyticsAggregationPeriodEnumParam = "TOTAL"
)

func (value BusinessadnetworkanalyticsAggregationPeriodEnumParam) String() string {
	return string(value)
}
