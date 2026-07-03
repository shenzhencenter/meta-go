package enums

type PagemessagesSenderActionEnumParam string

const (
	PagemessagesSenderActionEnumParamMarkSeen  PagemessagesSenderActionEnumParam = "MARK_SEEN"
	PagemessagesSenderActionEnumParamReact     PagemessagesSenderActionEnumParam = "REACT"
	PagemessagesSenderActionEnumParamTypingOff PagemessagesSenderActionEnumParam = "TYPING_OFF"
	PagemessagesSenderActionEnumParamTypingOn  PagemessagesSenderActionEnumParam = "TYPING_ON"
	PagemessagesSenderActionEnumParamUnreact   PagemessagesSenderActionEnumParam = "UNREACT"
)

func (value PagemessagesSenderActionEnumParam) String() string {
	return string(value)
}
