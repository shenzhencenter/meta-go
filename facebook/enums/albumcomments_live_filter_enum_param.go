package enums

type AlbumcommentsLiveFilterEnumParam string

const (
	AlbumcommentsLiveFilterEnumParamFilterLowQuality AlbumcommentsLiveFilterEnumParam = "filter_low_quality"
	AlbumcommentsLiveFilterEnumParamNoFilter         AlbumcommentsLiveFilterEnumParam = "no_filter"
)

func (value AlbumcommentsLiveFilterEnumParam) String() string {
	return string(value)
}
