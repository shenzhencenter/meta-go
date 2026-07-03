package enums

type AdaccountcustomaudiencesCustomerFileSourceEnumParam string

const (
	AdaccountcustomaudiencesCustomerFileSourceEnumParamBothUserAndPartnerProvided AdaccountcustomaudiencesCustomerFileSourceEnumParam = "BOTH_USER_AND_PARTNER_PROVIDED"
	AdaccountcustomaudiencesCustomerFileSourceEnumParamPartnerProvidedOnly        AdaccountcustomaudiencesCustomerFileSourceEnumParam = "PARTNER_PROVIDED_ONLY"
	AdaccountcustomaudiencesCustomerFileSourceEnumParamUserProvidedOnly           AdaccountcustomaudiencesCustomerFileSourceEnumParam = "USER_PROVIDED_ONLY"
)

func (value AdaccountcustomaudiencesCustomerFileSourceEnumParam) String() string {
	return string(value)
}
