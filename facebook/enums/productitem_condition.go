package enums

type ProductitemCondition string

const (
	ProductitemConditionCpo         ProductitemCondition = "cpo"
	ProductitemConditionNew         ProductitemCondition = "new"
	ProductitemConditionOpenBoxNew  ProductitemCondition = "open_box_new"
	ProductitemConditionRefurbished ProductitemCondition = "refurbished"
	ProductitemConditionUsed        ProductitemCondition = "used"
	ProductitemConditionUsedFair    ProductitemCondition = "used_fair"
	ProductitemConditionUsedGood    ProductitemCondition = "used_good"
	ProductitemConditionUsedLikeNew ProductitemCondition = "used_like_new"
)

func (value ProductitemCondition) String() string {
	return string(value)
}
