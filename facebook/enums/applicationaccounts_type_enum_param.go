package enums

type ApplicationaccountsTypeEnumParam string

const (
	ApplicationaccountsTypeEnumParamTestUsers ApplicationaccountsTypeEnumParam = "test-users"
)

func (value ApplicationaccountsTypeEnumParam) String() string {
	return string(value)
}
