package enums

type ProductcatalogcategoriesCategorizationCriteriaEnumParam string

const (
	ProductcatalogcategoriesCategorizationCriteriaEnumParamBrand       ProductcatalogcategoriesCategorizationCriteriaEnumParam = "BRAND"
	ProductcatalogcategoriesCategorizationCriteriaEnumParamCategory    ProductcatalogcategoriesCategorizationCriteriaEnumParam = "CATEGORY"
	ProductcatalogcategoriesCategorizationCriteriaEnumParamProductType ProductcatalogcategoriesCategorizationCriteriaEnumParam = "PRODUCT_TYPE"
)

func (value ProductcatalogcategoriesCategorizationCriteriaEnumParam) String() string {
	return string(value)
}
