package enums

type ProductfeedDelimiter string

const (
	ProductfeedDelimiterAutodetect ProductfeedDelimiter = "AUTODETECT"
	ProductfeedDelimiterBar        ProductfeedDelimiter = "BAR"
	ProductfeedDelimiterComma      ProductfeedDelimiter = "COMMA"
	ProductfeedDelimiterSemicolon  ProductfeedDelimiter = "SEMICOLON"
	ProductfeedDelimiterTab        ProductfeedDelimiter = "TAB"
	ProductfeedDelimiterTilde      ProductfeedDelimiter = "TILDE"
)

func (value ProductfeedDelimiter) String() string {
	return string(value)
}
