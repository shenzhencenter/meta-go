package enums

type AdcampaignFullFunnelExplorationMode string

const (
	AdcampaignFullFunnelExplorationModeExtendedExploration AdcampaignFullFunnelExplorationMode = "EXTENDED_EXPLORATION"
	AdcampaignFullFunnelExplorationModeLimitedExploration  AdcampaignFullFunnelExplorationMode = "LIMITED_EXPLORATION"
	AdcampaignFullFunnelExplorationModeNoneExploration     AdcampaignFullFunnelExplorationMode = "NONE_EXPLORATION"
)

func (value AdcampaignFullFunnelExplorationMode) String() string {
	return string(value)
}
