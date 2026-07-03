package enums

type TransactableitemoverrideDetailsTypeEnumParam string

const (
	TransactableitemoverrideDetailsTypeEnumParamCountry            TransactableitemoverrideDetailsTypeEnumParam = "COUNTRY"
	TransactableitemoverrideDetailsTypeEnumParamLanguage           TransactableitemoverrideDetailsTypeEnumParam = "LANGUAGE"
	TransactableitemoverrideDetailsTypeEnumParamLanguageAndCountry TransactableitemoverrideDetailsTypeEnumParam = "LANGUAGE_AND_COUNTRY"
)

func (value TransactableitemoverrideDetailsTypeEnumParam) String() string {
	return string(value)
}
