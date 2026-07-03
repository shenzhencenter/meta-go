package enums

type PostcommentsFilterEnumParam string

const (
	PostcommentsFilterEnumParamStream   PostcommentsFilterEnumParam = "stream"
	PostcommentsFilterEnumParamToplevel PostcommentsFilterEnumParam = "toplevel"
)

func (value PostcommentsFilterEnumParam) String() string {
	return string(value)
}
