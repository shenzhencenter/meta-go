package enums

type AdcampaignCostBiddingMode string

const (
	AdcampaignCostBiddingModeBalanced      AdcampaignCostBiddingMode = "BALANCED"
	AdcampaignCostBiddingModeCostFocused   AdcampaignCostBiddingMode = "COST_FOCUSED"
	AdcampaignCostBiddingModeVolumeFocused AdcampaignCostBiddingMode = "VOLUME_FOCUSED"
)

func (value AdcampaignCostBiddingMode) String() string {
	return string(value)
}
