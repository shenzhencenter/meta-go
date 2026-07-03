package enums

type ProductfeedQuotedFieldsMode string

const (
	ProductfeedQuotedFieldsModeAutodetect ProductfeedQuotedFieldsMode = "autodetect"
	ProductfeedQuotedFieldsModeOff        ProductfeedQuotedFieldsMode = "off"
	ProductfeedQuotedFieldsModeOn         ProductfeedQuotedFieldsMode = "on"
)

func (value ProductfeedQuotedFieldsMode) String() string {
	return string(value)
}
