package enums

type GrouppictureTypeEnumParam string

const (
	GrouppictureTypeEnumParamAlbum  GrouppictureTypeEnumParam = "album"
	GrouppictureTypeEnumParamLarge  GrouppictureTypeEnumParam = "large"
	GrouppictureTypeEnumParamNormal GrouppictureTypeEnumParam = "normal"
	GrouppictureTypeEnumParamSmall  GrouppictureTypeEnumParam = "small"
	GrouppictureTypeEnumParamSquare GrouppictureTypeEnumParam = "square"
)

func (value GrouppictureTypeEnumParam) String() string {
	return string(value)
}
