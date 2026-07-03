package enums

type GroupvideosOriginalProjectionTypeEnumParam string

const (
	GroupvideosOriginalProjectionTypeEnumParamCubemap             GroupvideosOriginalProjectionTypeEnumParam = "cubemap"
	GroupvideosOriginalProjectionTypeEnumParamEquirectangular     GroupvideosOriginalProjectionTypeEnumParam = "equirectangular"
	GroupvideosOriginalProjectionTypeEnumParamHalfEquirectangular GroupvideosOriginalProjectionTypeEnumParam = "half_equirectangular"
)

func (value GroupvideosOriginalProjectionTypeEnumParam) String() string {
	return string(value)
}
