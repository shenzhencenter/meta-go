package enums

type ApplicationadnetworkanalyticsOrderingColumnEnumParam string

const (
	ApplicationadnetworkanalyticsOrderingColumnEnumParamMetric ApplicationadnetworkanalyticsOrderingColumnEnumParam = "METRIC"
	ApplicationadnetworkanalyticsOrderingColumnEnumParamTime   ApplicationadnetworkanalyticsOrderingColumnEnumParam = "TIME"
	ApplicationadnetworkanalyticsOrderingColumnEnumParamValue  ApplicationadnetworkanalyticsOrderingColumnEnumParam = "VALUE"
)

func (value ApplicationadnetworkanalyticsOrderingColumnEnumParam) String() string {
	return string(value)
}
