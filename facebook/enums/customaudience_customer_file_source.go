package enums

type CustomaudienceCustomerFileSource string

const (
	CustomaudienceCustomerFileSourceBothUserAndPartnerProvided CustomaudienceCustomerFileSource = "BOTH_USER_AND_PARTNER_PROVIDED"
	CustomaudienceCustomerFileSourcePartnerProvidedOnly        CustomaudienceCustomerFileSource = "PARTNER_PROVIDED_ONLY"
	CustomaudienceCustomerFileSourceUserProvidedOnly           CustomaudienceCustomerFileSource = "USER_PROVIDED_ONLY"
)

func (value CustomaudienceCustomerFileSource) String() string {
	return string(value)
}
