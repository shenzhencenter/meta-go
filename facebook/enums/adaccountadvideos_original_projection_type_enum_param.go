package enums

type AdaccountadvideosOriginalProjectionTypeEnumParam string

const (
	AdaccountadvideosOriginalProjectionTypeEnumParamCubemap             AdaccountadvideosOriginalProjectionTypeEnumParam = "cubemap"
	AdaccountadvideosOriginalProjectionTypeEnumParamEquirectangular     AdaccountadvideosOriginalProjectionTypeEnumParam = "equirectangular"
	AdaccountadvideosOriginalProjectionTypeEnumParamHalfEquirectangular AdaccountadvideosOriginalProjectionTypeEnumParam = "half_equirectangular"
)

func (value AdaccountadvideosOriginalProjectionTypeEnumParam) String() string {
	return string(value)
}
