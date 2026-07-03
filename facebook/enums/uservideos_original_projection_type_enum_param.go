package enums

type UservideosOriginalProjectionTypeEnumParam string

const (
	UservideosOriginalProjectionTypeEnumParamCubemap             UservideosOriginalProjectionTypeEnumParam = "cubemap"
	UservideosOriginalProjectionTypeEnumParamEquirectangular     UservideosOriginalProjectionTypeEnumParam = "equirectangular"
	UservideosOriginalProjectionTypeEnumParamHalfEquirectangular UservideosOriginalProjectionTypeEnumParam = "half_equirectangular"
)

func (value UservideosOriginalProjectionTypeEnumParam) String() string {
	return string(value)
}
