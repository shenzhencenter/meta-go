package enums

type PhotocommentsOrderEnumParam string

const (
	PhotocommentsOrderEnumParamChronological        PhotocommentsOrderEnumParam = "chronological"
	PhotocommentsOrderEnumParamReverseChronological PhotocommentsOrderEnumParam = "reverse_chronological"
)

func (value PhotocommentsOrderEnumParam) String() string {
	return string(value)
}
