package enums

type PagepostcommentsOrderEnumParam string

const (
	PagepostcommentsOrderEnumParamChronological        PagepostcommentsOrderEnumParam = "chronological"
	PagepostcommentsOrderEnumParamReverseChronological PagepostcommentsOrderEnumParam = "reverse_chronological"
)

func (value PagepostcommentsOrderEnumParam) String() string {
	return string(value)
}
