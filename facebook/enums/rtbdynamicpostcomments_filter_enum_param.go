package enums

type RtbdynamicpostcommentsFilterEnumParam string

const (
	RtbdynamicpostcommentsFilterEnumParamStream   RtbdynamicpostcommentsFilterEnumParam = "stream"
	RtbdynamicpostcommentsFilterEnumParamToplevel RtbdynamicpostcommentsFilterEnumParam = "toplevel"
)

func (value RtbdynamicpostcommentsFilterEnumParam) String() string {
	return string(value)
}
