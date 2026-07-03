package enums

type PagepostFeedStoryVisibility string

const (
	PagepostFeedStoryVisibilityHidden  PagepostFeedStoryVisibility = "hidden"
	PagepostFeedStoryVisibilityVisible PagepostFeedStoryVisibility = "visible"
)

func (value PagepostFeedStoryVisibility) String() string {
	return string(value)
}
