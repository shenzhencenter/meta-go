package enums

type UserfeedTargetSurfaceEnumParam string

const (
	UserfeedTargetSurfaceEnumParamStory    UserfeedTargetSurfaceEnumParam = "STORY"
	UserfeedTargetSurfaceEnumParamTimeline UserfeedTargetSurfaceEnumParam = "TIMELINE"
)

func (value UserfeedTargetSurfaceEnumParam) String() string {
	return string(value)
}
