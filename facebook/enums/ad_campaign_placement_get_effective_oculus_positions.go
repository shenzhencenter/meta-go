package enums

type AdcampaignplacementgetEffectiveOculusPositions string

const (
	AdcampaignplacementgetEffectiveOculusPositionsTwilightDeveloperUpdate AdcampaignplacementgetEffectiveOculusPositions = "TWILIGHT_DEVELOPER_UPDATE"
	AdcampaignplacementgetEffectiveOculusPositionsTwilightFeed            AdcampaignplacementgetEffectiveOculusPositions = "TWILIGHT_FEED"
	AdcampaignplacementgetEffectiveOculusPositionsTwilightFeedSpotlight   AdcampaignplacementgetEffectiveOculusPositions = "TWILIGHT_FEED_SPOTLIGHT"
	AdcampaignplacementgetEffectiveOculusPositionsTwilightSearch          AdcampaignplacementgetEffectiveOculusPositions = "TWILIGHT_SEARCH"
	AdcampaignplacementgetEffectiveOculusPositionsTwilightSearchNullState AdcampaignplacementgetEffectiveOculusPositions = "TWILIGHT_SEARCH_NULL_STATE"
	AdcampaignplacementgetEffectiveOculusPositionsVrApps                  AdcampaignplacementgetEffectiveOculusPositions = "VR_APPS"
	AdcampaignplacementgetEffectiveOculusPositionsVrRewardedVideo         AdcampaignplacementgetEffectiveOculusPositions = "VR_REWARDED_VIDEO"
)

func (value AdcampaignplacementgetEffectiveOculusPositions) String() string {
	return string(value)
}
