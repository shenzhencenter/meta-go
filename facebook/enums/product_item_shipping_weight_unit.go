package enums

type ProductitemShippingWeightUnit string

const (
	ProductitemShippingWeightUnitG  ProductitemShippingWeightUnit = "g"
	ProductitemShippingWeightUnitKg ProductitemShippingWeightUnit = "kg"
	ProductitemShippingWeightUnitLb ProductitemShippingWeightUnit = "lb"
	ProductitemShippingWeightUnitOz ProductitemShippingWeightUnit = "oz"
)

func (value ProductitemShippingWeightUnit) String() string {
	return string(value)
}
