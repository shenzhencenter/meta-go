package enums

type BusinessvideosFormattingEnumParam string

const (
	BusinessvideosFormattingEnumParamMarkdown  BusinessvideosFormattingEnumParam = "MARKDOWN"
	BusinessvideosFormattingEnumParamPlaintext BusinessvideosFormattingEnumParam = "PLAINTEXT"
)

func (value BusinessvideosFormattingEnumParam) String() string {
	return string(value)
}
