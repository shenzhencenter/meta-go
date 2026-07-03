package enums

type AdaccountadcreativesbylabelsOperatorEnumParam string

const (
	AdaccountadcreativesbylabelsOperatorEnumParamAll AdaccountadcreativesbylabelsOperatorEnumParam = "ALL"
	AdaccountadcreativesbylabelsOperatorEnumParamAny AdaccountadcreativesbylabelsOperatorEnumParam = "ANY"
)

func (value AdaccountadcreativesbylabelsOperatorEnumParam) String() string {
	return string(value)
}
