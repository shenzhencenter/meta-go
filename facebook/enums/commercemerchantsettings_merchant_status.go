package enums

type CommercemerchantsettingsMerchantStatus string

const (
	CommercemerchantsettingsMerchantStatusEnabled            CommercemerchantsettingsMerchantStatus = "ENABLED"
	CommercemerchantsettingsMerchantStatusExternallyDisabled CommercemerchantsettingsMerchantStatus = "EXTERNALLY_DISABLED"
)

func (value CommercemerchantsettingsMerchantStatus) String() string {
	return string(value)
}
