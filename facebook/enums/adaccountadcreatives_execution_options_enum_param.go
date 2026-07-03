package enums

type AdaccountadcreativesExecutionOptionsEnumParam string

const (
	AdaccountadcreativesExecutionOptionsEnumParamValidateOnly AdaccountadcreativesExecutionOptionsEnumParam = "validate_only"
)

func (value AdaccountadcreativesExecutionOptionsEnumParam) String() string {
	return string(value)
}
