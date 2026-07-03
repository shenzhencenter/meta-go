package enums

type AutomotivemodeloverrideDetailsTypeEnumParam string

const (
	AutomotivemodeloverrideDetailsTypeEnumParamCountry            AutomotivemodeloverrideDetailsTypeEnumParam = "COUNTRY"
	AutomotivemodeloverrideDetailsTypeEnumParamLanguage           AutomotivemodeloverrideDetailsTypeEnumParam = "LANGUAGE"
	AutomotivemodeloverrideDetailsTypeEnumParamLanguageAndCountry AutomotivemodeloverrideDetailsTypeEnumParam = "LANGUAGE_AND_COUNTRY"
)

func (value AutomotivemodeloverrideDetailsTypeEnumParam) String() string {
	return string(value)
}
