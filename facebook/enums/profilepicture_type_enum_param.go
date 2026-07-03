package enums

type ProfilepictureTypeEnumParam string

const (
	ProfilepictureTypeEnumParamAlbum  ProfilepictureTypeEnumParam = "album"
	ProfilepictureTypeEnumParamLarge  ProfilepictureTypeEnumParam = "large"
	ProfilepictureTypeEnumParamNormal ProfilepictureTypeEnumParam = "normal"
	ProfilepictureTypeEnumParamSmall  ProfilepictureTypeEnumParam = "small"
	ProfilepictureTypeEnumParamSquare ProfilepictureTypeEnumParam = "square"
)

func (value ProfilepictureTypeEnumParam) String() string {
	return string(value)
}
