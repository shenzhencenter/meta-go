package enums

type LivevideocommentsOrderEnumParam string

const (
	LivevideocommentsOrderEnumParamChronological        LivevideocommentsOrderEnumParam = "chronological"
	LivevideocommentsOrderEnumParamReverseChronological LivevideocommentsOrderEnumParam = "reverse_chronological"
)

func (value LivevideocommentsOrderEnumParam) String() string {
	return string(value)
}
