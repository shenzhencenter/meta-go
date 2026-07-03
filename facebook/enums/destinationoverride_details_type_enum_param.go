package enums

type DestinationoverrideDetailsTypeEnumParam string

const (
	DestinationoverrideDetailsTypeEnumParamCountry            DestinationoverrideDetailsTypeEnumParam = "COUNTRY"
	DestinationoverrideDetailsTypeEnumParamLanguage           DestinationoverrideDetailsTypeEnumParam = "LANGUAGE"
	DestinationoverrideDetailsTypeEnumParamLanguageAndCountry DestinationoverrideDetailsTypeEnumParam = "LANGUAGE_AND_COUNTRY"
)

func (value DestinationoverrideDetailsTypeEnumParam) String() string {
	return string(value)
}
