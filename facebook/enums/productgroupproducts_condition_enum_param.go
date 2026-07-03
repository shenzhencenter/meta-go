package enums

type ProductgroupproductsConditionEnumParam string

const (
	ProductgroupproductsConditionEnumParamCpo         ProductgroupproductsConditionEnumParam = "cpo"
	ProductgroupproductsConditionEnumParamNew         ProductgroupproductsConditionEnumParam = "new"
	ProductgroupproductsConditionEnumParamOpenBoxNew  ProductgroupproductsConditionEnumParam = "open_box_new"
	ProductgroupproductsConditionEnumParamRefurbished ProductgroupproductsConditionEnumParam = "refurbished"
	ProductgroupproductsConditionEnumParamUsed        ProductgroupproductsConditionEnumParam = "used"
	ProductgroupproductsConditionEnumParamUsedFair    ProductgroupproductsConditionEnumParam = "used_fair"
	ProductgroupproductsConditionEnumParamUsedGood    ProductgroupproductsConditionEnumParam = "used_good"
	ProductgroupproductsConditionEnumParamUsedLikeNew ProductgroupproductsConditionEnumParam = "used_like_new"
)

func (value ProductgroupproductsConditionEnumParam) String() string {
	return string(value)
}
