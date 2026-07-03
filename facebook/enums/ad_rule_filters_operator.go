package enums

type AdrulefiltersOperator string

const (
	AdrulefiltersOperatorAll         AdrulefiltersOperator = "ALL"
	AdrulefiltersOperatorAny         AdrulefiltersOperator = "ANY"
	AdrulefiltersOperatorContain     AdrulefiltersOperator = "CONTAIN"
	AdrulefiltersOperatorEqual       AdrulefiltersOperator = "EQUAL"
	AdrulefiltersOperatorGreaterThan AdrulefiltersOperator = "GREATER_THAN"
	AdrulefiltersOperatorIn          AdrulefiltersOperator = "IN"
	AdrulefiltersOperatorInRange     AdrulefiltersOperator = "IN_RANGE"
	AdrulefiltersOperatorLessThan    AdrulefiltersOperator = "LESS_THAN"
	AdrulefiltersOperatorNone        AdrulefiltersOperator = "NONE"
	AdrulefiltersOperatorNotContain  AdrulefiltersOperator = "NOT_CONTAIN"
	AdrulefiltersOperatorNotEqual    AdrulefiltersOperator = "NOT_EQUAL"
	AdrulefiltersOperatorNotIn       AdrulefiltersOperator = "NOT_IN"
	AdrulefiltersOperatorNotInRange  AdrulefiltersOperator = "NOT_IN_RANGE"
)

func (value AdrulefiltersOperator) String() string {
	return string(value)
}
