package enums

type PagenotificationMessageTokensOptInSourceEnumParam string

const (
	PagenotificationMessageTokensOptInSourceEnumParamCommentAutomation PagenotificationMessageTokensOptInSourceEnumParam = "COMMENT_AUTOMATION"
	PagenotificationMessageTokensOptInSourceEnumParamCtm               PagenotificationMessageTokensOptInSourceEnumParam = "CTM"
	PagenotificationMessageTokensOptInSourceEnumParamReplyAutomation   PagenotificationMessageTokensOptInSourceEnumParam = "REPLY_AUTOMATION"
	PagenotificationMessageTokensOptInSourceEnumParamSubscriberList    PagenotificationMessageTokensOptInSourceEnumParam = "SUBSCRIBER_LIST"
)

func (value PagenotificationMessageTokensOptInSourceEnumParam) String() string {
	return string(value)
}
