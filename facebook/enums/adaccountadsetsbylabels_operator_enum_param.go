package enums

type AdaccountadsetsbylabelsOperatorEnumParam string

const (
	AdaccountadsetsbylabelsOperatorEnumParamAll AdaccountadsetsbylabelsOperatorEnumParam = "ALL"
	AdaccountadsetsbylabelsOperatorEnumParamAny AdaccountadsetsbylabelsOperatorEnumParam = "ANY"
)

func (value AdaccountadsetsbylabelsOperatorEnumParam) String() string {
	return string(value)
}
