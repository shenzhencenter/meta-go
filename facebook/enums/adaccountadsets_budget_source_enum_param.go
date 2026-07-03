package enums

type AdaccountadsetsBudgetSourceEnumParam string

const (
	AdaccountadsetsBudgetSourceEnumParamNone AdaccountadsetsBudgetSourceEnumParam = "NONE"
	AdaccountadsetsBudgetSourceEnumParamRmn  AdaccountadsetsBudgetSourceEnumParam = "RMN"
)

func (value AdaccountadsetsBudgetSourceEnumParam) String() string {
	return string(value)
}
