package enums

type ProductsetproductsErrorPriorityEnumParam string

const (
	ProductsetproductsErrorPriorityEnumParamHigh   ProductsetproductsErrorPriorityEnumParam = "HIGH"
	ProductsetproductsErrorPriorityEnumParamLow    ProductsetproductsErrorPriorityEnumParam = "LOW"
	ProductsetproductsErrorPriorityEnumParamMedium ProductsetproductsErrorPriorityEnumParam = "MEDIUM"
)

func (value ProductsetproductsErrorPriorityEnumParam) String() string {
	return string(value)
}
