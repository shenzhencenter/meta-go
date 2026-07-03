package enums

type ProductcatalogproductsErrorPriorityEnumParam string

const (
	ProductcatalogproductsErrorPriorityEnumParamHigh   ProductcatalogproductsErrorPriorityEnumParam = "HIGH"
	ProductcatalogproductsErrorPriorityEnumParamLow    ProductcatalogproductsErrorPriorityEnumParam = "LOW"
	ProductcatalogproductsErrorPriorityEnumParamMedium ProductcatalogproductsErrorPriorityEnumParam = "MEDIUM"
)

func (value ProductcatalogproductsErrorPriorityEnumParam) String() string {
	return string(value)
}
