package enums

type WhatsappbusinessaccountcallAnalyticsDirectionsEnumParam string

const (
	WhatsappbusinessaccountcallAnalyticsDirectionsEnumParamBusinessInitiated WhatsappbusinessaccountcallAnalyticsDirectionsEnumParam = "BUSINESS_INITIATED"
	WhatsappbusinessaccountcallAnalyticsDirectionsEnumParamUnknown           WhatsappbusinessaccountcallAnalyticsDirectionsEnumParam = "UNKNOWN"
	WhatsappbusinessaccountcallAnalyticsDirectionsEnumParamUserInitiated     WhatsappbusinessaccountcallAnalyticsDirectionsEnumParam = "USER_INITIATED"
)

func (value WhatsappbusinessaccountcallAnalyticsDirectionsEnumParam) String() string {
	return string(value)
}
