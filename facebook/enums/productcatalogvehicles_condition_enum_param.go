package enums

type ProductcatalogvehiclesConditionEnumParam string

const (
	ProductcatalogvehiclesConditionEnumParamExcellent ProductcatalogvehiclesConditionEnumParam = "EXCELLENT"
	ProductcatalogvehiclesConditionEnumParamFair      ProductcatalogvehiclesConditionEnumParam = "FAIR"
	ProductcatalogvehiclesConditionEnumParamGood      ProductcatalogvehiclesConditionEnumParam = "GOOD"
	ProductcatalogvehiclesConditionEnumParamNone      ProductcatalogvehiclesConditionEnumParam = "NONE"
	ProductcatalogvehiclesConditionEnumParamOther     ProductcatalogvehiclesConditionEnumParam = "OTHER"
	ProductcatalogvehiclesConditionEnumParamPoor      ProductcatalogvehiclesConditionEnumParam = "POOR"
	ProductcatalogvehiclesConditionEnumParamVeryGood  ProductcatalogvehiclesConditionEnumParam = "VERY_GOOD"
)

func (value ProductcatalogvehiclesConditionEnumParam) String() string {
	return string(value)
}
