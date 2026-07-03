package enums

type AdruleevaluationspecEvaluationType string

const (
	AdruleevaluationspecEvaluationTypeSchedule AdruleevaluationspecEvaluationType = "SCHEDULE"
	AdruleevaluationspecEvaluationTypeTrigger  AdruleevaluationspecEvaluationType = "TRIGGER"
)

func (value AdruleevaluationspecEvaluationType) String() string {
	return string(value)
}
