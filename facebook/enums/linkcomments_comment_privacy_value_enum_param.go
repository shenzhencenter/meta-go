package enums

type LinkcommentsCommentPrivacyValueEnumParam string

const (
	LinkcommentsCommentPrivacyValueEnumParamDeclinedByAdminAssistant         LinkcommentsCommentPrivacyValueEnumParam = "DECLINED_BY_ADMIN_ASSISTANT"
	LinkcommentsCommentPrivacyValueEnumParamDefaultPrivacy                   LinkcommentsCommentPrivacyValueEnumParam = "DEFAULT_PRIVACY"
	LinkcommentsCommentPrivacyValueEnumParamFriendsAndPostOwner              LinkcommentsCommentPrivacyValueEnumParam = "FRIENDS_AND_POST_OWNER"
	LinkcommentsCommentPrivacyValueEnumParamFriendsOnly                      LinkcommentsCommentPrivacyValueEnumParam = "FRIENDS_ONLY"
	LinkcommentsCommentPrivacyValueEnumParamGraphqlMultipleValueHackDoNotUse LinkcommentsCommentPrivacyValueEnumParam = "GRAPHQL_MULTIPLE_VALUE_HACK_DO_NOT_USE"
	LinkcommentsCommentPrivacyValueEnumParamOwnerOrCommenter                 LinkcommentsCommentPrivacyValueEnumParam = "OWNER_OR_COMMENTER"
	LinkcommentsCommentPrivacyValueEnumParamPendingApproval                  LinkcommentsCommentPrivacyValueEnumParam = "PENDING_APPROVAL"
	LinkcommentsCommentPrivacyValueEnumParamRemovedByAdminAssistant          LinkcommentsCommentPrivacyValueEnumParam = "REMOVED_BY_ADMIN_ASSISTANT"
	LinkcommentsCommentPrivacyValueEnumParamSideConversation                 LinkcommentsCommentPrivacyValueEnumParam = "SIDE_CONVERSATION"
	LinkcommentsCommentPrivacyValueEnumParamSideConversationAndPostOwner     LinkcommentsCommentPrivacyValueEnumParam = "SIDE_CONVERSATION_AND_POST_OWNER"
	LinkcommentsCommentPrivacyValueEnumParamSpotlightTab                     LinkcommentsCommentPrivacyValueEnumParam = "SPOTLIGHT_TAB"
)

func (value LinkcommentsCommentPrivacyValueEnumParam) String() string {
	return string(value)
}
