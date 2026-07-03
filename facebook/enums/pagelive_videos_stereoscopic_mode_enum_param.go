package enums

type PageliveVideosStereoscopicModeEnumParam string

const (
	PageliveVideosStereoscopicModeEnumParamLeftRight PageliveVideosStereoscopicModeEnumParam = "LEFT_RIGHT"
	PageliveVideosStereoscopicModeEnumParamMono      PageliveVideosStereoscopicModeEnumParam = "MONO"
	PageliveVideosStereoscopicModeEnumParamMultiView PageliveVideosStereoscopicModeEnumParam = "MULTI_VIEW"
	PageliveVideosStereoscopicModeEnumParamTopBottom PageliveVideosStereoscopicModeEnumParam = "TOP_BOTTOM"
)

func (value PageliveVideosStereoscopicModeEnumParam) String() string {
	return string(value)
}
