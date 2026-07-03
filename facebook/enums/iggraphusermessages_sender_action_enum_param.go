package enums

type IggraphusermessagesSenderActionEnumParam string

const (
	IggraphusermessagesSenderActionEnumParamMarkSeen  IggraphusermessagesSenderActionEnumParam = "MARK_SEEN"
	IggraphusermessagesSenderActionEnumParamReact     IggraphusermessagesSenderActionEnumParam = "REACT"
	IggraphusermessagesSenderActionEnumParamTypingOff IggraphusermessagesSenderActionEnumParam = "TYPING_OFF"
	IggraphusermessagesSenderActionEnumParamTypingOn  IggraphusermessagesSenderActionEnumParam = "TYPING_ON"
	IggraphusermessagesSenderActionEnumParamUnreact   IggraphusermessagesSenderActionEnumParam = "UNREACT"
)

func (value IggraphusermessagesSenderActionEnumParam) String() string {
	return string(value)
}
