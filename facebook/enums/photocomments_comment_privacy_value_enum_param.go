package enums

type PhotocommentsCommentPrivacyValueEnumParam string

const (
	PhotocommentsCommentPrivacyValueEnumParamDeclinedByAdminAssistant         PhotocommentsCommentPrivacyValueEnumParam = "DECLINED_BY_ADMIN_ASSISTANT"
	PhotocommentsCommentPrivacyValueEnumParamDefaultPrivacy                   PhotocommentsCommentPrivacyValueEnumParam = "DEFAULT_PRIVACY"
	PhotocommentsCommentPrivacyValueEnumParamFriendsAndPostOwner              PhotocommentsCommentPrivacyValueEnumParam = "FRIENDS_AND_POST_OWNER"
	PhotocommentsCommentPrivacyValueEnumParamFriendsOnly                      PhotocommentsCommentPrivacyValueEnumParam = "FRIENDS_ONLY"
	PhotocommentsCommentPrivacyValueEnumParamGraphqlMultipleValueHackDoNotUse PhotocommentsCommentPrivacyValueEnumParam = "GRAPHQL_MULTIPLE_VALUE_HACK_DO_NOT_USE"
	PhotocommentsCommentPrivacyValueEnumParamOwnerOrCommenter                 PhotocommentsCommentPrivacyValueEnumParam = "OWNER_OR_COMMENTER"
	PhotocommentsCommentPrivacyValueEnumParamPendingApproval                  PhotocommentsCommentPrivacyValueEnumParam = "PENDING_APPROVAL"
	PhotocommentsCommentPrivacyValueEnumParamRemovedByAdminAssistant          PhotocommentsCommentPrivacyValueEnumParam = "REMOVED_BY_ADMIN_ASSISTANT"
	PhotocommentsCommentPrivacyValueEnumParamSideConversation                 PhotocommentsCommentPrivacyValueEnumParam = "SIDE_CONVERSATION"
	PhotocommentsCommentPrivacyValueEnumParamSideConversationAndPostOwner     PhotocommentsCommentPrivacyValueEnumParam = "SIDE_CONVERSATION_AND_POST_OWNER"
	PhotocommentsCommentPrivacyValueEnumParamSpotlightTab                     PhotocommentsCommentPrivacyValueEnumParam = "SPOTLIGHT_TAB"
)

func (value PhotocommentsCommentPrivacyValueEnumParam) String() string {
	return string(value)
}
