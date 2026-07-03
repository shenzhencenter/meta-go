package enums

type WhatsappbusinessaccountgroupAnalyticsMetricTypesEnumParam string

const (
	WhatsappbusinessaccountgroupAnalyticsMetricTypesEnumParamClicks             WhatsappbusinessaccountgroupAnalyticsMetricTypesEnumParam = "CLICKS"
	WhatsappbusinessaccountgroupAnalyticsMetricTypesEnumParamCost               WhatsappbusinessaccountgroupAnalyticsMetricTypesEnumParam = "COST"
	WhatsappbusinessaccountgroupAnalyticsMetricTypesEnumParamDelivered          WhatsappbusinessaccountgroupAnalyticsMetricTypesEnumParam = "DELIVERED"
	WhatsappbusinessaccountgroupAnalyticsMetricTypesEnumParamParticipantsJoined WhatsappbusinessaccountgroupAnalyticsMetricTypesEnumParam = "PARTICIPANTS_JOINED"
	WhatsappbusinessaccountgroupAnalyticsMetricTypesEnumParamParticipantsLeft   WhatsappbusinessaccountgroupAnalyticsMetricTypesEnumParam = "PARTICIPANTS_LEFT"
	WhatsappbusinessaccountgroupAnalyticsMetricTypesEnumParamRead               WhatsappbusinessaccountgroupAnalyticsMetricTypesEnumParam = "READ"
	WhatsappbusinessaccountgroupAnalyticsMetricTypesEnumParamReplies            WhatsappbusinessaccountgroupAnalyticsMetricTypesEnumParam = "REPLIES"
	WhatsappbusinessaccountgroupAnalyticsMetricTypesEnumParamSent               WhatsappbusinessaccountgroupAnalyticsMetricTypesEnumParam = "SENT"
)

func (value WhatsappbusinessaccountgroupAnalyticsMetricTypesEnumParam) String() string {
	return string(value)
}
