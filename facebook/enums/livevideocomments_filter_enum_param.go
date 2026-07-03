package enums

type LivevideocommentsFilterEnumParam string

const (
	LivevideocommentsFilterEnumParamStream   LivevideocommentsFilterEnumParam = "stream"
	LivevideocommentsFilterEnumParamToplevel LivevideocommentsFilterEnumParam = "toplevel"
)

func (value LivevideocommentsFilterEnumParam) String() string {
	return string(value)
}
