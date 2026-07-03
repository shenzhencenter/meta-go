package enums

type LivevideocommentsLiveFilterEnumParam string

const (
	LivevideocommentsLiveFilterEnumParamFilterLowQuality LivevideocommentsLiveFilterEnumParam = "filter_low_quality"
	LivevideocommentsLiveFilterEnumParamNoFilter         LivevideocommentsLiveFilterEnumParam = "no_filter"
)

func (value LivevideocommentsLiveFilterEnumParam) String() string {
	return string(value)
}
