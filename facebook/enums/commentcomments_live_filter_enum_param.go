package enums

type CommentcommentsLiveFilterEnumParam string

const (
	CommentcommentsLiveFilterEnumParamFilterLowQuality CommentcommentsLiveFilterEnumParam = "filter_low_quality"
	CommentcommentsLiveFilterEnumParamNoFilter         CommentcommentsLiveFilterEnumParam = "no_filter"
)

func (value CommentcommentsLiveFilterEnumParam) String() string {
	return string(value)
}
