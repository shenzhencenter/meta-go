package enums

type ProductfeedQuotedFieldsModeEnum string

const (
	ProductfeedQuotedFieldsModeEnumAutodetect ProductfeedQuotedFieldsModeEnum = "AUTODETECT"
	ProductfeedQuotedFieldsModeEnumOff        ProductfeedQuotedFieldsModeEnum = "OFF"
	ProductfeedQuotedFieldsModeEnumOn         ProductfeedQuotedFieldsModeEnum = "ON"
)

func (value ProductfeedQuotedFieldsModeEnum) String() string {
	return string(value)
}
