package enums

type ProductcatalogproductFeedsQuotedFieldsModeEnumParam string

const (
	ProductcatalogproductFeedsQuotedFieldsModeEnumParamAutodetect ProductcatalogproductFeedsQuotedFieldsModeEnumParam = "autodetect"
	ProductcatalogproductFeedsQuotedFieldsModeEnumParamOff        ProductcatalogproductFeedsQuotedFieldsModeEnumParam = "off"
	ProductcatalogproductFeedsQuotedFieldsModeEnumParamOn         ProductcatalogproductFeedsQuotedFieldsModeEnumParam = "on"
)

func (value ProductcatalogproductFeedsQuotedFieldsModeEnumParam) String() string {
	return string(value)
}
