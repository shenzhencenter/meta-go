package enums

type RtbdynamicpostcommentsOrderEnumParam string

const (
	RtbdynamicpostcommentsOrderEnumParamChronological        RtbdynamicpostcommentsOrderEnumParam = "chronological"
	RtbdynamicpostcommentsOrderEnumParamReverseChronological RtbdynamicpostcommentsOrderEnumParam = "reverse_chronological"
)

func (value RtbdynamicpostcommentsOrderEnumParam) String() string {
	return string(value)
}
