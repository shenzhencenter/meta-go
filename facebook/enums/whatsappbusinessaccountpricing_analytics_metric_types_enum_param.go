package enums

type WhatsappbusinessaccountpricingAnalyticsMetricTypesEnumParam string

const (
	WhatsappbusinessaccountpricingAnalyticsMetricTypesEnumParamCost   WhatsappbusinessaccountpricingAnalyticsMetricTypesEnumParam = "COST"
	WhatsappbusinessaccountpricingAnalyticsMetricTypesEnumParamVolume WhatsappbusinessaccountpricingAnalyticsMetricTypesEnumParam = "VOLUME"
)

func (value WhatsappbusinessaccountpricingAnalyticsMetricTypesEnumParam) String() string {
	return string(value)
}
