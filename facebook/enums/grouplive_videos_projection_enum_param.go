package enums

type GroupliveVideosProjectionEnumParam string

const (
	GroupliveVideosProjectionEnumParamCubemap             GroupliveVideosProjectionEnumParam = "CUBEMAP"
	GroupliveVideosProjectionEnumParamEquirectangular     GroupliveVideosProjectionEnumParam = "EQUIRECTANGULAR"
	GroupliveVideosProjectionEnumParamHalfEquirectangular GroupliveVideosProjectionEnumParam = "HALF_EQUIRECTANGULAR"
)

func (value GroupliveVideosProjectionEnumParam) String() string {
	return string(value)
}
