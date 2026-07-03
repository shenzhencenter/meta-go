package enums

type VideocommentsOrderEnumParam string

const (
	VideocommentsOrderEnumParamChronological        VideocommentsOrderEnumParam = "chronological"
	VideocommentsOrderEnumParamReverseChronological VideocommentsOrderEnumParam = "reverse_chronological"
)

func (value VideocommentsOrderEnumParam) String() string {
	return string(value)
}
