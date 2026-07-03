package enums

type CommentcommentsCommentPrivacyValueEnumParam string

const (
	CommentcommentsCommentPrivacyValueEnumParamDeclinedByAdminAssistant         CommentcommentsCommentPrivacyValueEnumParam = "DECLINED_BY_ADMIN_ASSISTANT"
	CommentcommentsCommentPrivacyValueEnumParamDefaultPrivacy                   CommentcommentsCommentPrivacyValueEnumParam = "DEFAULT_PRIVACY"
	CommentcommentsCommentPrivacyValueEnumParamFriendsAndPostOwner              CommentcommentsCommentPrivacyValueEnumParam = "FRIENDS_AND_POST_OWNER"
	CommentcommentsCommentPrivacyValueEnumParamFriendsOnly                      CommentcommentsCommentPrivacyValueEnumParam = "FRIENDS_ONLY"
	CommentcommentsCommentPrivacyValueEnumParamGraphqlMultipleValueHackDoNotUse CommentcommentsCommentPrivacyValueEnumParam = "GRAPHQL_MULTIPLE_VALUE_HACK_DO_NOT_USE"
	CommentcommentsCommentPrivacyValueEnumParamOwnerOrCommenter                 CommentcommentsCommentPrivacyValueEnumParam = "OWNER_OR_COMMENTER"
	CommentcommentsCommentPrivacyValueEnumParamPendingApproval                  CommentcommentsCommentPrivacyValueEnumParam = "PENDING_APPROVAL"
	CommentcommentsCommentPrivacyValueEnumParamRemovedByAdminAssistant          CommentcommentsCommentPrivacyValueEnumParam = "REMOVED_BY_ADMIN_ASSISTANT"
	CommentcommentsCommentPrivacyValueEnumParamSideConversation                 CommentcommentsCommentPrivacyValueEnumParam = "SIDE_CONVERSATION"
	CommentcommentsCommentPrivacyValueEnumParamSideConversationAndPostOwner     CommentcommentsCommentPrivacyValueEnumParam = "SIDE_CONVERSATION_AND_POST_OWNER"
	CommentcommentsCommentPrivacyValueEnumParamSpotlightTab                     CommentcommentsCommentPrivacyValueEnumParam = "SPOTLIGHT_TAB"
)

func (value CommentcommentsCommentPrivacyValueEnumParam) String() string {
	return string(value)
}
