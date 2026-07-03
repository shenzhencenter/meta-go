package enums

type HighdemandperiodBudgetValueType string

const (
	HighdemandperiodBudgetValueTypeAbsolute   HighdemandperiodBudgetValueType = "ABSOLUTE"
	HighdemandperiodBudgetValueTypeMultiplier HighdemandperiodBudgetValueType = "MULTIPLIER"
)

func (value HighdemandperiodBudgetValueType) String() string {
	return string(value)
}
