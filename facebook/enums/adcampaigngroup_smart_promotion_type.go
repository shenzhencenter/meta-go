package enums

type AdcampaigngroupSmartPromotionType string

const (
	AdcampaigngroupSmartPromotionTypeGuidedCreation    AdcampaigngroupSmartPromotionType = "GUIDED_CREATION"
	AdcampaigngroupSmartPromotionTypeSmartAppPromotion AdcampaigngroupSmartPromotionType = "SMART_APP_PROMOTION"
)

func (value AdcampaigngroupSmartPromotionType) String() string {
	return string(value)
}
