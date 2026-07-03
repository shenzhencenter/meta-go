package enums

type ProductfeedproductsErrorPriorityEnumParam string

const (
	ProductfeedproductsErrorPriorityEnumParamHigh   ProductfeedproductsErrorPriorityEnumParam = "HIGH"
	ProductfeedproductsErrorPriorityEnumParamLow    ProductfeedproductsErrorPriorityEnumParam = "LOW"
	ProductfeedproductsErrorPriorityEnumParamMedium ProductfeedproductsErrorPriorityEnumParam = "MEDIUM"
)

func (value ProductfeedproductsErrorPriorityEnumParam) String() string {
	return string(value)
}
