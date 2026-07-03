package enums

type PagepostTimelineVisibility string

const (
	PagepostTimelineVisibilityForcedAllow PagepostTimelineVisibility = "forced_allow"
	PagepostTimelineVisibilityHidden      PagepostTimelineVisibility = "hidden"
	PagepostTimelineVisibilityNormal      PagepostTimelineVisibility = "normal"
)

func (value PagepostTimelineVisibility) String() string {
	return string(value)
}
