package enums

type ProductitemoverrideDetailsTypeEnumParam string

const (
	ProductitemoverrideDetailsTypeEnumParamCountry            ProductitemoverrideDetailsTypeEnumParam = "COUNTRY"
	ProductitemoverrideDetailsTypeEnumParamLanguage           ProductitemoverrideDetailsTypeEnumParam = "LANGUAGE"
	ProductitemoverrideDetailsTypeEnumParamLanguageAndCountry ProductitemoverrideDetailsTypeEnumParam = "LANGUAGE_AND_COUNTRY"
)

func (value ProductitemoverrideDetailsTypeEnumParam) String() string {
	return string(value)
}
