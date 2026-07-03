package enums

type PagefeedTargetSurfaceEnumParam string

const (
	PagefeedTargetSurfaceEnumParamStory    PagefeedTargetSurfaceEnumParam = "STORY"
	PagefeedTargetSurfaceEnumParamTimeline PagefeedTargetSurfaceEnumParam = "TIMELINE"
)

func (value PagefeedTargetSurfaceEnumParam) String() string {
	return string(value)
}
