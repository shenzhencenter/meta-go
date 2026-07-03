package enums

type PagevideosOriginalProjectionTypeEnumParam string

const (
	PagevideosOriginalProjectionTypeEnumParamCubemap             PagevideosOriginalProjectionTypeEnumParam = "cubemap"
	PagevideosOriginalProjectionTypeEnumParamEquirectangular     PagevideosOriginalProjectionTypeEnumParam = "equirectangular"
	PagevideosOriginalProjectionTypeEnumParamHalfEquirectangular PagevideosOriginalProjectionTypeEnumParam = "half_equirectangular"
)

func (value PagevideosOriginalProjectionTypeEnumParam) String() string {
	return string(value)
}
