package enums

type VehicleoverrideDetailsTypeEnumParam string

const (
	VehicleoverrideDetailsTypeEnumParamCountry            VehicleoverrideDetailsTypeEnumParam = "COUNTRY"
	VehicleoverrideDetailsTypeEnumParamLanguage           VehicleoverrideDetailsTypeEnumParam = "LANGUAGE"
	VehicleoverrideDetailsTypeEnumParamLanguageAndCountry VehicleoverrideDetailsTypeEnumParam = "LANGUAGE_AND_COUNTRY"
)

func (value VehicleoverrideDetailsTypeEnumParam) String() string {
	return string(value)
}
