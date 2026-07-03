package enums

type HoteloverrideDetailsTypeEnumParam string

const (
	HoteloverrideDetailsTypeEnumParamCountry            HoteloverrideDetailsTypeEnumParam = "COUNTRY"
	HoteloverrideDetailsTypeEnumParamLanguage           HoteloverrideDetailsTypeEnumParam = "LANGUAGE"
	HoteloverrideDetailsTypeEnumParamLanguageAndCountry HoteloverrideDetailsTypeEnumParam = "LANGUAGE_AND_COUNTRY"
)

func (value HoteloverrideDetailsTypeEnumParam) String() string {
	return string(value)
}
