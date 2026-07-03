package enums

type PagepictureTypeEnumParam string

const (
	PagepictureTypeEnumParamAlbum  PagepictureTypeEnumParam = "album"
	PagepictureTypeEnumParamLarge  PagepictureTypeEnumParam = "large"
	PagepictureTypeEnumParamNormal PagepictureTypeEnumParam = "normal"
	PagepictureTypeEnumParamSmall  PagepictureTypeEnumParam = "small"
	PagepictureTypeEnumParamSquare PagepictureTypeEnumParam = "square"
)

func (value PagepictureTypeEnumParam) String() string {
	return string(value)
}
