package enums

type EventliveVideosProjectionEnumParam string

const (
	EventliveVideosProjectionEnumParamCubemap             EventliveVideosProjectionEnumParam = "CUBEMAP"
	EventliveVideosProjectionEnumParamEquirectangular     EventliveVideosProjectionEnumParam = "EQUIRECTANGULAR"
	EventliveVideosProjectionEnumParamHalfEquirectangular EventliveVideosProjectionEnumParam = "HALF_EQUIRECTANGULAR"
)

func (value EventliveVideosProjectionEnumParam) String() string {
	return string(value)
}
