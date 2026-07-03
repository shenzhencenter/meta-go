package enums

type UserfeedFormattingEnumParam string

const (
	UserfeedFormattingEnumParamMarkdown  UserfeedFormattingEnumParam = "MARKDOWN"
	UserfeedFormattingEnumParamPlaintext UserfeedFormattingEnumParam = "PLAINTEXT"
)

func (value UserfeedFormattingEnumParam) String() string {
	return string(value)
}
