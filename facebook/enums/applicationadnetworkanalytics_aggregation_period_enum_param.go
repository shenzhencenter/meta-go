package enums

type ApplicationadnetworkanalyticsAggregationPeriodEnumParam string

const (
	ApplicationadnetworkanalyticsAggregationPeriodEnumParamDay   ApplicationadnetworkanalyticsAggregationPeriodEnumParam = "DAY"
	ApplicationadnetworkanalyticsAggregationPeriodEnumParamTotal ApplicationadnetworkanalyticsAggregationPeriodEnumParam = "TOTAL"
)

func (value ApplicationadnetworkanalyticsAggregationPeriodEnumParam) String() string {
	return string(value)
}
