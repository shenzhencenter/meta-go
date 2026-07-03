package enums

type BusinessvideosOriginalProjectionTypeEnumParam string

const (
	BusinessvideosOriginalProjectionTypeEnumParamCubemap             BusinessvideosOriginalProjectionTypeEnumParam = "cubemap"
	BusinessvideosOriginalProjectionTypeEnumParamEquirectangular     BusinessvideosOriginalProjectionTypeEnumParam = "equirectangular"
	BusinessvideosOriginalProjectionTypeEnumParamHalfEquirectangular BusinessvideosOriginalProjectionTypeEnumParam = "half_equirectangular"
)

func (value BusinessvideosOriginalProjectionTypeEnumParam) String() string {
	return string(value)
}
