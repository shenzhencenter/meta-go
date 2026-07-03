package enums

type AdBidType string

const (
	AdBidTypeAbsoluteOcpm AdBidType = "ABSOLUTE_OCPM"
	AdBidTypeCpa          AdBidType = "CPA"
	AdBidTypeCpc          AdBidType = "CPC"
	AdBidTypeCpm          AdBidType = "CPM"
	AdBidTypeMultiPremium AdBidType = "MULTI_PREMIUM"
)

func (value AdBidType) String() string {
	return string(value)
}
