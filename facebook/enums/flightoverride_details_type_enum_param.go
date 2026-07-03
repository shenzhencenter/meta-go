package enums

type FlightoverrideDetailsTypeEnumParam string

const (
	FlightoverrideDetailsTypeEnumParamCountry            FlightoverrideDetailsTypeEnumParam = "COUNTRY"
	FlightoverrideDetailsTypeEnumParamLanguage           FlightoverrideDetailsTypeEnumParam = "LANGUAGE"
	FlightoverrideDetailsTypeEnumParamLanguageAndCountry FlightoverrideDetailsTypeEnumParam = "LANGUAGE_AND_COUNTRY"
)

func (value FlightoverrideDetailsTypeEnumParam) String() string {
	return string(value)
}
