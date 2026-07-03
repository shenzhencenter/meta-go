package enums

type BusinesspictureTypeEnumParam string

const (
	BusinesspictureTypeEnumParamAlbum  BusinesspictureTypeEnumParam = "album"
	BusinesspictureTypeEnumParamLarge  BusinesspictureTypeEnumParam = "large"
	BusinesspictureTypeEnumParamNormal BusinesspictureTypeEnumParam = "normal"
	BusinesspictureTypeEnumParamSmall  BusinesspictureTypeEnumParam = "small"
	BusinesspictureTypeEnumParamSquare BusinesspictureTypeEnumParam = "square"
)

func (value BusinesspictureTypeEnumParam) String() string {
	return string(value)
}
