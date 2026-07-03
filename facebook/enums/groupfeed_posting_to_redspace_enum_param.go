package enums

type GroupfeedPostingToRedspaceEnumParam string

const (
	GroupfeedPostingToRedspaceEnumParamDisabled GroupfeedPostingToRedspaceEnumParam = "disabled"
	GroupfeedPostingToRedspaceEnumParamEnabled  GroupfeedPostingToRedspaceEnumParam = "enabled"
)

func (value GroupfeedPostingToRedspaceEnumParam) String() string {
	return string(value)
}
