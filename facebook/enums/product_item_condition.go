package enums

type ProductitemConditionEnum string

const (
	ProductitemConditionEnumCpo         ProductitemConditionEnum = "cpo"
	ProductitemConditionEnumNew         ProductitemConditionEnum = "new"
	ProductitemConditionEnumOpenBoxNew  ProductitemConditionEnum = "open_box_new"
	ProductitemConditionEnumRefurbished ProductitemConditionEnum = "refurbished"
	ProductitemConditionEnumUsed        ProductitemConditionEnum = "used"
	ProductitemConditionEnumUsedFair    ProductitemConditionEnum = "used_fair"
	ProductitemConditionEnumUsedGood    ProductitemConditionEnum = "used_good"
	ProductitemConditionEnumUsedLikeNew ProductitemConditionEnum = "used_like_new"
)

func (value ProductitemConditionEnum) String() string {
	return string(value)
}
