package enums

type UsereventsTypeEnumParam string

const (
	UsereventsTypeEnumParamAttending  UsereventsTypeEnumParam = "attending"
	UsereventsTypeEnumParamCreated    UsereventsTypeEnumParam = "created"
	UsereventsTypeEnumParamDeclined   UsereventsTypeEnumParam = "declined"
	UsereventsTypeEnumParamMaybe      UsereventsTypeEnumParam = "maybe"
	UsereventsTypeEnumParamNotReplied UsereventsTypeEnumParam = "not_replied"
)

func (value UsereventsTypeEnumParam) String() string {
	return string(value)
}
