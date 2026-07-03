package enums

type VehicleofferoverrideDetailsTypeEnumParam string

const (
	VehicleofferoverrideDetailsTypeEnumParamCountry            VehicleofferoverrideDetailsTypeEnumParam = "COUNTRY"
	VehicleofferoverrideDetailsTypeEnumParamLanguage           VehicleofferoverrideDetailsTypeEnumParam = "LANGUAGE"
	VehicleofferoverrideDetailsTypeEnumParamLanguageAndCountry VehicleofferoverrideDetailsTypeEnumParam = "LANGUAGE_AND_COUNTRY"
)

func (value VehicleofferoverrideDetailsTypeEnumParam) String() string {
	return string(value)
}
