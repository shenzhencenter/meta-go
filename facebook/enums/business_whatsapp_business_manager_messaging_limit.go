package enums

type BusinessWhatsappBusinessManagerMessagingLimit string

const (
	BusinessWhatsappBusinessManagerMessagingLimitTierX100K     BusinessWhatsappBusinessManagerMessagingLimit = "TIER_100K"
	BusinessWhatsappBusinessManagerMessagingLimitTierX10K      BusinessWhatsappBusinessManagerMessagingLimit = "TIER_10K"
	BusinessWhatsappBusinessManagerMessagingLimitTierX250      BusinessWhatsappBusinessManagerMessagingLimit = "TIER_250"
	BusinessWhatsappBusinessManagerMessagingLimitTierX2K       BusinessWhatsappBusinessManagerMessagingLimit = "TIER_2K"
	BusinessWhatsappBusinessManagerMessagingLimitTierX50       BusinessWhatsappBusinessManagerMessagingLimit = "TIER_50"
	BusinessWhatsappBusinessManagerMessagingLimitTierUnlimited BusinessWhatsappBusinessManagerMessagingLimit = "TIER_UNLIMITED"
	BusinessWhatsappBusinessManagerMessagingLimitUntiered      BusinessWhatsappBusinessManagerMessagingLimit = "UNTIERED"
)

func (value BusinessWhatsappBusinessManagerMessagingLimit) String() string {
	return string(value)
}
