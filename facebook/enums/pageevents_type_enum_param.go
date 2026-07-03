package enums

type PageeventsTypeEnumParam string

const (
	PageeventsTypeEnumParamAttending  PageeventsTypeEnumParam = "attending"
	PageeventsTypeEnumParamCreated    PageeventsTypeEnumParam = "created"
	PageeventsTypeEnumParamDeclined   PageeventsTypeEnumParam = "declined"
	PageeventsTypeEnumParamMaybe      PageeventsTypeEnumParam = "maybe"
	PageeventsTypeEnumParamNotReplied PageeventsTypeEnumParam = "not_replied"
)

func (value PageeventsTypeEnumParam) String() string {
	return string(value)
}
