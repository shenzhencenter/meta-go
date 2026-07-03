package enums

type AdaccountadrulesHistoryEvaluationTypeEnumParam string

const (
	AdaccountadrulesHistoryEvaluationTypeEnumParamSchedule AdaccountadrulesHistoryEvaluationTypeEnumParam = "SCHEDULE"
	AdaccountadrulesHistoryEvaluationTypeEnumParamTrigger  AdaccountadrulesHistoryEvaluationTypeEnumParam = "TRIGGER"
)

func (value AdaccountadrulesHistoryEvaluationTypeEnumParam) String() string {
	return string(value)
}
