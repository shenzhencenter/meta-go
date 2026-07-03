package enums

type AlbumcommentsOrderEnumParam string

const (
	AlbumcommentsOrderEnumParamChronological        AlbumcommentsOrderEnumParam = "chronological"
	AlbumcommentsOrderEnumParamReverseChronological AlbumcommentsOrderEnumParam = "reverse_chronological"
)

func (value AlbumcommentsOrderEnumParam) String() string {
	return string(value)
}
