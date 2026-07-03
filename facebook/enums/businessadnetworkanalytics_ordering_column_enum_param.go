package enums

type BusinessadnetworkanalyticsOrderingColumnEnumParam string

const (
	BusinessadnetworkanalyticsOrderingColumnEnumParamMetric BusinessadnetworkanalyticsOrderingColumnEnumParam = "METRIC"
	BusinessadnetworkanalyticsOrderingColumnEnumParamTime   BusinessadnetworkanalyticsOrderingColumnEnumParam = "TIME"
	BusinessadnetworkanalyticsOrderingColumnEnumParamValue  BusinessadnetworkanalyticsOrderingColumnEnumParam = "VALUE"
)

func (value BusinessadnetworkanalyticsOrderingColumnEnumParam) String() string {
	return string(value)
}
