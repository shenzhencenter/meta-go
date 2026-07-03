package enums

type PageliveVideosProjectionEnumParam string

const (
	PageliveVideosProjectionEnumParamCubemap             PageliveVideosProjectionEnumParam = "CUBEMAP"
	PageliveVideosProjectionEnumParamEquirectangular     PageliveVideosProjectionEnumParam = "EQUIRECTANGULAR"
	PageliveVideosProjectionEnumParamHalfEquirectangular PageliveVideosProjectionEnumParam = "HALF_EQUIRECTANGULAR"
)

func (value PageliveVideosProjectionEnumParam) String() string {
	return string(value)
}
