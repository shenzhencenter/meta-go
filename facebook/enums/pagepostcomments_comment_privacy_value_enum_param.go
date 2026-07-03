package enums

type PagepostcommentsCommentPrivacyValueEnumParam string

const (
	PagepostcommentsCommentPrivacyValueEnumParamDeclinedByAdminAssistant         PagepostcommentsCommentPrivacyValueEnumParam = "DECLINED_BY_ADMIN_ASSISTANT"
	PagepostcommentsCommentPrivacyValueEnumParamDefaultPrivacy                   PagepostcommentsCommentPrivacyValueEnumParam = "DEFAULT_PRIVACY"
	PagepostcommentsCommentPrivacyValueEnumParamFriendsAndPostOwner              PagepostcommentsCommentPrivacyValueEnumParam = "FRIENDS_AND_POST_OWNER"
	PagepostcommentsCommentPrivacyValueEnumParamFriendsOnly                      PagepostcommentsCommentPrivacyValueEnumParam = "FRIENDS_ONLY"
	PagepostcommentsCommentPrivacyValueEnumParamGraphqlMultipleValueHackDoNotUse PagepostcommentsCommentPrivacyValueEnumParam = "GRAPHQL_MULTIPLE_VALUE_HACK_DO_NOT_USE"
	PagepostcommentsCommentPrivacyValueEnumParamOwnerOrCommenter                 PagepostcommentsCommentPrivacyValueEnumParam = "OWNER_OR_COMMENTER"
	PagepostcommentsCommentPrivacyValueEnumParamPendingApproval                  PagepostcommentsCommentPrivacyValueEnumParam = "PENDING_APPROVAL"
	PagepostcommentsCommentPrivacyValueEnumParamRemovedByAdminAssistant          PagepostcommentsCommentPrivacyValueEnumParam = "REMOVED_BY_ADMIN_ASSISTANT"
	PagepostcommentsCommentPrivacyValueEnumParamSideConversation                 PagepostcommentsCommentPrivacyValueEnumParam = "SIDE_CONVERSATION"
	PagepostcommentsCommentPrivacyValueEnumParamSideConversationAndPostOwner     PagepostcommentsCommentPrivacyValueEnumParam = "SIDE_CONVERSATION_AND_POST_OWNER"
	PagepostcommentsCommentPrivacyValueEnumParamSpotlightTab                     PagepostcommentsCommentPrivacyValueEnumParam = "SPOTLIGHT_TAB"
)

func (value PagepostcommentsCommentPrivacyValueEnumParam) String() string {
	return string(value)
}
