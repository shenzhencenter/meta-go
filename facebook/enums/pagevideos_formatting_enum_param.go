package enums

type PagevideosFormattingEnumParam string

const (
	PagevideosFormattingEnumParamMarkdown  PagevideosFormattingEnumParam = "MARKDOWN"
	PagevideosFormattingEnumParamPlaintext PagevideosFormattingEnumParam = "PLAINTEXT"
)

func (value PagevideosFormattingEnumParam) String() string {
	return string(value)
}
