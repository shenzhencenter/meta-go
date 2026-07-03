package enums

type VideocommentsFilterEnumParam string

const (
	VideocommentsFilterEnumParamStream   VideocommentsFilterEnumParam = "stream"
	VideocommentsFilterEnumParamToplevel VideocommentsFilterEnumParam = "toplevel"
)

func (value VideocommentsFilterEnumParam) String() string {
	return string(value)
}
