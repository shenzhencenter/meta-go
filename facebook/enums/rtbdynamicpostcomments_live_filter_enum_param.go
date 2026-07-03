package enums

type RtbdynamicpostcommentsLiveFilterEnumParam string

const (
	RtbdynamicpostcommentsLiveFilterEnumParamFilterLowQuality RtbdynamicpostcommentsLiveFilterEnumParam = "filter_low_quality"
	RtbdynamicpostcommentsLiveFilterEnumParamNoFilter         RtbdynamicpostcommentsLiveFilterEnumParam = "no_filter"
)

func (value RtbdynamicpostcommentsLiveFilterEnumParam) String() string {
	return string(value)
}
