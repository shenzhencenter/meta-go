package enums

type LocalservicebusinessoverrideDetailsTypeEnumParam string

const (
	LocalservicebusinessoverrideDetailsTypeEnumParamCountry            LocalservicebusinessoverrideDetailsTypeEnumParam = "COUNTRY"
	LocalservicebusinessoverrideDetailsTypeEnumParamLanguage           LocalservicebusinessoverrideDetailsTypeEnumParam = "LANGUAGE"
	LocalservicebusinessoverrideDetailsTypeEnumParamLanguageAndCountry LocalservicebusinessoverrideDetailsTypeEnumParam = "LANGUAGE_AND_COUNTRY"
)

func (value LocalservicebusinessoverrideDetailsTypeEnumParam) String() string {
	return string(value)
}
