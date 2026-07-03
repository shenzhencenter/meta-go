package enums

type AlbumcommentsFilterEnumParam string

const (
	AlbumcommentsFilterEnumParamStream   AlbumcommentsFilterEnumParam = "stream"
	AlbumcommentsFilterEnumParamToplevel AlbumcommentsFilterEnumParam = "toplevel"
)

func (value AlbumcommentsFilterEnumParam) String() string {
	return string(value)
}
