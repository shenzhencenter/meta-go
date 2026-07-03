package enums

type WhatsappbusinessaccountconversationAnalyticsDimensionsEnumParam string

const (
	WhatsappbusinessaccountconversationAnalyticsDimensionsEnumParamConversationCategory  WhatsappbusinessaccountconversationAnalyticsDimensionsEnumParam = "CONVERSATION_CATEGORY"
	WhatsappbusinessaccountconversationAnalyticsDimensionsEnumParamConversationDirection WhatsappbusinessaccountconversationAnalyticsDimensionsEnumParam = "CONVERSATION_DIRECTION"
	WhatsappbusinessaccountconversationAnalyticsDimensionsEnumParamConversationType      WhatsappbusinessaccountconversationAnalyticsDimensionsEnumParam = "CONVERSATION_TYPE"
	WhatsappbusinessaccountconversationAnalyticsDimensionsEnumParamCountry               WhatsappbusinessaccountconversationAnalyticsDimensionsEnumParam = "COUNTRY"
	WhatsappbusinessaccountconversationAnalyticsDimensionsEnumParamPhone                 WhatsappbusinessaccountconversationAnalyticsDimensionsEnumParam = "PHONE"
	WhatsappbusinessaccountconversationAnalyticsDimensionsEnumParamUnknown               WhatsappbusinessaccountconversationAnalyticsDimensionsEnumParam = "UNKNOWN"
)

func (value WhatsappbusinessaccountconversationAnalyticsDimensionsEnumParam) String() string {
	return string(value)
}
