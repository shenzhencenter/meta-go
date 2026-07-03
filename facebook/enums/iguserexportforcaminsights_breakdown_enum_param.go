package enums

type IguserexportforcaminsightsBreakdownEnumParam string

const (
	IguserexportforcaminsightsBreakdownEnumParamAge          IguserexportforcaminsightsBreakdownEnumParam = "AGE"
	IguserexportforcaminsightsBreakdownEnumParamFollowType   IguserexportforcaminsightsBreakdownEnumParam = "FOLLOW_TYPE"
	IguserexportforcaminsightsBreakdownEnumParamGender       IguserexportforcaminsightsBreakdownEnumParam = "GENDER"
	IguserexportforcaminsightsBreakdownEnumParamMediaType    IguserexportforcaminsightsBreakdownEnumParam = "MEDIA_TYPE"
	IguserexportforcaminsightsBreakdownEnumParamTopCities    IguserexportforcaminsightsBreakdownEnumParam = "TOP_CITIES"
	IguserexportforcaminsightsBreakdownEnumParamTopCountries IguserexportforcaminsightsBreakdownEnumParam = "TOP_COUNTRIES"
)

func (value IguserexportforcaminsightsBreakdownEnumParam) String() string {
	return string(value)
}
