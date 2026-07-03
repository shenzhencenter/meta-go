package enums

type PagemessagesMessagingTypeEnumParam string

const (
	PagemessagesMessagingTypeEnumParamMessageTag PagemessagesMessagingTypeEnumParam = "MESSAGE_TAG"
	PagemessagesMessagingTypeEnumParamResponse   PagemessagesMessagingTypeEnumParam = "RESPONSE"
	PagemessagesMessagingTypeEnumParamUpdate     PagemessagesMessagingTypeEnumParam = "UPDATE"
	PagemessagesMessagingTypeEnumParamUtility    PagemessagesMessagingTypeEnumParam = "UTILITY"
)

func (value PagemessagesMessagingTypeEnumParam) String() string {
	return string(value)
}
