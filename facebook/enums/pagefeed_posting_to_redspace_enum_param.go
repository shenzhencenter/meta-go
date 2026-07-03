package enums

type PagefeedPostingToRedspaceEnumParam string

const (
	PagefeedPostingToRedspaceEnumParamDisabled PagefeedPostingToRedspaceEnumParam = "disabled"
	PagefeedPostingToRedspaceEnumParamEnabled  PagefeedPostingToRedspaceEnumParam = "enabled"
)

func (value PagefeedPostingToRedspaceEnumParam) String() string {
	return string(value)
}
