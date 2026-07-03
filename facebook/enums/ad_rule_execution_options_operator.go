package enums

type AdruleexecutionoptionsOperator string

const (
	AdruleexecutionoptionsOperatorEqual AdruleexecutionoptionsOperator = "EQUAL"
	AdruleexecutionoptionsOperatorIn    AdruleexecutionoptionsOperator = "IN"
)

func (value AdruleexecutionoptionsOperator) String() string {
	return string(value)
}
