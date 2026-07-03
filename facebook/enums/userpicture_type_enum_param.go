package enums

type UserpictureTypeEnumParam string

const (
	UserpictureTypeEnumParamAlbum  UserpictureTypeEnumParam = "album"
	UserpictureTypeEnumParamLarge  UserpictureTypeEnumParam = "large"
	UserpictureTypeEnumParamNormal UserpictureTypeEnumParam = "normal"
	UserpictureTypeEnumParamSmall  UserpictureTypeEnumParam = "small"
	UserpictureTypeEnumParamSquare UserpictureTypeEnumParam = "square"
)

func (value UserpictureTypeEnumParam) String() string {
	return string(value)
}
