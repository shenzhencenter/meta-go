package enums

type PostTimelineVisibility string

const (
	PostTimelineVisibilityForcedAllow PostTimelineVisibility = "forced_allow"
	PostTimelineVisibilityHidden      PostTimelineVisibility = "hidden"
	PostTimelineVisibilityNormal      PostTimelineVisibility = "normal"
)

func (value PostTimelineVisibility) String() string {
	return string(value)
}
