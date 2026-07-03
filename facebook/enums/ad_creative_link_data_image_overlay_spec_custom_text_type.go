package enums

type AdcreativelinkdataimageoverlayspecCustomTextType string

const (
	AdcreativelinkdataimageoverlayspecCustomTextTypeFreeShipping AdcreativelinkdataimageoverlayspecCustomTextType = "free_shipping"
	AdcreativelinkdataimageoverlayspecCustomTextTypePopular      AdcreativelinkdataimageoverlayspecCustomTextType = "popular"
	AdcreativelinkdataimageoverlayspecCustomTextTypeSale         AdcreativelinkdataimageoverlayspecCustomTextType = "sale"
)

func (value AdcreativelinkdataimageoverlayspecCustomTextType) String() string {
	return string(value)
}
