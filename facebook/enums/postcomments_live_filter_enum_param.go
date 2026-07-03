package enums

type PostcommentsLiveFilterEnumParam string

const (
	PostcommentsLiveFilterEnumParamFilterLowQuality PostcommentsLiveFilterEnumParam = "filter_low_quality"
	PostcommentsLiveFilterEnumParamNoFilter         PostcommentsLiveFilterEnumParam = "no_filter"
)

func (value PostcommentsLiveFilterEnumParam) String() string {
	return string(value)
}
