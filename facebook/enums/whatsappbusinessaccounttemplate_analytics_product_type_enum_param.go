package enums

type WhatsappbusinessaccounttemplateAnalyticsProductTypeEnumParam string

const (
	WhatsappbusinessaccounttemplateAnalyticsProductTypeEnumParamCampaignInsights         WhatsappbusinessaccounttemplateAnalyticsProductTypeEnumParam = "CAMPAIGN_INSIGHTS"
	WhatsappbusinessaccounttemplateAnalyticsProductTypeEnumParamCloudAPI                 WhatsappbusinessaccounttemplateAnalyticsProductTypeEnumParam = "CLOUD_API"
	WhatsappbusinessaccounttemplateAnalyticsProductTypeEnumParamMarketingMessagesLiteAPI WhatsappbusinessaccounttemplateAnalyticsProductTypeEnumParam = "MARKETING_MESSAGES_LITE_API"
)

func (value WhatsappbusinessaccounttemplateAnalyticsProductTypeEnumParam) String() string {
	return string(value)
}
