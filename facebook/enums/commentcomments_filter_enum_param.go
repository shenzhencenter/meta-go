package enums

type CommentcommentsFilterEnumParam string

const (
	CommentcommentsFilterEnumParamStream   CommentcommentsFilterEnumParam = "stream"
	CommentcommentsFilterEnumParamToplevel CommentcommentsFilterEnumParam = "toplevel"
)

func (value CommentcommentsFilterEnumParam) String() string {
	return string(value)
}
