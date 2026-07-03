package enums

type WhatsappbusinessaccountpricingAnalyticsDimensionsEnumParam string

const (
	WhatsappbusinessaccountpricingAnalyticsDimensionsEnumParamCountry         WhatsappbusinessaccountpricingAnalyticsDimensionsEnumParam = "COUNTRY"
	WhatsappbusinessaccountpricingAnalyticsDimensionsEnumParamPhone           WhatsappbusinessaccountpricingAnalyticsDimensionsEnumParam = "PHONE"
	WhatsappbusinessaccountpricingAnalyticsDimensionsEnumParamPricingCategory WhatsappbusinessaccountpricingAnalyticsDimensionsEnumParam = "PRICING_CATEGORY"
	WhatsappbusinessaccountpricingAnalyticsDimensionsEnumParamPricingType     WhatsappbusinessaccountpricingAnalyticsDimensionsEnumParam = "PRICING_TYPE"
	WhatsappbusinessaccountpricingAnalyticsDimensionsEnumParamTier            WhatsappbusinessaccountpricingAnalyticsDimensionsEnumParam = "TIER"
)

func (value WhatsappbusinessaccountpricingAnalyticsDimensionsEnumParam) String() string {
	return string(value)
}
