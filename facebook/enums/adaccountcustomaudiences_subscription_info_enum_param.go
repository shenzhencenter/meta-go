package enums

type AdaccountcustomaudiencesSubscriptionInfoEnumParam string

const (
	AdaccountcustomaudiencesSubscriptionInfoEnumParamMessenger AdaccountcustomaudiencesSubscriptionInfoEnumParam = "MESSENGER"
	AdaccountcustomaudiencesSubscriptionInfoEnumParamWhatsapp  AdaccountcustomaudiencesSubscriptionInfoEnumParam = "WHATSAPP"
)

func (value AdaccountcustomaudiencesSubscriptionInfoEnumParam) String() string {
	return string(value)
}
