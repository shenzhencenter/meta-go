package enums

type PaymentenginepaymentdisputeReasonEnumParam string

const (
	PaymentenginepaymentdisputeReasonEnumParamBannedUser             PaymentenginepaymentdisputeReasonEnumParam = "BANNED_USER"
	PaymentenginepaymentdisputeReasonEnumParamDeniedRefund           PaymentenginepaymentdisputeReasonEnumParam = "DENIED_REFUND"
	PaymentenginepaymentdisputeReasonEnumParamGrantedReplacementItem PaymentenginepaymentdisputeReasonEnumParam = "GRANTED_REPLACEMENT_ITEM"
)

func (value PaymentenginepaymentdisputeReasonEnumParam) String() string {
	return string(value)
}
