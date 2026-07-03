package enums

type WhatsappbusinessaccountpricingAnalyticsPricingTypesEnumParam string

const (
	WhatsappbusinessaccountpricingAnalyticsPricingTypesEnumParamFreeCustomerService      WhatsappbusinessaccountpricingAnalyticsPricingTypesEnumParam = "FREE_CUSTOMER_SERVICE"
	WhatsappbusinessaccountpricingAnalyticsPricingTypesEnumParamFreeEntryPoint           WhatsappbusinessaccountpricingAnalyticsPricingTypesEnumParam = "FREE_ENTRY_POINT"
	WhatsappbusinessaccountpricingAnalyticsPricingTypesEnumParamFreeGroupCustomerService WhatsappbusinessaccountpricingAnalyticsPricingTypesEnumParam = "FREE_GROUP_CUSTOMER_SERVICE"
	WhatsappbusinessaccountpricingAnalyticsPricingTypesEnumParamRegular                  WhatsappbusinessaccountpricingAnalyticsPricingTypesEnumParam = "REGULAR"
)

func (value WhatsappbusinessaccountpricingAnalyticsPricingTypesEnumParam) String() string {
	return string(value)
}
