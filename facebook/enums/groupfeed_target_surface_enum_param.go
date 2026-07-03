package enums

type GroupfeedTargetSurfaceEnumParam string

const (
	GroupfeedTargetSurfaceEnumParamStory    GroupfeedTargetSurfaceEnumParam = "STORY"
	GroupfeedTargetSurfaceEnumParamTimeline GroupfeedTargetSurfaceEnumParam = "TIMELINE"
)

func (value GroupfeedTargetSurfaceEnumParam) String() string {
	return string(value)
}
