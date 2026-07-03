package enums

type PagefeedFormattingEnumParam string

const (
	PagefeedFormattingEnumParamMarkdown  PagefeedFormattingEnumParam = "MARKDOWN"
	PagefeedFormattingEnumParamPlaintext PagefeedFormattingEnumParam = "PLAINTEXT"
)

func (value PagefeedFormattingEnumParam) String() string {
	return string(value)
}
