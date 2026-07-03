package enums

type ProductfeedDelimiterEnum string

const (
	ProductfeedDelimiterEnumAutodetect ProductfeedDelimiterEnum = "AUTODETECT"
	ProductfeedDelimiterEnumBar        ProductfeedDelimiterEnum = "BAR"
	ProductfeedDelimiterEnumComma      ProductfeedDelimiterEnum = "COMMA"
	ProductfeedDelimiterEnumSemicolon  ProductfeedDelimiterEnum = "SEMICOLON"
	ProductfeedDelimiterEnumTab        ProductfeedDelimiterEnum = "TAB"
	ProductfeedDelimiterEnumTilde      ProductfeedDelimiterEnum = "TILDE"
)

func (value ProductfeedDelimiterEnum) String() string {
	return string(value)
}
