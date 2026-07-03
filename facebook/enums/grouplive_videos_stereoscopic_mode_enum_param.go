package enums

type GroupliveVideosStereoscopicModeEnumParam string

const (
	GroupliveVideosStereoscopicModeEnumParamLeftRight GroupliveVideosStereoscopicModeEnumParam = "LEFT_RIGHT"
	GroupliveVideosStereoscopicModeEnumParamMono      GroupliveVideosStereoscopicModeEnumParam = "MONO"
	GroupliveVideosStereoscopicModeEnumParamMultiView GroupliveVideosStereoscopicModeEnumParam = "MULTI_VIEW"
	GroupliveVideosStereoscopicModeEnumParamTopBottom GroupliveVideosStereoscopicModeEnumParam = "TOP_BOTTOM"
)

func (value GroupliveVideosStereoscopicModeEnumParam) String() string {
	return string(value)
}
