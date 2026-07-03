package enums

type ProductfeeduploaderrorsErrorPriorityEnumParam string

const (
	ProductfeeduploaderrorsErrorPriorityEnumParamHigh   ProductfeeduploaderrorsErrorPriorityEnumParam = "HIGH"
	ProductfeeduploaderrorsErrorPriorityEnumParamLow    ProductfeeduploaderrorsErrorPriorityEnumParam = "LOW"
	ProductfeeduploaderrorsErrorPriorityEnumParamMedium ProductfeeduploaderrorsErrorPriorityEnumParam = "MEDIUM"
)

func (value ProductfeeduploaderrorsErrorPriorityEnumParam) String() string {
	return string(value)
}
