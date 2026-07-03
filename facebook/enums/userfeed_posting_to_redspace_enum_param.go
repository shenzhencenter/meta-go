package enums

type UserfeedPostingToRedspaceEnumParam string

const (
	UserfeedPostingToRedspaceEnumParamDisabled UserfeedPostingToRedspaceEnumParam = "disabled"
	UserfeedPostingToRedspaceEnumParamEnabled  UserfeedPostingToRedspaceEnumParam = "enabled"
)

func (value UserfeedPostingToRedspaceEnumParam) String() string {
	return string(value)
}
