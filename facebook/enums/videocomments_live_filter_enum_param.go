package enums

type VideocommentsLiveFilterEnumParam string

const (
	VideocommentsLiveFilterEnumParamFilterLowQuality VideocommentsLiveFilterEnumParam = "filter_low_quality"
	VideocommentsLiveFilterEnumParamNoFilter         VideocommentsLiveFilterEnumParam = "no_filter"
)

func (value VideocommentsLiveFilterEnumParam) String() string {
	return string(value)
}
