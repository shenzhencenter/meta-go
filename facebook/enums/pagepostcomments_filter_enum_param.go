package enums

type PagepostcommentsFilterEnumParam string

const (
	PagepostcommentsFilterEnumParamStream   PagepostcommentsFilterEnumParam = "stream"
	PagepostcommentsFilterEnumParamToplevel PagepostcommentsFilterEnumParam = "toplevel"
)

func (value PagepostcommentsFilterEnumParam) String() string {
	return string(value)
}
