package enums

type AdaccountadsetsAttributionCountTypeEnumParam string

const (
	AdaccountadsetsAttributionCountTypeEnumParamAllConversions  AdaccountadsetsAttributionCountTypeEnumParam = "ALL_CONVERSIONS"
	AdaccountadsetsAttributionCountTypeEnumParamFirstConversion AdaccountadsetsAttributionCountTypeEnumParam = "FIRST_CONVERSION"
)

func (value AdaccountadsetsAttributionCountTypeEnumParam) String() string {
	return string(value)
}
