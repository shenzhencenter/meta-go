package enums

type UserliveVideosProjectionEnumParam string

const (
	UserliveVideosProjectionEnumParamCubemap             UserliveVideosProjectionEnumParam = "CUBEMAP"
	UserliveVideosProjectionEnumParamEquirectangular     UserliveVideosProjectionEnumParam = "EQUIRECTANGULAR"
	UserliveVideosProjectionEnumParamHalfEquirectangular UserliveVideosProjectionEnumParam = "HALF_EQUIRECTANGULAR"
)

func (value UserliveVideosProjectionEnumParam) String() string {
	return string(value)
}
