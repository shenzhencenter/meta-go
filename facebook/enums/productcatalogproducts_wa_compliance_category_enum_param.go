package enums

type ProductcatalogproductsWaComplianceCategoryEnumParam string

const (
	ProductcatalogproductsWaComplianceCategoryEnumParamCountryOriginExempt ProductcatalogproductsWaComplianceCategoryEnumParam = "COUNTRY_ORIGIN_EXEMPT"
	ProductcatalogproductsWaComplianceCategoryEnumParamDefault             ProductcatalogproductsWaComplianceCategoryEnumParam = "DEFAULT"
)

func (value ProductcatalogproductsWaComplianceCategoryEnumParam) String() string {
	return string(value)
}
