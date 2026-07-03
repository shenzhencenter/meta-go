package enums

type TransactionProductType string

const (
	TransactionProductTypeCpReturnLabel TransactionProductType = "cp_return_label"
	TransactionProductTypeFacebookAd    TransactionProductType = "facebook_ad"
	TransactionProductTypeIgAd          TransactionProductType = "ig_ad"
	TransactionProductTypeWhatsapp      TransactionProductType = "whatsapp"
	TransactionProductTypeWorkplace     TransactionProductType = "workplace"
)

func (value TransactionProductType) String() string {
	return string(value)
}
