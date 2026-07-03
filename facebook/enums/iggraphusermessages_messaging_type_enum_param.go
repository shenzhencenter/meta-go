package enums

type IggraphusermessagesMessagingTypeEnumParam string

const (
	IggraphusermessagesMessagingTypeEnumParamMessageTag IggraphusermessagesMessagingTypeEnumParam = "MESSAGE_TAG"
	IggraphusermessagesMessagingTypeEnumParamResponse   IggraphusermessagesMessagingTypeEnumParam = "RESPONSE"
	IggraphusermessagesMessagingTypeEnumParamUpdate     IggraphusermessagesMessagingTypeEnumParam = "UPDATE"
	IggraphusermessagesMessagingTypeEnumParamUtility    IggraphusermessagesMessagingTypeEnumParam = "UTILITY"
)

func (value IggraphusermessagesMessagingTypeEnumParam) String() string {
	return string(value)
}
