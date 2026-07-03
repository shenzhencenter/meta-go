package enums

type AdcampaignplacementgetEffectiveDevicePlatforms string

const (
	AdcampaignplacementgetEffectiveDevicePlatformsConnectedTv AdcampaignplacementgetEffectiveDevicePlatforms = "CONNECTED_TV"
	AdcampaignplacementgetEffectiveDevicePlatformsDesktop     AdcampaignplacementgetEffectiveDevicePlatforms = "DESKTOP"
	AdcampaignplacementgetEffectiveDevicePlatformsMobile      AdcampaignplacementgetEffectiveDevicePlatforms = "MOBILE"
)

func (value AdcampaignplacementgetEffectiveDevicePlatforms) String() string {
	return string(value)
}
