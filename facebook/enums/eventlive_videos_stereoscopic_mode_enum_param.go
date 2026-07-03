package enums

type EventliveVideosStereoscopicModeEnumParam string

const (
	EventliveVideosStereoscopicModeEnumParamLeftRight EventliveVideosStereoscopicModeEnumParam = "LEFT_RIGHT"
	EventliveVideosStereoscopicModeEnumParamMono      EventliveVideosStereoscopicModeEnumParam = "MONO"
	EventliveVideosStereoscopicModeEnumParamMultiView EventliveVideosStereoscopicModeEnumParam = "MULTI_VIEW"
	EventliveVideosStereoscopicModeEnumParamTopBottom EventliveVideosStereoscopicModeEnumParam = "TOP_BOTTOM"
)

func (value EventliveVideosStereoscopicModeEnumParam) String() string {
	return string(value)
}
