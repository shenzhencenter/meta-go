package enums

type CommercemerchantsettingsreturnsStatusesEnumParam string

const (
	CommercemerchantsettingsreturnsStatusesEnumParamApproved                CommercemerchantsettingsreturnsStatusesEnumParam = "APPROVED"
	CommercemerchantsettingsreturnsStatusesEnumParamDisapproved             CommercemerchantsettingsreturnsStatusesEnumParam = "DISAPPROVED"
	CommercemerchantsettingsreturnsStatusesEnumParamMerchantMarkedCompleted CommercemerchantsettingsreturnsStatusesEnumParam = "MERCHANT_MARKED_COMPLETED"
	CommercemerchantsettingsreturnsStatusesEnumParamRefunded                CommercemerchantsettingsreturnsStatusesEnumParam = "REFUNDED"
	CommercemerchantsettingsreturnsStatusesEnumParamRequested               CommercemerchantsettingsreturnsStatusesEnumParam = "REQUESTED"
)

func (value CommercemerchantsettingsreturnsStatusesEnumParam) String() string {
	return string(value)
}
