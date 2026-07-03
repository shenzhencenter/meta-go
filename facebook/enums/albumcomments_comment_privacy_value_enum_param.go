package enums

type AlbumcommentsCommentPrivacyValueEnumParam string

const (
	AlbumcommentsCommentPrivacyValueEnumParamDeclinedByAdminAssistant         AlbumcommentsCommentPrivacyValueEnumParam = "DECLINED_BY_ADMIN_ASSISTANT"
	AlbumcommentsCommentPrivacyValueEnumParamDefaultPrivacy                   AlbumcommentsCommentPrivacyValueEnumParam = "DEFAULT_PRIVACY"
	AlbumcommentsCommentPrivacyValueEnumParamFriendsAndPostOwner              AlbumcommentsCommentPrivacyValueEnumParam = "FRIENDS_AND_POST_OWNER"
	AlbumcommentsCommentPrivacyValueEnumParamFriendsOnly                      AlbumcommentsCommentPrivacyValueEnumParam = "FRIENDS_ONLY"
	AlbumcommentsCommentPrivacyValueEnumParamGraphqlMultipleValueHackDoNotUse AlbumcommentsCommentPrivacyValueEnumParam = "GRAPHQL_MULTIPLE_VALUE_HACK_DO_NOT_USE"
	AlbumcommentsCommentPrivacyValueEnumParamOwnerOrCommenter                 AlbumcommentsCommentPrivacyValueEnumParam = "OWNER_OR_COMMENTER"
	AlbumcommentsCommentPrivacyValueEnumParamPendingApproval                  AlbumcommentsCommentPrivacyValueEnumParam = "PENDING_APPROVAL"
	AlbumcommentsCommentPrivacyValueEnumParamRemovedByAdminAssistant          AlbumcommentsCommentPrivacyValueEnumParam = "REMOVED_BY_ADMIN_ASSISTANT"
	AlbumcommentsCommentPrivacyValueEnumParamSideConversation                 AlbumcommentsCommentPrivacyValueEnumParam = "SIDE_CONVERSATION"
	AlbumcommentsCommentPrivacyValueEnumParamSideConversationAndPostOwner     AlbumcommentsCommentPrivacyValueEnumParam = "SIDE_CONVERSATION_AND_POST_OWNER"
	AlbumcommentsCommentPrivacyValueEnumParamSpotlightTab                     AlbumcommentsCommentPrivacyValueEnumParam = "SPOTLIGHT_TAB"
)

func (value AlbumcommentsCommentPrivacyValueEnumParam) String() string {
	return string(value)
}
