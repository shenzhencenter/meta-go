package enums

type PhotocommentsLiveFilterEnumParam string

const (
	PhotocommentsLiveFilterEnumParamFilterLowQuality PhotocommentsLiveFilterEnumParam = "filter_low_quality"
	PhotocommentsLiveFilterEnumParamNoFilter         PhotocommentsLiveFilterEnumParam = "no_filter"
)

func (value PhotocommentsLiveFilterEnumParam) String() string {
	return string(value)
}
