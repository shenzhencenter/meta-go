package enums

type HomelistingoverrideDetailsTypeEnumParam string

const (
	HomelistingoverrideDetailsTypeEnumParamCountry            HomelistingoverrideDetailsTypeEnumParam = "COUNTRY"
	HomelistingoverrideDetailsTypeEnumParamLanguage           HomelistingoverrideDetailsTypeEnumParam = "LANGUAGE"
	HomelistingoverrideDetailsTypeEnumParamLanguageAndCountry HomelistingoverrideDetailsTypeEnumParam = "LANGUAGE_AND_COUNTRY"
)

func (value HomelistingoverrideDetailsTypeEnumParam) String() string {
	return string(value)
}
