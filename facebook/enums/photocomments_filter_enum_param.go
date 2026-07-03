package enums

type PhotocommentsFilterEnumParam string

const (
	PhotocommentsFilterEnumParamStream   PhotocommentsFilterEnumParam = "stream"
	PhotocommentsFilterEnumParamToplevel PhotocommentsFilterEnumParam = "toplevel"
)

func (value PhotocommentsFilterEnumParam) String() string {
	return string(value)
}
