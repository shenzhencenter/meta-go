package enums

type AdaccountcampaignsSmartPromotionTypeEnumParam string

const (
	AdaccountcampaignsSmartPromotionTypeEnumParamGuidedCreation    AdaccountcampaignsSmartPromotionTypeEnumParam = "GUIDED_CREATION"
	AdaccountcampaignsSmartPromotionTypeEnumParamSmartAppPromotion AdaccountcampaignsSmartPromotionTypeEnumParam = "SMART_APP_PROMOTION"
)

func (value AdaccountcampaignsSmartPromotionTypeEnumParam) String() string {
	return string(value)
}
