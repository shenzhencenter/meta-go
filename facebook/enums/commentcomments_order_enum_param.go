package enums

type CommentcommentsOrderEnumParam string

const (
	CommentcommentsOrderEnumParamChronological        CommentcommentsOrderEnumParam = "chronological"
	CommentcommentsOrderEnumParamReverseChronological CommentcommentsOrderEnumParam = "reverse_chronological"
)

func (value CommentcommentsOrderEnumParam) String() string {
	return string(value)
}
