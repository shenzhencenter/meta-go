package enums

type AdcampaignplacementgetEffectivePublisherPlatforms string

const (
	AdcampaignplacementgetEffectivePublisherPlatformsAudienceNetwork AdcampaignplacementgetEffectivePublisherPlatforms = "AUDIENCE_NETWORK"
	AdcampaignplacementgetEffectivePublisherPlatformsFacebook        AdcampaignplacementgetEffectivePublisherPlatforms = "FACEBOOK"
	AdcampaignplacementgetEffectivePublisherPlatformsInstagram       AdcampaignplacementgetEffectivePublisherPlatforms = "INSTAGRAM"
	AdcampaignplacementgetEffectivePublisherPlatformsMessenger       AdcampaignplacementgetEffectivePublisherPlatforms = "MESSENGER"
	AdcampaignplacementgetEffectivePublisherPlatformsOculus          AdcampaignplacementgetEffectivePublisherPlatforms = "OCULUS"
	AdcampaignplacementgetEffectivePublisherPlatformsThreads         AdcampaignplacementgetEffectivePublisherPlatforms = "THREADS"
	AdcampaignplacementgetEffectivePublisherPlatformsWhatsapp        AdcampaignplacementgetEffectivePublisherPlatforms = "WHATSAPP"
)

func (value AdcampaignplacementgetEffectivePublisherPlatforms) String() string {
	return string(value)
}
