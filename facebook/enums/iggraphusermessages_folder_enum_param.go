package enums

type IggraphusermessagesFolderEnumParam string

const (
	IggraphusermessagesFolderEnumParamPartnership IggraphusermessagesFolderEnumParam = "PARTNERSHIP"
)

func (value IggraphusermessagesFolderEnumParam) String() string {
	return string(value)
}
