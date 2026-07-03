package enums

type PagemessagesFolderEnumParam string

const (
	PagemessagesFolderEnumParamPartnership PagemessagesFolderEnumParam = "PARTNERSHIP"
)

func (value PagemessagesFolderEnumParam) String() string {
	return string(value)
}
