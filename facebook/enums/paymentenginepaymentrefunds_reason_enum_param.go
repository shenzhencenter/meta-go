package enums

type PaymentenginepaymentrefundsReasonEnumParam string

const (
	PaymentenginepaymentrefundsReasonEnumParamCustomerService PaymentenginepaymentrefundsReasonEnumParam = "CUSTOMER_SERVICE"
	PaymentenginepaymentrefundsReasonEnumParamFriendlyFraud   PaymentenginepaymentrefundsReasonEnumParam = "FRIENDLY_FRAUD"
	PaymentenginepaymentrefundsReasonEnumParamMaliciousFraud  PaymentenginepaymentrefundsReasonEnumParam = "MALICIOUS_FRAUD"
)

func (value PaymentenginepaymentrefundsReasonEnumParam) String() string {
	return string(value)
}
