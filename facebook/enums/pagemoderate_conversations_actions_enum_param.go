package enums

type PagemoderateConversationsActionsEnumParam string

const (
	PagemoderateConversationsActionsEnumParamBanUser     PagemoderateConversationsActionsEnumParam = "BAN_USER"
	PagemoderateConversationsActionsEnumParamBlockUser   PagemoderateConversationsActionsEnumParam = "BLOCK_USER"
	PagemoderateConversationsActionsEnumParamMoveToSpam  PagemoderateConversationsActionsEnumParam = "MOVE_TO_SPAM"
	PagemoderateConversationsActionsEnumParamUnbanUser   PagemoderateConversationsActionsEnumParam = "UNBAN_USER"
	PagemoderateConversationsActionsEnumParamUnblockUser PagemoderateConversationsActionsEnumParam = "UNBLOCK_USER"
)

func (value PagemoderateConversationsActionsEnumParam) String() string {
	return string(value)
}
