package enums

type GroupvideosFormattingEnumParam string

const (
	GroupvideosFormattingEnumParamMarkdown  GroupvideosFormattingEnumParam = "MARKDOWN"
	GroupvideosFormattingEnumParamPlaintext GroupvideosFormattingEnumParam = "PLAINTEXT"
)

func (value GroupvideosFormattingEnumParam) String() string {
	return string(value)
}
