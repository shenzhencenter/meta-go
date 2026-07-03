package enums

type PagemessageTemplatesParameterFormatEnumParam string

const (
	PagemessageTemplatesParameterFormatEnumParamNamed      PagemessageTemplatesParameterFormatEnumParam = "NAMED"
	PagemessageTemplatesParameterFormatEnumParamPositional PagemessageTemplatesParameterFormatEnumParam = "POSITIONAL"
)

func (value PagemessageTemplatesParameterFormatEnumParam) String() string {
	return string(value)
}
