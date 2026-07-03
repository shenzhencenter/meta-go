package enums

type PagepostcommentsLiveFilterEnumParam string

const (
	PagepostcommentsLiveFilterEnumParamFilterLowQuality PagepostcommentsLiveFilterEnumParam = "filter_low_quality"
	PagepostcommentsLiveFilterEnumParamNoFilter         PagepostcommentsLiveFilterEnumParam = "no_filter"
)

func (value PagepostcommentsLiveFilterEnumParam) String() string {
	return string(value)
}
