package enums

type AdaccountadcreativesCategorizationCriteriaEnumParam string

const (
	AdaccountadcreativesCategorizationCriteriaEnumParamBrand       AdaccountadcreativesCategorizationCriteriaEnumParam = "brand"
	AdaccountadcreativesCategorizationCriteriaEnumParamCategory    AdaccountadcreativesCategorizationCriteriaEnumParam = "category"
	AdaccountadcreativesCategorizationCriteriaEnumParamProductType AdaccountadcreativesCategorizationCriteriaEnumParam = "product_type"
)

func (value AdaccountadcreativesCategorizationCriteriaEnumParam) String() string {
	return string(value)
}
