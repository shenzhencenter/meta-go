package enums

type UserliveVideosStereoscopicModeEnumParam string

const (
	UserliveVideosStereoscopicModeEnumParamLeftRight UserliveVideosStereoscopicModeEnumParam = "LEFT_RIGHT"
	UserliveVideosStereoscopicModeEnumParamMono      UserliveVideosStereoscopicModeEnumParam = "MONO"
	UserliveVideosStereoscopicModeEnumParamMultiView UserliveVideosStereoscopicModeEnumParam = "MULTI_VIEW"
	UserliveVideosStereoscopicModeEnumParamTopBottom UserliveVideosStereoscopicModeEnumParam = "TOP_BOTTOM"
)

func (value UserliveVideosStereoscopicModeEnumParam) String() string {
	return string(value)
}
