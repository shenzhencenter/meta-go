package enums

type LocalservicebusinessCondition string

const (
	LocalservicebusinessConditionPcCpo         LocalservicebusinessCondition = "PC_CPO"
	LocalservicebusinessConditionPcNew         LocalservicebusinessCondition = "PC_NEW"
	LocalservicebusinessConditionPcOpenBoxNew  LocalservicebusinessCondition = "PC_OPEN_BOX_NEW"
	LocalservicebusinessConditionPcRefurbished LocalservicebusinessCondition = "PC_REFURBISHED"
	LocalservicebusinessConditionPcUsed        LocalservicebusinessCondition = "PC_USED"
	LocalservicebusinessConditionPcUsedFair    LocalservicebusinessCondition = "PC_USED_FAIR"
	LocalservicebusinessConditionPcUsedGood    LocalservicebusinessCondition = "PC_USED_GOOD"
	LocalservicebusinessConditionPcUsedLikeNew LocalservicebusinessCondition = "PC_USED_LIKE_NEW"
)

func (value LocalservicebusinessCondition) String() string {
	return string(value)
}
