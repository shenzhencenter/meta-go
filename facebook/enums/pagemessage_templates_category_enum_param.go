package enums

type PagemessageTemplatesCategoryEnumParam string

const (
	PagemessageTemplatesCategoryEnumParamUtility PagemessageTemplatesCategoryEnumParam = "UTILITY"
)

func (value PagemessageTemplatesCategoryEnumParam) String() string {
	return string(value)
}
