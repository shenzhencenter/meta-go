package enums

type AdaccountadsbylabelsOperatorEnumParam string

const (
	AdaccountadsbylabelsOperatorEnumParamAll AdaccountadsbylabelsOperatorEnumParam = "ALL"
	AdaccountadsbylabelsOperatorEnumParamAny AdaccountadsbylabelsOperatorEnumParam = "ANY"
)

func (value AdaccountadsbylabelsOperatorEnumParam) String() string {
	return string(value)
}
