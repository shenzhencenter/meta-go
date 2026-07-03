package enums

type AlbumpictureTypeEnumParam string

const (
	AlbumpictureTypeEnumParamAlbum     AlbumpictureTypeEnumParam = "album"
	AlbumpictureTypeEnumParamSmall     AlbumpictureTypeEnumParam = "small"
	AlbumpictureTypeEnumParamThumbnail AlbumpictureTypeEnumParam = "thumbnail"
)

func (value AlbumpictureTypeEnumParam) String() string {
	return string(value)
}
