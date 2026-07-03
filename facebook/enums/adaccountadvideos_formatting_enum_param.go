package enums

type AdaccountadvideosFormattingEnumParam string

const (
	AdaccountadvideosFormattingEnumParamMarkdown  AdaccountadvideosFormattingEnumParam = "MARKDOWN"
	AdaccountadvideosFormattingEnumParamPlaintext AdaccountadvideosFormattingEnumParam = "PLAINTEXT"
)

func (value AdaccountadvideosFormattingEnumParam) String() string {
	return string(value)
}
