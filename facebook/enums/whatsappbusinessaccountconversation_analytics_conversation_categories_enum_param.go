package enums

type WhatsappbusinessaccountconversationAnalyticsConversationCategoriesEnumParam string

const (
	WhatsappbusinessaccountconversationAnalyticsConversationCategoriesEnumParamAuthentication              WhatsappbusinessaccountconversationAnalyticsConversationCategoriesEnumParam = "AUTHENTICATION"
	WhatsappbusinessaccountconversationAnalyticsConversationCategoriesEnumParamAuthenticationInternational WhatsappbusinessaccountconversationAnalyticsConversationCategoriesEnumParam = "AUTHENTICATION_INTERNATIONAL"
	WhatsappbusinessaccountconversationAnalyticsConversationCategoriesEnumParamMarketing                   WhatsappbusinessaccountconversationAnalyticsConversationCategoriesEnumParam = "MARKETING"
	WhatsappbusinessaccountconversationAnalyticsConversationCategoriesEnumParamMarketingLite               WhatsappbusinessaccountconversationAnalyticsConversationCategoriesEnumParam = "MARKETING_LITE"
	WhatsappbusinessaccountconversationAnalyticsConversationCategoriesEnumParamService                     WhatsappbusinessaccountconversationAnalyticsConversationCategoriesEnumParam = "SERVICE"
	WhatsappbusinessaccountconversationAnalyticsConversationCategoriesEnumParamUtility                     WhatsappbusinessaccountconversationAnalyticsConversationCategoriesEnumParam = "UTILITY"
)

func (value WhatsappbusinessaccountconversationAnalyticsConversationCategoriesEnumParam) String() string {
	return string(value)
}
