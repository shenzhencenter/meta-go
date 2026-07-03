package enums

type PostcommentsCommentPrivacyValueEnumParam string

const (
	PostcommentsCommentPrivacyValueEnumParamDeclinedByAdminAssistant         PostcommentsCommentPrivacyValueEnumParam = "DECLINED_BY_ADMIN_ASSISTANT"
	PostcommentsCommentPrivacyValueEnumParamDefaultPrivacy                   PostcommentsCommentPrivacyValueEnumParam = "DEFAULT_PRIVACY"
	PostcommentsCommentPrivacyValueEnumParamFriendsAndPostOwner              PostcommentsCommentPrivacyValueEnumParam = "FRIENDS_AND_POST_OWNER"
	PostcommentsCommentPrivacyValueEnumParamFriendsOnly                      PostcommentsCommentPrivacyValueEnumParam = "FRIENDS_ONLY"
	PostcommentsCommentPrivacyValueEnumParamGraphqlMultipleValueHackDoNotUse PostcommentsCommentPrivacyValueEnumParam = "GRAPHQL_MULTIPLE_VALUE_HACK_DO_NOT_USE"
	PostcommentsCommentPrivacyValueEnumParamOwnerOrCommenter                 PostcommentsCommentPrivacyValueEnumParam = "OWNER_OR_COMMENTER"
	PostcommentsCommentPrivacyValueEnumParamPendingApproval                  PostcommentsCommentPrivacyValueEnumParam = "PENDING_APPROVAL"
	PostcommentsCommentPrivacyValueEnumParamRemovedByAdminAssistant          PostcommentsCommentPrivacyValueEnumParam = "REMOVED_BY_ADMIN_ASSISTANT"
	PostcommentsCommentPrivacyValueEnumParamSideConversation                 PostcommentsCommentPrivacyValueEnumParam = "SIDE_CONVERSATION"
	PostcommentsCommentPrivacyValueEnumParamSideConversationAndPostOwner     PostcommentsCommentPrivacyValueEnumParam = "SIDE_CONVERSATION_AND_POST_OWNER"
	PostcommentsCommentPrivacyValueEnumParamSpotlightTab                     PostcommentsCommentPrivacyValueEnumParam = "SPOTLIGHT_TAB"
)

func (value PostcommentsCommentPrivacyValueEnumParam) String() string {
	return string(value)
}
