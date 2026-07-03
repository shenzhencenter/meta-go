package enums

type WhatsappbusinessaccountcallAnalyticsMetricTypesEnumParam string

const (
	WhatsappbusinessaccountcallAnalyticsMetricTypesEnumParamAverageDuration WhatsappbusinessaccountcallAnalyticsMetricTypesEnumParam = "AVERAGE_DURATION"
	WhatsappbusinessaccountcallAnalyticsMetricTypesEnumParamCost            WhatsappbusinessaccountcallAnalyticsMetricTypesEnumParam = "COST"
	WhatsappbusinessaccountcallAnalyticsMetricTypesEnumParamCount           WhatsappbusinessaccountcallAnalyticsMetricTypesEnumParam = "COUNT"
	WhatsappbusinessaccountcallAnalyticsMetricTypesEnumParamUnknown         WhatsappbusinessaccountcallAnalyticsMetricTypesEnumParam = "UNKNOWN"
)

func (value WhatsappbusinessaccountcallAnalyticsMetricTypesEnumParam) String() string {
	return string(value)
}
