package enums

type GroupfeedFormattingEnumParam string

const (
	GroupfeedFormattingEnumParamMarkdown  GroupfeedFormattingEnumParam = "MARKDOWN"
	GroupfeedFormattingEnumParamPlaintext GroupfeedFormattingEnumParam = "PLAINTEXT"
)

func (value GroupfeedFormattingEnumParam) String() string {
	return string(value)
}
