package enums

type PostcommentsOrderEnumParam string

const (
	PostcommentsOrderEnumParamChronological        PostcommentsOrderEnumParam = "chronological"
	PostcommentsOrderEnumParamReverseChronological PostcommentsOrderEnumParam = "reverse_chronological"
)

func (value PostcommentsOrderEnumParam) String() string {
	return string(value)
}
