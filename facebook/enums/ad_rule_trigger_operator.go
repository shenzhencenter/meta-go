package enums

type AdruletriggerOperator string

const (
	AdruletriggerOperatorAll         AdruletriggerOperator = "ALL"
	AdruletriggerOperatorAny         AdruletriggerOperator = "ANY"
	AdruletriggerOperatorContain     AdruletriggerOperator = "CONTAIN"
	AdruletriggerOperatorEqual       AdruletriggerOperator = "EQUAL"
	AdruletriggerOperatorGreaterThan AdruletriggerOperator = "GREATER_THAN"
	AdruletriggerOperatorIn          AdruletriggerOperator = "IN"
	AdruletriggerOperatorInRange     AdruletriggerOperator = "IN_RANGE"
	AdruletriggerOperatorLessThan    AdruletriggerOperator = "LESS_THAN"
	AdruletriggerOperatorNone        AdruletriggerOperator = "NONE"
	AdruletriggerOperatorNotContain  AdruletriggerOperator = "NOT_CONTAIN"
	AdruletriggerOperatorNotEqual    AdruletriggerOperator = "NOT_EQUAL"
	AdruletriggerOperatorNotIn       AdruletriggerOperator = "NOT_IN"
	AdruletriggerOperatorNotInRange  AdruletriggerOperator = "NOT_IN_RANGE"
)

func (value AdruletriggerOperator) String() string {
	return string(value)
}
