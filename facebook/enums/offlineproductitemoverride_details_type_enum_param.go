package enums

type OfflineproductitemoverrideDetailsTypeEnumParam string

const (
	OfflineproductitemoverrideDetailsTypeEnumParamCountry            OfflineproductitemoverrideDetailsTypeEnumParam = "COUNTRY"
	OfflineproductitemoverrideDetailsTypeEnumParamLanguage           OfflineproductitemoverrideDetailsTypeEnumParam = "LANGUAGE"
	OfflineproductitemoverrideDetailsTypeEnumParamLanguageAndCountry OfflineproductitemoverrideDetailsTypeEnumParam = "LANGUAGE_AND_COUNTRY"
)

func (value OfflineproductitemoverrideDetailsTypeEnumParam) String() string {
	return string(value)
}
