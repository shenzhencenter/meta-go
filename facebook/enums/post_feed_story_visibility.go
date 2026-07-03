package enums

type PostFeedStoryVisibility string

const (
	PostFeedStoryVisibilityHidden  PostFeedStoryVisibility = "hidden"
	PostFeedStoryVisibilityVisible PostFeedStoryVisibility = "visible"
)

func (value PostFeedStoryVisibility) String() string {
	return string(value)
}
