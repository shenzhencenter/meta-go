package enums

type UservideosFormattingEnumParam string

const (
	UservideosFormattingEnumParamMarkdown  UservideosFormattingEnumParam = "MARKDOWN"
	UservideosFormattingEnumParamPlaintext UservideosFormattingEnumParam = "PLAINTEXT"
)

func (value UservideosFormattingEnumParam) String() string {
	return string(value)
}
