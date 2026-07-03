package enums

type ProductcatalogproductsConditionEnumParam string

const (
	ProductcatalogproductsConditionEnumParamCpo         ProductcatalogproductsConditionEnumParam = "cpo"
	ProductcatalogproductsConditionEnumParamNew         ProductcatalogproductsConditionEnumParam = "new"
	ProductcatalogproductsConditionEnumParamOpenBoxNew  ProductcatalogproductsConditionEnumParam = "open_box_new"
	ProductcatalogproductsConditionEnumParamRefurbished ProductcatalogproductsConditionEnumParam = "refurbished"
	ProductcatalogproductsConditionEnumParamUsed        ProductcatalogproductsConditionEnumParam = "used"
	ProductcatalogproductsConditionEnumParamUsedFair    ProductcatalogproductsConditionEnumParam = "used_fair"
	ProductcatalogproductsConditionEnumParamUsedGood    ProductcatalogproductsConditionEnumParam = "used_good"
	ProductcatalogproductsConditionEnumParamUsedLikeNew ProductcatalogproductsConditionEnumParam = "used_like_new"
)

func (value ProductcatalogproductsConditionEnumParam) String() string {
	return string(value)
}
