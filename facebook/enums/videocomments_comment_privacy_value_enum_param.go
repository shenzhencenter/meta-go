package enums

type VideocommentsCommentPrivacyValueEnumParam string

const (
	VideocommentsCommentPrivacyValueEnumParamDeclinedByAdminAssistant         VideocommentsCommentPrivacyValueEnumParam = "DECLINED_BY_ADMIN_ASSISTANT"
	VideocommentsCommentPrivacyValueEnumParamDefaultPrivacy                   VideocommentsCommentPrivacyValueEnumParam = "DEFAULT_PRIVACY"
	VideocommentsCommentPrivacyValueEnumParamFriendsAndPostOwner              VideocommentsCommentPrivacyValueEnumParam = "FRIENDS_AND_POST_OWNER"
	VideocommentsCommentPrivacyValueEnumParamFriendsOnly                      VideocommentsCommentPrivacyValueEnumParam = "FRIENDS_ONLY"
	VideocommentsCommentPrivacyValueEnumParamGraphqlMultipleValueHackDoNotUse VideocommentsCommentPrivacyValueEnumParam = "GRAPHQL_MULTIPLE_VALUE_HACK_DO_NOT_USE"
	VideocommentsCommentPrivacyValueEnumParamOwnerOrCommenter                 VideocommentsCommentPrivacyValueEnumParam = "OWNER_OR_COMMENTER"
	VideocommentsCommentPrivacyValueEnumParamPendingApproval                  VideocommentsCommentPrivacyValueEnumParam = "PENDING_APPROVAL"
	VideocommentsCommentPrivacyValueEnumParamRemovedByAdminAssistant          VideocommentsCommentPrivacyValueEnumParam = "REMOVED_BY_ADMIN_ASSISTANT"
	VideocommentsCommentPrivacyValueEnumParamSideConversation                 VideocommentsCommentPrivacyValueEnumParam = "SIDE_CONVERSATION"
	VideocommentsCommentPrivacyValueEnumParamSideConversationAndPostOwner     VideocommentsCommentPrivacyValueEnumParam = "SIDE_CONVERSATION_AND_POST_OWNER"
	VideocommentsCommentPrivacyValueEnumParamSpotlightTab                     VideocommentsCommentPrivacyValueEnumParam = "SPOTLIGHT_TAB"
)

func (value VideocommentsCommentPrivacyValueEnumParam) String() string {
	return string(value)
}
